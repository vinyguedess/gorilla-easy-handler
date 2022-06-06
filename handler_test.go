package geh

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	handler := Handler(
		func(
			request Request, response Response, arguments ...interface{},
		) *Response {
			return response.Json(
				map[string]string{
					"hello": "world",
				},
			)
		},
	)

	httpRequest, _ := http.NewRequest("GET", "/resource", nil)
	responseWriter := httptest.NewRecorder()
	handler(responseWriter, httpRequest)

	assert.Equal(t, "{\"hello\":\"world\"}", responseWriter.Body.String())
	assert.Equal(t, "application/json", responseWriter.Header().Get("Content-Type"))
	assert.Equal(t, 200, responseWriter.Code)
}

func TestHandlerPassingArguments(t *testing.T) {
	handler := Handler(
		func(
			request Request, response Response, arguments ...interface{},
		) *Response {
			returnsSameString := arguments[0].(func(string) string)

			return response.Json(
				map[string]string{
					"hello": returnsSameString("world"),
				},
			)
		},
		func(something string) string { return something },
	)

	httpRequest, _ := http.NewRequest("GET", "/resource", nil)
	responseWriter := httptest.NewRecorder()
	handler(responseWriter, httpRequest)

	assert.Equal(t, "{\"hello\":\"world\"}", responseWriter.Body.String())
	assert.Equal(t, "application/json", responseWriter.Header().Get("Content-Type"))
	assert.Equal(t, 200, responseWriter.Code)
}
