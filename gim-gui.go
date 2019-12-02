package gimgui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	gim "github.com/ozankasikci/go-image-merge"
	"github.com/sqweek/dialog"
)

type Grid struct {
	ImageFilePath string
	Image         *canvas.Image
	Index         int
}

type Gim struct {
	gim             *gim.MergeImage
	Window          *fyne.Window
	ImagesSection   *fyne.Container
	GridColumnCount int
	GridRowCount    int
	GridSizeX       int
	GridSizeY       int
	Grids           []*Grid
}

func (t *Gim) generateGrids() {
	t.Grids = nil
	for i := 0; i < t.GridColumnCount*t.GridRowCount; i++ {
		grid := &Grid{
			Index: i,
		}
		t.Grids = append(t.Grids, grid)
	}
}

func NewGim(w *fyne.Window) *Gim {
	gimInstance := gim.New(nil, DefaultGridCountX, DefaultGridCountY)

	gim := &Gim{
		GridColumnCount: DefaultGridCountX,
		GridRowCount:    DefaultGridCountY,
		GridSizeX:       DefaultGridSizeX,
		GridSizeY:       DefaultGridSizeY,
		Window:          w,
		gim:             gimInstance,
	}
	gim.generateGrids()

	return gim
}

func Start() {
	app := app.New()
	w := app.NewWindow("GIM")
	w.Resize(fyne.Size{800, 600})
	gim := NewGim(&w)

	w.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewVBoxLayout(),
			gim.gridOptionsSection(),
			gim.gridImagesSection(),
			layout.NewSpacer(),
			gim.actionsSection(),
		),
	)

	w.ShowAndRun()
}

func (t *Gim) generateCanvasObjectsFromGrids() {
	imageSelectFunc := func(index int) func() {
		return func() {
			imgPath, _ := dialog.File().Title("Select an image file").Load()
			if imgPath == "" {
				(*t.Window).RequestFocus()
				return
			}

			img, _ := t.gim.ReadImageFile(imgPath)
			t.Grids[index].Image = canvas.NewImageFromImage(img)
			t.Grids[index].ImageFilePath = imgPath
			(*t.Window).RequestFocus()
			t.generateCanvasObjectsFromGrids()
		}
	}

	t.ImagesSection.Objects = nil

	for i := 0; i < t.GridRowCount; i++ {
		row := fyne.NewContainerWithLayout(
			layout.NewFixedGridLayout(fyne.NewSize(t.GridSizeX, t.GridSizeY)),
		)

		for j := 0; j < t.GridColumnCount; j++ {
			var obj fyne.CanvasObject
			index := i*t.GridColumnCount + j
			grid := t.Grids[index]
			obj = widget.NewButton("", imageSelectFunc(grid.Index))

			if grid.Image != nil {
				obj = imageBoxResized(grid.Image.Image, t.GridSizeX, t.GridSizeY)
			}
			row.AddObject(obj)
		}

		t.ImagesSection.AddObject(row)
	}
}
