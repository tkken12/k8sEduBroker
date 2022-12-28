package node

import (
	"k8sEduBroker/common"
	"k8sEduBroker/logger"
)

type NodeResourceCapacity struct {
	CpuCore     int64
	MemoryTotal int64
}

const (
	NODE_PRESSURE_MEMORY  = "MemoryPressure"
	NODE_PRESSURE_DISK    = "DiskPressure"
	NODE_PRESSURE_PID     = "PIDPressure"
	NODE_PRESSURE_NETWORK = "NetworkUnavailable"
	NODE_PRESSURE_READY   = "Ready"
)

func FindNodeRole(labels map[string]string) string {

	var role string = common.ROLE_WORKER

	for k, _ := range labels {
		if k == common.LABEL_MASTER_ROLE {
			role = common.ROLE_MASTER
			break
		}
	}

	return role
}

func GetNodeCapacity() map[string]NodeResourceCapacity {

	newMap := make(map[string]NodeResourceCapacity)

	nodes, err := GetNodeList()
	if err != nil {
		logger.Warn(err.Error())
		return newMap
	}

	for _, node := range nodes.Items {

		newMap[node.Name] = NodeResourceCapacity{
			CpuCore: func() int64 {
				cpu, isOk := node.Status.Capacity.Cpu().AsInt64()
				if isOk == false {
					return 0
				}
				return cpu
			}(),
			MemoryTotal: node.Status.Capacity.Memory().Value(),
		}
	}

	return newMap
}
