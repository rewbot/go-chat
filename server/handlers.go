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
		//conn.ReadMessage waits here until a message is received (in this case the name of the person joining)
		_, msg, err := conn.ReadMessage()
		client := chatRoom.Join(string(msg), conn)
		if client == nil || err != nil {
			fmt.Println(err.Error())
			return
		}

		for {
			//read any messages sent by client after joining
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
