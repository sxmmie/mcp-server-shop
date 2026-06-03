package jsonrpc

import (
	"encoding/json"
	"fmt"
)

type Request struct {
	JSONRPC string          `json:"jsonrpc"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params,omitempty"`
	ID      any             `json:"id,omitempty"` // can a string
}

type Response struct {
	JSONRPC string `json:"jsonrpc"`
	Result  any    `json:"result,omitempty"`
	Error   *Error `json:"error,omitempty"`
	ID      any    `json:"id,omitempty"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

// this is to conform to the Go Error interface
func (e *Error) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s, Data: %v", e.Code, e.Message, e.Data)
}

// When an ID is nil, request is considered a notification (thats when we consider a request a notification)
func (r *Request) IsNotification() bool {
	return r.ID == nil
}

// Validate the request
func (r *Request) Validate() error {
	if r.JSONRPC != "2.0" {
		return &Error{
			Code:    ErrorInvalidRequest,
			Message: "Invalid JSON-RPC version, must be 2.0",
		}
	}

	if r.Method == "" {
		return &Error{
			Code:    ErrorInvalidRequest,
			Message: "Method is required",
		}
	}

	return nil
}

// Creates a new JSON-RPC response
func NewResponse(result any, id any, err *Error) *Response {
	return &Response{
		JSONRPC: "2.0",
		Result:  result,
		Error:   err,
		ID:      id,
	}
}

// NewErrorResponse creates a new JSON-RPC error response
func NewErrorResponse(id any, err *Error) *Response {
	return NewResponse(nil, id, err)
}

// NewSuccessResponse create a new JSON-RPC success response
func NewSuccessResponse(result, id any) *Response {
	return NewResponse(result, id, nil)
}
