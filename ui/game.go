package ui

import (
	"fmt"
	"fyne.io/fyne/v2/theme"
)

var thisGame *game

func gridClick(ci *customIcon, h, w int) {
	if ci.Resource != theme.ViewFullScreenIcon() {
		return
	}
	var turnInt int
	if thisGame.turn {
		turnInt = 1
	}

	if thisGame.turn {
		ci.SetResource(theme.CancelIcon())
	} else {
		ci.SetResource(theme.RadioButtonIcon())
	}

	thisGame.players[turnInt].grid[h][w] = true
	if checkGrid(thisGame.players[turnInt].grid) {

	}

	thisGame.turn = !thisGame.turn
	changeTurnText(thisGame.turn)
}

func checkGrid(grid [][]bool) bool {
	return false
}

func changeTurnText(turn bool) {
	text := ""
	if turn {
		text = thisGame.players[0].name + "'s"
	} else {
		text = thisGame.players[1].name + "'s"
	}
	thisGame.turnText.Text = text + " turn"
	thisGame.turnText.Refresh()
}

func newEmptyPlayer(num, size int) *player {
	p := &player{
		name: fmt.Sprintf("Player%d", num),
		grid: make([][]bool, size),
	}
	for idx := range p.grid {
		p.grid[idx] = make([]bool, size)
	}
	return p
}
