package node

import "k8sEduBroker/common"

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
