package gimgui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/sqweek/dialog"
)

type Grid struct {
	Image *canvas.Image
	Index int
}

type Gim struct {
	GridSizeX int
	GridSizeY int
	Grids     []*Grid
}

func NewGim() *Gim {
	gim := &Gim{GridSizeX: 2, GridSizeY: 1}
	for i := 0; i < gim.GridSizeX * gim.GridSizeY; i++ {
		grid := &Grid{
			Index: i,
		}
		gim.Grids = append(gim.Grids, grid)
	}

	return gim
}

func Start() {
	app := app.New()
	w := app.NewWindow("GIM")
	w.Resize(fyne.Size{800, 800})
	gim := NewGim()

	w.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewVBoxLayout(),
			gim.GridOptionsSection(),
			//gim.ImagesSection(),
			gim.GridImagesSection(),
		),
	)

	w.ShowAndRun()
}

//func (t *Gim) ImagesSection() *fyne.Container {
//	imagesBox := fyne.NewContainerWithLayout(
//		layout.NewFixedGridLayout(fyne.NewSize(75, 75)),
//	)
//
//	addImageButton := widget.NewButton("Add Image", func() {
//		imgPath, _ := dialog.File().Title("Select an image file").Load()
//		img := canvas.NewImageFromFile(imgPath)
//		img.Resize(fyne.NewSize(75, 75))
//		imagesBox.AddObject(img)
//	})
//
//	return fyne.NewContainerWithLayout(
//		layout.NewVBoxLayout(),
//		widget.NewGroup("Images",
//			addImageButton,
//			imagesBox,
//		),
//	)
//}

func (t *Gim) generateCanvasObjectsFromGrids(container *fyne.Container) []fyne.CanvasObject {
	imageSelectFunc := func(index int) func() {
		return func() {
			imgPath, _ := dialog.File().Title("Select an image file").Load()
			img := canvas.NewImageFromFile(imgPath)
			img.Resize(fyne.NewSize(75, 75))
			t.Grids[index].Image = img
			t.generateCanvasObjectsFromGrids(container)
		}
	}

	var objs []fyne.CanvasObject
	container.Objects = objs
	for _, grid := range t.Grids {
		var obj fyne.CanvasObject
		obj = widget.NewButton("", imageSelectFunc(grid.Index))

		if grid.Image != nil {
			obj = grid.Image
		}
		container.AddObject(obj)
	}

	return objs
}

func (t *Gim) GridImagesSection() *fyne.Container {

	images := fyne.NewContainerWithLayout(
		layout.NewFixedGridLayout(fyne.NewSize(75, 75)),
	)

	t.generateCanvasObjectsFromGrids(images)

	return fyne.NewContainerWithLayout(
		layout.NewGridLayout(1),
		widget.NewGroup("Grids",
			images,
		),
	)
}

func (t *Gim) GridSize() *widget.Box {
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

func (t *Gim) GridOptionsSection() *fyne.Container {
	return fyne.NewContainerWithLayout(layout.NewGridLayout(1), widget.NewGroup("Grid Options", t.GridSize()))
}
