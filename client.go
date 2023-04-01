package jsonrpc

import "context"

type EndpointType int

const (
	HTTP EndpointType = iota
	WEBSOCKET
)

type Opts struct {
	Type EndpointType
	Url  string
}

type Client interface {
	Call(ctx context.Context, req Request, res *Response) error
}
