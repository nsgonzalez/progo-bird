package models

import (
	"encoding/csv"
	"image"
	"io"
	"os"
	"strconv"

	"github.com/faiface/pixel"
)

type Character struct {
	Name     string
	Nickname string

	Gravity   float64
	RunSpeed  float64
	JumpSpeed float64

	Rect   pixel.Rect
	Vel    pixel.Vec
	Ground bool

	Sheet pixel.Picture
	Anim  *Animation
}

func (character *Character) LoadAnimationSheet(frameWidth float64, frameHeight float64) {

	var sheet pixel.Picture
	var anims map[string][]pixel.Rect
	var err error

	sheetPath := CHARACTERS_DIR + "/" + character.Nickname + ".png"
	descPath := CHARACTERS_DIR + "/" + character.Nickname + ".csv"

	sheetFile, err := os.Open(sheetPath)
	if err != nil {
		panic(err)
	}
	defer sheetFile.Close()
	sheetImg, _, err := image.Decode(sheetFile)
	if err != nil {
		panic(err)
	}
	sheet = pixel.PictureDataFromImage(sheetImg)

	var frames []pixel.Rect
	for y := 0.0; y+frameHeight <= sheet.Bounds().Max.Y; y += frameHeight {
		for x := 0.0; x+frameWidth <= sheet.Bounds().Max.X; x += frameWidth {
			frames = append(frames, pixel.R(
				x,
				y,
				x+frameWidth,
				y+frameHeight,
			))
		}
	}

	descFile, err := os.Open(descPath)
	if err != nil {
		panic(err)
	}
	defer descFile.Close()

	anims = make(map[string][]pixel.Rect)
	desc := csv.NewReader(descFile)
	for {
		anim, err := desc.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		name := anim[0]
		start, _ := strconv.Atoi(anim[1])
		end, _ := strconv.Atoi(anim[2])

		anims[name] = frames[start : end+1]
	}

	character.Sheet = sheet
	character.Anim = &Animation{
		Sheet: sheet,
		Anims: anims,
		Rate:  1.0 / 10,
		Dir:   +1,
	}
}

func (character *Character) Update(dt float64, gameState *GameState, ctrl pixel.Vec, platforms []Platform) {

	if *gameState == GAME_ENDED || *gameState == GAME_WON {
		return
	}

	// apply controls
	switch {
	case ctrl.X < 0:
		character.Vel.X = -character.RunSpeed
	case ctrl.X > 0:
		character.Vel.X = +character.RunSpeed
	default:
		character.Vel.X = 0
	}

	// apply gravity and velocity
	character.Vel.Y += character.Gravity * dt
	character.Rect = character.Rect.Moved(character.Vel.Scaled(dt))

	// check collisions against each platform
	character.Ground = false
	for _, p := range platforms {
		cond0 := (character.Rect.Max.X >= p.Rect.Min.X && character.Rect.Max.X <= p.Rect.Max.X) ||
			(character.Rect.Min.X <= p.Rect.Min.X && character.Rect.Max.X >= p.Rect.Max.X) ||
			(character.Rect.Min.X >= p.Rect.Min.X && character.Rect.Min.X <= p.Rect.Max.X)
		cond1 := p.Type == PLATFORM_TOP && cond0 && character.Rect.Max.Y >= p.Rect.Min.Y
		cond2 := p.Type == PLATFORM_BOTTOM && cond0 && character.Rect.Min.Y <= p.Rect.Max.Y
		if !cond1 && !cond2 {
			continue
		}

		character.Ground = true
	}

	// jump if on the ground and the player wants to jump
	if ctrl.Y > 0 {
		character.Vel.Y = character.JumpSpeed
	}
}
