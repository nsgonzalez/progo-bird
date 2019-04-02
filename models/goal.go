package models

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/nsgonzalez/progo-bird/utils"
)

type Goal struct {
	Pos    pixel.Vec
	Radius float64
	Step   float64

	Counter float64
	Cols    [5]pixel.RGBA
}

func (g *Goal) Update(dt float64) {
	g.Counter += dt
	for g.Counter > g.Step {
		g.Counter -= g.Step
		for i := len(g.Cols) - 2; i >= 0; i-- {
			g.Cols[i+1] = g.Cols[i]
		}
		g.Cols[0] = utils.RandomNiceColor()
	}
}

func (g *Goal) Draw(imd *imdraw.IMDraw) {
	for i := len(g.Cols) - 1; i >= 0; i-- {
		imd.Color = g.Cols[i]
		imd.Push(g.Pos)
		imd.Circle(float64(i+1)*g.Radius/float64(len(g.Cols)), 0)
	}
}
