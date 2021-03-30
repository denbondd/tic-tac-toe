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

func getGameContent(a fyne.App, w fyne.Window, size int) (c *fyne.Container) {
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
		getMainGame(size, w),
		layout.NewSpacer(),
		getExitBtn(a, w))
	return
}

func getTurnText() fyne.CanvasObject {
	t := canvas.NewText("Player1's turn", textColor)
	thisGame.turnText = t
	thisGame.turnText.TextSize = 32
	c := container.NewCenter(t)
	return c
}

func getMainGame(size int, w fyne.Window) *fyne.Container {
	c := container.New(layout.NewGridLayoutWithColumns(3),
		getPlayerContent(0),
		getGameGrid(size, w),
		getPlayerContent(1))
	c = container.NewCenter(c)
	return c
}

func getGameGrid(size int, w fyne.Window) *fyne.Container {
	g := container.NewGridWithColumns(size)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			g.Add(newGameBtn(i, j, w))
		}
	}
	return g
}

func newGameBtn(h, w int, wind fyne.Window) *customIcon {
	btn := newCustomIcon(emptyRes, fyne.NewSize(72, 72), func(ci *customIcon) {
		gridClick(
			ci, h, w,
			func(playerNum int) {
				message := thisGame.players[playerNum].name + " win!"
				dialog.ShowInformation("Congratulations", message, wind)
				thisGame.players[playerNum].Win()
				unclickableBtns()
			},
			func() {
				dialog.ShowConfirm("Block!", "No one wins, restart?", func(b bool) {
					if b {
						restartGame()
					}
				}, wind)
			})
	})
	thisGame.btns[h][w] = btn
	return btn
}

func getPlayerContent(plNum int) *fyne.Container {
	icnRes := *new(fyne.Resource)
	if plNum == 0 {
		icnRes = theme.CancelIcon()
	} else {
		icnRes = theme.RadioButtonIcon()
	}
	icn := newCustomIconWithoutFunc(icnRes, fyne.NewSize(64, 64))
	et := widget.NewEntry()
	et.Text = "Player" + strconv.Itoa(plNum+1)
	btn := widget.NewButton("Save name", func() {
		thisGame.players[plNum].name = et.Text
		changeTurnText(thisGame.turn)
	})
	score := canvas.NewText("0", textColor)
	thisGame.players[plNum].winsText = score
	scoreC := container.NewCenter(score)
	c := container.New(layout.NewVBoxLayout(),
		icn,
		et,
		btn,
		scoreC)
	return c
}

func getExitBtn(a fyne.App, w fyne.Window) *fyne.Container {
	restartBtn := widget.NewButton("Restart", func() {
		restartGame()
	})
	exitBtn := widget.NewButton("Exit", func() {
		dialog.ShowConfirm("Exit", "Are you sure you want to exit?", func(b bool) {
			if b {
				w.SetContent(GetStartContent(a, w))
			}
		}, w)
	})
	c := container.NewHBox(restartBtn, exitBtn)
	cCenter := container.NewCenter(c)
	return cCenter
}
