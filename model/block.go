package model

/**
 * Block.
 */
type Block int

/**
 * Kinds of block.
 */
const (
	FLOOR Block = iota
	WALL
	LANDING
)
