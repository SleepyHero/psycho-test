package text_exam

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/psycho-test/pages"
)

var ()

type StartTrails func(isTest bool)

func InitBeginPage(w fyne.Window, titleStr binding.String, contentStr binding.String, trails StartTrails, backInit StartTrails) *pages.DefaultPage {
	center := container.NewHBox()
	lable := widget.NewLabelWithData(titleStr)
	lable.Alignment = fyne.TextAlignCenter

	content := widget.NewLabelWithData(contentStr)
	lable.Alignment = fyne.TextAlignCenter

	testButton := widget.NewButton("练习任务", func() {
		center.Hide()
		trails(true)
	})
	examButton := widget.NewButton("正式实验", func() {
		center.Hide()
		trails(false)
	})

	returnButton := widget.NewButton("返回主页", func() {
		backInit(true)
	})

	hl := container.NewHBox(testButton, layout.NewSpacer(), examButton, layout.NewSpacer(), returnButton)
	//button.Alignment = widget.ButtonAlignCenter

	bg := container.NewVBox(layout.NewSpacer(), lable, content, layout.NewSpacer(), hl, layout.NewSpacer())

	bg.Resize(fyne.NewSize(500, 100))

	center.Add(layout.NewSpacer())
	center.Add(bg)
	center.Add(layout.NewSpacer())

	//c.SetContent(center)
	center.Hide()
	center.Refresh()
	return pages.WrapPage(w, center)
}
