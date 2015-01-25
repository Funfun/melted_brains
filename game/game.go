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
type KillChannel chan bool

type Game struct {
	Id string
	Status
	Clients
	MessageChannel
	KillChannel
}

func (g *Game) KillBroadCast() {
	g.KillChannel <- true
}

func (g *Game) Add(conn *websocket.Conn) {
	g.Clients = append(g.Clients, NewClient(conn))
}
func (g *Game) RemoveClients(toRemove Clients) {
	newClients := Clients{}
	for _, client := range g.Clients {
		if !toRemove.Contains(client) {
			newClients = append(newClients, client)
		}
	}
	g.Clients = newClients
}

func (g *Game) Publish(event string) {
	g.MessageChannel <- event
}

func (g *Game) BroadCast() {
	for {
		select {
		case msg := <-g.MessageChannel:
			erroredClients := Clients{}
			for _, client := range g.Clients {
				if err := websocket.Message.Send(client.Conn, msg); err != nil {
					erroredClients = append(erroredClients, client)
				}
			}
			if len(erroredClients) > 0 {
				g.RemoveClients(erroredClients)
			}

		case _ = <-time.After(10 * time.Minute):
			log.Printf("Timing out game %s\n", g.Id)
			return
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
