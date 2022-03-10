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
	response.responseWriter.Header().Set(key, value)
	return response
}

func (response *Response) SetStatus(status int) *Response {
	response.responseWriter.WriteHeader(status)
	return response
}

func (response *Response) Json(data interface{}) *Response {
	response.SetHeader("Content-Type", "application/json")

	responseData, _ := json.Marshal(data)
	response.responseWriter.Write(responseData)

	return response
}

func (response *Response) Html(data string) *Response {
	response.SetHeader("Content-Type", "text/html")
	response.responseWriter.Write([]byte(data))

	return response
}

func (response *Response) Text(data string) *Response {
	response.SetHeader("Content-Type", "plain/text")
	response.responseWriter.Write([]byte(data))

	return response
}
