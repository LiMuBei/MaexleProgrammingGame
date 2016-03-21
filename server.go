package main

import (
	"net/http"
	"golang.org/x/net/websocket"
	"log"
	"fmt"
	"bytes"
)

// Echo the data received on the WebSocket.
func EchoServer(ws *websocket.Conn) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(ws)
	log.Println("Received something:", buf.String())
	buf.WriteTo(ws)
}

// This example demonstrates a trivial echo server.
func main() {
	const port int = 8080
	log.SetFlags(1)
	log.Println("Listening on port", port)
	http.Handle("/echo", websocket.Handler(EchoServer))
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
