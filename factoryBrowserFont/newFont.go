package factoryBrowserFont

import (
	"github.com/helmutkemper/iotmaker.platform.webbrowser/font"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/fontFamily"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/fontStyle"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/fontVariant"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/fontWeight"
)

func NewFont(size int, sizeUnit string, family fontFamily.FontFamily, style fontStyle.FontStyle, variant fontVariant.FontVariant, weight fontWeight.FontWeight) string {
	f := font.Font{
		Size:     size,
		SizeUnit: sizeUnit,
		Family:   family,
		Style:    style,
		Variant:  variant,
		Weight:   weight,
	}
	return f.String()
}
