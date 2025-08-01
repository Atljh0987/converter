package view

import (
	"converter/view/layout"
	"converter/view/theme"
	cwidget "converter/view/widget"

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
	draganddrop := cwidget.NewDragAndDrop(window, func(uris []fyne.URI) {
		for _, value := range uris {
			println(value.Path())
		}
	})

	// options := []string{"Вариант 1", "Вариант 2", "Вариант 3"}

	// selectWidget := widget.NewSelect(options, func(selected string) {
	// 	println("Выбрано:", selected)
	// })

	// selectWidget.Move(fyne.NewPos(50, 50))

	window.SetContent(container.New(&layout.TopLeftLayout{} /*selectWidget,*/, draganddrop))
}

func RunWindow() {
	initWindow().ShowAndRun()
}
