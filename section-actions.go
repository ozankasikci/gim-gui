package gimgui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	fyneDialog "fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	gim "github.com/ozankasikci/go-image-merge"
	"github.com/sqweek/dialog"
	"image"
	"image/jpeg"
	"os"
)

func (t *Gim) merge() {
	mergeFilePath, _ := dialog.File().Title("Merge Image Path").Filter("", "jpg", "png").Save()

	if mergeFilePath == "" {
		(*t.Window).RequestFocus()
		return
	}

	if t.GridSizeX != 0 && t.GridSizeY != 0 {
		gim.OptGridSize(t.GridSizeX, t.GridSizeY)(t.gim)
	}

	rgba, err := t.generateRGBA()
	if err != nil {
		println(err.Error())
	}
	file, _ := os.Create(mergeFilePath)
	jpeg.Encode(file, rgba, &jpeg.Options{Quality: 80})
	(*t.Window).RequestFocus()
}

func (t *Gim) generateRGBA() (*image.RGBA, error) {
	var gimGrids []*gim.Grid
	for _, grid := range t.Grids {
		gimGrids = append(gimGrids, &gim.Grid{
			ImageFilePath: grid.ImageFilePath,
		})
	}

	t.gim.Grids = gimGrids
	return t.gim.Merge()
}

func (t *Gim) preview() {
	rgba, err := t.generateRGBA()
	if err != nil {
		fyneDialog.ShowError(err, *t.Window)
		return
	}

	previewImage := canvas.NewImageFromImage(rgba)

	ratio := 1.0
	if rgba.Bounds().Dx() > PreviewImageMaxSize || rgba.Bounds().Dy() > PreviewImageMaxSize {
		max := rgba.Bounds().Dx()
		if max < rgba.Bounds().Dy() {
			max = rgba.Bounds().Dy()
		}

		ratio = float64(max) / PreviewImageMaxSize
	}

	previewImage.SetMinSize(fyne.NewSize(
		int(float64(rgba.Bounds().Dx())/ratio),
		int(float64(rgba.Bounds().Dy())/ratio),
	))

	container := widget.NewHBox(layout.NewSpacer(), previewImage, layout.NewSpacer())
	box := widget.NewVBox(
		layout.NewSpacer(),
		container,
		layout.NewSpacer(),
	)
	fyneDialog.ShowCustom("Preview", "OK", box, *t.Window)
}

func (t *Gim) actionsSection() fyne.CanvasObject {
	buttons := fyne.NewContainerWithLayout(
		layout.NewFixedGridLayout(fyne.NewSize(150, 35)),
		widget.NewButton(TitleButtonMerge, t.merge),
		widget.NewButton(TitleButtonPreview, t.preview),
	)

	return fyne.NewContainerWithLayout(
		layout.NewVBoxLayout(),
		widget.NewGroup(TitleSectionActions,
			buttons,
		),
	)
}
