package block_cluster_resources

import (
	"context"
	"fmt"

	"github.com/phoenixking25/kubectl-mtb/pkg/benchmark"
	authorizationv1 "k8s.io/api/authorization/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes"
)

type groupResource struct {
	APIGroup    string
	APIResource metav1.APIResource
}

func init() {
	err := BCRbenchmark.ReadConfig("/home/phoenix/GO/src/github.com/phoenixking25/kubectl-mtb/benchmarks/block_cluster_resources/config.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
}

var (
	verbs = []string{"get", "list", "create", "update", "patch", "watch", "delete", "deletecollection"}

	BCRbenchmark = &benchmark.Benchmark{
		Run: func(tenant, tenantNamespace string, kclient, tclient *kubernetes.Clientset) (bool, error) {
			resources := []groupResource{}

			lists, err := kclient.Discovery().ServerPreferredResources()
			if err != nil {
				return false, err
			}

			for _, list := range lists {
				if len(list.APIResources) == 0 {
					continue
				}
				gv, err := schema.ParseGroupVersion(list.GroupVersion)
				if err != nil {
					continue
				}
				for _, resource := range list.APIResources {
					if len(resource.Verbs) == 0 {
						continue
					}

					if resource.Namespaced {
						continue
					}
					resources = append(resources, groupResource{
						APIGroup:    gv.Group,
						APIResource: resource,
					})
				}
			}

			err = RunAccessCheck(tclient, resources)
			if err != nil {
				return false, err
			}

			return true, nil
		},
	}
)

func RunAccessCheck(tclient *kubernetes.Clientset, resources []groupResource) error {
	var sar *authorizationv1.SelfSubjectAccessReview

	// Todo for non resource url
	for _, resource := range resources {
		for _, verb := range verbs {
			sar = &authorizationv1.SelfSubjectAccessReview{
				Spec: authorizationv1.SelfSubjectAccessReviewSpec{
					ResourceAttributes: &authorizationv1.ResourceAttributes{
						Namespace:   "",
						Verb:        verb,
						Group:       resource.APIGroup,
						Resource:    resource.APIResource.Name,
						Subresource: "",
						Name:        "",
					},
				},
			}

			response, err := tclient.AuthorizationV1().SelfSubjectAccessReviews().Create(context.TODO(), sar, metav1.CreateOptions{DryRun: []string{metav1.DryRunAll}})
			if err != nil {
				return err
			}

			if response.Status.Allowed {
				return fmt.Errorf("Tenant can %s %s", verb, resource.APIResource.Name)
			} else {
				return nil
			}
		}
	}
	return nil
}
