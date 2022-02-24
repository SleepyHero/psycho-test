package process

import (
	"bufio"
	"fmt"
	"github.com/psycho-test/config"
	"io"
	"math/rand"
	"os"
	"strings"
)

type TargetWordGen struct {
	currentStep        int
	currentTextList    []string
	currentTargetIndex int
	list               []map[int]Column
	totalWords         []string
	curWordPool        []string
}

func NewTargetWordGen() *TargetWordGen {
	t := &TargetWordGen{}
	t.totalWords = make([]string, 0, 800)
	t.currentTextList = make([]string, 9, 9)
	//t.GenNextTrail()
	t.loadWordList()
	t.curWordPool = make([]string, len(t.totalWords), len(t.totalWords))
	return t
}

func (t *TargetWordGen) loadWordList() {
	file, _ := os.Open(config.Dir + string(os.PathSeparator) + "words.txt")
	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		fmt.Println(line)
		t.totalWords = append(t.totalWords, line)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return
			}
		}
	}
}

func (t *TargetWordGen) InitWordGen() {
	t.curWordPool = make([]string, len(t.totalWords), len(t.totalWords))
	for k, v := range t.totalWords {
		t.curWordPool[k] = v
	}
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

func (t *TargetWordGen) GenNextTrail(isTest bool) {
	t.currentStep++
	selectMap := make(map[int]string, 9)

	for {
		if len(selectMap) >= 9 {
			break
		}
		index := rand.Intn(len(t.curWordPool))
		if _, ok := selectMap[index]; ok {
			continue
		}
		//key是index，value是string
		selectMap[index] = t.curWordPool[index]
	}
	i := 0
	//随机一组target
	t.currentTargetIndex = rand.Intn(len(selectMap))
	for k, v := range selectMap {
		t.currentTextList[i] = v
		if i == t.currentTargetIndex {
			//不放回随机，所以把被选中的值从列表中删除。测试不删除
			if !isTest {
				t.curWordPool = append(t.curWordPool[:k], t.curWordPool[k+1:]...)
			}
		}
		i++
	}

	//t.currentTextList = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	//t.currentTargetIndex = rand.Intn(len(t.currentTextList))
}
