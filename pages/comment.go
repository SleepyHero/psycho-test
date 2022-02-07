package pages

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/psycho-test/config"
)

type CommendPage struct {
	w         fyne.Window
	o         fyne.CanvasObject
	nextpage  NextPage
	onCommend commendFunc
}

func (p *CommendPage) SetActive() {
	p.w.SetContent(p.o)
	p.w.Canvas().SetOnTypedKey(func(event *fyne.KeyEvent) {
		switch event.Name {
		case fyne.Key1:
			fallthrough
		case fyne.Key2:
			fallthrough
		case fyne.Key3:
			fallthrough
		case fyne.Key4:
			fallthrough
		case fyne.Key5:
			p.onCommend(string(event.Name))
			p.nextpage(3)
			fmt.Println(event.Name)
		}
	})
}

type commendFunc func(str string)

func InitCommentPage(w fyne.Window, page NextPage, onCommend commendFunc) Page {

	center := container.NewHBox(layout.NewSpacer())

	lable := widget.NewLabel(config.ConfigData.CommendText)
	lable.Alignment = fyne.TextAlignLeading

	lable2 := widget.NewLabel(config.CommentHint)
	lable2.Alignment = fyne.TextAlignCenter

	bg := container.NewVBox(layout.NewSpacer(), lable, lable2, layout.NewSpacer())

	center.Add(bg)
	center.Add(layout.NewSpacer())

	//Window.SetContent(center)
	return &CommendPage{
		w:         w,
		o:         center,
		nextpage:  page,
		onCommend: onCommend,
	}
}
