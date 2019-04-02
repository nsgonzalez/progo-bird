package models

import (
	"math"

	"github.com/faiface/pixel"
)

type animState int

const (
	AS_IDLE animState = iota
	AS_RUNNING
	AS_JUMPING
	AS_FLYING
)

type Animation struct {
	Sheet pixel.Picture
	Anims map[string][]pixel.Rect
	Rate  float64

	State   animState
	Counter float64
	Dir     float64

	Frame pixel.Rect

	Sprite *pixel.Sprite
}

func (anim *Animation) Update(dt float64, char *Character) {
	anim.Counter += dt

	var newState animState

	switch {
	case !char.Ground:
		newState = AS_FLYING
	case char.Vel.Len() == 0:
		newState = AS_IDLE
	}

	if anim.State != newState {
		anim.State = newState
		anim.Counter = 0
	}

	switch anim.State {
	case AS_IDLE:
		anim.Frame = anim.Anims["Fly"][0]
	case AS_FLYING:
		i := int(math.Floor(anim.Counter / anim.Rate))
		anim.Frame = anim.Anims["Fly"][i%len(anim.Anims["Fly"])]
	}
}

func (anim *Animation) Draw(t pixel.Target, char *Character) {
	if anim.Sprite == nil {
		anim.Sprite = pixel.NewSprite(nil, pixel.Rect{})
	}

	anim.Sprite.Set(anim.Sheet, anim.Frame)
	anim.Sprite.Draw(t, pixel.IM.
		ScaledXY(pixel.ZV, pixel.V(
			char.Rect.W()/anim.Sprite.Frame().W(),
			char.Rect.H()/anim.Sprite.Frame().H(),
		)).
		ScaledXY(pixel.ZV, pixel.V(anim.Dir, 1)).
		Moved(char.Rect.Center()))
}
