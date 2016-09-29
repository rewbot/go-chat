package client

import "github.com/gorilla/websocket"

type Client struct {
	Name            string
	conn            *websocket.Conn
	currentChatRoom *ChatRoom
}

func (client *Client) NewMsg(msg string) {
	client.currentChatRoom.AddMsg("<B>" + client.Name + ":</B> " + msg)
}

func (client *Client) Exit() {
	client.currentChatRoom.Leave(client.Name)
}

func (client *Client) Send(msgs string) {
	client.conn.WriteMessage(websocket.TextMessage, []byte(msgs))
}
