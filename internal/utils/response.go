package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	types "go-crud/internal/types"
	"io"
	"net/http"
)

type Response struct {
	W http.ResponseWriter
}

func (r Response) ErrorMap(err error) Response {

	var syntaxError *json.SyntaxError
	var typeError *json.UnmarshalTypeError
	var maxBytesError *http.MaxBytesError

	var status int
	var ProblemJson types.ProblemJson
	switch {
	case errors.As(err, &syntaxError):
		status = http.StatusBadRequest
		ProblemJson = types.ProblemJson{
			Type:   "example.com/json-syntax-error",
			Status: status,
			Title:  "JSON syntax error",
			Detail: fmt.Sprintf("Syntax error at offset %d", syntaxError.Offset),
		}
	case errors.As(err, &typeError):
		status = http.StatusBadRequest
		ProblemJson = types.ProblemJson{
			Type:   "example.com/json-type-mismatch",
			Status: status,
			Title:  "JSON type mismatch",
			Detail: fmt.Sprintf("Incorrect type for field '%s', expected %s", typeError.Field, typeError.Type),
		}
	case errors.Is(err, io.EOF): // Empty body
		status = http.StatusBadRequest
		ProblemJson = types.ProblemJson{
			Type:   "example.com/json-empty-body",
			Status: http.StatusBadRequest,
			Title:  "Empty Request Body",
			Detail: "The request body must not be empty",
		}
	case errors.As(err, &maxBytesError):
		status = http.StatusRequestEntityTooLarge
		ProblemJson = types.ProblemJson{
			Type:   "example.com/json-body-too-large",
			Status: status,
			Title:  "Request body too large",
			Detail: fmt.Sprintf("Request body exceeds %d bytes", maxBytesError.Limit),
		}
	default:
		status = http.StatusInternalServerError
		ProblemJson = types.ProblemJson{
			Type:   "example.com/unknown-error",
			Status: status,
			Title:  "Unknown error",
			Detail: "Something went wrong while processing the request",
		}
	}

	return r.Status(status).ProblemJson(ProblemJson)
}

func (r Response) Status(status int) Response {
	r.W.WriteHeader(status)
	return r
}

func (r Response) ProblemJson(v types.ProblemJson) Response {
	r.W.Header().Set("Content-Type", "application/problem+json")
	json.NewEncoder(r.W).Encode(v)
	return r
}
