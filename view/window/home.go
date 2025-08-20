package window

import (
	"converter/view/layout"
	cwidget "converter/view/widget"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func MainWindow(app fyne.App) fyne.Window {
	window := app.NewWindow("Converter")
	window.SetPadded(false)
	window.Resize(fyne.NewSize(600, 400))
	window.SetContent(CreateHomePage(window))

	return window
}

func DragAndDrop(window fyne.Window) *cwidget.DragAndDropWidget {
	return cwidget.NewDragAndDrop(window, func(uris []fyne.URI) {
		for _, value := range uris {
			inserter(value.Path())
		}
	})
}

func SettingsButton(window fyne.Window) fyne.CanvasObject {
	return widget.NewButton("Settings", func() {
		window.SetContent(CreateSettingsPage(window))
	})
}

func CreateHomePage(window fyne.Window) fyne.CanvasObject {
	return container.New(
		&layout.TopLeftLayout{},
		SettingsButton(window),
		DragAndDrop(window),
	)
}

func inserter(url string) {
	print(url)
}
