package client

import (
	"fmt"
	"github.com/gorilla/websocket"
	"time"
	"sync"
)

type ChatRoom struct {
	clients map[string]Client
	clientsMtx sync.Mutex
	queue   chan string
}

func (cr *ChatRoom) Init() {
	fmt.Println("Chatroom init")
	cr.queue = make(chan string, 5)
	cr.clients = make(map[string]Client)

	go func() {
		for {
			cr.BroadCast()
			time.Sleep(100 * time.Millisecond)
		}
	}()
}

//registering a new client
//returns pointer to a Client, or Nil, if the name is already taken
func (cr *ChatRoom) Join(name string, conn *websocket.Conn) *Client {
	defer cr.clientsMtx.Unlock();

	cr.clientsMtx.Lock(); //preventing simultaneous access to the `clients` map
	if _, exists := cr.clients[name]; exists {
		return nil
	}
	client := Client{
		name:      name,
		conn:      conn,
		belongsTo: cr,
	}
	cr.clients[name] = client

	cr.AddMsg("<B>" + name + "</B> has joined the chat.")
	return &client
}

func (cr *ChatRoom) Leave(name string) {
	cr.clientsMtx.Lock(); //preventing simultaneous access to the `clients` map
	delete(cr.clients, name)
	cr.clientsMtx.Unlock();
	cr.AddMsg("<B>" + name + "</B> has left the chat.")
}

func (cr *ChatRoom) AddMsg(msg string) {
	cr.queue <- msg
}

func (cr *ChatRoom) BroadCast() {
	msgBlock := ""
	infLoop:
	for {
		select {
		case m := <-cr.queue:
			fmt.Println("we have a queued item")
			fmt.Println(m)
			msgBlock += m + "<BR>"
		default:
			break infLoop
		}
	}
	if len(msgBlock) > 0 {
		for _, client := range cr.clients {
			client.Send(msgBlock)
		}
	}
}
