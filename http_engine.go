package jsonrpc

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
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
	responseLength, err := strconv.Atoi(httpRes.Header.Get("Content-Length"))

	if err != nil {
		responseLength = MaxBufferSize
	}

	b := make([]byte, responseLength)

	n, err := httpRes.Body.Read(b)
	defer func() { err = httpRes.Body.Close() }()
	if n == 0 && err != io.EOF {
		return Response{}, err
	}
	err = json.Unmarshal(b[:n], &res)
	return
}
