package pages

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type StartTrails func(isTest bool)

func InitBeginPage(w fyne.Window, titleStr binding.String, contentStr binding.String, trails StartTrails, backInit StartTrails) *DefaultPage {
	center := container.NewHBox()
	label := widget.NewLabelWithData(titleStr)
	label.Alignment = fyne.TextAlignCenter

	content := widget.NewLabelWithData(contentStr)
	label.Alignment = fyne.TextAlignCenter

	testButton := widget.NewButton("练习任务", func() {
		trails(true)
	})
	examButton := widget.NewButton("正式实验", func() {
		trails(false)
	})

	returnButton := widget.NewButton("返回主页", func() {
		backInit(true)
	})

	hl := container.NewHBox(testButton, layout.NewSpacer(), examButton, layout.NewSpacer(), returnButton)

	bg := container.NewVBox(layout.NewSpacer(), label, content, hl, layout.NewSpacer())

	bg.Resize(fyne.NewSize(500, 100))

	center.Add(layout.NewSpacer())
	center.Add(bg)
	center.Add(layout.NewSpacer())

	center.Hide()
	center.Refresh()
	return WrapPage(w, center)
}
