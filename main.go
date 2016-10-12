package main

import (
	"fmt"
	"net/http"
	"github.com/rewbotV86/go-chat/server"
)

/*
Channels
Named for loop label
websocket.Upgrader: Upgrades http connection to websocket connection
conn.ReadMessage(): waits until next message received before continuing
 */

func main() {
	fmt.Println("Server running...")
	server.AttachHandlers()
	server.Init()
	http.ListenAndServe(":8000", nil)
}