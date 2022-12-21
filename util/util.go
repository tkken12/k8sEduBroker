package util

import (
	"encoding/json"
	"io/ioutil"
	"k8sEduBroker/logger"
	"math"
	"net/http"
	"os"

	v1 "k8s.io/api/core/v1"
)

type BrokerConfig struct {
	K8sConfigPath     string `json:"kubeConfigPath"`
	BrokerListenPort  int    `json:"brokerListenPort"`
	BrokerProtocol    string `json:"brokerProtocol"`
	PromAddress       string `json:"prometheusAddress"`
	PromPort          int    `json:"prometheusPort"`
	PromProtocol      string `json:"prometheusProtocol"`
	PromTimeoutSecond int    `json:"prometheusTimeoutSecond"`
}

type ResponseBody struct {
	Code    int
	Message interface{}
}

var G_BrokerConf BrokerConfig

func ReadBrokerConfig(configPath *string) BrokerConfig {

	var brokerConfig BrokerConfig

	configFile, err := os.Open(*configPath)
	if err != nil {
		logger.Fatal("kubernetes education broker configuration file is required " + err.Error())
	}

	byteFile, err := ioutil.ReadAll(configFile)

	err = json.Unmarshal(byteFile, &brokerConfig)
	if err != nil {
		logger.Fatal(err.Error())
	}

	SetBrokerConf(brokerConfig)

	return brokerConfig
}

func GetBrokerConf() BrokerConfig         { return G_BrokerConf }
func SetBrokerConf(readConf BrokerConfig) { G_BrokerConf = readConf }

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

func KitoGI(decimalSize int64) int64 {
	return decimalSize / (1024 * 1024 * 1024)
}

func ByteToGi(decimalSize int64) int64 {
	return decimalSize / (1024 * 1024 * 1024 * 1024)
}

func RoundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func FindNodeNameByAddress(address string, nodeItem v1.Node) string {
	for _, inf := range nodeItem.Status.Addresses {
		if inf.Address == address {
			return nodeItem.Name
		}
	}

	return ""
}
