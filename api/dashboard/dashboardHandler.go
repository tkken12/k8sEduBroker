package dashboard

import (
	"k8sEduBroker/common"
	"k8sEduBroker/util"
	"net/http"
)

func DashboardGetHandler(w http.ResponseWriter, r *http.Request) {

	body := DashboardBody{}
	body.InfoMerger()

	util.SendResponse(w, util.ResponseBody{
		Code:    common.HTTP_OK,
		Message: body,
	})
}
