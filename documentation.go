package geh

import (
	"io/ioutil"
	"path"
	"path/filepath"
	"runtime"
)

type DocEndpointHttpProtocol map[string]DocEndpoint
type DocEndpointProduces string
type DocEndpointParameterIn string
type DocEndpointType string

const (
	DocEndpointProducesJson DocEndpointProduces = "application/json"
	DocEndpointProducesXml  DocEndpointProduces = "application/xml"
	DocEndpointProducesYaml DocEndpointProduces = "application/yaml"
	DocEndpointProducesText DocEndpointProduces = "text/plain"
)

const (
	DocEndpointParameterInQuery  DocEndpointParameterIn = "query"
	DocEndpointParameterInPath   DocEndpointParameterIn = "path"
	DocEndpointParameterInBody   DocEndpointParameterIn = "body"
	DocEndpointParameterInHeader DocEndpointParameterIn = "header"
)

const (
	DocEndpointTypeString  DocEndpointType = "string"
	DocEndpointTypeInteger DocEndpointType = "integer"
)

type DocAppInfo struct {
	Title       string `json:"title"`
	Version     string `json:"version,omitempty"`
	Description string `json:"description,omitempty"`
}

type DocApp struct {
	Swagger     string                             `json:"swagger"`
	Info        DocAppInfo                         `json:"info"`
	Tags        []DocEndpointTag                   `json:"tags,omitempty"`
	Paths       map[string]DocEndpointHttpProtocol `json:"paths"`
	Definitions map[string]DocEndpointDefinition   `json:"definitions"`
}

type DocEndpointTag struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

type DocEndpoint struct {
	Tags        []DocEndpointTag                 `json:"-"`
	Summary     string                           `json:"summary"`
	Description string                           `json:"description"`
	OperationId string                           `json:"operationId"`
	Responses   map[int]DocEndpointResponse      `json:"responses"`
	Parameters  []DocEndpointParameter           `json:"parameters,omitempty"`
	Produces    []DocEndpointProduces            `json:"produces,omitempty"`
	Definitions map[string]DocEndpointDefinition `json:"-"`

	ParsedTags []string `json:"tags,omitempty"`
}

type DocEndpointResponse struct {
	Description string                          `json:"description"`
	Schema      DocEndpointSchema               `json:"schema,omitempty"`
	Headers     map[string]DocEndpointParameter `json:"headers,omitempty"`
}

type DocEndpointParameter struct {
	In          DocEndpointParameterIn   `json:"in,omitempty"`
	Name        string                   `json:"name,omitempty"`
	Description string                   `json:"description,omitempty"`
	Required    bool                     `json:"required,omitempty"`
	Type        DocEndpointType          `json:"type,omitempty"`
	Enum        []string                 `json:"enum,omitempty"`
	Format      string                   `json:"format,omitempty"`
	Items       DocEndpointParameterItem `json:"items,omitempty"`
	Schema      DocEndpointSchema        `json:"schema,omitempty"`
}

type DocEndpointDefinition struct {
	Type       string                          `json:"type"`
	Properties map[string]DocEndpointParameter `json:"properties"`
}

type DocEndpointParameterItem struct {
	Type   DocEndpointType `json:"type,omitempty"`
	Enum   []string        `json:"enum,omitempty"`
	Format string          `json:"format,omitempty"`
	Ref    string          `json:"$ref,omitempty"`
}

type DocEndpointSchema struct {
	Type                 string `json:"type,omitempty"`
	Ref                  string `json:"$ref,omitempty"`
	AdditionalProperties bool   `json:"additionalProperties,omitempty"`
}

func NewDocApp(title string, version string) DocApp {
	return DocApp{
		Swagger: "2.0",
		Info: DocAppInfo{
			Title:   title,
			Version: version,
		},
		Paths:       map[string]DocEndpointHttpProtocol{},
		Definitions: map[string]DocEndpointDefinition{},
	}
}

func addEndpointToDocs(
	method string, endpoint string, docs DocEndpoint,
) {
	if _, ok := endpointDocs.Paths[endpoint]; !ok {
		endpointDocs.Paths[endpoint] = DocEndpointHttpProtocol{}
	}

	addEndpointTags(&docs, endpoint, method)
	addEndpointDefinitions(&docs)

	endpointDocs.Paths[endpoint][method] = docs
}

func addEndpointTags(
	docs *DocEndpoint, endpoint string, method string,
) {
	for _, tag := range docs.Tags {
		endpointDocs.Tags = append(
			endpointDocs.Tags, tag,
		)
		docs.ParsedTags = append(docs.ParsedTags, tag.Name)
	}
}

func addEndpointDefinitions(docs *DocEndpoint) {
	for schemaName, schemaDeclaration := range docs.Definitions {
		if _, ok := endpointDocs.Definitions[schemaName]; !ok {
			endpointDocs.Definitions[schemaName] = schemaDeclaration
		}
	}
}

func docsHandler(
	request Request, response Response, _ ...interface{},
) *Response {
	_, filename, _, _ := runtime.Caller(0)
	currentDirectory := filepath.Dir(filename)

	fileContent, _ := ioutil.ReadFile(
		path.Join(
			currentDirectory,
			"templates/swagger-ui/index.html",
		),
	)
	return response.Html(string(fileContent))
}

func docsSwaggerHandler(
	request Request, response Response, _ ...interface{},
) *Response {
	return response.Json(endpointDocs)
}
