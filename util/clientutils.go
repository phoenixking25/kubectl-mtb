package util

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	authorizationv1 "k8s.io/api/authorization/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"
	clientcmd "k8s.io/client-go/tools/clientcmd"
)

type GroupResource struct {
	APIGroup    string
	APIResource metav1.APIResource
}

func getKubeConfig() string {
	return filepath.Join(
		os.Getenv("HOME"), ".kube", "config",
	)
}

func GetContextFromKubeconfig(kubeconfigpath string) (string, error) {
	apiConfig := clientcmd.GetConfigFromFileOrDie(kubeconfigpath)

	if apiConfig.CurrentContext == "" {
		return "", fmt.Errorf("current-context is not set in %s", kubeconfigpath)
	}

	return apiConfig.CurrentContext, nil
}

func NewKubeClient() (*kubernetes.Clientset, error) {
	kubeconfig := getKubeConfig()

	clientConfig, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, err
	}

	kclient, err := kubernetes.NewForConfig(clientConfig)
	if err != nil {
		return nil, err
	}

	return kclient, nil
}

func ImpersonateWithUserClient(serviceaccountname string, namespace string) (*kubernetes.Clientset, error) {
	kubeconfig := getKubeConfig()

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, err
	}

	impersonateUsername := "system:serviceaccount:" + namespace + ":" + serviceaccountname
	config.Impersonate.UserName = impersonateUsername

	kclient, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return kclient, nil
}

func RunAccessCheck(client *kubernetes.Clientset, namespace string, resource GroupResource, verb string) (bool, string, error) {
	var sar *authorizationv1.SelfSubjectAccessReview

	// Todo for non resource url
	sar = &authorizationv1.SelfSubjectAccessReview{
		Spec: authorizationv1.SelfSubjectAccessReviewSpec{
			ResourceAttributes: &authorizationv1.ResourceAttributes{
				Namespace:   namespace,
				Verb:        verb,
				Group:       resource.APIGroup,
				Resource:    resource.APIResource.Name,
				Subresource: "",
				Name:        "",
			},
		},
	}

	response, err := client.AuthorizationV1().SelfSubjectAccessReviews().Create(context.TODO(), sar, metav1.CreateOptions{DryRun: []string{metav1.DryRunAll}})
	if err != nil {
		return false, "", err
	}

	if response.Status.Allowed {
		return true, fmt.Sprintf("User can %s %s", verb, resource.APIResource.Name), nil
	}

	return false, fmt.Sprintf("User cannot %s %s", verb, resource.APIResource.Name), nil
}
