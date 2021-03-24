package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type customIcon struct {
	widget.Icon
	Function func(ci *customIcon)
	size     fyne.Size
}

func newCustomIcon(resource fyne.Resource, size fyne.Size, function func(ci *customIcon)) *customIcon {
	ci := &customIcon{
		Function: function,
		size:     size,
	}
	ci.SetResource(resource)
	ci.ExtendBaseWidget(ci)
	return ci
}

func newCustomIconWithoutFunc(resource fyne.Resource, size fyne.Size) *customIcon {
	ci := &customIcon{
		size: size,
	}
	ci.SetResource(resource)
	ci.ExtendBaseWidget(ci)
	return ci
}

func (ci *customIcon) Tapped(_ *fyne.PointEvent) {
	ci.Function(ci)
}

func (ci *customIcon) MinSize() fyne.Size {
	return ci.size
}
