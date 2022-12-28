package kubernetes

import (
	"k8sEduBroker/logger"
	"k8sEduBroker/util"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	metrics "k8s.io/metrics/pkg/client/clientset/versioned"
)

var G_k8sClient *kubernetes.Clientset
var G_k8sMetricClient *metrics.Clientset

func GetK8sClient() *kubernetes.Clientset                   { return G_k8sClient }
func GetMetricClient() *metrics.Clientset                   { return G_k8sMetricClient }
func setK8sClient(clientSet *kubernetes.Clientset)          { G_k8sClient = clientSet }
func setK8sMetricClient(metricClientSet *metrics.Clientset) { G_k8sMetricClient = metricClientSet }

func buildConfig(configPath string) *rest.Config {
	kubeConfig, err := clientcmd.BuildConfigFromFlags("", configPath)
	if err != nil {
		logger.Fatal("invalid config file " + err.Error())
	}

	return kubeConfig
}

func NewClient() {
	k8sConfigPath := util.GetBrokerConf().K8sConfigPath
	if k8sConfigPath == "" {
		k8sConfigPath = "/root/.kube/config"
	}

	config := buildConfig(k8sConfigPath)
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		logger.Fatal("failed to set kubernetes client. invalid config " + err.Error())
	}

	setK8sClient(clientSet)
	newMetricClient(config)
}

func newMetricClient(config *rest.Config) {
	metric, err := metrics.NewForConfig(config)
	if err != nil {
		logger.Fatal("failed to set metric client." + err.Error())
	}

	setK8sMetricClient(metric)
}
