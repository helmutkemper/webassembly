package iotmaker_platform_webbrowser

type CanvasFontWeightRule int

var CanvasFontWeightRules = [...]string{
	"",
	"normal",
	"bold",
	"bolder",
	"lighter",
	"100",
	"200",
	"300",
	"400",
	"500",
	"600",
	"700",
	"800",
	"900",
}

func (el CanvasFontWeightRule) String() string {
	return CanvasFontWeightRules[el]
}

const (
	// en: Specifies the font weight normal.
	KFontWeightRuleNormal CanvasFontWeightRule = iota + 1

	// en: Specifies the font weight bold.
	KFontWeightRuleBold

	// en: Specifies the font weight bolder.
	KFontWeightRuleBolder

	// en: Specifies the font weight lighter.
	KFontWeightRuleLighter

	// en: Specifies the font weight 100.
	KFontWeightRule100

	// en: Specifies the font weight 200.
	KFontWeightRule200

	// en: Specifies the font weight 300.
	KFontWeightRule300

	// en: Specifies the font weight 400.
	KFontWeightRule400

	// en: Specifies the font weight 500.
	KFontWeightRule500

	// en: Specifies the font weight 600.
	KFontWeightRule600

	// en: Specifies the font weight 700.
	KFontWeightRule700

	// en: Specifies the font weight 800.
	KFontWeightRule800

	// en: Specifies the font weight 900.
	KFontWeightRule900
)
