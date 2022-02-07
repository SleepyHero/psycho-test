package pages

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"github.com/psycho-test/config"
	"image/color"
	"time"
)

var text = "数据"

type targetPage struct {
	w        fyne.Window
	object   fyne.CanvasObject
	nextPage NextPage
}

func NewTargetPage(w fyne.Window, textStr string, x, y float32, nextpage NextPage, textSize float32) *targetPage {
	self := container.NewWithoutLayout()
	targetContainer := container.NewMax()
	targetContainer.Resize(fyne.NewSize(config.ConfigData.TargetWidth, config.ConfigData.TargetHeight))
	targetContainer.Move(fyne.NewPos(x, y))

	outside := container.NewWithoutLayout()
	center := container.NewCenter()

	drawTarget(outside, 0, 0, config.ConfigData.TargetWidth, config.ConfigData.TargetHeight, 10)
	label := canvas.NewText(textStr, color.White)

	label.TextSize = textSize

	center.Add(label)
	targetContainer.Add(outside)
	targetContainer.Add(center)
	self.Add(targetContainer)

	return &targetPage{
		w:        w,
		object:   self,
		nextPage: nextpage,
	}
}

func NewTargetNumPage(w fyne.Window, x, y float32, nextpage NextPage) *targetPage {
	self := container.NewWithoutLayout()
	targetContainer := container.NewMax()
	targetContainer.Resize(fyne.NewSize(config.ConfigData.TargetWidth, config.ConfigData.TargetHeight))
	targetContainer.Move(fyne.NewPos(x, y))

	outside := container.NewWithoutLayout()

	drawTarget(outside, 0, 0, config.ConfigData.TargetWidth, config.ConfigData.TargetHeight, 10)

	targetContainer.Add(outside)
	self.Add(targetContainer)

	return &targetPage{
		w:        w,
		object:   self,
		nextPage: nextpage,
	}
}

func (p *targetPage) SetActive() {
	p.w.SetContent(p.object)
	go func() {
		time.Sleep(time.Duration(config.ConfigData.SleepTime) * time.Millisecond)
		p.nextPage(1)
	}()
}

func drawTarget(container *fyne.Container, x, y float32, width float32, height float32, len float32) {
	//width -= 5
	//height -= 5
	container.Add(drawLine(x, y, x+len, y))
	container.Add(drawLine(x, y, x, y+len))

	container.Add(drawLine(x+width, y, x+width-len, y))
	container.Add(drawLine(x+width, y, x+width, y+len))

	container.Add(drawLine(x, y+height, x+len, y+height))
	container.Add(drawLine(x, y+height, x, y+height-len))

	container.Add(drawLine(x+width, y+height, x+width-len, y+height))
	container.Add(drawLine(x+width, y+height, x+width, y-len+height))
}

func drawLine(x1, y1, x2, y2 float32) *canvas.Line {
	line := canvas.NewLine(color.White)
	line.Position1 = fyne.NewPos(x1, y1)
	line.Position2 = fyne.NewPos(x2, y2)
	line.StrokeWidth = 2
	return line
}
