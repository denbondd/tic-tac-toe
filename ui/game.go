package ui

import (
	"fmt"
	"fyne.io/fyne/v2/theme"
)

var thisGame *game

func gridClick(ci *customIcon, h, w int, onWin func(playerNum int), onBlock func()) {
	if ci.Resource != theme.ViewFullScreenIcon() {
		return
	}
	var turnInt int
	if thisGame.turn {
		turnInt = 1
	}

	if thisGame.turn {
		ci.SetResource(xRes)
	} else {
		ci.SetResource(oRes)
	}

	thisGame.players[turnInt].grid[h][w] = true
	if checkGrid(thisGame.players[turnInt].grid) {
		if thisGame.turn {
			turnInt = 0
		} else {
			turnInt = 1
		}
		onWin(turnInt)
	} else if checkBlock(thisGame.players[0].grid, thisGame.players[1].grid) {
		onBlock()
	}
	thisGame.turn = !thisGame.turn
	changeTurnText(thisGame.turn)
}

func checkGrid(grid [][]bool) bool {
	answ := grid[0][0] && grid[1][0] && grid[2][0] ||
		grid[0][1] && grid[1][1] && grid[2][1] ||
		grid[0][2] && grid[1][2] && grid[2][2] ||
		grid[0][0] && grid[0][1] && grid[0][2] ||
		grid[1][0] && grid[1][1] && grid[1][2] ||
		grid[2][0] && grid[2][1] && grid[2][2] ||
		grid[0][0] && grid[1][1] && grid[2][2] ||
		grid[0][2] && grid[1][1] && grid[2][0]
	return answ
}

func checkBlock(pl1, pl2 [][]bool) bool {
	block := make([][]bool, len(pl1))
	for i := range block {
		block[i] = make([]bool, len(pl1[0]))
	}
	for i, sl := range pl1 {
		for j, elem := range sl {
			if elem {
				block[i][j] = elem
			}
			if pl2[i][j] {
				block[i][j] = pl2[i][j]
			}
		}
	}
	for _, sl := range block {
		for _, elem := range sl {
			if !elem {
				return false
			}
		}
	}
	return true
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

func restartGame() {
	for _, sl := range thisGame.btns {
		for _, elem := range sl {
			elem.Resource = emptyRes
			elem.Refresh()
			elem.clickable = true
		}
	}
	thisGame.players[0].grid = cleanSlice(thisGame.players[0].grid)
	thisGame.players[1].grid = cleanSlice(thisGame.players[1].grid)
}

func unclickableBtns() {
	for i := range thisGame.btns {
		for j := range thisGame.btns[i] {
			thisGame.btns[i][j].clickable = false
		}
	}
}

func cleanSlice(grid [][]bool) [][]bool {
	answ := make([][]bool, len(grid))
	for idx := range answ {
		answ[idx] = make([]bool, len(grid[0]))
	}
	return answ
}
