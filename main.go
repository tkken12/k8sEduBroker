package main

import (
	"context"
	"k8sEduBroker/util"
	"log"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
	// homeDir, err := os.UserHomeDir()
	// if err != nil {
	// 	log.Println(err)
	// }

	// kubeConfigPath := filepath.Join(homeDir, ".kube", "config")

	// kubeConfig, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// clientSet, err := kubernetes.NewForConfig(kubeConfig)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// pods, _ := clientSet.CoreV1().Pods("kube-system").List(context.Background(), v1.ListOptions{})
	// log.Println(pods)
	nodes, _ := util.GetK8sClient().CoreV1().Nodes().List(context.Background(), v1.ListOptions{})
	log.Println(nodes)

}
