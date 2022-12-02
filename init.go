package main

import (
	"k8sEduBroker/logger"
	"k8sEduBroker/util"

	api "k8sEduBroker/api"
	kClient "k8sEduBroker/kubernetes/client"
	pClient "k8sEduBroker/monitoring/prometheus/client"

	"github.com/gorilla/mux"
)

func init() {

	logger.LoggerInit()
	logger.Info("logger initialize done.")

	util.SetBrokerConf(util.ReadBrokerConfig())
	logger.Info("configuration read succeed")

	kClient.NewClient()
	logger.Info("k8s initialize done.")

	pClient.NewPrometheusClient()
	logger.Info("prometheus initialize done.")
}

func GetServer() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, handler := range api.Handlers {
		for _, elem := range handler {
			router.HandleFunc(elem.Path, elem.HandleFunc).Methods(elem.RestMethod)
		}
	}

	return router
}
