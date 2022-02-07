package pages

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/psycho-test/config"
	"time"
)

type resultPage struct {
	w        fyne.Window
	object   fyne.CanvasObject
	nextPage NextPage
}

func InitResultPage(w fyne.Window, suc bool, page NextPage) Page {

	center := container.NewHBox(layout.NewSpacer())

	text := "错误"
	if suc {
		text = "正确"
	}

	lable := widget.NewLabel(text)
	lable.Alignment = fyne.TextAlignLeading

	bg := container.NewVBox(layout.NewSpacer(), lable, layout.NewSpacer())

	center.Add(bg)
	center.Add(layout.NewSpacer())

	//Window.SetContent(center)
	return &resultPage{
		w:        w,
		object:   center,
		nextPage: page,
	}
}

func (p *resultPage) SetActive() {
	p.w.SetContent(p.object)
	go func() {
		time.Sleep(time.Duration(config.ConfigData.SleepTime) * time.Millisecond)
		p.nextPage(5)
	}()
}
