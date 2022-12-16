package dashboard

import (
	kNode "k8sEduBroker/kubernetes/node"
	"k8sEduBroker/logger"
)

func (dashboardBody *DashboardBody) TableInfoMerger() {

	nodes, err := kNode.GetNodeList()
	if err != nil {
		logger.Warn(err.Error())
		return
	}

	for _, node := range nodes.Items {
		dashboardBody.TableInfo = append(dashboardBody.TableInfo, DashboardTableBody{
			NodeName:   node.Name,
			Role:       kNode.FindNodeRole(node.Labels),
			Address:    node.Status.Addresses[0].Address,
			OS:         node.Status.NodeInfo.OSImage,
			Kernel:     node.Status.NodeInfo.KernelVersion,
			K8sVersion: node.Status.NodeInfo.KubeletVersion,
		})
	}

	return
}
