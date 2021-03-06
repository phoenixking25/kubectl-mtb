package block_privileged_containers

import (
	"context"
	"fmt"

	"github.com/phoenixking25/kubectl-mtb/pkg/benchmark"
	"github.com/phoenixking25/kubectl-mtb/util"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
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
		podSpec := &util.PodSpec{NS: tenantNamespace, IsPrivileged: true}
		err := podSpec.SetDefaults()
		if err != nil {
			return false, err
		}
		pod := util.MakeSecPod(*podSpec)
		_, err = tclient.CoreV1().Pods(tenantNamespace).Create(context.TODO(), pod, metav1.CreateOptions{DryRun: []string{metav1.DryRunAll}})
		if err == nil {
			return false, fmt.Errorf("Tenant must be unable to create pod that sets privileged to true")
		}
		return true, nil
	},
}
