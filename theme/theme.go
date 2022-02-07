package theme

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"github.com/psycho-test/config"
	"image/color"
)

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
		return color.Gray{100}
	case theme.ColorNameForeground:
		return color.White
	}
	return theme.DefaultTheme().Color(colorName, themeVariant)
}
func (*MyTheme) Font(style fyne.TextStyle) fyne.Resource {
	res, _ := fyne.LoadResourceFromPath(fmt.Sprintf("%s/font.ttf", config.Dir))
	return res
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
