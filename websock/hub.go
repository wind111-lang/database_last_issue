package websock

import (
	"chat/structs"
	"encoding/json"
	"log"
)

type Hub struct {
	clients map[*Client]bool

	broadcast chan structs.JsonReturn

	register chan *Client

	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan structs.JsonReturn),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.RegisterClient(client)
		case client := <-h.unregister:
			h.UnregisterClient(client)
		case msg := <-h.broadcast:
			h.BroadcastClient(msg)
		}
	}
}

func (h *Hub) RegisterClient(client *Client) {
	h.clients[client] = true
}

func (h *Hub) UnregisterClient(client *Client) {
	if _, ok := h.clients[client]; ok {
		//close(client.send)
		delete(h.clients, client)
	}
}

func (h *Hub) BroadcastClient(msg structs.JsonReturn) {
	jsonR, err := json.Marshal(msg)
	if err != nil {
		log.Fatal(err)
	}
	for client := range h.clients {
		client.send <- jsonR
	}
}
