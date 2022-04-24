package html

type FontStyleRule string

func (e FontStyleRule) String() string {
	return string(e)
}

const (
	// KFontStyleRuleNormal
	//
	// English:
	//
	//  Specifies the font style normal.
	//
	// Português:
	//
	//  Especifica o estilo de fonte normal.
	KFontStyleRuleNormal FontStyleRule = "normal"

	// KFontStyleRuleItalic
	//
	// English:
	//
	//  Specifies the font style italic.
	//
	// Português:
	//
	//  Especifica o estilo de fonte em itálico.
	KFontStyleRuleItalic FontStyleRule = "italic"

	// KFontStyleRuleOblique
	//
	// English:
	//
	//  Specifies the font style oblique.
	//
	// Português:
	//
	//  Especifica o estilo de fonte oblíquo.
	KFontStyleRuleOblique FontStyleRule = "oblique"
)
