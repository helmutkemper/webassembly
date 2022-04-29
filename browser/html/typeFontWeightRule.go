package html

type FontWeightRule string

func (e FontWeightRule) String() string {
	return string(e)
}

const (
	// KFontWeightRuleNormal
	//
	// English:
	//
	//  Specifies the font weight normal.
	//
	// Português:
	//
	//  Especifica o peso da fonte normal.
	KFontWeightRuleNormal FontWeightRule = "normal"

	// KFontWeightRuleBold
	//
	// English:
	//
	//  Specifies the font weight bold.
	//
	// Português:
	//
	//  Especifica a espessura da fonte em negrito.
	KFontWeightRuleBold FontWeightRule = "bold"

	// KFontWeightRuleBolder
	//
	// English:
	//
	//  Specifies the font weight bolder.
	//
	// Português:
	//
	//  Especifica o peso da fonte em negrito.
	KFontWeightRuleBolder FontWeightRule = "bolder"

	// KFontWeightRuleLighter
	//
	// English:
	//
	//  Specifies the font weight lighter.
	//
	// Português:
	//
	//  Especifica o peso da fonte mais leve.
	KFontWeightRuleLighter FontWeightRule = "lighter"

	// KFontWeightRule100
	//
	// English:
	//
	//  Specifies the font weight 100.
	//
	// Português:
	//
	//  Especifica o peso da fonte 100.
	KFontWeightRule100 FontWeightRule = "100"

	// KFontWeightRule200
	//
	// English:
	//
	//  Specifies the font weight 200.
	//
	// Português:
	//
	//  Especifica o peso da fonte 200.
	KFontWeightRule200 FontWeightRule = "200"

	// KFontWeightRule300
	//
	// English:
	//
	//  Specifies the font weight 300.
	//
	// Português:
	//
	//  Especifica o peso da fonte 300.
	KFontWeightRule300 FontWeightRule = "300"

	// KFontWeightRule400
	//
	// English:
	//
	//  Specifies the font weight 400.
	//
	// Português:
	//
	//  Especifica o peso da fonte 400.
	KFontWeightRule400 FontWeightRule = "400"

	// KFontWeightRule500
	//
	// English:
	//
	//  Specifies the font weight 500.
	//
	// Português:
	//
	//  Especifica o peso da fonte 500.
	KFontWeightRule500 FontWeightRule = "500"

	// KFontWeightRule600
	//
	// English:
	//
	//  Specifies the font weight 600.
	//
	// Português:
	//
	//  Especifica o peso da fonte 600.
	KFontWeightRule600 FontWeightRule = "600"

	// KFontWeightRule700
	//
	// English:
	//
	//  Specifies the font weight 700.
	//
	// Português:
	//
	//  Especifica o peso da fonte 700.
	KFontWeightRule700 FontWeightRule = "700"

	// KFontWeightRule800
	//
	// English:
	//
	//  Specifies the font weight 800.
	//
	// Português:
	//
	//  Especifica o peso da fonte 800.
	KFontWeightRule800 FontWeightRule = "800"

	// KFontWeightRule900
	//
	// English:
	//
	//  Specifies the font weight 900.
	//
	// Português:
	//
	//  Especifica o peso da fonte 900.
	KFontWeightRule900 FontWeightRule = "900"
)
