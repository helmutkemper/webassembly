package font

import (
	"image/color"
	"strconv"
)

type Font struct {
	Size     float64
	SizeUnit string
	Color    color.RGBA
	Family   string
	Style    string
	Variant  string
	Weight   string
}

// en: Format the browser canvas font string as w3school rules
//     Note: browser canvas don't use color
//
// pt_br: Formata a string de fonte para elemento canvas no formato do w3school
//     Nota: o elemento canvas do navegador não usa cor
func (el *Font) String() string {
	return el.Style + " " + el.Variant + el.Weight + strconv.FormatFloat(el.Size, 'E', -1, 64) + el.SizeUnit + " " + el.Family
}
