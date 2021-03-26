package ui

import (
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var languages = []string{"English", "Russian", "Ukrainian"}
var currentFields *fields

func GetSettingsContent(a fyne.App, w fyne.Window) (c *fyne.Container) {
	currentFields = new(fields)
	themeIsDark := textColor == a.Settings().Theme().Color("foreground", 0)
	c = container.New(
		layout.NewVBoxLayout(),
		container.NewHBox(getSettingsTitle()),
		layout.NewSpacer(),
		getSettingsList(themeIsDark),
		layout.NewSpacer(),
		getCanSaveBtns(themeIsDark, a, w))
	return
}

func getSettingsTitle() fyne.CanvasObject {
	t := canvas.NewText("Settings", textColor)
	t.TextSize = 32
	return t
}

func getSettingsList(themeIsDark bool) fyne.CanvasObject {
	dt := widget.NewCheck("Dark Theme", func(b bool) {
		currentFields.dark = b
	})
	dt.Checked = themeIsDark

	langTitle := canvas.NewText("Language", textColor)
	langSelect := widget.NewSelect(languages, func(s string) {
		for idx, elem := range languages {
			if elem == s {
				currentFields.language = idx
				return
			}
		}
		panic(errors.New("no such language"))
	})
	langSelect.Selected = languages[0]
	lang := container.New(layout.NewHBoxLayout(), langTitle, langSelect)

	setList := container.NewVBox(dt, lang)
	return setList
}

func getCanSaveBtns(themeIsDark bool, a fyne.App, w fyne.Window) fyne.CanvasObject {
	cancelBtn := widget.NewButton("Cancel", func() {
		w.SetContent(GetStartContent(a, w))
	})
	saveBtn := widget.NewButton("Save", func() {
		if currentFields.dark && !themeIsDark {
			a.Settings().SetTheme(theme.DarkTheme())
		} else if !currentFields.dark && themeIsDark {
			a.Settings().SetTheme(theme.LightTheme())
		}
		//TODO add language change
		w.SetContent(GetStartContent(a, w))
	})
	btns := container.New(layout.NewHBoxLayout(),
		layout.NewSpacer(),
		cancelBtn,
		saveBtn)
	return btns
}

type fields struct {
	dark     bool
	language int
}
