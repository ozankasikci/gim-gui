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

func imageBox(img *canvas.Image) *widget.Box {
	container := widget.NewHBox(layout.NewSpacer(), img, layout.NewSpacer())
	return widget.NewVBox(
		layout.NewSpacer(),
		container,
		layout.NewSpacer(),
	)
}

func imageBoxResized(rgba image.Image, maxSizeX, maxSizeY int) *widget.Box {
	previewImage := canvas.NewImageFromImage(rgba)

	ratio := 1.0
	if rgba.Bounds().Dx() > maxSizeX || rgba.Bounds().Dy() > maxSizeY {
		max := rgba.Bounds().Dx()
		if max < rgba.Bounds().Dy() {
			max = rgba.Bounds().Dy()
		}

		maxDimensionSize := maxSizeX
		if maxDimensionSize < maxSizeY {
			maxDimensionSize = maxSizeY
		}
		ratio = float64(max) / float64(maxDimensionSize)
	}

	previewImage.SetMinSize(fyne.NewSize(
		int(float64(rgba.Bounds().Dx())/ratio),
		int(float64(rgba.Bounds().Dy())/ratio),
	))

	return imageBox(previewImage)
}

func (t *Gim) preview() {
	rgba, err := t.generateRGBA()
	if err != nil {
		fyneDialog.ShowError(err, *t.Window)
		return
	}

	previewBox := imageBoxResized(rgba, PreviewImageMaxSize, PreviewImageMaxSize)
	fyneDialog.ShowCustom("Preview", "OK", previewBox, *t.Window)
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
