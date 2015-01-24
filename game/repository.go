package game

type GameRepository struct {
	Games []*Game
}

func (gs *GameRepository) Find(id string) *Game {
	return NewGame()
}

var Repository = &GameRepository{}
