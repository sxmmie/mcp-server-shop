package jsonrpc

const (
	ErrorParseError     = -32700 // Invalid JSON
	ErrorInvalidRequest = -32600 // Invalid request object
	ErrorMethodNotFound = -32601 // Mthod does not exists
	ErrorInvalidParams  = -32602 // Invalid method parameters
	ErrorInternal       = -32603 // Internal JSON-RPC error
)
