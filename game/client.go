package game

import "golang.org/x/net/websocket"

type Client struct {
	*websocket.Conn
}

func NewClient(conn *websocket.Conn) *Client {
	return &Client{conn}
}

type Clients []*Client

func (clients Clients) Contains(c *Client) bool {
	for _, b := range clients {
		if b == c {
			return true
		}
	}
	return false
}
