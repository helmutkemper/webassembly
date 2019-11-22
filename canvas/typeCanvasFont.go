package canvas

import (
	iotmaker_types "github.com/helmutkemper/iotmaker.types"
	"strings"
)

type Font struct {
	// en: Specifies the font style.
	Style CanvasFontStyleRule

	// en: Specifies the font variant.
	Variant CanvasFontVariantRule

	// en: Specifies the font weight.
	Weight CanvasFontWeightRule

	// en: Specifies the font size and the line-height, in pixels
	Size iotmaker_types.Pixel

	// en: Specifies the font family
	Family string

	// en: Use the font captioned controls (like buttons, drop-downs, etc.)
	Caption string

	// en: 	Use the font used to label icons
	Icon string

	// en: Use the font used in menus (drop-down menus and menu lists)
	Menu string

	// en: Use the font used in dialog boxes
	MessageBox string

	// en: Use the font used for labeling small controls
	SmallCaption string

	// en: Use the fonts used in window status bar
	StatusBar string
}

func (el *Font) String() string {
	var ret string
	if el.Style != 0 {
		ret += el.Style.String()
		ret += " "
	}

	if el.Variant != 0 {
		ret += el.Variant.String()
		ret += " "
	}

	if el.Weight != 0 {
		ret += el.Weight.String()
		ret += " "
	}

	if el.Size != 0 {
		ret += el.Size.String()
		ret += " "
	}

	if el.Family != "" {
		ret += el.Family
		ret += " "
	}

	if el.Caption != "" {
		ret += el.Caption
		ret += " "
	}

	if el.Icon != "" {
		ret += el.Icon
		ret += " "
	}

	if el.Menu != "" {
		ret += el.Menu
		ret += " "
	}

	if el.MessageBox != "" {
		ret += el.MessageBox
		ret += " "
	}

	if el.SmallCaption != "" {
		ret += el.SmallCaption
		ret += " "
	}

	if el.StatusBar != "" {
		ret += el.StatusBar
		ret += " "
	}

	return strings.TrimSpace(ret)
}
