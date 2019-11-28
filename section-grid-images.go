package gimgui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func (t *Gim) gridImagesSection() *fyne.Container {

	images := fyne.NewContainerWithLayout(
		layout.NewVBoxLayout(),
	)

	t.ImagesSection = images
	t.generateGrids()
	t.generateCanvasObjectsFromGrids()

	return fyne.NewContainerWithLayout(
		layout.NewGridLayout(1),
		widget.NewGroup("Grids",
			images,
		),
	)
}
