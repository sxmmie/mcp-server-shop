package jsonrpc

const (
	ErrorParseError     = -32700 // Invalid JSON
	ErrorInvalidRequest = -32600 // Invalid request object
	ErrorMethodNotFound = -32601 // Mthod does not exists
	ErrorInvalidParams  = -32602 // Invalid method parameters
	ErrorInternal       = -32603 // Internal JSON-RPC error
)

// NewError creates a new JSON-RPC error
func NewError(code int, message string, data any) *Error {
	return &Error{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func NewParseError(data any) *Error {
	return NewError(ErrorParseError, "Parse error", data)
}

// creates a method not found error
func NewMethodNotFoundError(method string) *Error {
	return NewError(ErrorMethodNotFound, "method not found", nil)
}

// creates an invalid params error
func NewInValidParameterError(message string) *Error {
	return NewError(ErrorInvalidParams, message, nil)
}

// creates an internal error
func NewInternalError(err error) *Error {
	return NewError(ErrorInvalidParams, "internal error", err.Error())
}
