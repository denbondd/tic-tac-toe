package ui

import (
	"../util"
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

func GetStartContent(a fyne.App, w fyne.Window) (c *fyne.Container) {
	textColor = theme.ForegroundColor()
	c = container.New(
		layout.NewVBoxLayout(),
		container.NewHBox(getSettingsBtn(a, w)),
		getStartTitle(),
		layout.NewSpacer(),
		getPlayButton(a, w),
		layout.NewSpacer(),
		getAppInfo())
	return
}

func getSettingsBtn(a fyne.App, w fyne.Window) fyne.CanvasObject {
	s := newCustomIcon(theme.SettingsIcon(), fyne.NewSize(48, 48), func(ci *customIcon) {
		w.SetContent(GetSettingsContent(a, w))
	})
	return s
}

func getPlayButton(a fyne.App, w fyne.Window) fyne.CanvasObject {
	playButton := widget.NewButton("", func() {
		w.SetContent(getGameContent(a, w, 3))
	})
	text := canvas.NewText("Play", textColor)
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
