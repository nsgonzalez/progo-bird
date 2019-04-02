package main

import (
	"github.com/faiface/pixel"
	models "github.com/nsgonzalez/progo-bird/models"
)

func loadCharacter(name string, nickname string) *models.Character {
	character := models.Character{
		Name:      name,
		Nickname:  nickname,
		Gravity:   CHAR_GRAV_INIT,
		RunSpeed:  CHAR_RUNSPEED,
		JumpSpeed: CHAR_JUMPSPEED,
		Rect:      pixel.R(-(CHAR_W / 2), -(CHAR_H / 2), (CHAR_W / 2), (CHAR_H / 2)),
	}
	character.LoadAnimationSheet(CHAR_W, CHAR_H)
	return &character
}

func loadMainCharacter() *models.Character {
	return loadCharacter("Yellow Bird", "yellowbird")
}
