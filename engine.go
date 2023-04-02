package jsonrpc

type Engine interface {
	MakeRequest(req []byte) (Response, error)
}
