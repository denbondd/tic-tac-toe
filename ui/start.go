package ui

import (
	"../assets"
	"../util"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"net/url"
)

func GetStartContent(w fyne.Window) (c *fyne.Container) {
	c = container.New(
		layout.NewVBoxLayout(),
		container.NewHBox(getSettingsBtn(w)),
		getStartTitle(),
		layout.NewSpacer(),
		getPlayButton(),
		layout.NewSpacer(),
		getAppInfo())
	return
}

func getSettingsBtn(w fyne.Window) fyne.CanvasObject {
	s := newCustomIcon(assets.SettingsIco, fyne.NewSize(48, 48), func() {
		w.SetContent(GetSettingsContent(w))
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

func getStartTitle() fyne.CanvasObject {
	t := canvas.NewText("Tic-Tac-Toe", color.Black)
	t.TextStyle = fyne.TextStyle{Italic: true, Bold: true}
	t.TextSize = 64
	return container.NewCenter(t)
}

func getAppInfo() fyne.CanvasObject {
	link, err := url.Parse("https://github.com/denbondd/tic-tac-toe")
	util.CheckError(err)
	return widget.NewHyperlinkWithStyle("GitHub", link, fyne.TextAlignCenter, fyne.TextStyle{Italic: true})
}
