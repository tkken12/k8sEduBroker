package kubernetes

import (
	"k8sEduBroker/logger"
	"k8sEduBroker/util"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var G_k8sClient *kubernetes.Clientset

func GetK8sClient() *kubernetes.Clientset          { return G_k8sClient }
func SetK8sClient(clientSet *kubernetes.Clientset) { G_k8sClient = clientSet }

func buildConfig(configPath string) *rest.Config {
	kubeConfig, err := clientcmd.BuildConfigFromFlags("", configPath)
	if err != nil {
		logger.Fatal("invalid config file " + err.Error())
	}

	return kubeConfig
}

func NewClient() *kubernetes.Clientset {
	k8sConfigPath := util.ReadBrokerConfig().K8sConfigPath
	if k8sConfigPath == "" {
		k8sConfigPath = "/root/.kube/config"
	}

	clientSet, err := kubernetes.NewForConfig(buildConfig(k8sConfigPath))
	if err != nil {
		logger.Fatal("Failed to set kubernetes client. invalid config " + err.Error())
	}

	return clientSet
}
