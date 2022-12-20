package dashboard

import (
	kNode "k8sEduBroker/kubernetes/node"

	v1 "k8s.io/api/core/v1"
)

func (dashboardBody *DashboardBody) TableInfoMerger(nodes *v1.NodeList, nodeErr error) {

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
