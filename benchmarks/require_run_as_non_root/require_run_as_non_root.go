package require_run_as_non_root

import (
	"context"
	"fmt"

	"github.com/phoenixking25/kubectl-mtb/pkg/benchmark"
	"github.com/phoenixking25/kubectl-mtb/util"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func init() {
	err := RRANRbenchmark.ReadConfig("/home/phoenix/GO/src/github.com/phoenixking25/kubectl-mtb/benchmarks/require_run_as_non_root/config.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
}

var RRANRbenchmark = &benchmark.Benchmark{
	Run: func(tenant string, tenantNamespace string, kclient, tclient *kubernetes.Clientset) (bool, error) {
		podSpec := &util.PodSpec{NS: tenantNamespace, RunAsNonRoot: false}
		err := podSpec.SetDefaults()
		if err != nil {
			return false, err
		}
		pod := util.MakeSecPod(*podSpec)
		_, err = tclient.CoreV1().Pods(tenantNamespace).Create(context.TODO(), pod, metav1.CreateOptions{DryRun: []string{metav1.DryRunAll}})
		if err == nil {
			return false, fmt.Errorf("Tenant must be unable to create pod with RunAsNonRoot set to false")
		}
		return true, nil
	},
}
