package pod

import (
	"context"
	"k8sEduBroker/logger"
	"k8sEduBroker/util"

	pHandler "k8sEduBroker/api/pod"

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

func GetPodList(podReqBody pHandler.PodRequestParams) (PodGetResBody, error) {

	podsResElem := PodsBody{}

	pods, err := util.GetK8sClient().CoreV1().Pods("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		logger.Warn(err.Error())
		return PodGetResBody{}, err
	}

	return podReqBody.extractPodList(pods, podsResElem, podReqBody.PodNamespaces), nil
}

func (podReqBody *PodRequestBody) extractPodList(pods *v1.PodList, filteredItems PodsBody, namespaces []string) PodGetResBody {

	for idx, _ := range namespaces {
		for _, pod := range pods.Items {
			if pod.Namespace == namespaces[idx] {
				podReqBody.PodName = pod.Name
				podReqBody.Na
			}
		}
	}

	return filteredItems
}
