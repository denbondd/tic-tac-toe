package main

import (
	"./ui"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := a.NewWindow("tic-tac-toe")
	w.Resize(fyne.NewSize(768, 512))
	w.SetFixedSize(true)

	ui.SetContentFields(w)
	w.SetContent(ui.GetStartContent())

	w.ShowAndRun()
}
