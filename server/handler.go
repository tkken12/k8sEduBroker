package server

import (
	"encoding/json"
	"k8sEduBroker/api/pod"
	"k8sEduBroker/logger"
	"net/http"
)

type ResponseBody struct {
	Code    int
	Message interface{}
}

var PodHandler = []RequestHandler{
	{"/api/v1/pod", pod.PodHandler, POST},
	{"/api/v1/pod", pod.PodHandler, PUT},
	{"/api/v1/pod", pod.PodHandler, DELETE},
	{"/api/v1/pod", pod.PodHandler, GET},
	// {"api/v1/node", }
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
