package block_host_pid

import (
	"fmt"
	"strings"

	"github.com/phoenixking25/kubectl-mtb/pkg/benchmark"
	"github.com/phoenixking25/kubectl-mtb/util"
	"k8s.io/client-go/kubernetes"
)

const (
	expectedVal = "Host PID is not allowed to be used"
)

func init() {
	err := BHPIDbenchmark.ReadConfig("/home/phoenix/GO/src/github.com/phoenixking25/kubectl-mtb/benchmarks/block_host_pid/config.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
}

var BHPIDbenchmark = &benchmark.Benchmark{
	Run: func(tenant string, tenantNamespace string, kclient, tclient *kubernetes.Clientset) (bool, error) {
		pod := util.MakeSecPod(tenantNamespace, nil, nil, true, "", false, true, nil, nil, true, "")
		_, err := tclient.CoreV1().Pods(tenantNamespace).Create(pod)
		if err == nil {
			return false, fmt.Errorf("Tenant must be unable to create pod with HostPID set to true")
		} else if !strings.Contains(err.Error(), expectedVal) {
			return false, err
		}
		return true, nil
	},
}
