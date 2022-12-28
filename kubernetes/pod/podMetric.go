package pod

import (
	"context"
	kClient "k8sEduBroker/kubernetes/client"
	"k8sEduBroker/kubernetes/node"
	"k8sEduBroker/logger"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/metrics/pkg/apis/metrics/v1beta1"
)

func GetPodMetricses() *v1beta1.PodMetricsList {
	metric, err := kClient.GetMetricClient().MetricsV1beta1().
		PodMetricses(metav1.NamespaceAll).List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		logger.Warn(err.Error())
		return metric
	}

	return metric
}

func GetPodMetric(podItem v1.Pod, nodeCapa map[string]node.NodeResourceCapacity, podMetricses *v1beta1.PodMetricsList) PodMetric {

	podMetric := PodMetric{}
	for _, item := range podMetricses.Items {
		for _, container := range item.Containers {
			// log.Println("ctrName", container.Name, "podName", podItem.Spec.Containers)
			if container.Name == podItem.Name {
				podMetric.Name = container.Name
				// nano to milli
				podMetric.CPU.Value = container.Usage.Cpu().AsApproximateFloat64() / float64(nodeCapa[podItem.Spec.NodeName].CpuCore) * 100
				// default unit kibibyte
				podMetric.Memory.Value = float64(container.Usage.Memory().MilliValue()) / float64(nodeCapa[podItem.Spec.NodeName].MemoryTotal) * 100

			}
		}
	}

	return podMetric
}
