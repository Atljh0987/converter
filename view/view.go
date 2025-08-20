package view

import (
	"converter/view/theme"
	window "converter/view/window"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func initApp() fyne.App {
	a := app.NewWithID("goracy.kubek")
	a.Settings().SetTheme(theme.NewDarkTheme())

	return a
}

func RunWindow() {
	window.MainWindow(initApp()).ShowAndRun()
}
