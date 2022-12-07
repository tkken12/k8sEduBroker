package node

import (
	"context"
	kClient "k8sEduBroker/kubernetes/client"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetNodeList() (*v1.NodeList, error) {
	nodes, err := kClient.GetK8sClient().CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return nodes, nil
}
