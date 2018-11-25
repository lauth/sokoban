package fixtures

import "github.com/lauth/sokoban/model"

func LoadMap() model.Map {
	return model.Map{
		{model.FLOOR, model.FLOOR, model.WALL, model.WALL, model.WALL, model.WALL, model.WALL, model.FLOOR},
		{model.WALL, model.WALL, model.WALL, model.FLOOR, model.FLOOR, model.FLOOR, model.WALL, model.FLOOR},
		{model.WALL, model.LANDING, model.FLOOR, model.FLOOR, model.FLOOR, model.FLOOR, model.WALL, model.FLOOR},
		{model.WALL, model.WALL, model.WALL, model.FLOOR, model.FLOOR, model.LANDING, model.WALL, model.FLOOR},
		{model.WALL, model.LANDING, model.WALL, model.WALL, model.FLOOR, model.FLOOR, model.WALL, model.FLOOR},
		{model.WALL, model.FLOOR, model.WALL, model.FLOOR, model.LANDING, model.FLOOR, model.WALL, model.WALL},
		{model.WALL, model.FLOOR, model.FLOOR, model.LANDING, model.FLOOR, model.FLOOR, model.LANDING, model.WALL},
		{model.WALL, model.FLOOR, model.FLOOR, model.FLOOR, model.LANDING, model.FLOOR, model.FLOOR, model.WALL},
		{model.WALL, model.WALL, model.WALL, model.WALL, model.WALL, model.WALL, model.WALL, model.WALL}}
}
