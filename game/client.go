package game

import "golang.org/x/net/websocket"

type Client struct {
	*websocket.Conn
}

func NewClient(conn *websocket.Conn) *Client {
	return &Client{conn}
}

type Clients []*Client
