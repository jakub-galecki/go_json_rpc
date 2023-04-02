package main

import (
	"fmt"
	"jsonrpc"
)

func main() {
	c := jsonrpc.NewClient(jsonrpc.NewClientOpts(jsonrpc.HTTP, "http://localhost:8080/"))
	req := jsonrpc.NewRequest("dodo", "1", nil)
	res := jsonrpc.Response{}
	if err := c.Call(*req, &res); err != nil {
		panic(err)
	}
	fmt.Println(res)
}
