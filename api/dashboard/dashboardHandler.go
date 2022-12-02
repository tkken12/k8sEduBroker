package dashboard

import (
	"k8sEduBroker/common"
	"k8sEduBroker/util"
	"net/http"
)

type DashboardBody struct {
	NodeTotal   int `json:"nodeTotal"`
	MasterTotal int `json:"masterTotal"`
	WorkerTotal int `json:"workerTotal"`
	CPUTotal    int `json:"cpuTotal"`
	MemoryTotal int `json:"memoryTotal"`
	PodTotal    int `json:"podTotal"`
}

func DashboardGetHandler(w http.ResponseWriter, r *http.Request) {

	body := DashboardBody{}
	body.InfoMerger()

	util.SendResponse(w, util.ResponseBody{
		Code:    common.HTTP_OK,
		Message: body,
	})
}
