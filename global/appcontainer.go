package global

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var appContainer *AppContainer

type AppContainer struct {
	App    fyne.App
	Window fyne.Window
}

func InitAppContainer() AppContainer {
	appContainer = &AppContainer{
		App: initApp(),
	}

	appContainer.Window = window(appContainer.App)

	return *appContainer
}

func initApp() fyne.App {
	a := app.NewWithID("goracy.kubek")
	// a.Settings().SetTheme(theme.NewDarkTheme())

	return a
}

func window(app fyne.App) fyne.Window {
	window := app.NewWindow("Converter")
	window.SetPadded(false)
	window.Resize(fyne.NewSize(600, 400))

	return window
}

func GetWindow() fyne.Window {
	return appContainer.Window
}

func GetApp() fyne.App {
	return appContainer.App
}
