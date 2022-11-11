package util

import (
	"encoding/json"
	"io/ioutil"
	"k8sEduBroker/logger"
	"net/http"
	"os"

	"k8s.io/client-go/kubernetes"
)

type BrokerConfig struct {
	K8sConfigPath    string `json:"kubeConfigPath"`
	BrokerListenPort int    `json:"brokerListenPort"`
	BrokerProtocol   string `json:"brokerProtocol"`
	PromAddress      string `json:"prometheusAddress"`
	PromPort         int    `json:"prometheusPort"`
	PromProtocol     string `json:"prometheusProtocol"`
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

type ResponseBody struct {
	Code    int
	Message interface{}
}

func SendResponse(w http.ResponseWriter, resBody ResponseBody) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	resBodyMarshal, err := json.Marshal(resBody)
	if err != nil {
		logger.Warn("failed to send message\n" + err.Error())
	}

	w.Write(resBodyMarshal)
	return
}
