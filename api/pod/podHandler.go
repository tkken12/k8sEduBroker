package pod

import (
	"encoding/json"
	"io/ioutil"
	"k8sEduBroker/logger"
	"k8sEduBroker/util"
	"net/http"

	v1 "k8s.io/api/core/v1"
)

type PodRequestStruct struct {
	Command string           `json:"command"`
	Type    string           `json:"type"`
	Params  PodRequestParams `json:"params"`
}

type PodRequestParams struct {
	PodName       string           `json:"podName"`
	PodNamespaces []string         `json:"podNamespaces"`
	PodAddress    string           `json:"podAddress"`
	VolumeMount   []v1.VolumeMount `json:"volumeMount"`
	Volumes       []v1.Volume      `json:"volumes"`
	NodeName      string           `json:"nodeName"`
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

/*func PodGetHandler(w http.ResponseWriter, r *http.Request) {
	podGetInfo := readGetRequest(w, r)

}*/

func PodHandler(w http.ResponseWriter, r *http.Request) {
	podInfo := readRequest(w, r)

	switch podInfo.Command {
	case CREATE:
		return

	case GET:
		return
	}
}

func PodGetHandler(w http.ResponseWriter, r *http.Request) {
	// podInfo := readGetRequest(w, r)

	util.SendResponse(w, util.ResponseBody{
		Code:    200,
		Message: "test",
	})
}
