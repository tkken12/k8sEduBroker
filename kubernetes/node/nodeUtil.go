package node

import "k8sEduBroker/common"

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
