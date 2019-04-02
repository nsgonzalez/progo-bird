package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/nsgonzalez/progo-bird/models"
)

func agHandleKeys(last time.Time, gameState *models.GameState, win *pixelgl.Window, character *models.Character, i int, actionsSeq []string) (time.Time, pixel.Vec, []string) {
	action := ""

	ctrl := pixel.ZV

	if *gameState == models.GAME_NO_STARTED && i == 60*TIME_AGENT_FACTOR*2 {
		character.Gravity = CHAR_GRAV
		*gameState = models.GAME_STARTED
	}

	if *gameState == models.GAME_STARTED {
		ctrl.X++

		if i > 0 && i%(60*TIME_AGENT_FACTOR) == 0 {
			dt := time.Since(last).Seconds()
			last = time.Now()
			x := character.Rect.Max.X - (CHAR_W / 2)
			y := character.Rect.Max.Y - (CHAR_H / 2)

			if DEBUG {
				fmt.Println("\ni: " + strconv.Itoa(i))
				fmt.Println("dt: " + fmt.Sprintf("%f", dt))
				fmt.Println("pos: " + fmt.Sprintf("%f", x) + ", " + fmt.Sprintf("%f", y))
			}

			lenActionsSeq := len(actionsSeq)
			if lenActionsSeq > 1 {
				action, actionsSeq = actionsSeq[0], actionsSeq[1:]
			}

			if lenActionsSeq > 0 {
				if DEBUG {
					fmt.Println("next: " + action)
				}

				if action == "jump" {
					ctrl.Y = 1
				}
			}
		}
	} else if *gameState == models.GAME_ENDED {
		if win.JustPressed(pixelgl.KeyEscape) {
			fireEventRestart(gameState, character)
		}
	}

	return last, ctrl, actionsSeq
}
