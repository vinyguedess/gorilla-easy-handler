package geh

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

var endpointDocs = map[string]interface{}{
	"swagger": "2.0",
	"info": map[string]interface{}{
		"title": "Gorilla Easy Handler",
	},
	"tags":        []DocEndpointTag{},
	"paths":       map[string]interface{}{},
	"definitions": map[string]interface{}{},
}

type GEHRouter struct {
	muxRouter *mux.Router
}

func (router *GEHRouter) GetMuxRouter() *mux.Router {
	return router.muxRouter
}

func (router *GEHRouter) Handle(
	method string,
	endpoint string,
	handler func(http.ResponseWriter, *http.Request),
	docs DocEndpoint,
) *mux.Route {
	addEndpointToDocs(strings.ToLower(method), endpoint, docs)

	return router.muxRouter.HandleFunc(endpoint, handler)
}

func NewRouter(
	appName string, version string,
) *GEHRouter {
	endpointDocs["info"].(map[string]interface{})["title"] = appName

	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc(
		"/docs",
		Handler(docsHandler),
	).Methods(http.MethodGet)
	muxRouter.HandleFunc(
		"/docs/swagger.json",
		Handler(docsSwaggerHandler),
	).Methods(http.MethodGet)

	muxRouter.PathPrefix("/assets/").Handler(
		http.FileServer(http.Dir("./")),
	)

	return &GEHRouter{
		muxRouter: muxRouter,
	}
}
