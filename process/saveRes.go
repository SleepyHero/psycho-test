package process

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
//saveBatch []*SaveItem
//Index     = 0
//f         *os.File
)

//func init() {
//	saveBatch = make([]*SaveItem, 8, 8)
//	for i := 0; i < 8; i++ {
//		saveBatch[i] = &SaveItem{}
//	}
//	ff, exist := openFile()
//	f = ff
//	if !exist {
//		writeHeader()
//	}
//}
//
//func SaveRes(kind string, dis string, region string, trailNo int, fontSize int,
//	targetZone string, showTimes int, reactTime int64, pressBtn string, suc bool,
//	CommentNum string, targetWord string, targetNum int, number string, gender string, age string) {
//	saveMsg := saveBatch[Index]
//	saveMsg.kind = kind
//	saveMsg.dis = dis
//	saveMsg.region = region
//	saveMsg.trailNo = trailNo
//	saveMsg.fontSize = fontSize
//	saveMsg.targetZone = targetZone
//	saveMsg.showTimes = showTimes
//	saveMsg.reactTime = reactTime
//	saveMsg.pressBtn = pressBtn
//	saveMsg.suc = suc
//	saveMsg.CommentNum = CommentNum
//	saveMsg.targetWord = targetWord
//	saveMsg.targetNum = targetNum
//	saveMsg.number = number
//	saveMsg.gender = gender
//	saveMsg.age = age
//	Index++
//	Index = Index % 8
//}
//
//func Push(commend string) {
//	Index = 0
//	for _, v := range saveBatch {
//		str := fmt.Sprintf("%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v\n",
//			v.number, v.gender, v.age, v.kind, v.dis, v.region, v.trailNo, v.fontSize, v.targetZone,
//			v.showTimes, v.reactTime, v.pressBtn, v.suc, commend, v.targetWord, v.targetNum)
//		writeFile(f, str)
//	}
//}
