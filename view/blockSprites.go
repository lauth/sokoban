package view

import (
	"github.com/faiface/pixel"
	"github.com/lauth/sokoban/model"
)

/**
 * Get sprite for a block.
 */
type BlockSprites map[model.Block]*pixel.Sprite

func LoadBlockSprites() BlockSprites {
	blockSprites := BlockSprites{}

	wallPicture, err := loadPicture("resources/wall.png")
	if err != nil {
		panic(err)
	}
	blockSprites[model.WALL] = pixel.NewSprite(wallPicture, wallPicture.Bounds())

	floorPicture, err := loadPicture("resources/floor.png")
	if err != nil {
		panic(err)
	}
	blockSprites[model.FLOOR] = pixel.NewSprite(floorPicture, floorPicture.Bounds())

	landingPicture, err := loadPicture("resources/landing.png")
	if err != nil {
		panic(err)
	}
	blockSprites[model.LANDING] = pixel.NewSprite(landingPicture, landingPicture.Bounds())

	return blockSprites
}
