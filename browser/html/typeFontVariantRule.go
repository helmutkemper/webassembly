package html

type FontVariantRule string

func (e FontVariantRule) String() string {
	return string(e)
}

const (
	// KFontVariantRuleNormal
	//
	// English:
	//
	//  Specifies the font variant normal.
	//
	// Português:
	//
	//  Especifica a variante de fonte normal.
	KFontVariantRuleNormal FontVariantRule = "normal"

	// KFontVariantRuleSmallCaps
	//
	// English:
	//
	//  Specifies the font variant small-caps.
	//
	// Português:
	//
	//  Especifica as letras minúsculas da variante da fonte.
	KFontVariantRuleSmallCaps FontVariantRule = "small-caps"
)
