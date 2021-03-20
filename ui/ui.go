package ui

import (
	"../assets"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

func GetContent() (c *fyne.Container) {
	c = container.New(layout.NewVBoxLayout(),
		container.NewHBox(getSettingsBtn()),
		layout.NewSpacer(),
		getPlayButton(),
		layout.NewSpacer())
	return
}

func getSettingsBtn() fyne.CanvasObject {
	s := newCustomIcon(assets.SettingsIco, fyne.NewSize(48, 48), func() {
		println("settings")
	})
	return s
}

func getPlayButton() fyne.CanvasObject {
	playButton := widget.NewButton("", func() {
		println("play")
	})
	text := canvas.NewText("Play", color.Black)
	text.TextSize = 36
	text.TextStyle = fyne.TextStyle{Bold: true, Italic: true}
	return container.NewMax(playButton, container.NewCenter(text))
}
