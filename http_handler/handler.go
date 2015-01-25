package http_handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/gophergala/melted_brains/game"
	"golang.org/x/net/websocket"
)

var showTemplate *template.Template
var newUserTemplate *template.Template

func init() {
	var err error
	showTemplate, err = template.ParseFiles("static/template/show.thtml")
	if err != nil {
		log.Panicf("Cant parse templates %v", err)
	}
	newUserTemplate, err = template.ParseFiles("static/template/newUser.thtml")
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

func getGame(id string) *game.Game {
	if id == "random" {
		return game.Repository.RandomJoinable()
	} else {
		return game.Repository.Find(id)
	}
}
func GameHandler(w http.ResponseWriter, req *http.Request) {
	id, action := parseGameRequest(req.URL.Path)
	currentGame := getGame(id)
	if currentGame == nil {
		http.NotFound(w, req)
	}
	currentUser := getUser(req)
	fmt.Printf("%v, %v\n", id, action)
	switch action {
	case "join":
		if currentUser == nil {
			newUserTemplate.Execute(w, currentGame)
		} else {
			http.Redirect(w, req, "/game/"+currentGame.Id+"/show", http.StatusFound)
		}
	case "show":
		if currentUser == nil {
			newUserTemplate.Execute(w, currentGame)
		} else {
			showTemplate.Execute(w, currentGame)
		}
	case "new_user":
		createUser(w, req)
		http.Redirect(w, req, "/game/"+currentGame.Id+"/show", http.StatusFound)
	}
}
func getUser(req *http.Request) *game.User {
	cookie, err := req.Cookie("username")
	if err != nil {
		return nil
	} else {
		return game.NewUser(cookie.Value)
	}
}
func setUser(user *game.User, w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{Name: "username", Value: user.Name})
}
func createUser(w http.ResponseWriter, req *http.Request) {
	username := req.FormValue("username")
	fmt.Printf("USERNAME: %v\n", username)
	var user *game.User
	if username == "" {
		user = game.UserWithRandomName()
	} else {
		user = game.NewUser(username)
	}
	setUser(user, w)
}

func EventsHandler(ws *websocket.Conn) {
	id, _ := parseGameRequest(ws.Request().URL.Path)
	currentGame := getGame(id)
	currentGame.Add(ws)

	for {
		var event string
		if err := websocket.Message.Receive(ws, &event); err != nil {
			//TODO: Remove client on error
			// currentGame.ClientLost()
			return
		}
		currentGame.Publish(event)
	}
}
