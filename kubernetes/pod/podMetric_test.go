package pod_test

import (
	"context"
	"flag"
	kClient "k8sEduBroker/kubernetes/client"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8sEduBroker/util"
	"testing"
)

func TestMetric(t *testing.T) {

	configPath := flag.String("config", "/.k8sEdu/config.json", "set broker config path")
	flag.Parse()

	util.SetBrokerConf(util.ReadBrokerConfig(configPath))
	kClient.NewClient()

	metric, err := kClient.GetMetricClient().MetricsV1beta1().PodMetricses("kube-system").List(context.TODO(), metav1.ListOptions{
		FieldSelector: "metadata.name=kube-apiserver-tykim",
	})
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range metric.Items {
		for _, container := range item.Containers {
			log.Println("a", container.Usage.Cpu(), container.Usage.Cpu().AsDec(), container.Usage.Cpu().AsDec().UnscaledBig().Int64(), container.Usage.Cpu().MilliValue(), container.Usage.Memory())
			log.Println("b", container.Usage.Cpu().AsApproximateFloat64()*1000000/(24*1000000)*100)
			log.Println("c", 179216000/(24*100000000))
		}
	}
}
