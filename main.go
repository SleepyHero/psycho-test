package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/psycho-test/config"
	"github.com/psycho-test/process"
	myTheme "github.com/psycho-test/theme"
	"image/color"
	"time"
)

var (
	TotalLength = 1248
	TotalHeight = 216
)

func main() {
	//os.Setenv("FYNE_FONT", "/Users/fanbochao/Downloads/微软雅黑.ttf")
	a := app.New()
	my := &MyTheme{}
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

var (
	_ fyne.Theme = (*MyTheme)(nil)
)
var (
	bgColor = color.NRGBA{R: 15, G: 28, B: 47, A: 255}
	green   = color.NRGBA{R: 0, G: 180, B: 0, A: 255}
	red     = color.NRGBA{R: 180, G: 0, B: 0, A: 255}
)

type MyTheme struct {
}

func (*MyTheme) Color(colorName fyne.ThemeColorName, themeVariant fyne.ThemeVariant) color.Color {
	switch colorName {
	case theme.ColorNameBackground:
		return bgColor
	case theme.ColorNameButton:
		return red
	case theme.ColorNameShadow:
		return green
	case theme.ColorNameInputBackground:
		return color.Gray{Y: 100}
	case theme.ColorNameForeground:
		return color.White
	}
	return theme.DefaultTheme().Color(colorName, themeVariant)
}
func (*MyTheme) Font(style fyne.TextStyle) fyne.Resource {
	//res, _ := fyne.LoadResourceFromPath(fmt.Sprintf("%s/font.ttf", config.Dir))
	return myTheme.ResourceFontTtf
	//return theme.DefaultTheme().Font(style)
}
func (*MyTheme) Icon(iconName fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(iconName)
}
func (*MyTheme) Size(sizeName fyne.ThemeSizeName) float32 {
	switch sizeName {
	case theme.SizeNameText:
		return 12
	case theme.SizeNameCaptionText:
		return 12
	}
	return theme.DefaultTheme().Size(sizeName)
}
