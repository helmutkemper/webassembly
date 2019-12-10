package fontVariant

import "strings"

type FontVariant int

func (el FontVariant) String() string {
	return fontVariantString[el]
}

func (el FontVariant) ToType(value string) FontVariant {
	return fontVariantStringStringToFontVariantStringMap[strings.ToLower(value)]
}

var fontVariantString = [...]string{
	"",
	"Small-Caps ",
}

var fontVariantStringStringToFontVariantStringMap = map[string]FontVariant{
	"normal":     KNormal,
	"small-caps": KSmallCaps,
}

const (
	KNormal FontVariant = iota
	KSmallCaps
)
