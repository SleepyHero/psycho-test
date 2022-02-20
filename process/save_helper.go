package process

import (
	"fmt"
	"github.com/psycho-test/config"
	"io"
	"os"
	"time"
)

type SaveHelper struct {
	f         *os.File
	dis       string
	saveBatch []*SaveItem
	Index     int
}

func newFile(dis string) *os.File {
	ff, exist := openFile(dis)
	if !exist {
		writeHeader(ff)
	}
	return ff
}

func openFile(dis string) (*os.File, bool) {
	var ff *os.File
	filename := config.Dir + "/" + config.Num + "_" + dis + "_" + time.Now().Format("2006-01-02") + ".csv"
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

func writeHeader(f *os.File) {
	str := fmt.Sprintf("%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v\n",
		"用户编号", "性别", "年龄", "实验类型", "距离", "呈现区域", "trail序号", "字号", "文字区域",
		"该区域第几次出现", "反应时（毫秒）", "用户按键", "选择是否正确", "主观评价", "目标词", "目标词位置")
	writeFile(f, str)
}

func checkFileIsExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

func writeFile(f *os.File, str string) {
	_, err1 := io.WriteString(f, str) //写入文件(字符串)
	if err1 != nil {
		panic(err1)
	}
}

func NewSaveHelper(dis string) *SaveHelper {
	f := newFile(dis)
	saveBatch := make([]*SaveItem, 8, 8)
	for i := 0; i < 8; i++ {
		saveBatch[i] = &SaveItem{}
	}
	sh := &SaveHelper{
		f:         f,
		dis:       dis,
		saveBatch: saveBatch,
		Index:     0,
	}
	return sh
}

func (sh *SaveHelper) SaveRes(kind string, dis string, region string, trailNo int, fontSize int,
	targetZone string, showTimes int, reactTime int64, pressBtn string, suc bool,
	CommentNum string, targetWord string, targetNum int, number string, gender string, age string) {
	saveMsg := sh.saveBatch[sh.Index]
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
	sh.Index++
	sh.Index = sh.Index % 8
}

func (sh *SaveHelper) Push(commend string) {
	sh.Index = 0
	for _, v := range sh.saveBatch {
		str := fmt.Sprintf("%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v\n",
			v.number, v.gender, v.age, v.kind, v.dis, v.region, v.trailNo, v.fontSize, v.targetZone,
			v.showTimes, v.reactTime, v.pressBtn, v.suc, commend, v.targetWord, v.targetNum)
		writeFile(sh.f, str)
	}
}

func (sh *SaveHelper) CloseFile() {
	_ = sh.f.Close()
}
