package node_test

import (
	kClient "k8sEduBroker/kubernetes/client"
	"k8sEduBroker/kubernetes/node"
	"k8sEduBroker/util"
	"log"
	"testing"
)

func TestList(t *testing.T) {
	util.SetBrokerConf(util.ReadBrokerConfig())
	kClient.NewClient()

	nodes, err := node.GetNodeList()
	if err != nil {
		log.Fatal(err)
	}

	for _, node := range nodes.Items {
		for k, v := range node.Labels {
			log.Println("k==>", k, "v ==>", v)
		}
	}
}
