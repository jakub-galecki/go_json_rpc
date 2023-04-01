package jsonrpc

type Notification struct {
	JsonRpc string
	Method  string
	Params  interface{}
}
