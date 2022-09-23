package pod

import (
	"encoding/json"
	"io/ioutil"
	"k8sEduBroker/logger"
	"net/http"
)

type PodStruct struct {
	Command string `json:"command"`
}

func readRequest(w http.ResponseWriter, r *http.Request) PodStruct {

	podInfo := PodStruct{}

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

	}
}
