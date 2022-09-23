package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

var Router *mux.Router

func GetRouter() *mux.Router          { return Router }
func SetRouter(newRouter *mux.Router) { Router = newRouter }

type RequestHandler struct {
	Path       string
	HandleFunc func(http.ResponseWriter, *http.Request)
	RestMethod string
}

func init() {
	router := mux.NewRouter().StrictSlash(true)

	for _, handler := range mergeHandlers {
		for _, elem := range handler {
			router.HandleFunc(elem.Path, elem.HandleFunc).Methods(elem.RestMethod)
		}
	}

	SetRouter(router)
}

var mergeHandlers = [][]RequestHandler{
	PodHandler,
}
