package block_cluster_resources

import (
	"fmt"

	"github.com/phoenixking25/kubectl-mtb/util"
	authorizationv1 "k8s.io/api/authorization/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes"
)

type groupResource struct {
	APIGroup    string
	APIResource metav1.APIResource
}

var verbs = []string{"get", "list", "create", "update", "patch", "watch", "delete", "deletecollection"}

func Run(tenant string, tenantNamespace string) (bool, error) {
	resources := []groupResource{}

	kclient, err := util.NewKubeClient()
	if err != nil {
		return false, err
	}

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

	tclient, err := util.ImpersonateWithUserClient(tenant, tenantNamespace)
	if err != nil {
		return false, err
	}

	err = RunAccessCheck(tclient, resources)
	if err != nil {
		return false, err
	}

	return true, nil
}

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

			response, err := tclient.AuthorizationV1().SelfSubjectAccessReviews().Create(sar)
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
