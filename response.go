package geh

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	headers        map[string]string
	responseWriter http.ResponseWriter
	status         int
	data           string
}

func NewResponse(responseWriter http.ResponseWriter) Response {
	response := Response{}
	response.SetResponseWriter(responseWriter)
	return response
}

func (response *Response) SetResponseWriter(responseWriter http.ResponseWriter) *Response {
	response.responseWriter = responseWriter
	return response
}

func (response *Response) SetHeader(key string, value string) *Response {
	if response.headers == nil {
		response.headers = map[string]string{}
	}

	response.headers[key] = value
	return response
}

func (response *Response) GetHeaders() map[string]string {
	return response.headers
}

func (response *Response) GetHeader(key string) string {
	return response.GetHeaderOrDefaultValue(key, "")
}

func (response *Response) GetHeaderOrDefaultValue(key string, defaultValue string) string {
	if value, ok := response.headers[key]; ok {
		return value
	}

	return defaultValue
}

func (response *Response) SetStatus(status int) *Response {
	response.status = status
	return response
}

func (response *Response) GetStatus() int {
	if response.status == 0 {
		response.status = http.StatusOK
	}

	return response.status
}

func (response *Response) GetData() string {
	return response.data
}

func (response *Response) Json(data interface{}) *Response {
	responseData, _ := json.Marshal(data)
	response.data = string(responseData)

	return response.SetHeader("Content-Type", "application/json")
}

func (response *Response) Html(data string) *Response {
	response.data = data
	return response.SetHeader("Content-Type", "text/html")
}

func (response *Response) Text(data string) *Response {
	response.data = data
	return response.SetHeader("Content-Type", "plain/text")
}
