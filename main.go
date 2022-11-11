package main

import (
	"k8sEduBroker/common"
	"k8sEduBroker/logger"
	"k8sEduBroker/server"
	"k8sEduBroker/util"
	"net/http"
	"strconv"
)

func main() {

	brokerConf := util.GetBrokerConf()

	listenServer := http.Server{}

	switch brokerConf.BrokerProtocol {
	case common.HTTPS:
		return

	case common.HTTP:
		listenServer = http.Server{
			Addr:    ":" + strconv.Itoa(util.GetBrokerConf().BrokerListenPort),
			Handler: server.GetRouter(),
		}

		err := listenServer.ListenAndServe()
		if err != nil {
			logger.Fatal(err.Error())
		}

	default:
		logger.Fatal("Please set valid web protocol. (https or http)")
	}

	logger.Info("broker is running")
}
