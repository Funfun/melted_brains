package game

import (
	"crypto/sha1"
	"encoding/base64"
	"strings"
	"time"
)

const (
	Created GameStatus = iota
	Started
	Ended
)

type Game struct {
	Id string
	GameStatus
}

type GameStatus int

func NewGame() *Game {

	return &Game{Id: NewId()}
}

func NewId() string {
	hash := sha1.Sum([]byte(time.Now().String()))
	return strings.Replace(base64.URLEncoding.EncodeToString(hash[:]), "=", "", -1)
}
