package main

import (
	"fmt"
	"github.com/rewbotV86/go-chat/server"
)

func main() {
	fmt.Println("Server running...")
	server.AttachHandlers()
	server.Init()
}