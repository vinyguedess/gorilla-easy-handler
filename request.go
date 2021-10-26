package geh

import (
	"encoding/json"
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

func (request *Request) GetMethod() string {
	return request.httpRequest.Method
}

func (request *Request) GetURI() string {
	return request.httpRequest.URL.RequestURI()
}

func (request *Request) GetURL() string {
	return request.httpRequest.URL.Path
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

func (request *Request) GetBody(typo interface{}) interface{} {
	decoder := json.NewDecoder(request.httpRequest.Body)
	decoder.Decode(&typo)

	return typo
}
