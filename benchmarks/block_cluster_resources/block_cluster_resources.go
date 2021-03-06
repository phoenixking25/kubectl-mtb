package block_cluster_resources

import (
	"fmt"

	"github.com/phoenixking25/kubectl-mtb/pkg/benchmark"
	"github.com/phoenixking25/kubectl-mtb/util"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes"
)

func init() {
	err := BCRbenchmark.ReadConfig("/home/phoenix/GO/src/github.com/phoenixking25/kubectl-mtb/benchmarks/block_cluster_resources/config.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
}

var (
	verbs = []string{"get", "update"}

	BCRbenchmark = &benchmark.Benchmark{
		Run: func(tenant, tenantNamespace string, kclient, tclient *kubernetes.Clientset) (bool, error) {
			resources := []util.GroupResource{}

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
					resources = append(resources, util.GroupResource{
						APIGroup:    gv.Group,
						APIResource: resource,
					})
				}
			}

			for _, resource := range resources {
				for _, verb := range verbs {
					access, msg, err := util.RunAccessCheck(tclient, tenantNamespace, resource, verb)
					if err != nil {
						return false, err
					}
					if access {
						return false, fmt.Errorf(msg)
					}

				}
			}
			return true, nil
		},
	}
)
