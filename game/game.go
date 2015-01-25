package game

import (
	"crypto/sha1"
	"encoding/base64"
	"log"
	"strings"
	"time"

	"golang.org/x/net/websocket"
)

const (
	Created Status = iota
	Started
	Ended
)

type MessageChannel chan string

type Game struct {
	Id string
	Status
	Clients
	MessageChannel
}

func (g *Game) Add(conn *websocket.Conn) {
	g.Clients = append(g.Clients, NewClient(conn))
}

func (g *Game) Publish(event string) {
	g.MessageChannel <- event
}

func (g *Game) BroadCast() {
	for msg := range g.MessageChannel {
		for _, client := range g.Clients {
			if err := websocket.Message.Send(client.Conn, msg); err != nil {
				// TODO: Remove conn on failure
				log.Println("Sending failed")
			}
		}
	}
}

type Status int

func NewGame() *Game {
	game := &Game{Id: NewId(), Clients: Clients{}, MessageChannel: make(MessageChannel, 100)}
	go game.BroadCast()
	return game
}

func NewId() string {
	hash := sha1.Sum([]byte(time.Now().String()))
	return strings.Replace(base64.URLEncoding.EncodeToString(hash[:]), "=", "", -1)
}
