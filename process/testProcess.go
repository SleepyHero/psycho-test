package process

//
//import (
//	"fyne.io/fyne/v2"
//	"fyne.io/fyne/v2/data/binding"
//	"github.com/psycho-test/config"
//	"github.com/psycho-test/pages"
//	"github.com/psycho-test/pages/text_exam"
//	"github.com/psycho-test/pages/trail"
//)
//
//package process
//
//import (
//"fyne.io/fyne/v2"
//"fyne.io/fyne/v2/data/binding"
//"github.com/psycho-test/config"
//"github.com/psycho-test/pages"
//"github.com/psycho-test/pages/text_exam"
//"github.com/psycho-test/pages/trail"
//)
//
//type ProcessHandler struct {
//	window fyne.Window
//	pages  []pages.Page
//	trails []pages.Page
//	Kind   string
//	Dis    string
//	Region string
//	IsTest bool
//	IsText bool
//
//	trailCount int
//
//	endText    binding.String
//	titleStr   binding.String
//	contentStr binding.String
//	targetStr  binding.String
//
//	trail      *TargetWordGen
//	tailGen    *TrailGen
//	x, y, size float32
//	suc        bool
//}
//
//func NewProcessHandler(w fyne.Window) *ProcessHandler {
//	res := &ProcessHandler{}
//	res.endText = binding.NewString()
//	res.titleStr = binding.NewString()
//	res.contentStr = binding.NewString()
//	res.targetStr = binding.NewString()
//	res.pages = make([]pages.Page, 0, 10)
//	res.trails = make([]pages.Page, 0, 10)
//	res.window = w
//	res.trail = NewTargetWordGen()
//	res.initProcess(w)
//	return res
//}
//
//func (p *ProcessHandler) initProcess(w fyne.Window) {
//	p.pages = append(p.pages, pages.NewInitPage(w, p.nextPage))
//	p.pages = append(p.pages, pages.NewStartPage(w, p.startLoop))
//	p.pages = append(p.pages, text_exam.InitBeginPage(w, p.titleStr, p.contentStr, p.startTrails))
//
//	p.trails = append(p.trails, trail.InitIntroPage(w, p.nextTrail))
//	//p.trails = append(p.trails, pages.NewTargetPage(w, p.targetStr, 100, 100, p.nextTrail))
//	//p.trails = append(p.trails, pages.NewSearchPage(w, p.targetStr, 100, 100, p.nextTrail, []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}, 2))
//	p.trails = append(p.trails, pages.InitCommentPage(w, p.nextTrail))
//	p.trails = append(p.trails, pages.InitGameFinishPage(w, p.endText, p.nextTrail))
//
//	p.pages[0].SetActive()
//}
//
//func (p *ProcessHandler) nextPage(pageNum int) {
//	p.pages[1].SetActive()
//}
//
//func (p *ProcessHandler) nextTrail(pageNum int) {
//	switch pageNum {
//	case 0:
//		//p.trail.GenNextTrail()
//		p.x, p.y, p.size = p.tailGen.NextPos()
//		np := pages.NewTargetPage(p.window, p.trail.GetCurrentTrailTextTargetStr(), p.x, p.y, p.nextTrail, p.size)
//		np.SetActive()
//	case 1:
//		np := pages.NewSearchPage(p.window, p.targetStr, p.x, p.y, p.nextTrail, p.trail.GetCurrentTrailTextList(), p.trail.GetCurrentTrailTextTarget(), p.size, p.recordResult)
//		np.SetActive()
//	case 2:
//		if !p.tailGen.IsCurrentFinish() {
//			p.trails[0].SetActive()
//		} else {
//			p.tailGen.MoveNext()
//			p.trail.GenNextTrail()
//			p.trails[1].SetActive()
//		}
//
//	case 3:
//		if p.tailGen.IsFinish() {
//			p.trails[2].SetActive()
//		} else {
//			p.trails[0].SetActive()
//			p.trail.GenNextTrail()
//		}
//
//	case 4:
//		p.pages[2].SetActive()
//	}
//
//	//nextPage := (pageNum + 1) % len(p.trails)
//	//p.refreshData()
//	//p.trails[nextPage].SetActive()
//}
//
//
//func
//
//func (p *ProcessHandler) refreshData() {
//	if p.Kind == "文字搜索" {
//
//	}
//}
//
//func (p *ProcessHandler) recordResult(suc bool, timeCost int64) {
//	p.suc = suc
//	if !p.IsTest{
//
//	}
//}
//
//func (p *ProcessHandler) startLoop(kind string, dis string, region string) {
//	p.Kind = kind
//	p.Dis = dis
//	p.Region = region
//	p.IsText = kind == "文字搜索"
//	p.tailGen = NewTailGen(dis, region)
//	if p.IsText {
//		_ = p.titleStr.Set(config.TextTitle)
//		_ = p.contentStr.Set(config.TextContent)
//		_ = p.targetStr.Set("数据")
//	} else {
//		_ = p.titleStr.Set(config.NumTitle)
//		_ = p.contentStr.Set(config.NumContent)
//	}
//	p.pages[2].SetActive()
//}
//
//func (p *ProcessHandler) startTrails(isTest bool) {
//	p.IsTest = isTest
//	if isTest {
//		_ = p.endText.Set(config.FinishTest)
//	} else {
//		_ = p.endText.Set(config.FinishGame)
//	}
//	p.trails[0].SetActive()
//}
