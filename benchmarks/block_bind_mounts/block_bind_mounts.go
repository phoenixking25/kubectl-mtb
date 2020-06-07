package block_bind_mounts

import (
	"context"
	"fmt"

	"github.com/phoenixking25/kubectl-mtb/pkg/benchmark"
	"github.com/phoenixking25/kubectl-mtb/util"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func init() {
	err := BBMbenchmark.ReadConfig("/home/phoenix/GO/src/github.com/phoenixking25/kubectl-mtb/benchmarks/block_bind_mounts/config.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
}

var BBMbenchmark = &benchmark.Benchmark{
	Run: func(tenant string, tenantNamespace string, kclient, tclient *kubernetes.Clientset) (bool, error) {

		// Host path
		inlineVolumeSources := []*v1.VolumeSource{
			{
				HostPath: &v1.HostPathVolumeSource{
					Path: "/tmp/busybox",
				},
			},
		}

		podSpec := &util.PodSpec{NS: tenantNamespace, InlineVolumeSources: inlineVolumeSources}
		err := podSpec.SetDefaults()
		if err != nil {
			return false, err
		}

		pod := util.MakeSecPod(*podSpec)
		_, err = tclient.CoreV1().Pods(tenantNamespace).Create(context.TODO(), pod, metav1.CreateOptions{DryRun: []string{metav1.DryRunAll}})
		if err == nil {
			return false, fmt.Errorf("Tenant must be unable to create pod with host-path volume")
		}
		return true, nil
	},
}
