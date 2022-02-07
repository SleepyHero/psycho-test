package process

import "math/rand"

type TargetWordGen struct {
	currentStep        int
	currentTextList    []string
	currentTargetIndex int
	list               []map[int]Column
}

func NewTargetWordGen() *TargetWordGen {
	t := &TargetWordGen{}
	t.GenNextTrail()
	return t
}

func (t *TargetWordGen) GetCurrentTrailTextList() []string {
	return t.currentTextList
}

func (t *TargetWordGen) GetCurrentTrailTextTarget() int {
	return t.currentTargetIndex
}

func (t *TargetWordGen) GetCurrentTrailTextTargetStr() string {
	return t.currentTextList[t.currentTargetIndex]
}

func (t *TargetWordGen) GenNextTrail() {
	t.currentStep++
	t.currentTextList = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	t.currentTargetIndex = rand.Intn(len(t.currentTextList))
}
