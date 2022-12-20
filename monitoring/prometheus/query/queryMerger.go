package query

const (
	NODE_DISK_TOTAL         = "sum(node_filesystem_size_bytes)"
	NODE_CPU_UTILIZATION    = "100 - (avg by (instance) (irate(node_cpu_seconds_total{mode='idle'}[2m])) * 100)"
	NODE_MEMORY_UTILIZATION = "100 * (1 - ((avg_over_time(node_memory_MemFree_bytes[2m]) + avg_over_time(node_memory_Cached_bytes[2m]) + avg_over_time(node_memory_Buffers_bytes[2m])) / avg_over_time(node_memory_MemTotal_bytes[2m])))"
	NODE_DISK_IO            = "avg by (instance) (irate(node_disk_io_time_seconds_total[2m]) ) * 100"
)
