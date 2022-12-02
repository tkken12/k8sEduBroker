package api

import (
	"encoding/json"
	"k8sEduBroker/api/pod"
	"k8sEduBroker/common"
	"k8sEduBroker/logger"
	"net/http"
)

type ResponseBody struct {
	Code    int
	Message interface{}
}

type RequestHandler struct {
	Path       string
	HandleFunc func(http.ResponseWriter, *http.Request)
	RestMethod string
}

var PodHandler = []RequestHandler{
	{"/api/v1/pod", pod.PodHandler, common.POST},
	{"/api/v1/pod", pod.PodHandler, common.PUT},
	{"/api/v1/pod", pod.PodHandler, common.DELETE},
	{"/api/v1/pod", pod.PodHandler, common.GET},
	// {"api/v1/node", }
}

var Handlers = [][]RequestHandler{
	PodHandler,
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
