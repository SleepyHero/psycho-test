package pages

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"github.com/psycho-test/config"
	"image/color"
	"math/rand"
	"time"
)

type Record func(suc bool, timeCost int64, keyPressed string,
	targetKey string, targetIndex int, fontSize int)

type searchPage struct {
	w          fyne.Window
	object     fyne.CanvasObject
	NextPage   NextPage
	RecordFuc  Record
	RightIndex int
	TargetKey  string
	fontSize   float32
	textStr    binding.String
}

func (p *searchPage) SetActive() {
	p.w.SetContent(p.object)
	startTime := time.Now().UnixMilli()
	p.w.Canvas().SetOnTypedKey(func(event *fyne.KeyEvent) {
		//提示
		if event.Name == fyne.Key0 {
			//p.textStr.Set("按1-9数字键选择目标词")
		}

		if _, ok := config.KeyMap[string(event.Name)]; !ok {
			return
		}
		suc := false
		if config.KeyMap[string(event.Name)] == p.RightIndex {
			suc = true
		}
		timeCost := time.Now().UnixMilli() - startTime
		p.w.Canvas().SetOnTypedKey(nil)
		p.RecordFuc(suc, timeCost, string(event.Name), p.TargetKey, p.RightIndex, int(p.fontSize))
		p.NextPage(2)
	})
}

func NewSearchPage(w fyne.Window, textStr binding.String, x, y float32, page NextPage, wordList []string, rightIndex int, textSize float32, record Record) Page {

	self := container.NewWithoutLayout()
	//Window.SetContent(center)

	targetContainer := container.NewMax()
	targetContainer.Resize(fyne.NewSize(config.ConfigData.TargetWidth, config.ConfigData.TargetHeight))
	targetContainer.Move(fyne.NewPos(x, y))

	//outside := container.NewWithoutLayout()
	labels := make([]fyne.CanvasObject, 0)
	for i := 0; i < 9; i++ {
		label := canvas.NewText(wordList[i], color.White)

		label.TextSize = textSize

		label.Alignment = fyne.TextAlignCenter
		labels = append(labels, label)
	}
	center := container.NewAdaptiveGrid(3, labels...)

	//center.Add(label)
	//targetContainer.Add(outside)
	targetContainer.Add(center)
	self.Add(targetContainer)

	return &searchPage{
		w:          w,
		object:     self,
		NextPage:   page,
		RightIndex: rightIndex,
		RecordFuc:  record,
		fontSize:   textSize,
		TargetKey:  wordList[rightIndex],
	}
}

type searchNumPage struct {
	w         fyne.Window
	object    fyne.CanvasObject
	NextPage  NextPage
	RecordFuc Record
	Num       int
	fontSize  float32
}

func (p *searchNumPage) SetActive() {
	p.w.SetContent(p.object)
	startTime := time.Now().UnixMilli()
	p.w.Canvas().SetOnTypedKey(func(event *fyne.KeyEvent) {
		suc := false
		if event.Name == fyne.KeyUp || event.Name == fyne.KeyDown {
			timeCost := time.Now().UnixMilli() - startTime
			if (event.Name == fyne.KeyUp && p.Num > 600) || (event.Name == fyne.KeyDown && p.Num < 600) {
				suc = true
			}
			p.w.Canvas().SetOnTypedKey(nil)
			p.RecordFuc(suc, timeCost, string(event.Name), fmt.Sprint(p.Num), 0, int(p.fontSize))
			p.NextPage(2)
		}

	})
}

func NewSearchNumPage(w fyne.Window, x, y float32, page NextPage, textSize float32, record Record) Page {

	self := container.NewWithoutLayout()
	//Window.SetContent(center)

	targetContainer := container.NewMax()
	targetContainer.Resize(fyne.NewSize(config.ConfigData.TargetWidth, config.ConfigData.TargetHeight))
	targetContainer.Move(fyne.NewPos(x, y))

	num := GenNum()
	label := canvas.NewText(fmt.Sprint(num), color.White)

	label.TextSize = textSize

	label.Alignment = fyne.TextAlignCenter

	center := container.NewCenter(label)

	//center.Add(label)
	//targetContainer.Add(outside)
	targetContainer.Add(center)
	self.Add(targetContainer)

	return &searchNumPage{
		w:         w,
		object:    self,
		NextPage:  page,
		RecordFuc: record,
		Num:       num,
		fontSize:  textSize,
	}
}

func GenNum() int {
	for {
		n := rand.Intn(1000)
		if n != 600 || n < 100 {
			return n
		}
	}

}
