package block_privileged_containers

import (
	"strings"

	"github.com/phoenixking25/kubectl-mtb/util"
	e2epod "k8s.io/kubernetes/test/e2e/framework/pod"
)

const (
	expectedVal = "Privileged containers are not allowed"
)

func Run(tenant string, tenantNamespace string) (bool, error) {
	kclient, err := util.ImpersonateWithUserClient(tenant, tenantNamespace)
	if err != nil {
		return false, err
	}
	// IsPrivileged set to true so that pod creation would fail
	pod := e2epod.MakeSecPod(tenantNamespace, nil, nil, true, "", false, false, nil, nil)
	_, err = kclient.CoreV1().Pods(tenantNamespace).Create(pod)
	if !strings.Contains(err.Error(), expectedVal) {
		return false, err
	}
	return true, nil
}
