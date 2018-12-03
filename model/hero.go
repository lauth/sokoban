package model

/**
 * Hero has coordinates.
 */
type Hero struct {
	X, Y int
}

/**
 * Move hero to given direction.
 * TODO: check bounds.
 */
func (hero *Hero) Move(direction string) {
	if direction == "left" {
		hero.X--
	} else if direction == "right" {
		hero.X++
	} else if direction == "up" {
		hero.Y--
	} else if direction == "down" {
		hero.Y++
	}
}
