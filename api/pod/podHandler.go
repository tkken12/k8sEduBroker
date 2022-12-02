package pod

import (
	"encoding/json"
	"io/ioutil"
	"k8sEduBroker/logger"
	"net/http"

	v1 "k8s.io/api/core/v1"
)

type PodRequestStruct struct {
	Command string           `json:"command"`
	Params  PodRequestParams `json:"params"`
}

type PodRequestParams struct {
	PodName       string           `json:"podName"`
	PodNamespaces []string         `json:"podNamespace"`
	PodAddress    string           `json:"podAddress"`
	VolumeMount   []v1.VolumeMount `json:"volumeMount"`
	Volumes       []v1.Volume      `json:"volumes"`
	NodeName      string           `json:"nodeName"`
}

func readGetRequest(w http.ResponseWriter, r *http.Request) PodRequestStruct {

	var podInfo = PodRequestStruct{}

	for k, v := range r.URL.Query() {
		switch k {
		case GET:
			for _, elem := range v {
				json.Unmarshal([]byte(elem), &podInfo)
			}

		}
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
