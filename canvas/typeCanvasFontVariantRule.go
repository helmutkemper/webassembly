package canvas

type CanvasFontVariantRule int

var CanvasFontVariantRules = [...]string{
	"",
	"normal",
	"small-caps",
}

func (el CanvasFontVariantRule) String() string {
	return CanvasFontVariantRules[el]
}

const (
	// en: Specifies the font variant normal.
	KFontVariantRuleNormal CanvasFontVariantRule = iota + 1

	// en: Specifies the font variant small-caps.
	KFontVariantRuleSmallCaps
)
