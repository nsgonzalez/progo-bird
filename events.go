package main

import (
	"github.com/faiface/pixel"
	"github.com/nsgonzalez/progo-bird/models"
)

func fireEventEnd(gameState *models.GameState) {
	*gameState = models.GAME_ENDED
}

func fireEventRestart(gameState *models.GameState, character *models.Character) {
	*gameState = models.GAME_NO_STARTED
	character.Gravity = 0
	character.Rect = character.Rect.Moved(character.Rect.Center().Scaled(-1))
	character.Vel = pixel.ZV
}

func fireEventWon(gameState *models.GameState) {
	*gameState = models.GAME_WON
}
