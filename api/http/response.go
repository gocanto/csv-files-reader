package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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

func (receiver Response) MethodNotAllowed(verb string) error {
	receiver.writer.WriteHeader(http.StatusMethodNotAllowed)

	return json.NewEncoder(receiver.writer).Encode(
		map[string]string{
			"status":      strconv.Itoa(http.StatusMethodNotAllowed),
			"message":     "The given method is not allowed",
			"description": fmt.Sprintf("The given methos has to be [%s]", verb),
		},
	)
}

func (receiver Response) BadRequest(message string) error {
	receiver.writer.WriteHeader(http.StatusBadRequest)

	return json.NewEncoder(receiver.writer).Encode(
		map[string]string{
			"status":  strconv.Itoa(http.StatusBadRequest),
			"message": message,
		},
	)
}

func (receiver Response) ServerError(message string) error {
	receiver.writer.WriteHeader(http.StatusInternalServerError)

	return json.NewEncoder(receiver.writer).Encode(
		map[string]string{
			"status":  strconv.Itoa(http.StatusInternalServerError),
			"message": message,
		},
	)
}
