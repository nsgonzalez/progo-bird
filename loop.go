package main

import (
	"fmt"
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/nsgonzalez/progo-bird/models"
	"golang.org/x/image/colornames"
)

func loopCalc(dt float64, gameState *models.GameState, ctrl *pixel.Vec, character *models.Character, platforms *[]models.Platform, goal *models.Goal) {
	character.Update(dt, gameState, *ctrl, *platforms)
	character.Anim.Update(dt, character)
	goal.Update(dt)

	if character.Ground {
		fireEventEnd(gameState)

		x := character.Rect.Max.X - (CHAR_W / 2)
		y := character.Rect.Max.Y - (CHAR_H / 2)

		if DEBUG {
			fmt.Println("pos: " + fmt.Sprintf("%f", x) + ", " + fmt.Sprintf("%f", y))
		}
	}

	cond0 := (character.Rect.Max.X >= goal.Pos.X-goal.Radius && character.Rect.Max.X < goal.Pos.X+goal.Radius) ||
		(character.Rect.Min.X >= goal.Pos.X-goal.Radius && character.Rect.Min.X <= goal.Pos.X+goal.Radius)
	cond1 := cond0 && character.Rect.Max.Y <= goal.Pos.Y+goal.Radius && character.Rect.Max.Y >= goal.Pos.Y-goal.Radius
	if cond1 {
		fireEventWon(gameState)
	}
}

func loopDraw(gameState *models.GameState, canvas *pixelgl.Canvas, imdCh *imdraw.IMDraw, character *models.Character, platforms *[]models.Platform, goal *models.Goal, atlas *text.Atlas) {
	canvas.Clear(colornames.Crimson)
	imdCh.Clear()

	for _, p := range *platforms {
		p.Draw(imdCh)
	}

	character.Anim.Draw(imdCh, character)
	goal.Draw(imdCh)
	imdCh.Draw(canvas)

	txt := text.New(character.Rect.Center(), atlas)
	txt.Color = colornames.Greenyellow

	if *gameState == models.GAME_ENDED {
		txt.WriteString("PERDISTE")
		txt.Draw(canvas, pixel.IM.Moved(character.Rect.Center().Sub(txt.Bounds().Center()).Add(pixel.Vec{X: 0, Y: 100})))
	}

	if *gameState == models.GAME_WON {
		txt.WriteString("GANASTE")
		txt.Draw(canvas, pixel.IM.Moved(character.Rect.Center().Sub(txt.Bounds().Center()).Add(pixel.Vec{X: 0, Y: 100})))
	}
}

func loopFitCanvas(win *pixelgl.Window, canvas *pixelgl.Canvas) {
	win.Clear(colornames.White)
	win.SetMatrix(pixel.IM.Scaled(pixel.ZV,
		math.Min(
			win.Bounds().W()/canvas.Bounds().W(),
			win.Bounds().H()/canvas.Bounds().H(),
		),
	).Moved(win.Bounds().Center()))
	canvas.Draw(win, pixel.IM.Moved(canvas.Bounds().Center()))
}
