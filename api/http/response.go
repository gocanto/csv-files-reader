package http

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	writer  http.ResponseWriter
	request *http.Request
}

func MakeResponse(w http.ResponseWriter, r *http.Request) Response {
	w.Header().Set("Content-Type", "application/json")

	return Response{
		writer:  w,
		request: r,
	}
}

func (receiver Response) Ok(payload interface{}) error {
	receiver.writer.WriteHeader(http.StatusOK)

	return json.NewEncoder(receiver.writer).Encode(payload)
}
