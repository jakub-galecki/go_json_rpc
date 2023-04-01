package jsonrpc

type Request struct {
	JsonRpc string
	Method  string
	Params  interface{}
	Id      string
}

func NewRequest(method, id string, params interface{}) *Request {
	return &Request{
		Method:  method,
		Id:      id,
		Params:  params,
		JsonRpc: "2.0",
	}
}
