package model

type Game struct {
	Map   [][]Block
	Boxes []Box
	Hero  Hero
}

func (game Game) IsWall(x int, y int) bool {
	return game.Map[y][x] == WALL
}
