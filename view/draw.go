package view

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/lauth/sokoban/model"
	"golang.org/x/image/colornames"
)

/**
 * Draw screen
 */
func Draw(window *pixelgl.Window, game model.Game, blockSprites BlockSprites) {
	window.Clear(colornames.Black)

	screenY := window.Bounds().Max.Y

	drawMap(window, game, blockSprites, screenY)
	drawBoxes(window, game, blockSprites, screenY)
	drawHero(window, game, blockSprites, screenY)
}

func drawMap(window *pixelgl.Window, game model.Game, blockTypes BlockSprites, screenY float64) {
	for y := range game.Map {
		for x, block := range game.Map[y] {
			position := pixel.IM.Moved(pixel.V(float64(x*32+16), screenY-float64(y*32+16)))
			blockTypes[block].Draw(window, position)
		}
	}
}

func drawBoxes(window *pixelgl.Window, game model.Game, blockTypes BlockSprites, screenY float64) {
	boxPicture, err := loadPicture("resources/box.png")
	if err != nil {
		panic(err)
	}
	boxSprite := pixel.NewSprite(boxPicture, boxPicture.Bounds())
	for _, box := range game.Boxes {
		position := pixel.IM.Moved(pixel.V(float64(box.X*32+16), screenY-float64(box.Y*32+16)))
		boxSprite.Draw(window, position)
	}
}

func drawHero(window *pixelgl.Window, game model.Game, blockTypes BlockSprites, screenY float64) {
	heroPicture, err := loadPicture("resources/hero.png")
	if err != nil {
		panic(err)
	}
	heroSprite := pixel.NewSprite(heroPicture, heroPicture.Bounds())
	position := pixel.IM.Moved(pixel.V(float64(game.Hero.X*32+16), screenY-float64(game.Hero.Y*32+16)))
	heroSprite.Draw(window, position)
}
