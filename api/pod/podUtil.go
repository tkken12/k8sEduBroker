package pod

import (
	"context"
	kClient "k8sEduBroker/kubernetes/client"
	"k8sEduBroker/logger"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetAllPods() *v1.PodList {
	pods, err := kClient.GetK8sClient().CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Warn(err.Error())
	}

	return pods
}

func GetPodAtNodes(nodeNames []string) map[string]*v1.PodList {

	podListByNodes := make(map[string]*v1.PodList)

	for _, nodeName := range nodeNames {
		pods, err := kClient.GetK8sClient().CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{
			FieldSelector: "spec.nodeName=" + nodeName,
		})
		if err != nil {
			logger.Warn(err.Error())
			continue
		}

		podListByNodes[nodeName] = pods
	}

	return podListByNodes
}
