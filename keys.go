package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/nsgonzalez/progo-bird/models"
)

func handleKeys(dt float64, gameState *models.GameState, win *pixelgl.Window, character *models.Character) (float64, pixel.Vec) {

	// acelerate
	if win.Pressed(pixelgl.KeyLeftShift) {
		dt *= 2
	}

	ctrl := pixel.ZV

	if *gameState == models.GAME_NO_STARTED {
		// before start
		if win.JustPressed(pixelgl.KeyEnter) {
			character.Gravity = CHAR_GRAV
			*gameState = models.GAME_STARTED
		}
	} else if *gameState == models.GAME_STARTED {
		// started
		ctrl.X++

		if win.JustPressed(pixelgl.KeyUp) || win.JustPressed(pixelgl.KeySpace) {
			ctrl.Y = 1
		}
	} else if *gameState == models.GAME_ENDED || *gameState == models.GAME_WON {
		// ended
		if win.JustPressed(pixelgl.KeyEscape) {
			fireEventRestart(gameState, character)
		}
	}

	return dt, ctrl
}
