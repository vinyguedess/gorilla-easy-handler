package geh

import (
	"net/http"
)

type EasyHandler func(
	request Request,
	response Response,
	arguments ...interface{},
) *Response

type MuxHandler func(w http.ResponseWriter, r *http.Request)

func Handler(handler EasyHandler, arguments ...interface{}) MuxHandler {
	return func(responseWriter http.ResponseWriter, httpRequest *http.Request) {
		request := NewRequest(httpRequest)
		response := NewResponse(responseWriter)

		receivedResponse := handler(request, response, arguments)
		for key, value := range receivedResponse.GetHeaders() {
			responseWriter.Header().Set(key, value)
		}
		responseWriter.WriteHeader(receivedResponse.GetStatus())
		responseWriter.Write([]byte(receivedResponse.GetData()))
	}
}
