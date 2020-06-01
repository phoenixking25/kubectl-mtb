package require_run_as_non_root

import (
	"fmt"
	"strings"

	"github.com/phoenixking25/kubectl-mtb/pkg/benchmark"
	"github.com/phoenixking25/kubectl-mtb/util"
	"k8s.io/client-go/kubernetes"
)

const (
	expectedVal = "spec.containers[0].securityContext.runAsNonRoot: Invalid value: false: must be true"
)

func init() {
	err := RRANRbenchmark.ReadConfig("/home/phoenix/GO/src/github.com/phoenixking25/kubectl-mtb/benchmarks/require_run_as_non_root/config.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
}

var RRANRbenchmark = &benchmark.Benchmark{
	Run: func(tenant string, tenantNamespace string, kclient, tclient *kubernetes.Clientset) (bool, error) {
		pod := util.MakeSecPod(tenantNamespace, nil, nil, false, "", false, false, nil, nil, false)
		_, err := tclient.CoreV1().Pods(tenantNamespace).Create(pod)
		if err == nil {
			return false, fmt.Errorf("Tenant must be unable to create pod with RunAsNonRoot set to false")
		} else if !strings.Contains(err.Error(), expectedVal) {
			return false, err
		}
		return true, nil
	},
}
