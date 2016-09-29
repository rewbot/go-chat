package client

import "github.com/gorilla/websocket"

type Client struct {
	name      string
	conn      *websocket.Conn
	belongsTo *ChatRoom
}

func (cl *Client) NewMsg(msg string) {
	cl.belongsTo.AddMsg("<B>" + cl.name + ":</B> " + msg)
}

func (cl *Client) Exit() {
	cl.belongsTo.Leave(cl.name)
}

func (cl *Client) Send(msgs string) {
	cl.conn.WriteMessage(websocket.TextMessage, []byte(msgs))
}
