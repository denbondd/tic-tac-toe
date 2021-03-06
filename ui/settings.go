package ui

import (
	"github.com/denbondd/tic-tac-toe/util"
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

func getSettingsContent() (c *fyne.Container) {
	themeIsDark := textColor == currApp.Settings().Theme().Color("foreground", 0)
	currentFields = &fields{
		dark: themeIsDark,
		language: lang.GetLangNum(),
	}
	c = container.New(
		layout.NewVBoxLayout(),
		container.NewHBox(getSettingsTitle()),
		layout.NewSpacer(),
		getSettingsList(themeIsDark),
		layout.NewSpacer(),
		getCanSaveBtns(themeIsDark))
	return
}

func getSettingsTitle() fyne.CanvasObject {
	t := canvas.NewText(lang.Settings(), textColor)
	t.TextSize = 32
	return t
}

func getSettingsList(themeIsDark bool) fyne.CanvasObject {
	dt := widget.NewCheck(lang.DarkTheme(), func(b bool) {
		currentFields.dark = b
	})
	dt.Checked = themeIsDark

	langTitle := canvas.NewText(lang.Language(), textColor)
	langSelect := widget.NewSelect(languages, func(s string) {
		for idx, elem := range languages {
			if elem == s {
				currentFields.language = idx
				return
			}
		}
		panic(errors.New("no such language"))
	})
	langSelect.Selected = languages[currentFields.language]
	lang := container.New(layout.NewHBoxLayout(), langTitle, langSelect)

	setList := container.NewVBox(dt, lang)
	return setList
}

func getCanSaveBtns(themeIsDark bool) fyne.CanvasObject {
	cancelBtn := widget.NewButton(lang.Cancel(), func() {
		currWindow.SetContent(GetStartContent())
	})
	saveBtn := widget.NewButton(lang.Save(), func() {
		if currentFields.dark && !themeIsDark {
			currApp.Settings().SetTheme(theme.DarkTheme())
		} else if !currentFields.dark && themeIsDark {
			currApp.Settings().SetTheme(theme.LightTheme())
		}
		switch currentFields.language {
		case 0:
			lang = util.English{}
		case 1:
			lang = util.Russian{}
		case 2:
			lang = util.Ukrainian{}
		}
		currWindow.SetContent(GetStartContent())
	})
	btns := container.New(layout.NewHBoxLayout(),
		layout.NewSpacer(),
		cancelBtn,
		saveBtn)
	return btns
}
