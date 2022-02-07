package trail

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/psycho-test/config"
	"github.com/psycho-test/pages"
)

type IntroPage struct {
	w        fyne.Window
	o        fyne.CanvasObject
	nextPage pages.NextPage
}

func InitIntroPage(w fyne.Window, nextPage pages.NextPage) pages.Page {

	lable := widget.NewLabel(config.IntroHint)
	lable.Alignment = fyne.TextAlignCenter
	center := container.NewCenter()

	//bg := container.NewVBox(layout.NewSpacer(), lable, lable2, layout.NewSpacer())

	center.Add(lable)

	return &IntroPage{
		w:        w,
		o:        center,
		nextPage: nextPage,
	}
}

func (p *IntroPage) SetActive() {
	p.w.SetContent(p.o)
	p.w.Canvas().SetOnTypedKey(func(event *fyne.KeyEvent) {
		switch event.Name {
		case fyne.KeySpace:
			p.w.Canvas().SetOnTypedKey(nil)
			p.nextPage(0)
		}
	})
}
