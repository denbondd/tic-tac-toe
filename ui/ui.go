package ui

import (
	"../assets"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func GetContent() (c *fyne.Container) {
	settingsButton := widget.NewButtonWithIcon("", assets.SettingsIco, func() {
		println("success")
	})
	c = container.New(layout.NewVBoxLayout(), settingsButton)
	return
}
