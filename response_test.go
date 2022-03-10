package geh

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponse_SetHeader(t *testing.T) {
	responseWriter := httptest.NewRecorder()

	response := Response{}
	response.SetResponseWriter(responseWriter).
		SetHeader("Hello", "world")

	assert.Equal(t, "world", response.GetHeader("Hello"))
}

func TestResponse_GetHeaders(t *testing.T) {
	responseWriter := httptest.NewRecorder()

	response := Response{}
	response.SetResponseWriter(responseWriter).
		SetHeader("Hello", "world")

	assert.Equal(
		t,
		map[string]string{
			"Hello": "world",
		},
		response.GetHeaders(),
	)
}

func TestResponse_SetStatus(t *testing.T) {
	responseWriter := httptest.NewRecorder()

	response := Response{}
	response.SetResponseWriter(responseWriter).
		SetStatus(http.StatusCreated)

	assert.Equal(t, http.StatusCreated, responseWriter.Code)
}

func TestResponse_Json(t *testing.T) {
	responseWriter := httptest.NewRecorder()

	response := Response{}
	response.SetResponseWriter(responseWriter).
		Json(map[string]string{
			"hello": "world",
		})

	assert.Equal(t, "{\"hello\":\"world\"}", responseWriter.Body.String())
	assert.Equal(t, "application/json", response.GetHeader("Content-Type"))
}

func TestResponse_Html(t *testing.T) {
	responseWriter := httptest.NewRecorder()

	response := Response{}
	response.SetResponseWriter(responseWriter).
		Html("<h1>Hello</h1>")

	assert.Equal(t, "<h1>Hello</h1>", responseWriter.Body.String())
	assert.Equal(t, "text/html", response.GetHeader("Content-Type"))
}

func TestResponse_Text(t *testing.T) {
	responseWriter := httptest.NewRecorder()

	response := Response{}
	response.SetResponseWriter(responseWriter).
		Text("oi ne")

	assert.Equal(t, "oi ne", responseWriter.Body.String())
	assert.Equal(t, "plain/text", response.GetHeader("Content-Type"))
}
