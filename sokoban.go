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
		Bounds: pixel.R(0, 0, 512, 576),
		VSync:  true,
	}
	window, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	// window.SetSmooth(true)

	blockSprites := view.LoadBlockSprites()

	game := model.Game{}
	// TODO: load game level from file.
	game.Map = fixtures.LoadMap()
	game.Boxes = fixtures.LoadBoxes()
	game.Hero = fixtures.LoadHero()

	for !window.Closed() {
		heroX, heroY := game.Hero.GetPosition()
		if window.JustPressed(pixelgl.KeyLeft) {
			heroX, heroY = game.Hero.Move("left")
		} else if window.JustPressed(pixelgl.KeyRight) {
			heroX, heroY = game.Hero.Move("right")
		} else if window.JustPressed(pixelgl.KeyUp) {
			heroX, heroY = game.Hero.Move("up")
		} else if window.JustPressed(pixelgl.KeyDown) {
			heroX, heroY = game.Hero.Move("down")
		}

		if game.IsWall(heroX, heroY) {
			// perdu

		} else {
			game.Hero.X = heroX
			game.Hero.Y = heroY
		}

		view.Draw(window, game, blockSprites)
		window.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
