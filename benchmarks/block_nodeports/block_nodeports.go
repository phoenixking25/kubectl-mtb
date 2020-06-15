package block_nodeports

import (
	"context"
	"fmt"

	"github.com/phoenixking25/kubectl-mtb/pkg/benchmark"
	"github.com/phoenixking25/kubectl-mtb/util"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/uuid"
	"k8s.io/client-go/kubernetes"
	e2edeployment "k8s.io/kubernetes/test/e2e/framework/deployment"
	imageutils "k8s.io/kubernetes/test/utils/image"
)

func init() {
	err := BNPbenchmark.ReadConfig("/home/phoenix/GO/src/github.com/phoenixking25/kubectl-mtb/benchmarks/block_nodeports/config.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
}

var BNPbenchmark = &benchmark.Benchmark{
	Run: func(tenant string, tenantNamespace string, kclient, tclient *kubernetes.Clientset) (bool, error) {

		podLabels := map[string]string{"test": "multi"}
		deploymentName := "deployment-" + string(uuid.NewUUID())
		imageName := "image-" + string(uuid.NewUUID())
		deployment := e2edeployment.NewDeployment(deploymentName, 1, podLabels, imageName, imageutils.GetE2EImage(imageutils.Nginx), "Recreate")

		_, err := tclient.AppsV1().Deployments(tenantNamespace).Create(context.TODO(), deployment, metav1.CreateOptions{DryRun: []string{metav1.DryRunAll}})
		if err != nil {
			return false, err
		}

		svcSpec := util.ServiceConfig{"NodePort", podLabels}
		svc := util.CreateServiceSpec(svcSpec)
		_, err = kclient.CoreV1().Services(tenantNamespace).Create(context.TODO(), svc, metav1.CreateOptions{DryRun: []string{metav1.DryRunAll}})

		if err == nil {
			return false, fmt.Errorf("Tenant must be unable to create service of type NodePort")
		}

		fmt.Println(err.Error())

		return true, nil
	},
}
