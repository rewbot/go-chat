package client

import (
	"github.com/gorilla/websocket"
	"time"
	"sync"
)

type ChatRoom struct {
	Clients map[string]Client
	clientsMtx sync.Mutex
	queue      chan string
}

func (chatRoom *ChatRoom) Init() {
	//The 5 is a buffer. The channel will block when the queue exceeds 5
	chatRoom.queue = make(chan string, 5)
	chatRoom.Clients = make(map[string]Client)

	go func() {
		for {
			chatRoom.BroadCast()
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
	chatRoom.AddMessage("<B>" + name + "</B> has joined the chat.")
	return &client
}

func (chatRoom *ChatRoom) Leave(name string) {
	chatRoom.clientsMtx.Lock(); //preventing simultaneous access to the `clients` map
	delete(chatRoom.Clients, name)
	chatRoom.clientsMtx.Unlock();
	chatRoom.AddMessage("<B>" + name + "</B> has left the chat.")
}

func (chatRoom *ChatRoom) AddMessage(msg string) {
	chatRoom.queue <- msg
}

func (chatRoom *ChatRoom) BroadCast() {
	messageBlock := ""
	infLoop:
		for {
			select {
			case m := <- chatRoom.queue:
				messageBlock += m + "<BR>"
			default:
				break infLoop
			}
		}

	if len(messageBlock) > 0 {
		for _, client := range chatRoom.Clients {
			client.Send(messageBlock)
		}
	}
}
