package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/nsgonzalez/progo-bird/models"
	"github.com/nsgonzalez/progo-bird/utils"
)

func initialize(character *models.Character) (*pixelgl.Window, *pixelgl.Canvas) {
	cfg := pixelgl.WindowConfig{
		Title:  TITLE,
		Bounds: pixel.R(0, 0, WINDOW_W, WINDOW_H),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	canvas := pixelgl.NewCanvas(pixel.R(-SCENARIO_W, -SCENARIO_H, SCENARIO_W, SCENARIO_H))

	return win, canvas
}

func loadFont() *text.Atlas {
	face, _ := utils.LoadTTF("./res/super-plumber-brothers.ttf", 80)
	return text.NewAtlas(face, text.ASCII)
}
