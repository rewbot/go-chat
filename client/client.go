package client

import "github.com/gorilla/websocket"

type Client struct {
	Name            string
	conn            *websocket.Conn
	currentChatRoom *ChatRoom
}

func (client *Client) NewMessage(message string) {
	//TODO: Uncomment this line after AddMessage has been implemented
	//client.currentChatRoom.AddMessage("<B>" + client.Name + ":</B> " + message)
}

func (client *Client) Exit() {
	client.currentChatRoom.Leave(client.Name)
}

func (client *Client) Send(messages string) {
	client.conn.WriteMessage(websocket.TextMessage, []byte(messages))
}
