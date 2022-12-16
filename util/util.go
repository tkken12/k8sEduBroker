package util

import (
	"encoding/json"
	"io/ioutil"
	"k8sEduBroker/logger"
	"math"
	"net/http"
	"os"
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
