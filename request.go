package jsonrpc

type Request struct {
	JsonRpc string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	Id      string      `json:"id"`
}

func NewRequest(method, id string, params interface{}) *Request {
	return &Request{
		Method:  method,
		Id:      id,
		Params:  params,
		JsonRpc: "2.0",
	}
}
