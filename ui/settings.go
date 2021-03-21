package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"image/color"
)

func GetSettingsContent() (c *fyne.Container) {
	c = container.New(
		layout.NewHBoxLayout(),
		container.NewVBox(getSettingsTitle()))
	return
}

func getSettingsTitle() fyne.CanvasObject {
	t := canvas.NewText("Settings", color.Black)
	t.TextSize = 32
	return t
}
