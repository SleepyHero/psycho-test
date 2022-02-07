package process

import (
	"fmt"
	"github.com/psycho-test/config"
	"math/rand"
)

type TrailGen struct {
	leftToRight  bool
	smallToBig   bool
	order        [][]*Column
	curColunm    int
	curIndex     int
	count        int
	lastRow      int
	lastRowTimes int
	isLeft       bool
}

func NewTailGen(dis string, region string) *TrailGen {
	tg := &TrailGen{}
	tg.smallToBig = rand.Intn(2) == 0
	tg.leftToRight = rand.Intn(2) == 0
	columnNum := 4
	if dis == "3m" {
		columnNum = 3
	}
	direction := 1
	if region == "左" {
		tg.isLeft = true
		direction = -1
	}
	tg.order = make([][]*Column, columnNum)
	for i := 0; i < columnNum; i++ {
		tg.order[i] = make([]*Column, len(config.ConfigData.FontSize))
		for j := 0; j < len(config.ConfigData.FontSize); j++ {
			tg.order[i][j] = NewColumn(3 + (i * direction))
		}
	}
	//if !tg.smallToBig {
	//	tg.curIndex = len(config.FontSize) - 1
	//}

	if !tg.leftToRight {
		tg.curColunm = len(tg.order) - 1
	}

	return tg
}

func (tg *TrailGen) GetTextSize() float32 {
	if tg.smallToBig {
		return config.ConfigData.FontSize[tg.curIndex]
	}
	return config.ConfigData.FontSize[len(config.ConfigData.FontSize)-1-tg.curIndex]
}

func (tg *TrailGen) NextPos() (float32, float32, float32) {
	x, y, row, times := tg.order[tg.curColunm][tg.curIndex].GetNewPos()
	tg.lastRow = row
	tg.lastRowTimes = times
	size := tg.GetTextSize()
	tg.count++
	return x, y, size
}

func (tg *TrailGen) GetCount() int {
	return tg.count
}

func (tg *TrailGen) GetLastTargetInfo() (string, int) {
	//return tg.count
	first := 0
	if tg.isLeft {
		first = 4 - tg.curColunm
	} else {
		first = 4 + tg.curColunm
	}

	return fmt.Sprintf("%v-%v", first, tg.lastRow), tg.lastRowTimes
}

func (tg *TrailGen) MoveNext() {

	tg.curIndex++
	if !tg.IsRoundFinish() {
		return
	}

	tg.curIndex = 0

	if tg.leftToRight {
		tg.curColunm++
	} else {
		tg.curColunm--
	}
}

func (tg *TrailGen) IsFinish() bool {
	return tg.curColunm < 0 || tg.curColunm >= len(tg.order)
}

func (tg *TrailGen) IsRoundFinish() bool {
	return tg.curIndex < 0 || tg.curIndex >= len(config.ConfigData.FontSize)
}

func (tg *TrailGen) IsCurrentFinish(isTest bool) bool {
	return tg.order[tg.curColunm][tg.curIndex].IsColumnFinish(isTest)
}

func (tg *TrailGen) MoveNextTest() {
	tg.curIndex += len(config.ConfigData.FontSize) - 1
	if !tg.IsRoundFinish() {
		return
	}

	tg.curIndex = 0

	if tg.leftToRight {
		tg.curColunm += 2
	} else {
		tg.curColunm -= 2
	}
}
