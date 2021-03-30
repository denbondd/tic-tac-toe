package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

type game struct {
	btns    [][]*customIcon
	players [2]*player
	turn    bool

	turnText *canvas.Text
}

type player struct {
	name string
	grid [][]bool
	wins int

	winsText *canvas.Text
}

func (pl *player) Win() {
	pl.wins++
	pl.winsText.Text = strconv.Itoa(pl.wins)
	pl.winsText.Refresh()
}

type fields struct {
	dark     bool
	language int
}

type customIcon struct {
	widget.Icon
	Function  func(ci *customIcon)
	size      fyne.Size
	clickable bool
}

func newCustomIcon(resource fyne.Resource, size fyne.Size, function func(ci *customIcon)) *customIcon {
	ci := &customIcon{
		Function:  function,
		size:      size,
		clickable: true,
	}
	ci.SetResource(resource)
	ci.ExtendBaseWidget(ci)
	return ci
}

func newCustomIconWithoutFunc(resource fyne.Resource, size fyne.Size) *customIcon {
	ci := &customIcon{
		size:      size,
		clickable: true,
	}
	ci.SetResource(resource)
	ci.ExtendBaseWidget(ci)
	return ci
}

func (ci *customIcon) Tapped(_ *fyne.PointEvent) {
	if ci.Function == nil {
		return
	}
	if !ci.clickable {
		return
	}
	ci.Function(ci)
}

func (ci *customIcon) MinSize() fyne.Size {
	return ci.size
}
