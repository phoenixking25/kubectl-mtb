package block_add_capabilities

import (
	"fmt"
	"strings"

	"github.com/phoenixking25/kubectl-mtb/pkg/benchmark"
	"github.com/phoenixking25/kubectl-mtb/util"
	"k8s.io/client-go/kubernetes"
)

const (
	expectedVal = "capability may not be added"
)

func init() {
	err := BACbenchmark.ReadConfig("/home/phoenix/GO/src/github.com/phoenixking25/kubectl-mtb/benchmarks/block_add_capabilities/config.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
}

var BACbenchmark = &benchmark.Benchmark{
	Run: func(tenant string, tenantNamespace string, kclient, tclient *kubernetes.Clientset) (bool, error) {
		pod := util.MakeSecPod(tenantNamespace, nil, nil, false, "", false, false, nil, nil, false, "SETPCAP")
		_, err := tclient.CoreV1().Pods(tenantNamespace).Create(pod)
		if err == nil {
			return false, fmt.Errorf("Tenant must be unable to create pod with add capabilities")
		} else if !strings.Contains(err.Error(), expectedVal) {
			return false, err
		}
		return true, nil
	},
}
