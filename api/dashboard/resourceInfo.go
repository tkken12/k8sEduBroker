package dashboard

import (
	"fmt"
	pQuery "k8sEduBroker/monitoring/prometheus/query"
	"k8sEduBroker/util"
	"strings"

	"github.com/prometheus/common/model"
	v1 "k8s.io/api/core/v1"
)

type NodeQuery struct {
	CPUQuery    pQuery.QueryResult
	MemoryQuery pQuery.QueryResult
	DiskQuery   pQuery.QueryResult
}

func (nodeInfo *DashboardNodeInfo) NodeInfoMerger(nodes *v1.NodeList, nodeErr error) {

	if nodeErr != nil {
		return
	}

	condition := NodeCondition{}

	query := NodeQuery{}
	query.CPUQuery.QueryCall(pQuery.NODE_CPU_UTILIZATION)
	query.MemoryQuery.QueryCall(pQuery.NODE_MEMORY_UTILIZATION)
	query.DiskQuery.QueryCall(pQuery.NODE_DISK_IO)

	for _, node := range nodes.Items {
		nodeInfo.NodeResource.utilizationMerger(node, query)
		condition.conditionMerger(node)
	}

}

func (nodeInfo *NodeResource) utilizationMerger(node v1.Node, queryResult NodeQuery) {

	nodeInfo.CPUUtilization = append(nodeInfo.CPUUtilization, getUtilization(node, queryResult.CPUQuery))
	nodeInfo.MemoryUtilization = append(nodeInfo.MemoryUtilization, getUtilization(node, queryResult.MemoryQuery))
	nodeInfo.DiskIO = append(nodeInfo.DiskIO, getUtilization(node, queryResult.DiskQuery))

}

func getUtilization(node v1.Node, queryResult pQuery.QueryResult) ResourceUtilization {

	resourceBody := ResourceUtilization{}

	for i := 0; i < queryResult.Value.(model.Vector).Len(); i++ {

		// 가독성이 안좋아서 별도의 변수로 선언함
		address := queryResult.Value.(model.Vector)[i].Metric["instance"]
		resourceBody.NodeName = util.FindNodeNameByAddress(
			fmt.Sprintf("%s", strings.Split(string(address), ":")[0]), node)
		resourceBody.Utilization = util.RoundFloat(float64(queryResult.Value.(model.Vector)[i].Value), 2)
	}

	return resourceBody
}

func getMemoryUtilization(node v1.Node, memoryQueryResult pQuery.QueryResult) ResourceUtilization {
	resourceBody := ResourceUtilization{}

	return resourceBody
}

func getDiskIO(node v1.Node, diskQueryResult pQuery.QueryResult) ResourceUtilization {
	resourceBody := ResourceUtilization{}

	return resourceBody
}

func (condition *NodeCondition) conditionMerger(node v1.Node) {
	getPressure()
}

func getPressure() {

}
