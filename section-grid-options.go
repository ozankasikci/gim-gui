package gimgui

import (
	"errors"
	"fyne.io/fyne"
	fyneDialog "fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"strconv"
)

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
