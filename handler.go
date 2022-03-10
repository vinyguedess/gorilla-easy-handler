package geh

import (
	"net/http"
)

type EasyHandler func(request Request, response Response) *Response

type MuxHandler func(w http.ResponseWriter, r *http.Request)

func Handler(handler EasyHandler) MuxHandler {
	return func(responseWriter http.ResponseWriter, httpRequest *http.Request) {
		request := NewRequest(httpRequest)
		response := NewResponse(responseWriter)

		handler(request, response)
	}
}
