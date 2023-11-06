package websocket

import (
	"fmt"
)

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case Client := <-pool.Register:
			pool.Clients[Client] = true
			fmt.Println("size of Connection Pool:", len(pool.Clients))
			for Client, _ := range pool.Clients {
				fmt.Println(Client)
				Client.Conn.WriteJSON(Message{Type: 1, Body: "New User Joined..."})
			}
			break

		case Client := <-pool.Unregister:
			delete(pool.Clients, Client)
			fmt.Println("size of Connection pool:", len(pool.Clients))
			for Client, _ := range pool.Clients {
				Client.Conn.WriteJSON(Message{Type: 1, Body: "user Disconnected"})

			}
			break

		case message := <-pool.Broadcast:
			fmt.Println("sending message to all clients in pool")
			for Client, _ := range pool.Clients {
				if err := Client.Conn.WriteJSON(message); err != nil {
					fmt.Println(err)
					return
				}
			}

		}
	}
}
