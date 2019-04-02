package main

import (
	"math"

	"github.com/faiface/pixel"
	models "github.com/nsgonzalez/progo-bird/models"
)

func getCamPos(camPos pixel.Vec, character *models.Character, dt float64) pixel.Vec {
	return pixel.Lerp(camPos, character.Rect.Center(), 1-math.Pow(1.0/128, dt))
}

func getCamera(camPos pixel.Vec) pixel.Matrix {
	return pixel.IM.Moved(camPos.Scaled(-1))
}
