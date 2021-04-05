package ui

import (
	"github.com/denbondd/tic-tac-toe/util"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"net/url"
)

var textColor color.Color
var currApp fyne.App
var currWindow fyne.Window
var lang util.Lang

func SetContentFields(w fyne.Window) {
	currWindow = w
	currApp = fyne.CurrentApp()
	lang = util.English{}
}

func GetStartContent() (c *fyne.Container) {
	textColor = theme.ForegroundColor()
	c = container.New(
		layout.NewVBoxLayout(),
		container.NewHBox(getSettingsBtn()),
		getStartTitle(),
		layout.NewSpacer(),
		getPlayButton(),
		layout.NewSpacer(),
		getAppInfo())
	return
}

func getSettingsBtn() fyne.CanvasObject {
	s := newCustomIcon(theme.SettingsIcon(), fyne.NewSize(48, 48), func(ci *customIcon) {
		currWindow.SetContent(getSettingsContent())
	})
	return s
}

func getPlayButton() fyne.CanvasObject {
	playButton := widget.NewButton("", func() {
		currWindow.SetContent(getGameContent(3))
	})
	text := canvas.NewText(lang.Play(), textColor)
	text.TextSize = 36
	text.TextStyle = fyne.TextStyle{Bold: true, Italic: true}
	return container.NewMax(playButton, container.NewCenter(text))
}

func getStartTitle() fyne.CanvasObject {
	t := canvas.NewText("Tic-Tac-Toe", textColor)
	t.TextStyle = fyne.TextStyle{Italic: true, Bold: true}
	t.TextSize = 64
	return container.NewCenter(t)
}

func getAppInfo() fyne.CanvasObject {
	link, err := url.Parse("https://github.com/denbondd/tic-tac-toe")
	util.CheckError(err)
	return widget.NewHyperlinkWithStyle("GitHub", link, fyne.TextAlignCenter, fyne.TextStyle{Italic: true})
}
