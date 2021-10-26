package geh

import (
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestRequest_GetMethod(t *testing.T) {
	httpRequest, _ := http.NewRequest("PATCH", "/v1/resource/1", nil)

	request := Request{}
	request.SetHttpRequest(httpRequest)

	assert.Equal(t, "PATCH", request.GetMethod())
}

func TestRequest_GetURI(t *testing.T) {
	httpRequest, _ := http.NewRequest("PATCH", "http://localhost:8080/v1/resource/1?query=string", nil)

	request := Request{}
	request.SetHttpRequest(httpRequest)

	assert.Equal(t, "/v1/resource/1?query=string", request.GetURI())
}

func TestRequest_GetURL(t *testing.T) {
	httpRequest, _ := http.NewRequest("PATCH", "http://localhost:8080/v1/resource/1?query=string", nil)

	request := Request{}
	request.SetHttpRequest(httpRequest)

	assert.Equal(t, "/v1/resource/1", request.GetURL())
}

func TestRequest_GetParams(t *testing.T) {
	httpRequest, _ := http.NewRequest("GET", "/resource/1", nil)
	httpRequest = mux.SetURLVars(httpRequest, map[string]string{
		"id": "1",
	})

	request := Request{}
	request.SetHttpRequest(httpRequest)

	assert.Equal(t, "1", request.GetParams("id"))
}

func TestRequest_GetParamsIfNotExists(t *testing.T) {
	httpRequest, _ := http.NewRequest("GET", "resource/1", nil)

	request := Request{}
	request.SetHttpRequest(httpRequest)

	assert.Equal(t, "", request.GetParams("id"))
}

func TestRequest_GetQueryString(t *testing.T) {
	httpRequest, _ := http.NewRequest("GET", "/?hello=world", nil)

	request := Request{}
	request.SetHttpRequest(httpRequest)

	assert.Equal(t, "world", request.GetQueryString("hello"))
}

func TestRequest_GetQueryStringIfDoesntExist(t *testing.T) {
	httpRequest, _ := http.NewRequest("GET", "/?hello=world", nil)

	request := Request{}
	request.SetHttpRequest(httpRequest)

	assert.Equal(t, "", request.GetQueryString("hello2"))
}

func TestRequest_GetHeader(t *testing.T) {
	httpRequest, _ := http.NewRequest("GET", "/?hello=world", nil)
	httpRequest.Header.Set("Content-Type", "application/json")

	request := Request{}
	request.SetHttpRequest(httpRequest)

	assert.Equal(t, "application/json", request.GetHeader("Content-Type"))
}
