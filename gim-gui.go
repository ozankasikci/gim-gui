package gimgui

import (
	"errors"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	fyneDialog "fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	gim "github.com/ozankasikci/go-image-merge"
	"github.com/sqweek/dialog"
	"image/jpeg"
	"os"
	"strconv"
)

const (
	DefaultGridCountX = 2
	DefaultGridCountY = 1
	DefaultGridSize = 75
)

type Grid struct {
	ImageFilePath string
	Image *canvas.Image
	Index int
}

type Gim struct {
	Window        *fyne.Window
	ImagesSection *fyne.Container
	GridCountX    int
	GridCountY    int
	Grids         []*Grid
}

func (t *Gim) generateGrids() {
	t.Grids = nil
	for i := 0; i < t.GridCountX* t.GridCountY; i++ {
		grid := &Grid{
			Index: i,
		}
		t.Grids = append(t.Grids, grid)
	}
}

func NewGim(w *fyne.Window) *Gim {
	gim := &Gim{GridCountX: DefaultGridCountX, GridCountY: DefaultGridCountY, Window: w}
	gim.generateGrids()

	return gim
}

func Start() {
	app := app.New()
	w := app.NewWindow("GIM")
	w.Resize(fyne.Size{500, 500})
	gim := NewGim(&w)

	w.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewVBoxLayout(),
			gim.GridOptionsSection(),
			//gim.ImagesSection(),
			gim.GridImagesSection(),
			gim.ActionsSection(),
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

			img := canvas.NewImageFromFile(imgPath)
			img.Resize(fyne.NewSize(DefaultGridSize, DefaultGridSize))
			t.Grids[index].Image = img
			t.Grids[index].ImageFilePath = imgPath
			t.generateCanvasObjectsFromGrids()
		}
	}

	t.ImagesSection.Objects = nil
	for i := 0; i < t.GridCountY; i++ {
		row := fyne.NewContainerWithLayout(
			layout.NewFixedGridLayout(fyne.NewSize(DefaultGridSize, DefaultGridSize)),
		)
		for j := 0; j < t.GridCountX; j++ {
			var obj fyne.CanvasObject
			index := i * DefaultGridCountX + j
			grid := t.Grids[index]
			obj = widget.NewButton("", imageSelectFunc(grid.Index))

			if grid.Image != nil {
				obj = grid.Image
			}
			row.AddObject(obj)
		}

		t.ImagesSection.AddObject(row)
	}
}

func (t *Gim) GridImagesSection() *fyne.Container {

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

func (t *Gim) GridSize() *widget.Box {
	onSizeChange := func(enum rune) func(string) {
		return func(s string) {
			i, err := strconv.Atoi(s)
			if s != "" && err != nil {
				fyneDialog.ShowError(errors.New("Please enter a digit"), *t.Window)
			}

			if enum == 'x' {
				if s == "" {
					i = DefaultGridCountX
				}
				t.GridCountX = i
			} else {
				if s == "" {
					i = DefaultGridCountY
				}
				t.GridCountY = i
			}

			t.generateGrids()
			t.generateCanvasObjectsFromGrids()
		}
	}

	sizeEntryX := widget.NewEntry()
	sizeEntryX.OnChanged = onSizeChange('x')
	sizeEntryX.SetPlaceHolder(strconv.Itoa(DefaultGridCountX))

	sizeEntryY := widget.NewEntry()
	sizeEntryY.OnChanged = onSizeChange('y')
	sizeEntryY.SetPlaceHolder(strconv.Itoa(DefaultGridCountY))

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

func (t *Gim) merge() {
	var gimGrids []*gim.Grid
	for _, grid := range t.Grids {
		gimGrids = append(gimGrids, &gim.Grid{
			ImageFilePath: grid.ImageFilePath,
		})
	}

	mergeFilePath, _ := dialog.File().Title("Merge Image Path").Save()

	if mergeFilePath == "" {
		(*t.Window).RequestFocus()
		return
	}

	rgba, _ := gim.New(gimGrids, t.GridCountX, t.GridCountY).Merge()
	file, _ := os.Create(mergeFilePath)
	jpeg.Encode(file, rgba, &jpeg.Options{Quality: 80})
	(*t.Window).RequestFocus()
}

func (t *Gim) ActionsSection() fyne.CanvasObject {
	return fyne.NewContainerWithLayout(
		layout.NewVBoxLayout(),
		widget.NewGroup("Actions",
			widget.NewButton("Merge",t.merge),
		),
	)
}
