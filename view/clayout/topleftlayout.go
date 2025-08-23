package clayout

import (
	"converter/view/cwidget"

	"fyne.io/fyne/v2"
)

var padding = float32(20)

type TopLeftLayout struct{}

func (l *TopLeftLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	for _, obj := range objects {
		if obj.Size().IsZero() {
			obj.Resize(obj.MinSize())
		}

		if _, isDnD := obj.(*cwidget.DragAndDropWidget); isDnD {
			obj.Move(fyne.NewPos(
				size.Width-obj.MinSize().Width-padding,
				size.Height-obj.MinSize().Height-padding,
			))

			continue
		}

		if obj.Position().IsZero() {
			obj.Move(fyne.NewPos(0, 0))
		}
	}
}

func (l *TopLeftLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(100, 100)
}
