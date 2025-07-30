package widget

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

var (
	_             desktop.Hoverable  = (*DragAndDropWidget)(nil)
	_             desktop.Cursorable = (*DragAndDropWidget)(nil)
	hoverDuration                    = time.Millisecond * 300
	hoverColor                       = color.NRGBA{R: 100, G: 100, B: 100, A: 150}
	label                            = widget.NewLabel("Drop files here or click on me!")
)

func (m *DragAndDropWidget) Cursor() desktop.Cursor {
	return desktop.PointerCursor
}

func (w *DragAndDropWidget) MouseMoved(e *desktop.MouseEvent) {}

func changeColor(w *DragAndDropWidget, from color.Color, to color.Color) {
	canvas.NewColorRGBAAnimation(
		from,
		to,
		hoverDuration,
		func(c color.Color) {
			w.dropZone.FillColor = c
			w.dropZone.Refresh()
		},
	).Start()
}

func (w *DragAndDropWidget) MouseIn(*desktop.MouseEvent) {
	changeColor(w, color.Transparent, hoverColor)
}

func (w *DragAndDropWidget) MouseOut() {
	changeColor(w, hoverColor, color.Transparent)
}

type DragAndDropWidget struct {
	widget.BaseWidget
	parentWindow fyne.Window
	label        *widget.Label
	dropZone     *canvas.Rectangle
	onDrop       func([]fyne.URI)
}

type dragAndDropRenderer struct {
	widget  *DragAndDropWidget
	objects []fyne.CanvasObject
}

func newRectangle() *canvas.Rectangle {
	r := canvas.NewRectangle(color.Transparent)
	r.StrokeColor = color.NRGBA{R: 80, G: 80, B: 80, A: 180}
	r.CornerRadius = 10
	r.StrokeWidth = 1

	return r
}

func NewDragAndDrop(parentWindow fyne.Window, onDrop func([]fyne.URI)) *DragAndDropWidget {
	d := &DragAndDropWidget{
		parentWindow: parentWindow,
		label:        label,
		dropZone:     newRectangle(),
		onDrop:       onDrop,
	}
	d.ExtendBaseWidget(d)

	return d
}

func (w *DragAndDropWidget) CreateRenderer() fyne.WidgetRenderer {
	return &dragAndDropRenderer{
		widget:  w,
		objects: []fyne.CanvasObject{w.dropZone, w.label},
	}
}

func (w *DragAndDropWidget) Tapped(*fyne.PointEvent) {
	dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
		if err != nil || reader == nil || w.parentWindow == nil {
			return
		}

		w.onDrop([]fyne.URI{reader.URI()})
	}, w.parentWindow)
}

func (w *dragAndDropRenderer) Destroy() {}

func (w *dragAndDropRenderer) Layout(size fyne.Size) {
	dropzone := w.widget.dropZone
	label := w.widget.label

	dw := dropzone.Size().Width
	dh := dropzone.Size().Height

	lw := label.MinSize().Width
	lh := label.MinSize().Height

	label.Move(fyne.NewPos(
		(dw-lw)/2,
		(dh-lh)/2,
	))

	w.widget.parentWindow.SetOnDropped(func(p fyne.Position, uris []fyne.URI) {
		wx := w.widget.Position().X
		wy := w.widget.Position().Y

		x := p.X >= wx && p.X <= wx+dw
		y := p.Y >= wy && p.Y <= wy+dw

		if x && y {
			w.widget.onDrop(uris)
		}
	})

	dropzone.Resize(size)
}

func (w *dragAndDropRenderer) MinSize() fyne.Size {
	return fyne.NewSize(300, 100)
}

func (w *dragAndDropRenderer) Objects() []fyne.CanvasObject {
	return w.objects
}

func (w *dragAndDropRenderer) Refresh() {
	w.widget.dropZone.Refresh()
	w.widget.label.Refresh()
}
