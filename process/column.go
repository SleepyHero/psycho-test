package process

import (
	"github.com/psycho-test/config"
	"math/rand"
	"time"
)

const RowNum = 4

func init() {
	rand.Seed(time.Now().Unix())
}

type Column struct {
	remainList []int
	x          float32
	timesCount []int
}

func NewColumn(index int) *Column {
	remainLen := RowNum * config.ConfigData.RepeatTimes
	list := make([]int, remainLen, remainLen)
	for i := 0; i < remainLen; i++ {
		list[i] = i
	}
	c := &Column{remainList: list, x: float32(index) * (config.ConfigData.TargetWidth + config.TargetDis)}
	c.timesCount = make([]int, RowNum, RowNum)
	return c
}

func (c *Column) GetNewPos() (float32, float32, int, int) {
	n := rand.Intn(len(c.remainList))
	num := c.remainList[n]
	c.remainList = append(c.remainList[:n], c.remainList[n+1:]...)
	row := num / config.ConfigData.RepeatTimes
	c.timesCount[row]++
	return c.x, float32(row)*config.ConfigData.TargetHeight + config.BaseHeight, row, c.timesCount[row]
}

func (c *Column) IsColumnFinish(isTest bool) bool {
	if !isTest {
		return len(c.remainList) == 0
	}
	//测试只取两个位置
	return RowNum*config.ConfigData.RepeatTimes-len(c.remainList) == 2
}
