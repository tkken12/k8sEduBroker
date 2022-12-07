package dashboard

import (
	"k8sEduBroker/common"
	reqNode "k8sEduBroker/kubernetes/node"
	reqPod "k8sEduBroker/kubernetes/pod"

	v1 "k8s.io/api/core/v1"
)

func (dashboardBody *DashboardBody) TotalMerger() {

	podErrFlag, nodeErrFlag := false, false // pod, nodes 모두 에러일 때

	pods, err := reqPod.GetPods()
	if err != nil {
		dashboardBody.PodTotal = 0
		podErrFlag = true
	} else {
		dashboardBody.PodTotal = len(pods.Items)
	}

	nodes, err := reqNode.GetNodeList()
	if err != nil {
		dashboardBody.NodeTotal = 0
		dashboardBody.MasterTotal = 0
		dashboardBody.WorkerTotal = 0
		nodeErrFlag = true
	} else {
		dashboardBody.NodeTotal = len(nodes.Items)
	}

	if podErrFlag && nodeErrFlag {
		return
	}

	dashboardBody.MasterTotal, dashboardBody.WorkerTotal = GetNodeRoleTotal(nodes)
}

func GetNodeRoleTotal(nodes *v1.NodeList) (int, int) {
	masterTotal := 0
	workerTotal := 0
	for _, node := range nodes.Items {
		for key, _ := range node.Labels {
			if key == common.LABEL_MASTER_ROLE {
				masterTotal++
			} else {
				workerTotal++
			}
		}
	}

	return masterTotal, workerTotal
}

func GetAllNodeProcessor(nodes *v1.NodeList) (int64, int64, int64) {

	var cpuTotal int64 = 0
	var memTotal int64 = 0
	var diskTotal int64 = 0

	for _, node := range nodes.Items {
		cpuCapa, _ := node.Status.Capacity.Cpu().AsInt64()
		memCapa, _ := node.Status.Capacity.Memory().AsInt64() // kibibyte

		cpuTotal += cpuCapa
		memTotal += memCapa
	}

}
