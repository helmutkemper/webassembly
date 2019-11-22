package canvas

type CanvasFontAlignRule int

var CanvasFontAlignRules = [...]string{
	"",
	"start",
	"end",
	"center",
	"left",
	"right",
}

func (el CanvasFontAlignRule) String() string {
	return CanvasFontAlignRules[el]
}

const (
	// en: Default. The text starts at the specified position
	KFontAlignRuleStart CanvasFontAlignRule = iota + 1

	// en: The text ends at the specified position
	KFontAlignRuleEnd

	// en: The center of the text is placed at the specified position
	KFontAlignRuleCenter

	// en: The text starts at the specified position
	KFontAlignRuleLeft

	// en: The text ends at the specified position
	KFontAlignRuleRight
)
