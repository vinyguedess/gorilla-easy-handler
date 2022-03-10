package geh

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	handler := Handler(func(request Request, response Response) *Response {
		return response.Json(map[string]string{
			"hello": "world",
		})
	})

	httpRequest, _ := http.NewRequest("GET", "/resource", nil)
	responseWriter := httptest.NewRecorder()
	handler(responseWriter, httpRequest)

	assert.Equal(t, "{\"hello\":\"world\"}", responseWriter.Body.String())
	assert.Equal(t, "application/json", responseWriter.Header().Get("Content-Type"))
	assert.Equal(t, 200, responseWriter.Code)
}
