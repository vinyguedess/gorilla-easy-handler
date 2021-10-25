package geh

import (
	"github.com/gorilla/mux"
	"net/http"
	"net/url"
)

type Request struct {
	httpRequest *http.Request
	params map[string]string
	queries url.Values
}

func NewRequest(httpRequest *http.Request) Request {
	request := Request{}
	request.SetHttpRequest(httpRequest)

	return request
}

func (request *Request) SetHttpRequest(httpRequest *http.Request) {
	request.httpRequest = httpRequest
	request.params = mux.Vars(httpRequest)
	request.queries = httpRequest.URL.Query()
}

func (request *Request) GetParams(key string) string {
	if value, ok := request.params[key]; ok {
		return value
	}

	return ""
}

func (request *Request) GetQueryString(key string) string {
	return request.queries.Get(key)
}

func (request *Request) GetHeader(key string) string {
	return request.httpRequest.Header.Get(key)
}
