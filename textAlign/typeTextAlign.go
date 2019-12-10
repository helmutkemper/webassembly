package textAlign

import "strings"

type TextAlign int

func (el TextAlign) String() string {
	return textAlignString[el]
}

func (el TextAlign) ToType(value string) TextAlign {
	return textAlignStringStringToTextAlignStringMap[strings.ToLower(value)]
}

var textAlignString = [...]string{
	"start",
	"end",
	"left",
	"center",
	"right",
}

var textAlignStringStringToTextAlignStringMap = map[string]TextAlign{
	"start":  KStart,
	"end":    KEnd,
	"left":   KLeft,
	"center": KCenter,
	"right":  KRight,
}

const (
	KStart TextAlign = iota
	KEnd
	KLeft
	KCenter
	KRight
)
