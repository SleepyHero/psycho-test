package pages

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"github.com/psycho-test/config"
	"image/color"
	"math/rand"
	"time"
)

type Record func(suc bool, timeCost int64, keyPressed string,
	targetKey string, targetIndex int, fontSize int)

type SearchBase struct {
	W         fyne.Window
	object    fyne.CanvasObject
	NextPage  NextPage
	RecordFuc Record
	fontSize  float32
	hintLabel *canvas.Text
	isTest    bool
}

type searchPage struct {
	SearchBase
	RightIndex int
	TargetKey  string
}

func (p *searchPage) SetActive() {
	p.W.SetContent(p.object)
	startTime := time.Now().UnixMilli()
	if p.isTest {
		p.hintLabel.Text = "按1-9数字键选择目标词"
		p.hintLabel.Refresh()
	}
	p.W.Canvas().SetOnTypedKey(func(event *fyne.KeyEvent) {
		//提示
		if event.Name == fyne.Key0 && !p.isTest {
			p.hintLabel.Text = fmt.Sprint(p.TargetKey)
			p.W.Canvas().Refresh(p.hintLabel)
			return
		}

		if _, ok := config.KeyMap[string(event.Name)]; !ok {
			return
		}
		suc := false
		if config.KeyMap[string(event.Name)] == p.RightIndex {
			suc = true
		}
		timeCost := time.Now().UnixMilli() - startTime
		p.W.Canvas().SetOnTypedKey(nil)
		p.RecordFuc(suc, timeCost, string(event.Name), p.TargetKey, p.RightIndex, int(p.fontSize))
		p.NextPage(2)
	})
}

func NewSearchPage(w fyne.Window, x, y float32, page NextPage, wordList []string, rightIndex int, textSize float32, record Record, isTest bool) Page {

	self := container.NewWithoutLayout()

	targetContainer := container.NewMax()
	targetContainer.Resize(fyne.NewSize(config.ConfigData.TargetWidth, config.ConfigData.TargetHeight))
	targetContainer.Move(fyne.NewPos(x, y))

	labels := make([]fyne.CanvasObject, 0)
	for i := 0; i < 9; i++ {
		label := canvas.NewText(wordList[i], color.White)

		label.TextSize = textSize

		label.Alignment = fyne.TextAlignCenter
		labels = append(labels, label)
	}
	center := container.NewAdaptiveGrid(3, labels...)

	hintLabel := canvas.NewText("", color.White)
	hintLabel.TextSize = textSize
	hintx, hinty := getHitPos(x, y)
	hintLabel.Move(fyne.NewPos(hintx, hinty))
	self.Add(hintLabel)

	targetContainer.Add(center)
	self.Add(targetContainer)

	return &searchPage{
		SearchBase: SearchBase{
			W:         w,
			object:    self,
			NextPage:  page,
			RecordFuc: record,
			fontSize:  textSize,
			isTest:    isTest,
			hintLabel: hintLabel,
		},
		RightIndex: rightIndex,
		TargetKey:  wordList[rightIndex],
	}
}

type searchNumPage struct {
	SearchBase
	Num int
}

func (p *searchNumPage) SetActive() {
	p.W.SetContent(p.object)
	startTime := time.Now().UnixMilli()
	if p.isTest {
		p.hintLabel.Text = "大于600按↑，小于600按↓"
		p.hintLabel.Refresh()
	}
	p.W.Canvas().SetOnTypedKey(func(event *fyne.KeyEvent) {
		//提示
		//if event.Name == fyne.Key0 && !p.isTest {
		//	p.hintLabel.Text = fmt.Sprint(p.Num)
		//	p.hintLabel.Refresh()
		//	return
		//}
		suc := false
		if event.Name == fyne.KeyUp || event.Name == fyne.KeyDown {
			timeCost := time.Now().UnixMilli() - startTime
			if (event.Name == fyne.KeyUp && p.Num > 600) || (event.Name == fyne.KeyDown && p.Num < 600) {
				suc = true
			}
			p.W.Canvas().SetOnTypedKey(nil)
			p.RecordFuc(suc, timeCost, string(event.Name), fmt.Sprint(p.Num), 0, int(p.fontSize))
			p.NextPage(2)
		}

	})
}

func NewSearchNumPage(w fyne.Window, x, y float32, page NextPage, textSize float32, record Record, isTest bool) Page {

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

	hintLabel := canvas.NewText("", color.White)
	hintLabel.TextSize = textSize
	hintx, hinty := getHitPos(x, y)
	hintLabel.Move(fyne.NewPos(hintx, hinty))
	self.Add(hintLabel)

	targetContainer.Add(center)
	self.Add(targetContainer)

	return &searchNumPage{
		SearchBase: SearchBase{
			W:         w,
			object:    self,
			NextPage:  page,
			RecordFuc: record,
			fontSize:  textSize,
			isTest:    isTest,
			hintLabel: hintLabel,
		},
		Num: num,
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

func getHitPos(x, y float32) (float32, float32) {
	if x < config.ConfigData.TotalWidth/2 {
		return config.ConfigData.TotalWidth/2 + 100, y
	}
	return config.ConfigData.TotalWidth/2 - 100, y
}
