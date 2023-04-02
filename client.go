package jsonrpc

import (
	"encoding/json"
)

type EndpointType int

const (
	HTTP EndpointType = iota
	WEBSOCKET
)

type Opts struct {
	Type EndpointType
	Url  string
}

func NewClientOpts(endpointType EndpointType, url string) *Opts {
	return &Opts{
		Type: endpointType,
		Url:  url,
	}
}

type Client interface {
	Call(req Request, res *Response) error
}

type rpcC struct {
	url    string
	eType  EndpointType
	engine Engine
}

func NewClient(opts *Opts) Client {
	var e Engine
	switch opts.Type {
	case HTTP:
		e = newHttpEngine(opts.Url)
	case WEBSOCKET:
		e = newWebsocketEngine(opts.Url)
	}
	return &rpcC{
		url:    opts.Url,
		eType:  opts.Type,
		engine: e,
	}
}

func (c *rpcC) Call(req Request, res *Response) error {
	return c.makeCall(req, res)
}

func (c *rpcC) makeCall(req Request, res *Response) error {
	bytes, err := json.Marshal(req)
	if err != nil {
		return err
	}
	r, err := c.engine.MakeRequest(bytes)
	*res = r
	if err != nil {
		return err
	}
	return nil
}
