package dashboard

type DashboardBody struct {
	NodeTotal   int `json:"nodeTotal"`
	MasterTotal int `json:"masterTotal"`
	WorkerTotal int `json:"workerTotal"`
	CPUTotal    int `json:"cpuTotal"`
	MemoryTotal int `json:"memoryTotal"`
	PodTotal    int `json:"podTotal"`
}

func (body *DashboardBody) InfoMerger() {
	body.TotalMerger()
}
