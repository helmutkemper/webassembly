package font

import (
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/fontFamily"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/fontStyle"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/fontVariant"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/fontWeight"
	"image/color"
	"log"
	"strconv"
)

type Font struct {
	Size     float64
	SizeUnit string
	Color    color.RGBA
	Family   interface{}
	Style    fontStyle.FontStyle
	Variant  fontVariant.FontVariant
	Weight   fontWeight.FontWeight
}

// en: Format the browser canvas font string as w3school rules
//     Note: browser canvas don't use color
//
// pt_br: Formata a string de fonte para elemento canvas no formato do w3school
//     Nota: o elemento canvas do navegador não usa cor
func (el *Font) String() string {
	switch converted := el.Family.(type) {
	case fontFamily.FontFamily:
		return el.Style.String() + " " + el.Variant.String() + el.Weight.String() + strconv.FormatFloat(el.Size, 'E', -1, 64) + el.SizeUnit + " " + converted.String()

	case string:
		return el.Style.String() + " " + el.Variant.String() + el.Weight.String() + strconv.FormatFloat(el.Size, 'E', -1, 64) + el.SizeUnit + " " + converted
	}

	log.Fatalf("error: font.Font.Family must be a string or a fontFamily.FontFamily constant")
	return "error: font.Font.Family must be a string or a fontFamily.FontFamily constant"
}
