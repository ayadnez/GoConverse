package main

import (
	"chat-app/pkg/websocket"
	"fmt"
	"net/http"
)

// define out websocket endpoint

func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	//fmt.Println(r.Host)
	fmt.Println("WebSocket Endpoint Hit")
	// upgrade this connection to a websocket connection
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+V\n", err)
	}
	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}
	pool.Register <- client
	client.Read()
	// listen indefintely for new message comming through on our websocket connection
	// go websocket.Reader(ws)
	// websocket.Reader(ws)
}

func setupRoutes() {
	pool := websocket.NewPool()
	go pool.Start()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})
	// map our '/ws' endpoint to the 'serveWs' function
	// http.HandleFunc("/ws", serveWs)
}
func main() {
	fmt.Println("Distributed Chat-App v0.01")
	setupRoutes()
	http.ListenAndServe(":8081", nil)
}
