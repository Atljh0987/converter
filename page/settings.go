package page

import (
	"converter/page/settings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var settingsContainer *fyne.Container

func NewSettingsPage(window fyne.Window) fyne.CanvasObject {
	settingsContainer = container.NewStack(settings.NewLinkSettings())

	return container.NewBorder(
		container.NewHBox(backButton(window), newLabel()),
		nil,
		settingsElements(),
		nil,
		settingsContainer,
	)
}

func settingsElements() fyne.CanvasObject {
	options := []string{"Links", "Rules"}

	group := widget.NewRadioGroup(options, func(value string) {
		println("Radio set to", value)

		var newPage fyne.CanvasObject
		switch value {
		case options[0]:
			newPage = settings.NewLinkSettings()
		case options[1]:
			newPage = settings.NewRuleSettings()
		}

		settingsContainer.Objects = []fyne.CanvasObject{newPage}
		settingsContainer.Refresh()
	})

	group.SetSelected(options[0])

	return group
}

func newLabel() *widget.Label {
	currWidget := widget.NewLabel("Settings")

	currWidget.Move(fyne.NewPos(50, 0))

	return currWidget
}

func backButton(window fyne.Window) fyne.CanvasObject {
	return widget.NewButtonWithIcon("", theme.NavigateBackIcon(), func() {
		window.SetContent(NewHomePage(window))
	})
}