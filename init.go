package main

import (
	"k8sEduBroker/logger"
	"k8sEduBroker/util"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func init() {

	logger.LoggerInit()

	kubeConfigPath := getKubeConfigPath()
	if kubeConfigPath == "" {
		kubeConfigPath = getRootKubeConfigPath()
	}

	util.SetK8sClient(newClient(buildConfig(kubeConfigPath)))
	logger.Info("set k8s client done... ")
}

func getKubeConfigPath() string {
	kubeConfigPath := util.ReadBrokerConfig().K8sConfigPath

	if kubeConfigPath != "" {
		return kubeConfigPath
	}

	return ""
}

func getRootKubeConfigPath() string { return "/root/.kube/config" }

func buildConfig(configPath string) *rest.Config {
	kubeConfig, err := clientcmd.BuildConfigFromFlags("", configPath)
	if err != nil {
		logger.Fatal("invalid config file " + err.Error())
	}

	return kubeConfig
}

func newClient(k8sConfig *rest.Config) *kubernetes.Clientset {
	clientSet, err := kubernetes.NewForConfig(k8sConfig)
	if err != nil {
		logger.Fatal("Failed to set kubernetes client. invalid config " + err.Error())
	}

	return clientSet
}
