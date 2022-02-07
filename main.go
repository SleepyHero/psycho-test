package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/psycho-test/config"
	"github.com/psycho-test/process"
	my_theme "github.com/psycho-test/theme"
	"time"
)

var (
	TotalLength = 1248
	TotalHeight = 216
)

func main() {
	//os.Setenv("FYNE_FONT", "/Users/fanbochao/Downloads/微软雅黑.ttf")
	theme.DefaultTextFont()
	a := app.New()
	my := &my_theme.MyTheme{}
	a.Settings().SetTheme(my)
	w := a.NewWindow("Hello")
	w.Resize(fyne.Size{
		Width:  config.ConfigData.TotalWidth + 5,
		Height: config.ConfigData.TotalHeight + 5,
	})
	w.CenterOnScreen()
	process.NewProcessHandler(w)
	w.Show()
	a.Run()
}

func updateTime(clock *widget.Label) {
	formatted := time.Now().Format("2006-01-02 15:04:05")
	clock.SetText(formatted)
}
