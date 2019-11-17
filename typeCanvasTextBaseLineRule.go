package iotmaker_platform_webbrowser

type CanvasTextBaseLineRule int

var CanvasTextBaseLineRules = [...]string{
	"",
	"alphabetic",
	"top",
	"hanging",
	"middle",
	"ideographic",
	"bottom",
}

func (el CanvasTextBaseLineRule) String() string {
	return CanvasTextBaseLineRules[el]
}

const (
	// en: Default. The text baseline is the normal alphabetic baseline
	KTextBaseLineRuleAlphabetic CanvasTextBaseLineRule = iota + 1

	// en: The text baseline is the top of the em square
	KTextBaseLineRuleTop

	// en: The text baseline is the hanging baseline
	KTextBaseLineRuleHanging

	// en: The text baseline is the middle of the em square
	KTextBaseLineRuleMiddle

	// en: The text baseline is the ideographic baseline
	KTextBaseLineRuleIdeographic

	// en: The text baseline is the bottom of the bounding box
	KTextBaseLineRuleBottom
)
