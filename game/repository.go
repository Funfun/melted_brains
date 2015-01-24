package game

import "math/rand"

type GameRepository struct {
	Games
}

func (gs *GameRepository) RandomJoinable() *Game {
	games := gs.Games.Joinable()
	if len(games) == 0 {
		return gs.Create()
	}
	return games[rand.Intn(len(games))]
}

func (gs *GameRepository) Find(id string) *Game {
	return gs.Games.Find(id)
}

func (gs *GameRepository) Create() *Game {
	game := NewGame()
	gs.Games = append(gs.Games, game)
	return game
}

var Repository = &GameRepository{}
