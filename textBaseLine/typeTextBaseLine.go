package textBaseLine

import "strings"

type TextBaseLine int

func (el TextBaseLine) String() string {
	return textBaseLineString[el]
}

func (el TextBaseLine) ToType(value string) TextBaseLine {
	return textBaseLineStringStringToTextBaseLineStringMap[strings.ToLower(value)]
}

var textBaseLineString = [...]string{
	"KAlphabetic",
	"KTop",
	"KHanging",
	"KMiddle",
	"KIdeographic",
	"KBottom",
}

var textBaseLineStringStringToTextBaseLineStringMap = map[string]TextBaseLine{
	"KAlphabetic":  KAlphabetic,
	"KTop":         KTop,
	"KHanging":     KHanging,
	"KMiddle":      KMiddle,
	"KIdeographic": KIdeographic,
	"KBottom":      KBottom,
}

const (
	// KAlphabetic
	// en: Default. The text baseline is the normal alphabetic baseline
	KAlphabetic TextBaseLine = iota

	// KTop
	// en: The text baseline is the top of the em square
	KTop

	// KHanging
	// en: The text baseline is the hanging baseline
	KHanging

	// KMiddle
	// en: The text baseline is the middle of the em square
	KMiddle

	// KIdeographic
	// en: The text baseline is the ideographic baseline
	KIdeographic

	// KBottom
	// en: The text baseline is the bottom of the bounding box
	KBottom
)
