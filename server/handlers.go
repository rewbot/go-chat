package server

import (
	"github.com/gorilla/websocket"
	"net/http"
	"fmt"
)

//What is this and why do I need it?
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true }, //not checking origin
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Entering websocket handler")
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		fmt.Println("Error upgrading to websocket:", err)
		return
	}
	go func() {
		//first message has to be the name
		_, msg, err := conn.ReadMessage()
		client := chatRoom.Join(string(msg), conn)
		if client == nil || err != nil {
			fmt.Println(err.Error())
			conn.Close() //closing connection to indicate failed Join
			return
		}

		//then watch for incoming messages
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil { //if error then assuming that the connection is closed
				client.Exit()
				return
			}
			fmt.Println("we are listening!")
			fmt.Println(string(msg))
			client.NewMsg(string(msg))
		}

	}()
}

func staticFiles(w http.ResponseWriter, r *http.Request) {
	//http.ServeFile(w, r, "./static/"+r.URL.Path)
	http.ServeFile(w, r, "./static/index.html")
}
