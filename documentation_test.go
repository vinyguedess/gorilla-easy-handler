package geh

import (
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AddEndpointToDocs_TitleAndVersion(t *testing.T) {
	NewRouter("title", "version")

	assert.Equal(
		t,
		endpointDocs.Info.Title,
		"title",
	)
	assert.Equal(
		t,
		endpointDocs.Info.Version,
		"version",
	)
}

func Test_AddEndpointToDocs_BasicEndpoint(t *testing.T) {
	router := NewRouter("title", "version")

	router.Handle(
		http.MethodGet,
		"/tests",
		Handler(
			func(
				request Request, response Response, _ ...interface{},
			) *Response {
				return response.Text("OK")
			},
		),
		DocEndpoint{
			OperationId: "Test endpoint",
			Summary:     "Test endpoint",
			Description: "Test endpoint",
		},
	)

	assert.Equal(
		t,
		endpointDocs.Paths,
		map[string]DocEndpointHttpProtocol{
			"/tests": {
				strings.ToLower(http.MethodGet): {
					OperationId: "Test endpoint",
					Summary:     "Test endpoint",
					Description: "Test endpoint",
				},
			},
		},
	)
}

func Test_AddEndpointToDocs_EndpointTags(t *testing.T) {
	router := NewRouter("title", "version")

	router.Handle(
		http.MethodGet,
		"/tests",
		Handler(
			func(
				request Request, response Response, _ ...interface{},
			) *Response {
				return response.Text("OK")
			},
		),
		DocEndpoint{
			OperationId: "Test endpoint",
			Tags: []DocEndpointTag{
				{
					Name: "Hello World",
				},
			},
		},
	)

	assert.Equal(
		t,
		endpointDocs.Tags,
		[]DocEndpointTag{
			{
				Name:        "Hello World",
				Description: "",
			},
		},
	)
	assert.Equal(
		t,
		endpointDocs.Paths,
		map[string]DocEndpointHttpProtocol{
			"/tests": {
				strings.ToLower(http.MethodGet): {
					OperationId: "Test endpoint",
					Tags: []DocEndpointTag{
						{
							Name:        "Hello World",
							Description: "",
						},
					},
					ParsedTags: []string{
						"Hello World",
					},
				},
			},
		},
	)
}

func Test_AddEndpointToDocs_EndpointDefinitions(t *testing.T) {
	router := NewRouter("title", "version")

	router.Handle(
		http.MethodPost,
		"/tests",
		Handler(
			func(
				request Request, response Response, _ ...interface{},
			) *Response {
				return response.Text("OK")
			},
		),
		DocEndpoint{
			OperationId: "Test endpoint",
			Parameters: []DocEndpointParameter{
				{
					In:   DocEndpointParameterInBody,
					Name: "body",
					Schema: map[string]string{
						"$ref": "#/definitions/TestObject",
					},
				},
			},
			Definitions: map[string]DocEndpointDefinition{
				"TestObject": {
					Type: "object",
					Properties: map[string]DocEndpointParameter{
						"name": {
							Type:     "string",
							Required: true,
						},
					},
				},
			},
		},
	)

	assert.Equal(
		t,
		map[string]DocEndpointHttpProtocol{
			"/tests": {
				strings.ToLower(http.MethodPost): {
					OperationId: "Test endpoint",
					Parameters: []DocEndpointParameter{
						{
							In:   DocEndpointParameterInBody,
							Name: "body",
							Schema: map[string]string{
								"$ref": "#/definitions/TestObject",
							},
						},
					},
					Definitions: map[string]DocEndpointDefinition{
						"TestObject": {
							Type: "object",
							Properties: map[string]DocEndpointParameter{
								"name": {
									Type:     "string",
									Required: true,
								},
							},
						},
					},
				},
			},
		},
		endpointDocs.Paths,
	)
	assert.Equal(
		t,
		map[string]DocEndpointDefinition{
			"TestObject": {
				Type: "object",
				Properties: map[string]DocEndpointParameter{
					"name": {
						Type:     "string",
						Required: true,
					},
				},
			},
		},
		endpointDocs.Definitions,
	)
}
