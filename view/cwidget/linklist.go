package cwidget

import (
	"converter/global"
	"converter/storage"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var currentStorage = storage.GetLinkStorage()
var currentUri = ""
var pathEntry *widget.Entry
var list *widget.List

func LinkList() fyne.CanvasObject {
	linkListV := container.NewBorder(
		addForm(),
		nil, nil, nil,
		listWidget(),
	)

	return linkListV
}

func addForm() fyne.CanvasObject {
	return container.NewBorder(
		nil, nil, nil,
		addButton(),
		folderSelector(),
	)
}

func folderSelector() fyne.CanvasObject {
	pathEntry = widget.NewEntry()
	pathEntry.SetPlaceHolder("Choose folder...")

	pathEntry.OnChanged = func(s string) {
		if len(s) != 0 {
			currentUri = s
		}
	}

	selectButton := widget.NewButton("Browse", func() {
		dialog.ShowFolderOpen(func(uri fyne.ListableURI, err error) {
			if err == nil && uri != nil {
				pathEntry.SetText(uri.Path())
				currentUri = uri.Path()
			}
		}, global.GetWindow())
	})

	return container.NewBorder(nil, nil, nil, selectButton, pathEntry)
}

func addButton() fyne.CanvasObject {
	text := "     Add     "
	button := widget.NewButton(text, func() {
		if len(currentUri) == 0 {
			return
		}

		currentStorage.Create(currentUri)
		currentUri = ""

		pathEntry.SetText(currentUri)
		list.Refresh()
	})

	button.Importance = widget.HighImportance

	return button
}

func newDeleteButton() *widget.Button {
	button := widget.NewButton("Удалить", nil)
	button.Importance = widget.DangerImportance

	return button
}

func listWidget() fyne.CanvasObject {
	list = widget.NewList(
		func() int {
			return currentStorage.GetCount()
		},
		func() fyne.CanvasObject {
			return container.NewHBox(
				widget.NewHyperlink("", nil),
				layout.NewSpacer(),
				newDeleteButton(),
			)
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			container := o.(*fyne.Container)
			link := container.Objects[0].(*widget.Hyperlink)
			button := container.Objects[2].(*widget.Button)

			element := currentStorage.Get(i)

			link.SetText(element)
			link.OnTapped = func() {
				println(element)
			}

			button.OnTapped = func() {
				currentStorage.Delete(i)
				list.Refresh()
			}
		})

	return list
}
