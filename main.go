package main

import (
	"fmt"
	"net/http"
	"github.com/rewbotV86/go-chat/server"
)

func main() {
	fmt.Println("Server running...")
	server.AttachHandlers()
	server.Init()
	http.ListenAndServe(":8000", nil)
}