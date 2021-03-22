package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

func GetSettingsContent(w fyne.Window) (c *fyne.Container) {
	c = container.New(
		layout.NewVBoxLayout(),
		container.NewHBox(getSettingsTitle()),
		layout.NewSpacer(),
		getSettingsList(),
		layout.NewSpacer(),
		getCanSaveBtns(w))
	return
}

func getSettingsTitle() fyne.CanvasObject {
	t := canvas.NewText("Settings", color.Black)
	t.TextSize = 32
	return t
}

func getSettingsList() fyne.CanvasObject {
	dt := widget.NewCheck("Dark Theme", func(b bool) {
		println("darktheme", b)
	})
	langTitle := canvas.NewText("Language", color.Black)
	langSelect := widget.NewSelect([]string{"English", "Russian", "Ukrainian"}, func(s string) {
		println("lang", s)
	})
	langSelect.Selected = "English"
	lang := container.New(layout.NewHBoxLayout(), langTitle, langSelect)
	setList := container.NewVBox(
		dt,
		lang)
	return setList
}

func getCanSaveBtns(w fyne.Window) fyne.CanvasObject {
	cancelBtn := widget.NewButton("Cancel", func() {
		w.SetContent(GetStartContent(w))
	})
	saveBtn := widget.NewButton("Save", func() {
		println("save")
	})
	btns := container.New(
		layout.NewHBoxLayout(),
		layout.NewSpacer(),
		cancelBtn,
		saveBtn)
	return btns
}
