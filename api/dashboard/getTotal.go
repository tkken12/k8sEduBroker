package dashboard

import (
	"k8sEduBroker/common"
	reqPod "k8sEduBroker/kubernetes/pod"
	"k8sEduBroker/util"

	pQuery "k8sEduBroker/monitoring/prometheus/query"

	"github.com/prometheus/common/model"
	v1 "k8s.io/api/core/v1"
)

func (dashboardBody *DashboardBody) TotalMerger(nodes *v1.NodeList, nodeErr error) {

	podErrFlag, nodeErrFlag := false, false // pod, nodes 모두 에러일 때

	pods, err := reqPod.GetPods()
	if err != nil {
		dashboardBody.Total.PodTotal = 0
		podErrFlag = true
	} else {
		dashboardBody.Total.PodTotal = len(pods.Items)
	}

	// nodes, err := reqNode.GetNodeList()
	if nodeErr != nil {
		dashboardBody.Total.NodeTotal = 0
		dashboardBody.Total.MasterTotal = 0
		dashboardBody.Total.WorkerTotal = 0
		nodeErrFlag = true

	} else {
		dashboardBody.Total.NodeTotal = len(nodes.Items)
	}

	if podErrFlag && nodeErrFlag {
		return
	}

	dashboardBody.Total.MasterTotal, dashboardBody.Total.WorkerTotal = GetNodeRoleTotal(nodes)
	dashboardBody.Total.CPUTotal, dashboardBody.Total.MemoryTotal, dashboardBody.Total.DiskTotal = GetAllNodeProcessor(nodes)
}

func GetNodeRoleTotal(nodes *v1.NodeList) (int, int) {
	masterTotal := 0
	workerTotal := 0
	for _, node := range nodes.Items {
		for key, _ := range node.Labels {
			if key == common.LABEL_MASTER_ROLE {
				masterTotal++
				break
			}
		}
		workerTotal++
	}

	workerTotal -= masterTotal

	return masterTotal, workerTotal
}

func GetAllNodeProcessor(nodes *v1.NodeList) (int64, int64, int64) {

	var cpuTotal int64 = 0
	var memTotal int64 = 0
	var diskTotal int64 = 0

	queryResult := pQuery.QueryResult{}
	queryResult.QueryCall(pQuery.NODE_DISK_TOTAL)

	for _, node := range nodes.Items {
		cpuCapa, _ := node.Status.Capacity.Cpu().AsInt64()
		memCapa, _ := node.Status.Capacity.Memory().AsInt64() // kibibyte

		cpuTotal += cpuCapa
		memTotal += memCapa
	}

	memTotal = util.KitoGI(memTotal)
	if queryResult.Value == nil {
		diskTotal = 0
	} else {
		diskTotal = util.KitoGI(int64(queryResult.Value.(model.Vector)[0].Value))
	}

	return cpuTotal, memTotal, diskTotal
}
