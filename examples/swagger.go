package main

import (
	"net/http"

	geh "github.com/vinyguedess/gorilla-easy-handler"
)

func main() {
	router := geh.NewRouter("GEH Swagger Example", "v1.0")

	router.Handle(
		http.MethodGet,
		"/tests",
		geh.Handler(
			func(
				request geh.Request,
				response geh.Response,
				_ ...interface{},
			) *geh.Response {
				return response.Text("OK")
			},
		),
		geh.DocEndpoint{
			OperationId: "Test endpoint",
			Summary:     "Test endpoint",
			Responses: map[int]geh.DocEndpointResponse{
				http.StatusOK: {
					Description: "OK",
				},
			},
		},
	)

	router.Handle(
		http.MethodPost,
		"/tests",
		geh.Handler(
			func(
				request geh.Request,
				response geh.Response,
				_ ...interface{},
			) *geh.Response {
				return response.Text("OK")
			},
		),
		geh.DocEndpoint{
			Tags: []geh.DocEndpointTag{
				{
					Name: "Hello World",
				},
			},
			OperationId: "Test endpoint",
			Summary:     "Test endpoint",
			Parameters: []geh.DocEndpointParameter{
				{
					Name: "body",
					In:   geh.DocEndpointParameterInBody,
					Schema: &geh.DocEndpointSchema{
						Ref: "#/definitions/TestObject",
					},
				},
			},
			Responses: map[int]geh.DocEndpointResponse{
				http.StatusCreated: {
					Description: "OK",
				},
			},
			Definitions: map[string]geh.DocEndpointDefinition{
				"TestObject": {
					Type: "object",
					Properties: map[string]geh.DocEndpointParameter{
						"name": {
							Type:     geh.DocEndpointTypeString,
							Required: true,
						},
						"age": {
							Type: geh.DocEndpointTypeInteger,
						},
						"birthday": {
							Type:   geh.DocEndpointTypeString,
							Format: "date",
						},
					},
				},
			},
		},
	)

	http.ListenAndServe(":8080", router.GetMuxRouter())
}
