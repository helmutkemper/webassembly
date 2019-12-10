package fontWeight

import "strings"

type FontWeight int

func (el FontWeight) String() string {
	return fontWeightString[el]
}

func (el FontWeight) ToType(value string) FontWeight {
	return fontWeightStringStringToFontWeightStringMap[strings.ToLower(value)]
}

var fontWeightString = [...]string{
	"",
	"Bold ",
	"Bolder ",
	"Lighter ",
	"100 ",
	"200 ",
	"300 ",
	"400 ",
	"500 ",
	"600 ",
	"700 ",
	"800 ",
	"900 ",
}

var fontWeightStringStringToFontWeightStringMap = map[string]FontWeight{
	"normal":  KNormal,
	"bold":    KBold,
	"bolder":  KBolder,
	"lighter": KLighter,
	"100":     K100,
	"200":     K200,
	"300":     K300,
	"400":     K400,
	"500":     K500,
	"600":     K600,
	"700":     K700,
	"800":     K800,
	"900":     K900,
}

const (
	KNormal FontWeight = iota
	KBold
	KBolder
	KLighter
	K100
	K200
	K300
	K400
	K500
	K600
	K700
	K800
	K900
)
