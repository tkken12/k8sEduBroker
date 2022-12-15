package dashboard

type DashboardBody struct {
	NodeTotal   int   `json:"nodeTotal"`
	MasterTotal int   `json:"masterTotal"`
	WorkerTotal int   `json:"workerTotal"`
	CPUTotal    int64 `json:"cpuTotal"`
	MemoryTotal int64 `json:"memoryTotal"`
	DiskTotal   int64
	PodTotal    int `json:"podTotal"`
}

func (body *DashboardBody) InfoMerger() {
	body.TotalMerger()
}
