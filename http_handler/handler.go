package http_handler

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gophergala/melted_brains/game"
	"golang.org/x/net/websocket"
)

var showTemplate *template.Template

func init() {
	var err error
	showTemplate, err = template.ParseFiles("static/template/show.thtml")
	if err != nil {
		log.Panicf("Cant parse templates %v", err)
	}
}

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

func GameHandler(w http.ResponseWriter, req *http.Request) {
	id, action := parseGameRequest(req.URL.Path)
	game := game.Repository.Find(id)

	switch action {
	case "join":
		http.Redirect(w, req, "/game/"+game.Id+"/show", http.StatusFound)
	case "show":
		showTemplate.Execute(w, game)
	}
}

func EventsHandler(ws *websocket.Conn) {
	fmt.Printf("In handler")
	io.Copy(ws, ws)
}
