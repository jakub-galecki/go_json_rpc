package jsonrpc

type Response struct {
	JsonRpc string      `json:"jsonrpc"`
	Result  interface{} `json:"result"`
	Error   *Error      `json:"error"`
	Id      string      `json:"id"`
}

func NewResponse(jsonRpc string, result interface{}, id string, error *Error) *Response {
	return &Response{JsonRpc: jsonRpc, Result: result, Error: error, Id: id}
}
