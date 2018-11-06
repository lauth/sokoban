package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type Block int

const (
	FLOOR Block = iota
	WALL
	LANDING
)

type Box struct {
	X, Y int
}

type Hero struct {
	X, Y int
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Sokoban",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	window, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	testMap := [][]Block{
		{FLOOR, FLOOR, WALL, WALL, WALL, WALL, WALL, FLOOR},
		{WALL, WALL, WALL, FLOOR, FLOOR, FLOOR, WALL, FLOOR},
		{WALL, LANDING, FLOOR, FLOOR, FLOOR, FLOOR, WALL, FLOOR},
		{WALL, WALL, WALL, FLOOR, FLOOR, LANDING, WALL, FLOOR},
		{WALL, LANDING, WALL, WALL, FLOOR, FLOOR, WALL, FLOOR},
		{WALL, FLOOR, WALL, FLOOR, LANDING, FLOOR, WALL, WALL},
		{WALL, FLOOR, FLOOR, LANDING, FLOOR, FLOOR, LANDING, WALL},
		{WALL, FLOOR, FLOOR, FLOOR, LANDING, FLOOR, FLOOR, WALL},
		{WALL, WALL, WALL, WALL, WALL, WALL, WALL, WALL}}

	boxes := []Box{{X: 1, Y: 2}, {X: 5, Y: 3}, {X: 4, Y: 4}, {X: 1, Y: 6}, {X: 3, Y: 6}, {X: 6, Y: 6}, {X: 4, Y: 7}}

	hero := Hero{X: 4, Y: 3}

	wallPicture, err := LoadPicture("resources/wall.png")
	if err != nil {
		panic(err)
	}
	wall := pixel.NewSprite(wallPicture, wallPicture.Bounds())

	floorPicture, err := LoadPicture("resources/floor.png")
	if err != nil {
		panic(err)
	}
	floor := pixel.NewSprite(floorPicture, floorPicture.Bounds())

	landingPicture, err := LoadPicture("resources/landing.png")
	if err != nil {
		panic(err)
	}
	landing := pixel.NewSprite(landingPicture, landingPicture.Bounds())

	window.Clear(colornames.Black)

	// Draw map
	for y, _ := range testMap {
		for x, block := range testMap[y] {
			position := pixel.IM.Moved(pixel.V(float64(x*32+16), 768-float64(y*32+16)))
			if block == WALL {
				wall.Draw(window, position)
			} else {
				if block == FLOOR {
					floor.Draw(window, position)
				} else {
					if block == LANDING {
						landing.Draw(window, position)
					}
				}

			}
		}
	}

	// Draw boxes
	boxPicture, err := LoadPicture("resources/box.png")
	if err != nil {
		panic(err)
	}
	boxSprite := pixel.NewSprite(boxPicture, boxPicture.Bounds())
	for _, box := range boxes {
		position := pixel.IM.Moved(pixel.V(float64(box.X*32+16), 768-float64(box.Y*32+16)))
		boxSprite.Draw(window, position)
	}

	// Draw hero
	heroPicture, err := LoadPicture("resources/hero.png")
	if err != nil {
		panic(err)
	}
	heroSprite := pixel.NewSprite(heroPicture, heroPicture.Bounds())
	position := pixel.IM.Moved(pixel.V(float64(hero.X*32+16), 768-float64(hero.Y*32+16)))
	heroSprite.Draw(window, position)

	for !window.Closed() {
		window.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
