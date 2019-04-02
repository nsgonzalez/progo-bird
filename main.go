package main

import (
	"fmt"
	"math/rand"
	"os/exec"
	"strconv"
	"strings"
	"time"

	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"

	"github.com/nsgonzalez/progo-bird/models"
)

type gameMode int

const (
	GM_MANUAL gameMode = iota
	GM_AGENT

	DEBUG = true
)

func run() {
	rand.Seed(time.Now().UnixNano())

	gameState := models.GAME_NO_STARTED
	gameMode := GM_AGENT

	character, platforms, goal := loadScenario()
	win, canvas := initialize(character)
	atlas := loadFont()

	imdCh := imdraw.New(character.Sheet)
	imdCh.Precision = 32
	camPos := pixel.ZV

	var linkerCmd *exec.Cmd
	actionsSeq := []string{}
	if gameMode == GM_AGENT {

		linkerCmd = linkerExec()
		if linkerCmd == nil {
			return
		}

		actionsSeqStr := linkerQuery(plGenConfigs(goal), plGenPlatforms(&platforms), plGetGoal())
		actionsSeqStr = strings.Replace(actionsSeqStr, "'", "", -1)
		actionsSeqStr = strings.Replace(actionsSeqStr, "[", "", -1)
		actionsSeqStr = strings.Replace(actionsSeqStr, "]", "", -1)
		actionsSeq = strings.Split(actionsSeqStr, ", ")

		if DEBUG {
			fmt.Print("actions: ")
			fmt.Println(actionsSeq)
			fmt.Println("actions: " + strconv.Itoa(len(actionsSeq)))
		}
	}

	// actionsSeq = []string{"jump", "jump", "jump", "jump"}

	i := int(60.0*TIME_AGENT_FACTOR + 3.0)

	last := time.Now()
	last2 := time.Now()
	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()

		camPos = getCamPos(camPos, character, dt)
		canvas.SetMatrix(getCamera(camPos))

		var ctrl pixel.Vec
		if gameMode == GM_MANUAL {
			dt, ctrl = handleKeys(dt, &gameState, win, character)

			if i > 0 && i%int(60.0*TIME_AGENT_FACTOR) == 0 {
				dt2 := time.Since(last2).Seconds()
				last2 = time.Now()
				x := character.Rect.Max.X - (CHAR_W / 2)
				y := character.Rect.Max.Y - (CHAR_H / 2)

				if DEBUG {
					fmt.Println("\ni: " + strconv.Itoa(i))
					fmt.Println("dt: " + fmt.Sprintf("%f", dt2))
					fmt.Println("pos: " + fmt.Sprintf("%f", x) + ", " + fmt.Sprintf("%f", y))
				}
			}
		} else if gameMode == GM_AGENT {
			last2, ctrl, actionsSeq = agHandleKeys(last2, &gameState, win, character, i, actionsSeq)
		}

		// Update the physics and animation
		if gameState != models.GAME_ENDED {
			loopCalc(dt, &gameState, &ctrl, character, &platforms, goal)
		}

		// Draw the scene to the canvas using imdDraw
		loopDraw(&gameState, canvas, imdCh, character, &platforms, goal, atlas)
		loopFitCanvas(win, canvas)
		win.Update()

		i++
	}

	linkerCmd.Process.Kill()
}

func main() {
	pixelgl.Run(run)
}
