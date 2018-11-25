package fixtures

import "github.com/lauth/sokoban/model"

func LoadBoxes() []model.Box {
	return []model.Box{{X: 1, Y: 2}, {X: 5, Y: 3}, {X: 4, Y: 4}, {X: 1, Y: 6}, {X: 3, Y: 6}, {X: 6, Y: 6}, {X: 4, Y: 7}}
}
