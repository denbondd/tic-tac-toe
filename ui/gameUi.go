package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

var xRes fyne.Resource
var oRes fyne.Resource
var emptyRes fyne.Resource

func getGameContent(size int) (c *fyne.Container) {
	xRes = theme.CancelIcon()
	oRes = theme.RadioButtonIcon()
	emptyRes = theme.ViewFullScreenIcon()

	thisGame = &game{
		turn:    true,
		btns:    make([][]*customIcon, size),
		players: [2]*player{newEmptyPlayer(1, size), newEmptyPlayer(2, size)},
	}
	for idx := range thisGame.btns {
		thisGame.btns[idx] = make([]*customIcon, size)
	}

	c = container.New(layout.NewVBoxLayout(),
		getTurnText(),
		layout.NewSpacer(),
		getMainGame(size),
		layout.NewSpacer(),
		getExitBtn())
	return
}

func getTurnText() fyne.CanvasObject {
	t := canvas.NewText("Player1's turn", textColor)
	t.TextSize = 32
	thisGame.turnText = t
	c := container.NewCenter(t)
	return c
}

func getMainGame(size int) *fyne.Container {
	c := container.New(layout.NewGridLayoutWithColumns(3),
		getPlayerContent(0),
		getGameGrid(size),
		getPlayerContent(1))
	c = container.NewCenter(c)
	return c
}

func getGameGrid(size int) *fyne.Container {
	g := container.NewGridWithColumns(size)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			g.Add(newGameBtn(i, j))
		}
	}
	return g
}

func newGameBtn(h, w int) *customIcon {
	btn := newCustomIcon(emptyRes, fyne.NewSize(72, 72), func(ci *customIcon) {
		gridClick(ci, h, w, onWin, onBlock)
	})
	thisGame.btns[h][w] = btn
	return btn
}

func onWin() {
	turnInt := 0
	if !thisGame.turn {
		turnInt = 1
	}
	message := thisGame.players[turnInt].name + " win!"
	dialog.ShowInformation("Congratulations", message, currWindow)
	thisGame.players[turnInt].Win()
	unclickableBtns()
}

func onBlock() {
	dialog.ShowConfirm("Block!", "No one wins, restart?", func(b bool) {
		if b {
			restartGame()
		}
	}, currWindow)
}

func getPlayerContent(plNum int) *fyne.Container {
	icnRes := *new(fyne.Resource)
	if plNum == 0 {
		icnRes = xRes
	} else {
		icnRes = oRes
	}
	icn := newCustomIcon(icnRes, fyne.NewSize(64, 64), nil)
	entry := widget.NewEntry()
	entry.Text = "Player" + strconv.Itoa(plNum+1)
	btn := widget.NewButton("Save name", func() {
		thisGame.players[plNum].name = entry.Text
		changeTurnText(thisGame.turn)
	})
	score := canvas.NewText("0", textColor)
	thisGame.players[plNum].winsText = score
	scoreCont := container.NewCenter(score)
	c := container.New(layout.NewVBoxLayout(),
		icn,
		entry,
		btn,
		scoreCont)
	return c
}

func getExitBtn() *fyne.Container {
	restartBtn := widget.NewButton("Restart", func() {
		restartGame()
	})
	exitBtn := widget.NewButton("Exit", func() {
		dialog.ShowConfirm("Exit", "Are you sure you want to exit?", func(b bool) {
			if b {
				currWindow.SetContent(GetStartContent())
			}
		}, currWindow)
	})
	c := container.NewHBox(restartBtn, exitBtn)
	cCenter := container.NewCenter(c)
	return cCenter
}
