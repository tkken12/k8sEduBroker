package main

import (
	"k8sEduBroker/util"
	"log"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func init() {

	kubeConfigPath := getKubeConfigPath()
	if kubeConfigPath == "" {
		kubeConfigPath = getRootKubeConfigPath()
	}

	util.SetK8sClient(newClient(buildConfig(kubeConfigPath)))
}

func getKubeConfigPath() string {
	return ""
}

func getRootKubeConfigPath() string { return "/root/.kube/config" }

func buildConfig(configPath string) *rest.Config {
	kubeConfig, err := clientcmd.BuildConfigFromFlags("", configPath)
	if err != nil {
		log.Fatal("Kubernetes configuration file is required.")
	}

	return kubeConfig
}

func newClient(k8sConfig *rest.Config) *kubernetes.Clientset {
	clientSet, err := kubernetes.NewForConfig(k8sConfig)
	if err != nil {
		log.Fatal("Failed to set kubernetes client. invalid config", err)
	}

	return clientSet
}

func initLogger() {

}
