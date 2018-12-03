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
func (hero *Hero) Move(direction string) (int, int) {
	X := hero.X
	Y := hero.Y

	if direction == "left" {
		X--
	} else if direction == "right" {
		X++
	} else if direction == "up" {
		Y--
	} else if direction == "down" {
		Y++
	}

	return X, Y
}

/**
 * Get hero position.
 */
func (hero Hero) GetPosition() (int, int) {
	return hero.X, hero.Y
}
