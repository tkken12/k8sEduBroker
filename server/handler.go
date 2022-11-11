package server

import (
	"k8sEduBroker/api/pod"
)

var PodHandler = []RequestHandler{
	{"/api/v1/pod", pod.PodHandler, POST},
	{"/api/v1/pod", pod.PodHandler, PUT},
	{"/api/v1/pod", pod.PodHandler, DELETE},
	{"/api/v1/pod", pod.PodGetHandler, GET},
	// {"api/v1/node", }
}
