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

func getGameContent(a fyne.App, w fyne.Window, size int) (c *fyne.Container) {
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
	btn := newCustomIcon(theme.ViewFullScreenIcon(), fyne.NewSize(72, 72), func(ci *customIcon) {
		gridClick(ci, h, w)
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
	scoreC := container.NewCenter(score)
	c := container.New(layout.NewVBoxLayout(),
		icn,
		et,
		btn,
		scoreC)
	return c
}

func getExitBtn(a fyne.App, w fyne.Window) *fyne.Container {
	btn := widget.NewButton("Exit", func() {
		dialog.ShowConfirm("Exit", "Are you sure you want to exit?", func(b bool) {
			if b {
				w.SetContent(GetStartContent(a, w))
			}
		}, w)
	})
	c := container.NewCenter(btn)
	return c
}
