package block_ns_quota

import (
	"fmt"

	"github.com/phoenixking25/kubectl-mtb/pkg/benchmark"
	"github.com/phoenixking25/kubectl-mtb/util"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func init() {
	err := BNQbenchmark.ReadConfig("/home/phoenix/GO/src/github.com/phoenixking25/kubectl-mtb/benchmarks/block_ns_quota/config.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
}

var BNQbenchmark = &benchmark.Benchmark{
	Run: func(tenant string, tenantNamespace string, kclient, tclient *kubernetes.Clientset) (bool, error) {
		resources := []util.GroupResource{
			{
				APIGroup: "",
				APIResource: metav1.APIResource{
					Name: "resourcequotas",
				},
			},
		}
		verbs := []string{"create", "update", "patch", "delete", "deletecollection"}
		for _, resource := range resources {
			for _, verb := range verbs {
				access, msg, err := util.RunAccessCheck(tclient, tenantNamespace, resource, verb)
				if err != nil {
					fmt.Println(err.Error())
				}
				if access {
					return false, fmt.Errorf(msg)
				}
			}
		}
		return true, nil
	},
}
