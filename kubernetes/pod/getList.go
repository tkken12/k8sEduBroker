package pod

import (
	"context"
	"k8sEduBroker/logger"

	kClient "k8sEduBroker/kubernetes/client"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PodGetResBody struct {
	PodName      string           `json:"podName"`
	PodNamespace string           `json:"podNamespace"`
	PodAddress   string           `json:"podAddress"`
	VolumeMount  []v1.VolumeMount `json:"volumeMount"`
	Volumes      []v1.Volume      `json:"volumes"`
	NodeName     string           `json:"nodeName"`
}

type PodsBody []PodGetResBody

func GetPods() (*v1.PodList, error) {
	pods, err := kClient.GetK8sClient().CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return pods, nil
}

func (podReqBody *PodGetResBody) extractPodList(pods *v1.PodList, filteredItems PodsBody, namespaces []string) {
	for idx, _ := range namespaces {
		for _, pod := range pods.Items {
			if pod.Namespace == namespaces[idx] {
				podReqBody.PodName = pod.Name
			}
		}
	}
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
