package geh

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	responseWriter http.ResponseWriter
	headers        map[string]string
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
	if len(response.headers) == 0 {
		response.headers = map[string]string{}
	}

	response.headers[key] = value
	return response
}

func (response *Response) GetHeader(key string) string {
	if value, ok := response.headers[key]; ok {
		return value
	}

	return ""
}

func (response *Response) GetHeaders() map[string]string {
	return response.headers
}

func (response *Response) SetStatus(status int) *Response {
	response.responseWriter.WriteHeader(status)
	return response
}

func (response *Response) Json(data interface{}) *Response {
	responseData, _ := json.Marshal(data)
	response.responseWriter.Write(responseData)

	return response.SetHeader("Content-Type", "application/json")
}

func (response *Response) Html(data string) *Response {
	response.responseWriter.Write([]byte(data))
	return response.SetHeader("Content-Type", "text/html")
}

func (response *Response) Text(data string) *Response {
	response.responseWriter.Write([]byte(data))
	return response.SetHeader("Content-Type", "plain/text")
}
