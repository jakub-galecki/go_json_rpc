package jsonrpc

type Response struct {
	JsonRpc string
	Result  interface{}
	Error   *Error
	Id      string
}

func NewResponse(jsonRpc string, result interface{}, id string, error *Error) *Response {
	return &Response{JsonRpc: jsonRpc, Result: result, Error: error, Id: id}
}
