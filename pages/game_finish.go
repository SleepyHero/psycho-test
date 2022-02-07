package pages

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func InitGameFinishPage(w fyne.Window, bindLabelStr binding.String, nextpage NextPage) *DefaultPage {

	center := container.NewHBox(layout.NewSpacer())

	lable := widget.NewLabelWithData(bindLabelStr)
	lable.Alignment = fyne.TextAlignCenter

	finishButton := widget.NewButton("返回", func() {
		center.Hide()
		nextpage(4)
	})

	bg := container.NewVBox(layout.NewSpacer(), lable, layout.NewSpacer(), finishButton, layout.NewSpacer())

	center.Add(bg)
	center.Add(layout.NewSpacer())
	//c.SetContent(center)
	center.Hide()
	center.Refresh()
	return WrapPage(w, center)
}
