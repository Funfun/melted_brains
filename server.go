package main

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

type Game struct {
	Id string
	GameStatus
}
type GameStatus int

type GameServer struct {
	Games []*Game
}

func (gs *GameServer) Find(id string) *Game {
	return NewGame()
}
func NewGame() *Game {

	return &Game{Id: NewId()}
}

func NewId() string {
	hash := sha1.Sum([]byte(time.Now().String()))
	return strings.Replace(base64.URLEncoding.EncodeToString(hash[:]), "=", "", -1)
}

var CurrentGameServer = &GameServer{}

const (
	Created GameStatus = iota
	Started
	Ended
)

func parseGameRequest(path string) (id string, action string) {
	parts := strings.Split(path, "/")
	id = "random"
	action = "join"

	if len(parts) >= 3 {
		id = parts[2]
	}
	if len(parts) >= 4 {
		action = parts[3]
	}
	return
}
func gameHandler(w http.ResponseWriter, req *http.Request) {
	id, action := parseGameRequest(req.URL.Path)
	game := CurrentGameServer.Find(id)
	if action == "join" {
		io.WriteString(w, game.Id)
	}

	fmt.Printf("%v %v\n", id, action)
	io.WriteString(w, "")
}

func main() {
	http.HandleFunc("/game/", gameHandler)
	http.Handle("/", http.FileServer(http.Dir("./static")))
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
