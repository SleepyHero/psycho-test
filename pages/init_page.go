package pages

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/psycho-test/config"
	"image/color"
)

var (
	textSize  = 24
	TextColor = color.White
	Title     = "大屏文字测试实验"
)

func NewInitPage(w fyne.Window, nextPage NextPage) *DefaultPage {

	lable := widget.NewLabel(Title)
	lable.Alignment = fyne.TextAlignCenter

	bg := container.NewVBox(lable)
	center := container.NewCenter(bg)
	numEntry := widget.NewEntry()
	num := widget.NewFormItem("编号", numEntry)

	selectEntry := widget.NewSelect([]string{"男", "女"}, func(s string) {

	})
	gender := widget.NewFormItem("性别", selectEntry)
	ageEntry := widget.NewEntry()
	age := widget.NewFormItem("年龄", ageEntry)
	form := widget.NewForm(num, gender, age)
	form.Resize(fyne.NewSize(10*config.ConfigData.Scale, 200*config.ConfigData.Scale))
	form.Refresh()
	button := widget.NewButton("进入实验", func() {
		fmt.Println(numEntry.Text)
		fmt.Println(selectEntry.Selected)
		fmt.Println(ageEntry.Text)
		//center.Hide()
		config.Num = numEntry.Text
		config.Gender = selectEntry.Selected
		config.Age = ageEntry.Text
		nextPage(0)
	})
	button.Alignment = widget.ButtonAlignCenter

	//bg.Add(nil)
	//bg.Add(numEntry)
	bg.Add(form)
	bg.Add(button)

	//bg.Add(button)
	button.Resize(fyne.NewSize(10, 30))
	button.Refresh()
	//c.SetContent(center)
	return WrapPage(w, center)
}
