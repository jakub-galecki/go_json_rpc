package main

import (
	"encoding/json"
	"fmt"
	"jsonrpc"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var req jsonrpc.Request
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res := jsonrpc.NewResponse("2.0", "DONE", "1", nil)
		data, _ := json.Marshal(res)

		_, err = w.Write(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
	}
}

func serverSetup() {
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func main() {
	serverSetup()

	c := jsonrpc.NewClient(jsonrpc.NewClientOpts(jsonrpc.HTTP, "http://localhost:8080/"))
	req := jsonrpc.NewRequest("dodo", "1", nil)
	res := jsonrpc.Response{}
	if err := c.Call(*req, &res); err != nil {
		panic(err)
	}
	fmt.Println(res)
}
