package page

import (
	"converter/view/clayout"
	"converter/view/cwidget"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewHomePage(window fyne.Window) fyne.CanvasObject {
	return container.New(
		&clayout.TopLeftLayout{},
		settingsButton(window),
		dragAndDrop(window),
	)
}

func dragAndDrop(window fyne.Window) *cwidget.DragAndDropWidget {
	return cwidget.NewDragAndDrop(window, func(uris []fyne.URI) {
		for _, value := range uris {
			inserter(value.Path())
		}
	})
}

func settingsButton(window fyne.Window) fyne.CanvasObject {
	return widget.NewButton("Settings", func() {
		window.SetContent(NewSettingsPage(window))
	})
}

func inserter(url string) {
	print(url)
}
