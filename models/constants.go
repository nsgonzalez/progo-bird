package models

type GameState int

const (
	GAME_NO_STARTED GameState = iota
	GAME_STARTED
	GAME_ENDED
	GAME_WON

	SPRITES_DIR    = "./sprites"
	CHARACTERS_DIR = SPRITES_DIR + "/characters"
)
