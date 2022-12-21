package pod

import (
	"encoding/json"
	"io/ioutil"
	"k8sEduBroker/common"
	"k8sEduBroker/kubernetes/pod"
	"k8sEduBroker/logger"
	"k8sEduBroker/util"
	"net/http"
)

type PodRequestStruct struct {
	Command string           `json:"command"`
	Type    string           `json:"type"`
	Params  PodRequestParams `json:"params"`
}

type PodRequestParams struct {
	PodName       string   `json:"podName"`
	PodNamespaces []string `json:"podNamespaces"`
	PodAddress    string   `json:"podAddress"`
	NodeName      string   `json:"nodeName"`
}

func readGetRequest(w http.ResponseWriter, r *http.Request) PodRequestStruct {

	var podInfo = PodRequestStruct{}

	switch r.URL.Query()["command"][0] {
	case GET:
		podInfo.Command = r.URL.Query()["command"][0]
		podInfo.Type = r.URL.Query()["type"][0]
		json.Unmarshal([]byte(r.URL.Query()["params"][0]), &podInfo.Params)
	}

	return podInfo
}

func readRequest(w http.ResponseWriter, r *http.Request) PodRequestStruct {

	podInfo := PodRequestStruct{}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Warn("failed to parse request")
	}

	json.Unmarshal(reqBody, &podInfo)

	return podInfo
}

func PodHandler(w http.ResponseWriter, r *http.Request) {
	podInfo := readRequest(w, r)

	switch podInfo.Command {
	case CREATE:
		return

	case DELETE:
		return

	case EDIT:
		return
	}
}

func PodGetHandler(w http.ResponseWriter, r *http.Request) {
	podInfo := readGetRequest(w, r)

	var code int
	var message pod.PodGetBody

	if len(podInfo.Params.PodNamespaces) == 0 {
		pods, err := pod.GetPods()
		if err != nil {
			code = common.HTTP_INTERNAL_ERROR
		}

		for _, pod := range pods.Items {
			message.PodParser(pod)
		}
	}

	util.SendResponse(w, util.ResponseBody{
		Code:    code,
		Message: message,
	})
}
