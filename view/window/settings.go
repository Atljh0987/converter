package window

import (
	"converter/storage"
	"converter/view/layout"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func CreateSettingsPage(window fyne.Window) fyne.CanvasObject {
	return container.New(
		&layout.TopLeftLayout{},
		NewLabel(),
		BackButton(window),
		SelectWidget(),
	)
}

func NewLabel() *widget.Label {
	currWidget := widget.NewLabel("Settings")

	currWidget.Move(fyne.NewPos(50, 50))

	return currWidget
}

func BackButton(window fyne.Window) fyne.CanvasObject {
	return widget.NewButton("Back", func() {
		window.SetContent(CreateHomePage(window))
	})
}

func SelectWidget() *widget.Select {
	options := storage.VideoExtensionsStrings()

	selectWidget := widget.NewSelect(options, func(selected string) {
		println("Выбрано:", selected)
	})

	return selectWidget
}
