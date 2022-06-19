package geh

import "io/ioutil"

type DocEndpointParameterIn string
type DocEndpointType string

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
	Definitions map[string]DocEndpointDefinition `json:"-"`

	ParsedTags []string `json:"tags,omitempty"`
}

type DocEndpointResponse struct {
	Description string            `json:"description"`
	Schema      map[string]string `json:"schema,omitempty"`
}

type DocEndpointParameter struct {
	In       DocEndpointParameterIn `json:"in,omitempty"`
	Name     string                 `json:"name,omitempty"`
	Required bool                   `json:"required,omitempty"`
	Type     DocEndpointType        `json:"type,omitempty"`
	Enum     []string               `json:"enum,omitempty"`
	Format   string                 `json:"format,omitempty"`
	Schema   map[string]string      `json:"schema,omitempty"`
}

type DocEndpointDefinition struct {
	Type       string                          `json:"type"`
	Properties map[string]DocEndpointParameter `json:"properties"`
}

func addEndpointToDocs(
	method string, endpoint string, docs DocEndpoint,
) {
	if _, ok := endpointDocs["paths"].(map[string]interface{})[endpoint]; !ok {
		endpointDocs["paths"].(map[string]interface{})[endpoint] = map[string]interface{}{}
	}

	if _, ok := endpointDocs["paths"].(map[string]interface{})[endpoint].(map[string]interface{})[method]; !ok {
		endpointDocs["paths"].(map[string]interface{})[endpoint].(map[string]interface{})[method] = map[string]interface{}{}
	}

	for _, tag := range docs.Tags {
		endpointDocs["tags"] = append(
			endpointDocs["tags"].([]DocEndpointTag), tag,
		)

		docs.ParsedTags = append(docs.ParsedTags, tag.Name)
	}

	endpointDocs["paths"].(map[string]interface{})[endpoint].(map[string]interface{})[method] = docs

	for schemaName, schemaDeclaration := range docs.Definitions {
		if _, ok := endpointDocs["definitions"].(map[string]interface{})[schemaName]; !ok {
			endpointDocs["definitions"].(map[string]interface{})[schemaName] = schemaDeclaration
		}
	}
}

func docsHandler(
	request Request, response Response, _ ...interface{},
) *Response {
	fileContent, _ := ioutil.ReadFile("./templates/swagger-ui/index.html")
	return response.Html(string(fileContent))
}

func docsSwaggerHandler(
	request Request, response Response, _ ...interface{},
) *Response {
	return response.Json(endpointDocs)
}
