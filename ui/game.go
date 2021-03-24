package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

var thisGame *game

func getGameContent(w fyne.Window, size int) (c *fyne.Container) {
	thisGame = &game{btns: make([]*customIcon, size*size)}
	c = container.New(layout.NewVBoxLayout(),
		getTurnText(),
		layout.NewSpacer(),
		getMainGame(size),
		layout.NewSpacer(),
		getExitBtn(w))
	return
}

func getTurnText() fyne.CanvasObject {
	t := canvas.NewText("Player1's turn", color.Black)
	thisGame.turnText = t
	thisGame.turnText.TextSize = 32
	c := container.NewCenter(t)
	return c
}

func getMainGame(size int) *fyne.Container {
	c := container.New(layout.NewGridLayoutWithColumns(3),
		getPlayerContent(),
		getGameGrid(size),
		getPlayerContent())
	c = container.NewCenter(c)
	return c
}

func getGameGrid(size int) *fyne.Container {
	g := container.NewGridWithColumns(size)
	for i := 0; i < size*size; i++ {
		btn := newCustomIcon(theme.CheckButtonIcon(), fyne.NewSize(72, 72), func(ci *customIcon) {
			ci.SetResource(theme.CancelIcon())
		})
		thisGame.btns[i] = btn
		g.Add(btn)
	}
	return g
}

func getPlayerContent() *fyne.Container {
	icn := newCustomIconWithoutFunc(theme.CancelIcon(), fyne.NewSize(64, 64))
	et := widget.NewEntry()
	et.Text = "Player1"
	btn := widget.NewButton("Save name", func() {

	})
	score := canvas.NewText("0", color.Black)
	scoreC := container.NewCenter(score)
	c := container.New(layout.NewVBoxLayout(),
		icn,
		et,
		btn,
		scoreC)
	return c
}

func getExitBtn(w fyne.Window) *fyne.Container {
	btn := widget.NewButton("Exit", func() {
		dialog.ShowConfirm("Exit", "Are you sure you want to exit?", func(b bool) {
			if b {
				w.SetContent(GetStartContent(w))
			}
		}, w)
	})
	c := container.NewCenter(btn)
	return c
}

type game struct {
	btns    []*customIcon
	players [2]string

	turnText *canvas.Text
}
