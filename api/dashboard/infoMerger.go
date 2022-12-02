package dashboard

import (
	reqPod "k8sEduBroker/api/pod"
)

func (body *DashboardBody) InfoMerger() {
	body.PodTotal = len(reqPod.GetAllPods().Items)
}
