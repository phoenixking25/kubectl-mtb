package util

import (
	"fmt"
	"os"
	"path/filepath"

	kubernetes "k8s.io/client-go/kubernetes"
	clientcmd "k8s.io/client-go/tools/clientcmd"
)

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
