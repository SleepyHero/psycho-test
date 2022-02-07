package process

import (
	"github.com/psycho-test/config"
	"math/rand"
	"time"
)

const randNum = 8

func init() {
	rand.Seed(time.Now().Unix())
}

type Column struct {
	remainList []int
	x          float32
	timesCount []int
}

func NewColumn(index int) *Column {
	c := &Column{remainList: []int{0, 1, 2, 3, 4, 5, 6, 7}, x: float32(index) * (config.ConfigData.TargetWidth + config.TargetDis)}
	c.timesCount = make([]int, 4, 4)
	return c
}

func (c *Column) GetNewPos() (float32, float32, int, int) {
	n := rand.Intn(len(c.remainList))
	num := c.remainList[n]
	c.remainList = append(c.remainList[:n], c.remainList[n+1:]...)
	row := num / 2
	c.timesCount[row]++
	return c.x, float32(row) * config.ConfigData.TargetHeight, row, c.timesCount[row]
}

func (c *Column) IsColumnFinish(isTest bool) bool {
	if !isTest {
		return len(c.remainList) == 0
	}
	return 8-len(c.remainList) == 2
}
