package fontFamily

import "strings"

type FontFamily int

func (el FontFamily) String() string {
	return fontFamilyString[el]
}

func (el FontFamily) ToType(value string) FontFamily {
	return fontFamilyStringStringToFontFamilyStringMap[strings.ToLower(value)]
}

var fontFamilyString = [...]string{
	"Arial ",
	"'Arial Black' ",
	"'Book Antiqua' ",
	"Charcoal ",
	"'Comic Sans MS' ",
	"Courier ",
	"'Courier New' ",
	"cursive ",
	"Gadget ",
	"Geneva ",
	"Georgia ",
	"Helvetica ",
	"Impact ",
	"'Lucida Console' ",
	"'Lucida Grande' ",
	"'Lucida Sans Unicode' ",
	"Monaco ",
	"monospace ",
	"Palatino ",
	"'Palatino Linotype' ",
	"'sans-serif' ",
	"serif ",
	"Tahoma ",
	"Times ",
	"'Times New Roman' ",
	"'Trebuchet MS' ",
	"Verdana ",
}

var fontFamilyStringStringToFontFamilyStringMap = map[string]FontFamily{
	"Arial":               KArial,
	"Arial Black":         KArialBlack,
	"Book Antiqua":        KBookAntiqua,
	"Charcoal":            KCharcoal,
	"Comic Sans MS":       KComicSansMS,
	"Courier":             KCourier,
	"Courier New":         KCourierNew,
	"cursive":             KCursive,
	"Gadget":              KGadget,
	"Geneva":              KGeneva,
	"Georgia":             KGeorgia,
	"Helvetica":           KHelvetica,
	"Impact":              KImpact,
	"Lucida Console":      KLucidaConsole,
	"Lucida Grande":       KLucidaGrande,
	"Lucida Sans Unicode": KLucidaSansUnicode,
	"Monaco":              KMonaco,
	"monospace":           KMonospace,
	"Palatino":            KPalatino,
	"Palatino Linotype":   KPalatinoLinotype,
	"sans-serif":          KSansSerif,
	"serif":               KSerif,
	"Tahoma":              KTahoma,
	"Times":               KTimes,
	"Times New Roman":     KTimesNewRoman,
	"Trebuchet MS":        KTrebuchetMs,
	"Verdana":             KVerdana,
}

const (
	KArial FontFamily = iota
	KArialBlack
	KBookAntiqua
	KCharcoal
	KComicSansMS
	KCourier
	KCourierNew
	KCursive
	KGadget
	KGeneva
	KGeorgia
	KHelvetica
	KImpact
	KLucidaConsole
	KLucidaGrande
	KLucidaSansUnicode
	KMonaco
	KMonospace
	KPalatino
	KPalatinoLinotype
	KSansSerif
	KSerif
	KTahoma
	KTimes
	KTimesNewRoman
	KTrebuchetMs
	KVerdana
)
