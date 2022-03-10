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

		receivedResponse := handler(request, response)
		responseWriter.WriteHeader(receivedResponse.GetStatus())
		for key, value := range receivedResponse.GetHeaders() {
			responseWriter.Header().Set(key, value)
		}
		responseWriter.Write([]byte(receivedResponse.GetData()))
	}
}
