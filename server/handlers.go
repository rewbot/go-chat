package server

import (
	"github.com/gorilla/websocket"
	"net/http"
	"fmt"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	//Upgrades http connection to websocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading to websocket:", err)
		return
	}

	go func() {
		//first message has to be the name, it waits here until message is received
		_, msg, err := conn.ReadMessage()
		client := chatRoom.Join(string(msg), conn)
		if client == nil || err != nil {
			fmt.Println(err.Error())
			return
		}

		//then watch for incoming messages
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				fmt.Println("Error occurred on conn.ReadMessage: " + err.Error())
				client.Exit()
				return
			}
			message := string(msg)
			client.NewMessage(message)
		}
	}()
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}
