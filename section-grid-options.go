package gimgui

import (
	"errors"
	"fyne.io/fyne"
	fyneDialog "fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	gim "github.com/ozankasikci/go-image-merge"
	"strconv"
)

func (t *Gim) gridColumnRowOptions() *widget.Box {
	onSizeChange := func(enum rune) func(string) {
		return func(s string) {
			i, err := strconv.Atoi(s)
			if s != "" && err != nil {
				fyneDialog.ShowError(errors.New(TextEnterDigitWarning), *t.Window)
			}

			if enum == 'x' {
				if s == "" {
					i = DefaultGridCountX
				}

				if i > MaxGridCountX {
					i = MaxGridCountX
				}

				t.GridColumnCount = i
			} else {
				if s == "" {
					i = DefaultGridCountY
				}

				if i > MaxGridCountY {
					i = MaxGridCountY
				}

				t.GridRowCount = i
			}

			t.generateGrids()
			t.generateCanvasObjectsFromGrids()
		}
	}

	sizeEntryColumn := widget.NewEntry()
	sizeEntryColumn.OnChanged = onSizeChange('x')
	sizeEntryColumn.SetPlaceHolder(strconv.Itoa(DefaultGridCountX))

	sizeEntryRow := widget.NewEntry()
	sizeEntryRow.OnChanged = onSizeChange('y')
	sizeEntryRow.SetPlaceHolder(strconv.Itoa(DefaultGridCountY))

	return widget.NewVBox(widget.NewHBox(
		widget.NewLabel(TextColumnCount),
		sizeEntryColumn,
		widget.NewLabel(TextRowCount),
		sizeEntryRow,
	))
}

func (t *Gim) gridSizeOptions() *widget.Box {
	onSizeChange := func(enum rune) func(string) {
		return func(s string) {
			if s == "" {
				return
			}

			i, err := strconv.Atoi(s)
			if err != nil {
				fyneDialog.ShowError(errors.New(TextEnterDigitWarning), *t.Window)
				return
			}

			println(i)
			if enum == 'x' {
				t.GridSizeX = i
			} else {
				t.GridSizeY = i
			}

			if t.GridSizeX != 0 && t.GridSizeY != 0 {
				gim.OptGridSize(t.GridSizeX, t.GridSizeY)(t.gim)
			}
		}
	}

	sizeEntryX := widget.NewEntry()
	sizeEntryX.OnChanged = onSizeChange('x')
	sizeEntryY := widget.NewEntry()
	sizeEntryY.OnChanged = onSizeChange('y')

	return widget.NewVBox(widget.NewHBox(
		widget.NewLabel(TextGridSizeX),
		sizeEntryX,
		widget.NewLabel(TextGridSizeY),
		sizeEntryY,
	))
}

func (t *Gim) gridOptionsSection() *fyne.Container {
	return fyne.NewContainerWithLayout(
		layout.NewGridLayout(1),
		widget.NewTabContainer(
			widget.NewTabItem(TitleTabItemGridCount, t.gridColumnRowOptions()),
			widget.NewTabItem(TitleTabItemGridSize, t.gridSizeOptions()),
		),
	)
}
