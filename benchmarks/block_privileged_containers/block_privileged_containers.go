package block_privileged_containers

import (
	"fmt"
	"strings"

	"github.com/phoenixking25/kubectl-mtb/pkg/benchmark"
	"github.com/phoenixking25/kubectl-mtb/util"
	"k8s.io/client-go/kubernetes"
)

const (
	expectedVal = "Privileged containers are not allowed"
)

func init() {
	err := BPCbenchmark.ReadConfig("/home/phoenix/GO/src/github.com/phoenixking25/kubectl-mtb/benchmarks/block_privileged_containers/config.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
}

var BPCbenchmark = &benchmark.Benchmark{
	Run: func(tenant string, tenantNamespace string, kclient, tclient *kubernetes.Clientset) (bool, error) {
		// IsPrivileged set to true so that pod creation would fail
		pod := util.MakeSecPod(tenantNamespace, nil, nil, true, "", false, false, nil, nil, true)
		_, err := tclient.CoreV1().Pods(tenantNamespace).Create(pod)
		if err == nil {
			return false, fmt.Errorf("Tenant must be unable to create pod that sets privileged to true")
		} else if !strings.Contains(err.Error(), expectedVal) {
			return false, err
		}
		return true, nil
	},
}
