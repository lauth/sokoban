package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/lauth/sokoban/fixtures"
	"github.com/lauth/sokoban/model"
	"github.com/lauth/sokoban/view"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Sokoban",
		Bounds: pixel.R(0, 0, 400, 450),
		VSync:  true,
	}
	window, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	blockSprites := view.LoadBlockSprites()

	game := model.Game{}
	// TODO: load game level from file.
	game.Map = fixtures.LoadMap()
	game.Boxes = fixtures.LoadBoxes()
	game.Hero = fixtures.LoadHero()

	view.Draw(window, game, blockSprites)

	for !window.Closed() {
		window.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
