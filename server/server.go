package server

import (
	"net/http"
	"github.com/rewbotV86/go-chat/client"
)

var chatRoom client.ChatRoom

func AttachHandlers() {
	http.HandleFunc("/ws", wsHandler)
	http.HandleFunc("/", indexHandler)
}

func Init() {
	chatRoom.Init()
}
