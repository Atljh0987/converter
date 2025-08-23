package settings

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewRuleSettings() fyne.CanvasObject {
	return container.NewVBox(
		widget.NewLabel("Rule settings"),
	)
}
