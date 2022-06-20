package geh

import (
	"net/http"
	"path/filepath"
	"runtime"
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

	_, filename, _, _ := runtime.Caller(0)
	currentDirectory := filepath.Dir(filename)

	muxRouter.PathPrefix("/assets/").Handler(
		http.FileServer(http.Dir(currentDirectory)),
	)

	return &GEHRouter{
		muxRouter: muxRouter,
	}
}
