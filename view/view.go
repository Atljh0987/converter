package view

import (
	"converter/view/theme"
	"converter/view/widget"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func initApp() fyne.App {
	a := app.NewWithID("goracy.kubek")
	a.Settings().SetTheme(theme.NewDarkTheme())

	return a
}

func initWindow() fyne.Window {
	window := initApp().NewWindow("Converter")
	window.SetPadded(false)
	window.Resize(fyne.NewSize(600, 400))

	widgets(window)

	return window
}

func widgets(window fyne.Window) {
	draganddrop := widget.NewDragAndDrop(window, func(uris []fyne.URI) {
		for _, value := range uris {
			println(value.Path())
		}
	})

	window.SetContent(container.NewStack(
		container.NewCenter(draganddrop),
	))
}

func RunWindow() {
	initWindow().ShowAndRun()
}
