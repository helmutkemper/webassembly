package canvas

type CanvasFontStyleRule int

var CanvasFontStyleRules = [...]string{
	"",
	"normal",
	"italic",
	"oblique",
}

func (el CanvasFontStyleRule) String() string {
	return CanvasFontStyleRules[el]
}

const (
	// en: 	Specifies the font style normal.
	KFontStyleRuleNormal CanvasFontStyleRule = iota + 1

	// en: 	Specifies the font style italic.
	KFontStyleRuleItalic

	// en: 	Specifies the font style oblique.
	KFontStyleRuleOblique
)
