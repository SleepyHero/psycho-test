package pages

import "fyne.io/fyne/v2"

type Page interface {
	SetActive()
}

type DefaultPage struct {
	Window fyne.Window
	Object fyne.CanvasObject
}

func (p *DefaultPage) SetActive() {
	p.Window.SetContent(p.Object)
}

func WrapPage(w fyne.Window, object fyne.CanvasObject) *DefaultPage {
	return &DefaultPage{
		Window: w,
		Object: object,
	}
}
