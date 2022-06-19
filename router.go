package geh

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

var endpointDocs DocApp

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
	title string, version string,
) *GEHRouter {
	endpointDocs = NewDocApp(title, version)

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
