package dashboard

import v1 "k8s.io/api/core/v1"

type DashboardBody struct {
	Total       DashboardTotal         `json:"total"`
	TableInfo   []DashboardTableBody   `json:"tableInfo"`
	Utilization []DashboardUtilization `json:"utilizaiton"`
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
	NodeName   string           `json:"nodeName"`
	Role       string           `json:"role"`
	Address    []v1.NodeAddress `json:"address"`
	OS         string           `json:"os"`
	Kernel     string           `json:"kernel"`
	K8sVersion string           `json:"k8sVersion"`
}

type DashboardUtilization struct {
}

func (body *DashboardBody) InfoMerger() {
	body.TotalMerger()
	body.TableInfoMerger()
}
