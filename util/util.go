package util

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"k8s.io/client-go/kubernetes"
)

type BrokerConfig struct {
	K8sConfigPath string `json:"kubeConfigPath"`
}

var G_k8sClient *kubernetes.Clientset

func GetK8sClient() *kubernetes.Clientset          { return G_k8sClient }
func SetK8sClient(clientSet *kubernetes.Clientset) { G_k8sClient = clientSet }

func ReadBrokerConfig() BrokerConfig {

	var brokerConfig BrokerConfig

	configFile, err := os.Open("/.k8sEdu/config.json")
	if err != nil {
		log.Fatal("kubernetes education broker configuration file is required ", err)
	}

	byteFile, err := ioutil.ReadAll(configFile)

	err = json.Unmarshal(byteFile, &brokerConfig)
	if err != nil {
		log.Println(err)
	}

	return brokerConfig
}
