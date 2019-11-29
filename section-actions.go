package gimgui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func (t *Gim) actionsSection() fyne.CanvasObject {
	return fyne.NewContainerWithLayout(
		layout.NewVBoxLayout(),
		widget.NewGroup(TitleSectionActions,
			widget.NewButton(TitleButtonMerge, t.merge),
		),
	)
}
