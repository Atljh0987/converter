package layout

import (
	"converter/view/widget"

	"fyne.io/fyne/v2"
)

var padding = float32(20)

type TopLeftLayout struct{}

func (l *TopLeftLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	for _, obj := range objects {
		obj.Resize(obj.MinSize())

		if _, isDnD := obj.(*widget.DragAndDropWidget); isDnD {
			obj.Move(fyne.NewPos(
				size.Width-obj.MinSize().Width-padding,
				size.Height-obj.MinSize().Height-padding,
			))

			continue
		}

		obj.Move(fyne.NewPos(0, 0))
	}
}

func (l *TopLeftLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(100, 100)
}
