package client

import (
	"sync"
)

type ChatRoom struct {
	Clients map[string]Client
	clientsMtx sync.Mutex
	queue      chan string
}

func (chatRoom *ChatRoom) Leave(name string) {
	chatRoom.clientsMtx.Lock(); //preventing simultaneous access to the `clients` map
	delete(chatRoom.Clients, name)
	chatRoom.clientsMtx.Unlock();
	//TODO: Uncomment this line after AddMessage has been implemented
	//chatRoom.AddMessage("<B>" + name + "</B> has left the chat.")
}
