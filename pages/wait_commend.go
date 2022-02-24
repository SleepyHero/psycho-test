package pages

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/psycho-test/config"
)

type WaitCommendPage struct {
	w        fyne.Window
	o        fyne.CanvasObject
	nextPage NextPage
}

func InitWaitCommendPage(w fyne.Window, nextPage NextPage) Page {

	lable := widget.NewLabel(config.WaitHint)
	lable.Alignment = fyne.TextAlignCenter
	center := container.NewCenter()

	//bg := container.NewVBox(layout.NewSpacer(), lable, lable2, layout.NewSpacer())

	center.Add(lable)

	return &WaitCommendPage{
		w:        w,
		o:        center,
		nextPage: nextPage,
	}
}

func (p *WaitCommendPage) SetActive() {
	p.w.SetContent(p.o)
	p.w.Canvas().SetOnTypedKey(func(event *fyne.KeyEvent) {
		switch event.Name {
		case fyne.KeySpace:
			p.w.Canvas().SetOnTypedKey(nil)
			p.nextPage(6)
		}
	})
}
