package util

import "k8s.io/client-go/kubernetes"

var G_k8sClient *kubernetes.Clientset

func GetK8sClient() *kubernetes.Clientset          { return G_k8sClient }
func SetK8sClient(clientSet *kubernetes.Clientset) { G_k8sClient = clientSet }
