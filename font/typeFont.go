package font

import (
	"github.com/helmutkemper/iotmaker.platform.webbrowser/fontFamily"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/fontStyle"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/fontVariant"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/fontWeight"
	"strconv"
)

type Font struct {
	Size     int
	SizeUnit string
	Family   fontFamily.FontFamily
	Style    fontStyle.FontStyle
	Variant  fontVariant.FontVariant
	Weight   fontWeight.FontWeight
}

func (el *Font) String() string {
	return el.Style.String() + " " + el.Variant.String() + el.Weight.String() + strconv.Itoa(el.Size) + el.SizeUnit + " " + el.Family.String()
}
