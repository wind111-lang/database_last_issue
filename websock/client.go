package websock

import (
	"bytes"
	"chat/api"
	"chat/structs"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait  = 10 * time.Second
	pongWait   = 60 * time.Second
	pingPeriod = (pongWait * 9) / 10
	maxMsgSize = 8192
)

type Client struct {
	hub  *Hub
	conn *websocket.Conn
	send chan []byte
}

func NewClient(hub *Hub, conn *websocket.Conn) *Client {
	return &Client{
		hub:  hub,
		conn: conn,
		send: make(chan []byte, 256),
	}
}

var (
	line = []byte{'\n'}
	spc  = []byte{' '}
)

var WsConnect = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func (cli *Client) ReadPump() {
	defer func() {
		cli.Disconnect()
	}()

	cli.conn.SetReadLimit(maxMsgSize)
	cli.conn.SetReadDeadline(time.Now().Add(pongWait))
	cli.conn.SetPongHandler(func(string) error { cli.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := cli.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Fatal(err)
			}
			break
		}

		var mess structs.JsonMessage

		if err := json.Unmarshal(message, &mess); err != nil {
			log.Fatal(err)
			return
		}

		//fmt.Print(mess.Message)
		//fmt.Printf(",")
		translate := api.Translate([]byte(mess.Message))

		message = bytes.TrimSpace(bytes.Replace(translate, line, spc, -1))

		var nam structs.JsonReturn

		nam.Name = mess.Name
		nam.Text = string(message)
		//fmt.Println(string(message))

		//dbctl.InsertNewTask(nam)

		cli.hub.broadcast <- nam
	}
}

func (cli *Client) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		cli.conn.Close()
	}()

	for {
		select {
		case message, ok := <-cli.send:
			cli.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				cli.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := cli.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}

			w.Write(message)

			n := len(cli.send)
			for i := 0; i < n; i++ {
				w.Write(line)
				w.Write(<-cli.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			cli.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := cli.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (cli *Client) Disconnect() {
	cli.hub.unregister <- cli
	close(cli.send)
	cli.conn.Close()
}

func ServeWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := WsConnect.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	client := NewClient(hub, conn)
	fmt.Println("New Client joined the hub.")

	go client.WritePump()
	go client.ReadPump()

	hub.register <- client
}
