package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type customIcon struct {
	widget.Icon
	Function func()
	size     fyne.Size
}

func newCustomIcon(resource *fyne.StaticResource, size fyne.Size, function func()) *customIcon {
	ci := &customIcon{
		Function: function,
		size:     size,
	}
	ci.SetResource(resource)
	ci.ExtendBaseWidget(ci)
	return ci
}

func (ci *customIcon) Tapped(_ *fyne.PointEvent) {
	ci.Function()
}

func (ci *customIcon) MinSize() fyne.Size {
	return ci.size
}
