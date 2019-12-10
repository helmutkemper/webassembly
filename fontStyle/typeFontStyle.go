package fontStyle

import "strings"

type FontStyle int

func (el FontStyle) String() string {
	return fontStyleString[el]
}

func (el FontStyle) ToType(value string) FontStyle {
	return fontStyleStringStringToFontStyleStringMap[strings.ToLower(value)]
}

var fontStyleString = [...]string{
	"",
	"Italic ",
	"Oblique ",
}

var fontStyleStringStringToFontStyleStringMap = map[string]FontStyle{
	"normal":  KNormal,
	"italic":  KItalic,
	"oblique": KOblique,
}

const (
	KNormal FontStyle = iota
	KItalic
	KOblique
)
