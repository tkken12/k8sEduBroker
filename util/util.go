package util

import (
	"encoding/json"
	"io/ioutil"
	"k8sEduBroker/logger"
	"os"

	"k8s.io/client-go/kubernetes"
)

type BrokerConfig struct {
	K8sConfigPath    string `json:"kubeConfigPath"`
	BrokerListenPort int    `json:"brokerListenPort"`
	BrokerProtocol   string `json:"brokerProtocol"`
}

var G_k8sClient *kubernetes.Clientset
var G_BrokerConf BrokerConfig

func GetK8sClient() *kubernetes.Clientset          { return G_k8sClient }
func SetK8sClient(clientSet *kubernetes.Clientset) { G_k8sClient = clientSet }

func ReadBrokerConfig() BrokerConfig {

	var brokerConfig BrokerConfig

	configFile, err := os.Open("/.k8sEdu/config.json")
	if err != nil {
		logger.Fatal("kubernetes education broker configuration file is required " + err.Error())
	}

	byteFile, err := ioutil.ReadAll(configFile)

	err = json.Unmarshal(byteFile, &brokerConfig)
	if err != nil {
		logger.Fatal(err.Error())
	}

	SetBrokerConf(brokerConfig)

	logger.Info("read broker config done... ")
	return brokerConfig
}

func GetBrokerConf() BrokerConfig         { return G_BrokerConf }
func SetBrokerConf(readConf BrokerConfig) { G_BrokerConf = readConf }
