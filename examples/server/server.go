package main

import (
	"encoding/json"
	"flag"
	"github.com/gorilla/websocket"
	"jsonrpc"
	"log"
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
	} else {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
	}
}

func serverHttpSetup() {
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

var upgrader = websocket.Upgrader{}

func handleWS(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		panic(err)
	}

	defer c.Close()

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		var req jsonrpc.Request
		err = json.Unmarshal(message, &req)
		if err != nil {
			panic(err)
		}
		log.Println(req)

		res := jsonrpc.NewResponse("2.0", "DONE", "1", nil)
		data, _ := json.Marshal(res)

		err = c.WriteMessage(mt, data)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}

}

func serverWsSetup() {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	http.HandleFunc("/", handleWS)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func main() {
	srvType := flag.String("serverType", "ws", "")
	flag.Parse()

	switch *srvType {
	case "ws":
		serverWsSetup()
	case "http":
		serverHttpSetup()
	}
}
