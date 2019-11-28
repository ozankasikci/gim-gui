package gimgui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func (t *Gim) ActionsSection() fyne.CanvasObject {
	return fyne.NewContainerWithLayout(
		layout.NewVBoxLayout(),
		widget.NewGroup("Actions",
			widget.NewButton("Merge",t.merge),
		),
	)
}
