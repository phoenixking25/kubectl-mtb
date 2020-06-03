package block_host_net_ports

import (
	"context"
	"fmt"
	"strings"

	"github.com/phoenixking25/kubectl-mtb/pkg/benchmark"
	"github.com/phoenixking25/kubectl-mtb/util"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

const (
	expectedHostNetworkVal = "Use of hostNetwork is not allowed"
	expectedHostPortVal    = "Host port is not allowed to be used"
)

func init() {
	err := BHNPbenchmark.ReadConfig("/home/phoenix/GO/src/github.com/phoenixking25/kubectl-mtb/benchmarks/block_host_net_ports/config.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
}

var BHNPbenchmark = &benchmark.Benchmark{
	Run: func(tenant string, tenantNamespace string, kclient, tclient *kubernetes.Clientset) (bool, error) {
		//Tenant containers cannot use host networking
		pod := util.MakeSecPod(util.PodSpec{NS: tenantNamespace, HostNetwork: true, Ports: nil})
		_, err := tclient.CoreV1().Pods(tenantNamespace).Create(context.TODO(), pod, metav1.CreateOptions{DryRun: []string{metav1.DryRunAll}})
		if err == nil {
			return false, fmt.Errorf("Tenant must be unable to create pod with host networking set to true")
		} else if !strings.Contains(err.Error(), expectedHostNetworkVal) {
			return false, err
		}

		//Tenant should not be allowed to use host ports
		ports := []v1.ContainerPort{
			{
				HostPort:      8086,
				ContainerPort: 8086,
			},
		}
		pod = util.MakeSecPod(util.PodSpec{NS: tenantNamespace, HostNetwork: false, Ports: ports})
		_, err = tclient.CoreV1().Pods(tenantNamespace).Create(context.TODO(), pod, metav1.CreateOptions{DryRun: []string{metav1.DryRunAll}})
		if err == nil {
			return false, fmt.Errorf("Tenant must be unable to create pod with defined host ports")
		} else if !strings.Contains(err.Error(), expectedHostNetworkVal) {
			return false, err
		}

		return true, nil
	},
}
