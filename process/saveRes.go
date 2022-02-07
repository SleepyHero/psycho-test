package process

import (
	"fmt"
	"github.com/psycho-test/config"
	"io"
	"os"
)

type SaveItem struct {
	kind       string
	dis        string
	region     string
	trailNo    int
	fontSize   int
	targetZone string
	showTimes  int
	reactTime  int64
	pressBtn   string
	suc        bool
	CommentNum string
	targetWord string
	targetNum  int
	number     string
	gender     string
	age        string
}

var (
	saveBatch []*SaveItem
	Index     = 0
	f         *os.File
)

func init() {
	saveBatch = make([]*SaveItem, 8, 8)
	for i := 0; i < 8; i++ {
		saveBatch[i] = &SaveItem{}
	}
	ff, exist := openFile()
	f = ff
	if !exist {
		writeHeader()
	}
}

func SaveRes(kind string, dis string, region string, trailNo int, fontSize int,
	targetZone string, showTimes int, reactTime int64, pressBtn string, suc bool,
	CommentNum string, targetWord string, targetNum int, number string, gender string, age string) {
	saveMsg := saveBatch[Index]
	saveMsg.kind = kind
	saveMsg.dis = dis
	saveMsg.region = region
	saveMsg.trailNo = trailNo
	saveMsg.fontSize = fontSize
	saveMsg.targetZone = targetZone
	saveMsg.showTimes = showTimes
	saveMsg.reactTime = reactTime
	saveMsg.pressBtn = pressBtn
	saveMsg.suc = suc
	saveMsg.CommentNum = CommentNum
	saveMsg.targetWord = targetWord
	saveMsg.targetNum = targetNum
	saveMsg.number = number
	saveMsg.gender = gender
	saveMsg.age = age
	Index++
	Index = Index % 8
}

func checkFileIsExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

func Push(commend string) {
	Index = 0
	for _, v := range saveBatch {
		str := fmt.Sprintf("%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v\n",
			v.number, v.gender, v.age, v.kind, v.dis, v.region, v.trailNo, v.fontSize, v.targetZone,
			v.showTimes, v.reactTime, v.pressBtn, v.suc, commend, v.targetWord, v.targetNum)
		writeFile(f, str)
	}
}

func writeHeader() {
	str := fmt.Sprintf("%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v\n",
		"用户编号", "性别", "年龄", "实验类型", "距离", "呈现区域", "trail序号", "字号", "文字区域",
		"该区域第几次出现", "反应时（毫秒）", "用户按键", "选择是否正确", "主观评价", "目标词", "目标词位置")
	writeFile(f, str)
}

func openFile() (*os.File, bool) {
	var ff *os.File
	filename := config.Dir + "/res.csv"
	exist := false
	if checkFileIsExist(filename) { //如果文件存在
		ff, _ = os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0666) //打开文件
		fmt.Println("文件存在")
		exist = true
	} else {
		ff, _ = os.Create(filename) //创建文件
		fmt.Println("文件不存在")
		exist = false
	}
	return ff, exist
}

func writeFile(f *os.File, str string) {
	_, err1 := io.WriteString(f, str) //写入文件(字符串)
	if err1 != nil {
		panic(err1)
	}
}
