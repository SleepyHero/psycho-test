package process

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"github.com/psycho-test/config"
	"github.com/psycho-test/pages"
)

type ProcessHandler struct {
	window fyne.Window
	pages  []pages.Page
	trails []pages.Page
	Kind   string
	Dis    string
	Region string
	IsTest bool
	IsText bool

	trailCount int

	endText    binding.String
	titleStr   binding.String
	contentStr binding.String

	wordGen    *TargetWordGen
	trailGen   *TrailGen
	x, y, size float32
	suc        bool
}

func NewProcessHandler(w fyne.Window) *ProcessHandler {
	res := &ProcessHandler{}
	res.endText = binding.NewString()
	res.titleStr = binding.NewString()
	res.contentStr = binding.NewString()
	res.pages = make([]pages.Page, 0, 10)
	res.trails = make([]pages.Page, 0, 10)
	res.window = w
	res.wordGen = NewTargetWordGen()
	res.initProcess(w)
	return res
}

func (p *ProcessHandler) initProcess(w fyne.Window) {
	p.pages = append(p.pages, pages.NewInitPage(w, p.nextPage))
	p.pages = append(p.pages, pages.NewStartPage(w, p.startLoop))
	p.pages = append(p.pages, pages.InitBeginPage(w, p.titleStr, p.contentStr, p.startTrails, p.backInit))

	p.trails = append(p.trails, pages.InitIntroPage(w, p.nextTrail))
	p.trails = append(p.trails, pages.InitCommentPage(w, p.nextTrail, p.onCommend))
	p.trails = append(p.trails, pages.InitGameFinishPage(w, p.endText, p.nextTrail))
	p.trails = append(p.trails, pages.InitResultPage(w, true, p.nextTrail))
	p.trails = append(p.trails, pages.InitResultPage(w, false, p.nextTrail))

	p.pages[0].SetActive()
}

func (p *ProcessHandler) nextPage(pageNum int) {
	p.pages[1].SetActive()
}

func (p *ProcessHandler) nextTrail(pageNum int) {
	switch pageNum {
	case 0: //开始实验结束
		p.x, p.y, p.size = p.trailGen.NextPos()
		var np pages.Page
		if p.IsText {
			np = pages.NewTargetPage(p.window, p.wordGen.GetCurrentTrailTextTargetStr(), p.x, p.y, p.nextTrail, p.size)
		} else {
			np = pages.NewTargetNumPage(p.window, p.x, p.y, p.nextTrail)
		}
		np.SetActive()
	case 1: //目标结束
		var np pages.Page
		if p.IsText {
			np = pages.NewSearchPage(p.window, p.x, p.y, p.nextTrail, p.wordGen.GetCurrentTrailTextList(), p.wordGen.GetCurrentTrailTextTarget(), p.size, p.recordResult, p.IsTest)
		} else {
			np = pages.NewSearchNumPage(p.window, p.x, p.y, p.nextTrail, p.size, p.recordResult, p.IsTest)
		}
		np.SetActive()
	case 2: //搜索结束
		if p.IsTest {
			if p.suc {
				p.trails[3].SetActive()
			} else {
				p.trails[4].SetActive()
			}
		} else {
			if !p.trailGen.IsCurrentFinish(p.IsTest) {

				p.trails[0].SetActive()
			} else {
				p.trailGen.MoveNext()
				p.trails[1].SetActive()
			}
			p.wordGen.GenNextTrail()
		}
	case 3: //评论结束
		if p.trailGen.IsFinish() {
			p.trails[2].SetActive()
		} else {
			p.trails[0].SetActive()
			p.wordGen.GenNextTrail()
		}
	case 4: //结束页结束
		p.pages[2].SetActive()
	case 5: //测试判断页结束
		if !p.trailGen.IsCurrentFinish(p.IsTest) {
			p.trails[0].SetActive()
		} else {
			p.trailGen.MoveNextTest()
			p.trails[1].SetActive()
		}
		p.wordGen.GenNextTrail()
	}
}

func (p *ProcessHandler) recordResult(suc bool, timeCost int64, keyPressed string,
	targetKey string, targetIndex int, fontSize int) {
	p.suc = suc
	if p.IsTest {
		return
	}
	targetPos, showTimes := p.trailGen.GetLastTargetInfo()

	SaveRes(p.Kind, p.Dis, p.Region, p.trailGen.GetCount(), fontSize, targetPos, showTimes, timeCost,
		keyPressed, suc, "", targetKey, targetIndex, config.Num, config.Gender, config.Age)
}

func (p *ProcessHandler) startLoop(kind string, dis string, region string) {
	p.Kind = kind
	p.Dis = dis
	p.Region = region
	p.IsText = kind == "文字搜索"

	if p.IsText {
		_ = p.titleStr.Set(config.TextTitle)
		_ = p.contentStr.Set(config.TextContent)
	} else {
		_ = p.titleStr.Set(config.NumTitle)
		_ = p.contentStr.Set(config.NumContent)
	}
	p.pages[2].SetActive()
}

func (p *ProcessHandler) startTrails(isTest bool) {
	p.IsTest = isTest
	if isTest {
		_ = p.endText.Set(config.FinishTest)
	} else {
		_ = p.endText.Set(config.FinishGame)
	}
	p.trailGen = NewTailGen(p.Dis, p.Region)
	p.trails[0].SetActive()
}

func (p *ProcessHandler) backInit(isTest bool) {
	p.pages[1].SetActive()
}

func (p *ProcessHandler) onCommend(res string) {
	Push(res)
}
