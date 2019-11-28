package gimgui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

//type Gim struct {
//	GridSizeX int
//	GridSizeY int
//}

func Start() {
	app := app.New()
	w := app.NewWindow("Hello")
	w.SetContent(fyne.NewContainerWithLayout(
		layout.NewGridLayout(1),
		GridOptionsSection(),
	))

	w.ShowAndRun()
}

func GridSize() *widget.Box {
	sizeEntryX := widget.NewEntry()
	sizeEntryX.SetPlaceHolder("2")
	sizeEntryY := widget.NewEntry()
	sizeEntryY.SetPlaceHolder("1")

	return widget.NewVBox(widget.NewHBox(
		widget.NewLabel("Horizontal Size:"),
		sizeEntryX,
		widget.NewLabel("Vertical Size:"),
		sizeEntryY,
	))
}

func GridOptionsSection() *fyne.Container {
	return fyne.NewContainerWithLayout(layout.NewGridLayout(1), widget.NewGroup("Grid Options", GridSize()))
}
