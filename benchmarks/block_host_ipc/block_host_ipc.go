package block_host_ipc

import (
	"fmt"
	"strings"

	"github.com/phoenixking25/kubectl-mtb/pkg/benchmark"
	"github.com/phoenixking25/kubectl-mtb/util"
	"k8s.io/client-go/kubernetes"
)

const (
	expectedVal = "Host IPC is not allowed to be used"
)

func init() {
	err := BHIPCbenchmark.ReadConfig("/home/phoenix/GO/src/github.com/phoenixking25/kubectl-mtb/benchmarks/block_host_ipc/config.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
}

var BHIPCbenchmark = &benchmark.Benchmark{
	Run: func(tenant string, tenantNamespace string, kclient, tclient *kubernetes.Clientset) (bool, error) {
		pod := util.MakeSecPod(tenantNamespace, nil, nil, true, "", true, false, nil, nil, true, "")
		_, err := tclient.CoreV1().Pods(tenantNamespace).Create(pod)
		if err == nil {
			return false, fmt.Errorf("Tenant must be unable to create pod with HostIPC set to true")
		} else if !strings.Contains(err.Error(), expectedVal) {
			return false, err
		}
		return true, nil
	},
}
