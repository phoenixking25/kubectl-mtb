package block_bind_mounts

import (
	"fmt"
	"strings"

	"github.com/phoenixking25/kubectl-mtb/pkg/benchmark"
	"github.com/phoenixking25/kubectl-mtb/util"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
)

const (
	expectedVal = "Host path volumes are not allowed"
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

		pod := util.MakeSecPod(tenantNamespace, nil, inlineVolumeSources, false, "", false, false, nil, nil, true, "")
		_, err := tclient.CoreV1().Pods(tenantNamespace).Create(pod)
		if err == nil {
			return false, fmt.Errorf("Tenant must be unable to create pod with host-path volume")
		} else if !strings.Contains(err.Error(), expectedVal) {
			return false, err
		}
		return true, nil
	},
}
