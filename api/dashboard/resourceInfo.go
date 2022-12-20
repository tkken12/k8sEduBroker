package dashboard

import (
	"fmt"
	"k8sEduBroker/kubernetes/node"
	"k8sEduBroker/logger"
	pQuery "k8sEduBroker/monitoring/prometheus/query"
	"k8sEduBroker/util"
	"strconv"
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

	query := NodeQuery{}
	query.CPUQuery.QueryCall(pQuery.NODE_CPU_UTILIZATION)
	query.MemoryQuery.QueryCall(pQuery.NODE_MEMORY_UTILIZATION)
	query.DiskQuery.QueryCall(pQuery.NODE_DISK_IO)

	for _, node := range nodes.Items {
		nodeInfo.NodeResource.utilizationMerger(node, query)
		nodeInfo.NodeCondition.conditionMerger(node)
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

func (condition *NodeCondition) conditionMerger(nodeItem v1.Node) {
	condition.MemoryPressure = append(condition.MemoryPressure, getPressure(nodeItem, node.NODE_PRESSURE_MEMORY))
	condition.PIDPressure = append(condition.PIDPressure, getPressure(nodeItem, node.NODE_PRESSURE_PID))
	condition.DiskPressure = append(condition.DiskPressure, getPressure(nodeItem, node.NODE_PRESSURE_DISK))
	condition.NetworkUnavailable = append(condition.NetworkUnavailable, getPressure(nodeItem, node.NODE_PRESSURE_NETWORK))
	condition.NodeReady = append(condition.NodeReady, getPressure(nodeItem, node.NODE_PRESSURE_READY))
}

func getPressure(node v1.Node, conditionType string) NodePressure {

	var nodeName string
	var pressure bool
	for _, condition := range node.Status.Conditions {
		if conditionType == string(condition.Type) {
			nodeName = string(condition.Type)
			pressure = func() bool {
				bool, err := strconv.ParseBool(string(condition.Status))
				if err != nil {
					logger.Warn("failed to parse bool of node pressure")
					return false
				}
				return bool
			}()
			break
		}
	}

	return NodePressure{
		NodeName: nodeName,
		Pressure: pressure,
	}
}
