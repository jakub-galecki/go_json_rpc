package jsonrpc

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

type httpEngine struct {
	url string
}

func newHttpEngine(url string) *httpEngine {
	return &httpEngine{url: url}
}

func (h *httpEngine) MakeRequest(req []byte) (res Response, err error) {
	bReader := bytes.NewReader(req)
	httpReq, err := http.NewRequest(http.MethodPost, h.url, bReader)
	if err != nil {
		return Response{}, err
	}

	c := http.Client{
		Timeout: 1 * time.Second,
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpRes, err := c.Do(httpReq)

	var b []byte
	_, err = httpRes.Body.Read(b)
	defer func() { err = httpRes.Body.Close() }()
	if err != nil {
		return Response{}, err
	}
	err = json.Unmarshal(b, &res)
	return
}
