package client

import (
	"fmt"
	"github.com/gorilla/websocket"
	"time"
	"sync"
)

type ChatRoom struct {
	Clients map[string]Client
	clientsMtx sync.Mutex
	queue      chan string
}

func (cr *ChatRoom) Init() {
	fmt.Println("Chatroom init")
	cr.queue = make(chan string, 5)
	cr.Clients = make(map[string]Client)

	go func() {
		for {
			cr.BroadCast()
			time.Sleep(100 * time.Millisecond)
		}
	}()
}

func (chatRoom *ChatRoom) Join(name string, conn *websocket.Conn) *Client {
	defer chatRoom.clientsMtx.Unlock();

	chatRoom.clientsMtx.Lock(); //preventing simultaneous access to the `clients` map
	if _, exists := chatRoom.Clients[name]; exists {
		return nil
	}
	client := Client{
		Name:      name,
		conn:      conn,
		currentChatRoom: chatRoom,
	}
	chatRoom.Clients[name] = client

	chatRoom.AddMsg("<B>" + name + "</B> has joined the chat.")
	return &client
}

func (chatRoom *ChatRoom) Leave(name string) {
	chatRoom.clientsMtx.Lock(); //preventing simultaneous access to the `clients` map
	delete(chatRoom.Clients, name)
	chatRoom.clientsMtx.Unlock();
	chatRoom.AddMsg("<B>" + name + "</B> has left the chat.")
}

func (chatRoom *ChatRoom) AddMsg(msg string) {
	chatRoom.queue <- msg
}

func (chatRoom *ChatRoom) BroadCast() {
	msgBlock := ""
	infLoop:
		for {
			select {
			case m := <-chatRoom.queue:
				msgBlock += m + "<BR>"
			default:
				break infLoop
			}
		}

	if len(msgBlock) > 0 {
		for _, client := range chatRoom.Clients {
			client.Send(msgBlock)
		}
	}
}
