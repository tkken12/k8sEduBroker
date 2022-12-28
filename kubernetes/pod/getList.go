package pod

import (
	"context"
	"k8sEduBroker/logger"

	kClient "k8sEduBroker/kubernetes/client"
	"k8sEduBroker/kubernetes/node"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/metrics/pkg/apis/metrics/v1beta1"
)

type PodGetBody struct {
	PodInfo []PodInfo `json:"podInfo"`
}

type PodInfo struct {
	PodName    string         `json:"podName"`
	Namespace  string         `json:"namespace"`
	Address    string         `json:"address"`
	Volume     []v1.Volume    `json:"volume"`
	Containers []v1.Container `json:"containers"`
	NodeName   string         `json:"nodeName"`
	HostName   string         `json:"hostName"`
	Phase      v1.PodPhase    `json:"phase"`
	// PodCondition PodCondition   `json:"condition"`
	CreationTime *metav1.Time `json:"creationTime"`
	StartTime    *metav1.Time `json:"startTime"`
	Metric       PodMetric    `json:"metric"`
}

type PodCondition struct {
	Type   v1.PodConditionType `json:"type"`
	Status v1.ConditionStatus  `json:"status"`
}

type PodMetric struct {
	Name   string         `json:"podName"`
	CPU    PodMetricValue `json:"cpu"`
	Memory PodMetricValue `json:"memory"`
}

type PodMetricValue struct {
	Value float64 `json:"value"`
}

func GetPods() (*v1.PodList, error) {
	pods, err := kClient.GetK8sClient().CoreV1().Pods(metav1.NamespaceAll).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return pods, nil
}

// func (podInfo *PodInfo) extractPodList(pods *v1.PodList, filteredItems PodsBody, namespaces []string) {
// 	for idx, _ := range namespaces {
// 		for _, pod := range pods.Items {
// 			if pod.Namespace == namespaces[idx] {
// 				podInfo.PodName = pod.Name
// 			}
// 		}
// 	}
// }

func GetPodAtNodes(nodeNames []string) map[string]*v1.PodList {

	podListByNodes := make(map[string]*v1.PodList)

	for _, nodeName := range nodeNames {
		pods, err := kClient.GetK8sClient().CoreV1().Pods(metav1.NamespaceAll).List(context.TODO(), metav1.ListOptions{
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

func (podGetBody *PodGetBody) PodParser(podItem v1.Pod,
	nodeCapa map[string]node.NodeResourceCapacity,
	podMetricses *v1beta1.PodMetricsList) {

	podGetBody.PodInfo = append(podGetBody.PodInfo, PodInfo{
		PodName:    podItem.Name,
		Namespace:  podItem.Namespace,
		Address:    podItem.Status.PodIP,
		Volume:     podItem.Spec.Volumes,
		Containers: podItem.Spec.Containers,
		NodeName:   podItem.Spec.NodeName,
		HostName:   podItem.Spec.Hostname,
		Phase:      podItem.Status.Phase,
		// PodCondition: PodCondition{
		// 	Type: podItem.Status.Conditions[],
		// },
		CreationTime: &podItem.CreationTimestamp,
		StartTime:    podItem.Status.StartTime,
		Metric:       GetPodMetric(podItem, nodeCapa, podMetricses),
	})
}
