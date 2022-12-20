package dashboard

import "k8sEduBroker/kubernetes/node"

type DashboardBody struct {
	Total     DashboardTotal       `json:"total"`
	TableInfo []DashboardTableBody `json:"tableInfo"`
	NodeInfo  DashboardNodeInfo    `json:"nodeInfo"`
}

type DashboardTotal struct {
	NodeTotal   int   `json:"nodeTotal"`
	MasterTotal int   `json:"masterTotal"`
	WorkerTotal int   `json:"workerTotal"`
	CPUTotal    int64 `json:"cpuTotal"`
	MemoryTotal int64 `json:"memoryTotal"`
	DiskTotal   int64 `json:"diskTotal"`
	PodTotal    int   `json:"podTotal"`
}

type DashboardTableBody struct {
	NodeName   string `json:"nodeName"`
	Role       string `json:"role"`
	Address    string `json:"address"`
	OS         string `json:"os"`
	Kernel     string `json:"kernelVersion"`
	K8sVersion string `json:"k8sVersion"`
}

type DashboardNodeInfo struct {
	NodeResource  NodeResource  `json:"nodeResource"`
	NodeCondition NodeCondition `json:"nodeCondition"`
}

type NodeResource struct {
	CPUUtilization    []ResourceUtilization `json:"nodeCpuUtilization"`
	MemoryUtilization []ResourceUtilization `json:"nodeMemoryUtilization"`
	DiskIO            []ResourceUtilization `json:"nodeDiskiO"`
}

type ResourceUtilization struct {
	NodeName    string  `json:"nodeName"`
	Utilization float64 `json:"utilization"`
}

type NodePressure struct {
	NodeName string `json:"nodeName"`
	Pressure bool   `json:"pressure"`
}
type NodeCondition struct {
	MemoryPressure     []NodePressure `json:"memoryPressure"`
	PIDPressure        []NodePressure `json:"pidPressure"`
	DiskPressure       []NodePressure `json:"diskPressure"`
	NetworkUnavailable []NodePressure `json:"networkUnavailable"`
}

// TODO
// 불필요하게 여러번 반복함
// 기타 모듈들 완성 후 for문 한번에 처리하도록 수정해야함
func (body *DashboardBody) InfoMerger() {
	nodes, nodeErr := node.GetNodeList()
	body.TotalMerger(nodes, nodeErr)
	body.TableInfoMerger(nodes, nodeErr)
	body.NodeInfo.NodeInfoMerger(nodes, nodeErr)
}
