package block_add_capabilities

import (
	"context"
	"fmt"

	"github.com/phoenixking25/kubectl-mtb/internal/box"
	"github.com/phoenixking25/kubectl-mtb/pkg/benchmark"
	"github.com/phoenixking25/kubectl-mtb/util"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func init() {
	err := BACbenchmark.ReadConfig(string(box.Get("/block_add_capabilities/config.yaml")))
	if err != nil {
		fmt.Println(err.Error())
	}
}

var BACbenchmark = &benchmark.Benchmark{
	Run: func(tenant string, tenantNamespace string, kclient, tclient *kubernetes.Clientset) (bool, error) {

		podSpec := &util.PodSpec{NS: tenantNamespace, Capability: []v1.Capability{"SETPCAP"}}
		err := podSpec.SetDefaults()
		if err != nil {
			return false, err
		}

		pod := util.MakeSecPod(*podSpec)
		_, err = tclient.CoreV1().Pods(tenantNamespace).Create(context.TODO(), pod, metav1.CreateOptions{DryRun: []string{metav1.DryRunAll}})
		if err == nil {
			return false, fmt.Errorf("Tenant must be unable to create pod with add capabilities")
		}
		return true, nil
	},
}
