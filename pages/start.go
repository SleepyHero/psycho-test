package pages

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var (
	kinds = []string{"文字搜索", "数字识别"}
)

type NextPage func(pageNum int)
type StartLoop func(kind string, dis string, region string)

func NewStartPage(w fyne.Window, nextPage StartLoop) *DefaultPage {

	center := container.NewHBox(layout.NewSpacer())

	lable := widget.NewLabel(Title)
	lable.Alignment = fyne.TextAlignCenter

	selectkind := widget.NewSelect([]string{"文字搜索", "数字识别"}, func(s string) {

	})
	num := widget.NewFormItem("实验类型", selectkind)

	selectDistance := widget.NewSelect([]string{"3m", "5m", "7m"}, func(s string) {

	})
	gender := widget.NewFormItem("距离", selectDistance)
	selectRegion := widget.NewSelect([]string{"左", "右"}, func(s string) {

	})
	age := widget.NewFormItem("呈现区域", selectRegion)
	form := widget.NewForm(num, gender, age)
	startButton := widget.NewButton("开始实验", func() {
		fmt.Println(selectkind.Selected)
		center.Hide()
		nextPage(selectkind.Selected, selectDistance.Selected, selectRegion.Selected)
	})
	exitButton := widget.NewButton("结束实验", func() {
		w.Close()
	})

	hl := container.NewHBox(startButton, layout.NewSpacer(), exitButton)
	bg := container.NewVBox(layout.NewSpacer(), lable, form, layout.NewSpacer(), hl, layout.NewSpacer())

	center.Add(bg)
	center.Add(layout.NewSpacer())

	//Window.SetContent(center)
	center.Hide()
	center.Refresh()
	return WrapPage(w, center)
}
