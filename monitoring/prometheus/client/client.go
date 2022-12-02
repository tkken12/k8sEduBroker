package prometheus

import (
	"fmt"
	"k8sEduBroker/logger"
	"k8sEduBroker/util"
	"strconv"

	prom "github.com/prometheus/client_golang/api"
)

var G_Prometheus prom.Client

func GetPrometheusClient() prom.Client       { return G_Prometheus }
func SetPrometheusClient(client prom.Client) { G_Prometheus = client }
func NewPrometheusClient() prom.Client {

	promClient, err := prom.NewClient(prom.Config{
		Address: fmt.Sprintf("%s://%s:%s",
			util.GetBrokerConf().PromProtocol,
			util.GetBrokerConf().PromAddress,
			strconv.Itoa(util.ReadBrokerConfig().PromPort),
		),
	})
	if err != nil {
		logger.Fatal("failed to set prometheus client\n" + err.Error())
	}

	logger.Info("set prometheus client done...")
	return promClient
}
