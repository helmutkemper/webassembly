package rulesIcon

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/examples/ide/rulesDensity"
	"image/color"
	"reflect"
)

const (
	KPipeLineNormal int = iota
	KPipeLineDisabled
	KPipeLineSelected
	KPipeLineAttention1
	KPipeLineAttention2
	KPipeLineAlert
)

type Data struct {
	Status          int
	IconX           rulesDensity.Density
	IconY           rulesDensity.Density
	IconWidth       rulesDensity.Density
	IconHeight      rulesDensity.Density
	IconViewBox     []int
	Label           string
	LabelFontSize   rulesDensity.Density
	LabelY          rulesDensity.Density
	Path            string
	ColorIcon       color.RGBA
	ColorBorder     color.RGBA
	ColorLabel      color.RGBA
	ColorBackground color.RGBA
	Name            string
	Category        string
}

var BorderColor = color.RGBA{R: 0x5F, G: 0x5F, B: 0x5F, A: 0xFF}
var BorderColorAttention2 = color.RGBA{R: 255, G: 0, B: 0, A: 255}
var FillColor = color.RGBA{R: 180, G: 180, B: 255, A: 255}
var TextColor = color.RGBA{R: 0, G: 0, B: 0, A: 255}
var TextColorSelected = color.RGBA{R: 128, G: 0, B: 0, A: 255}
var TextColorAttention2 = color.RGBA{R: 255, G: 180, B: 180, A: 255}
var TextColorDisabled = color.RGBA{R: 0, G: 0, B: 0, A: 0x6f}
var CategoryIconColorSelected = color.RGBA{R: 128, G: 0, B: 0, A: 255}
var CategoryIconColor = color.RGBA{R: 0xf8, G: 0xf8, B: 0xef, A: 0xff}
var CategoryIconColorAttention1 = color.RGBA{R: 255, G: 180, B: 180, A: 255}
var CategoryIconColorDisabled = color.RGBA{R: 0xb4, G: 0xb4, B: 0xff, A: 0x6f}
var BorderWidth = rulesDensity.Density(4)
var TextY = rulesDensity.Density(160)
var FontFamily = "Helvetica"
var FontWeight = "normal"
var FontStyle = "normal"
var FontSize = rulesDensity.Density(20)
var Width = rulesDensity.Density(200)
var Height = rulesDensity.Density(200)
var SizeRatio = rulesDensity.Density(0.5)
var FilterIcon *html.TagSvgFilter
var FilterText *html.TagSvgFilter
var IconX = rulesDensity.Density(60)
var IconY = rulesDensity.Density(40)
var IconWidth = rulesDensity.Density(80)
var IconHeight = rulesDensity.Density(80)

func init() {
	FilterIcon = factoryBrowser.NewTagSvgFilter().Id("iconBlur").Append(
		//factoryBrowser.NewTagSvgFeOffset().Dx(1).Dy(1),
		factoryBrowser.NewTagSvgFeBlend().In2(html.KSvgIn2SourceAlpha),
		factoryBrowser.NewTagSvgFeGaussianBlur().StdDeviation(5).In(html.KSvgInStrokePaint),
	)
	FilterText = factoryBrowser.NewTagSvgFilter().Id("textBlur").Append(
		//factoryBrowser.NewTagSvgFeOffset().Dx(1).Dy(1),
		factoryBrowser.NewTagSvgFeBlend().In2(html.KSvgIn2SourceAlpha),
		factoryBrowser.NewTagSvgFeGaussianBlur().StdDeviation(0.5).In(html.KSvgInFillPaint),
	)
}

func DataVerifySystemIcon(data Data) Data {
	if data.IconViewBox == nil {
		data.IconViewBox = []int{0, 0, 512, 512}
	}

	if data.IconX == 0 {
		data.IconX = IconX
	}

	if data.IconY == 0 {
		data.IconY = IconY
	}

	if data.IconWidth == 0 {
		data.IconWidth = IconWidth
	}

	if data.IconHeight == 0 {
		data.IconHeight = IconHeight
	}

	if data.LabelFontSize == 0 {
		data.LabelFontSize = FontSize
	}

	if data.LabelY == 0 {
		data.LabelY = TextY
	}

	switch data.Status {
	case KPipeLineDisabled:
		if reflect.DeepEqual(data.ColorIcon, color.RGBA{}) {
			data.ColorIcon = CategoryIconColorDisabled
		}
	case KPipeLineSelected:
		if reflect.DeepEqual(data.ColorIcon, color.RGBA{}) {
			data.ColorIcon = CategoryIconColorSelected
		}
	default:
		if reflect.DeepEqual(data.ColorIcon, color.RGBA{}) {
			data.ColorIcon = FillColor
		}
	}

	switch data.Status {
	case KPipeLineAttention2:
		if reflect.DeepEqual(data.ColorBorder, color.RGBA{}) {
			data.ColorBorder = BorderColorAttention2
		}
	default:
		if reflect.DeepEqual(data.ColorBorder, color.RGBA{}) {
			data.ColorBorder = BorderColor
		}
	}

	switch data.Status {
	case KPipeLineDisabled:
		if reflect.DeepEqual(data.ColorLabel, color.RGBA{}) {
			data.ColorLabel = TextColorDisabled
		}
	case KPipeLineSelected:
		if reflect.DeepEqual(data.ColorLabel, color.RGBA{}) {
			data.ColorLabel = TextColorSelected
		}
	case KPipeLineAttention2:
		if reflect.DeepEqual(data.ColorLabel, color.RGBA{}) {
			data.ColorLabel = TextColorAttention2
		}
	default:
		if reflect.DeepEqual(data.ColorLabel, color.RGBA{}) {
			data.ColorLabel = TextColor
		}
	}

	switch data.Status {
	case KPipeLineAttention1:
		if reflect.DeepEqual(data.ColorBackground, color.RGBA{}) {
			data.ColorBackground = CategoryIconColorAttention1
		}
	default:
		if reflect.DeepEqual(data.ColorBackground, color.RGBA{}) {
			data.ColorBackground = CategoryIconColor
		}
	}

	return data
}

func DataVerifyElementIcon(data Data) Data {
	if data.IconViewBox == nil {
		data.IconViewBox = []int{0, 0, 512, 512}
	}

	if data.IconX == 0 {
		data.IconX = IconX
	}

	if data.IconY == 0 {
		data.IconY = IconY
	}

	if data.IconWidth == 0 {
		data.IconWidth = IconWidth
	}

	if data.IconHeight == 0 {
		data.IconHeight = IconHeight
	}

	if data.LabelFontSize == 0 {
		data.LabelFontSize = FontSize
	}

	if data.LabelY == 0 {
		data.LabelY = TextY
	}

	switch data.Status {
	case KPipeLineDisabled:
		if reflect.DeepEqual(data.ColorIcon, color.RGBA{}) {
			data.ColorIcon = CategoryIconColorDisabled
		}
	case KPipeLineSelected:
		if reflect.DeepEqual(data.ColorIcon, color.RGBA{}) {
			data.ColorIcon = CategoryIconColorSelected
		}
	default:
		if reflect.DeepEqual(data.ColorIcon, color.RGBA{}) {
			data.ColorIcon = FillColor
		}
	}

	switch data.Status {
	case KPipeLineSelected:
		if reflect.DeepEqual(data.ColorBorder, color.RGBA{}) {
			data.ColorBorder = color.RGBA{R: 255, G: 0, B: 0, A: 255}
		}
	case KPipeLineAttention2:
		if reflect.DeepEqual(data.ColorBorder, color.RGBA{}) {
			data.ColorBorder = color.RGBA{R: 255, G: 0, B: 0, A: 255}
		}
	default:
		if reflect.DeepEqual(data.ColorBorder, color.RGBA{}) {
			data.ColorBorder = BorderColor
		}
	}

	switch data.Status {
	case KPipeLineDisabled:
		if reflect.DeepEqual(data.ColorLabel, color.RGBA{}) {
			data.ColorLabel = TextColorDisabled
		}
	case KPipeLineSelected:
		if reflect.DeepEqual(data.ColorLabel, color.RGBA{}) {
			data.ColorLabel = color.RGBA{R: 255, G: 0, B: 0, A: 255}
		}
	case KPipeLineAttention2:
		if reflect.DeepEqual(data.ColorLabel, color.RGBA{}) {
			data.ColorLabel = color.RGBA{R: 255, G: 0, B: 0, A: 255}
		}
	default:
		if reflect.DeepEqual(data.ColorLabel, color.RGBA{}) {
			data.ColorLabel = TextColor
		}
	}

	switch data.Status {
	case KPipeLineNormal:
		if reflect.DeepEqual(data.ColorBackground, color.RGBA{}) {
			data.ColorBackground = color.RGBA{R: 220, G: 220, B: 255, A: 255}
		}
	case KPipeLineDisabled:
		if reflect.DeepEqual(data.ColorBackground, color.RGBA{}) {
			data.ColorBackground = color.RGBA{R: 230, G: 230, B: 230, A: 255}
		}
	case KPipeLineSelected:
		if reflect.DeepEqual(data.ColorBackground, color.RGBA{}) {
			data.ColorBackground = color.RGBA{R: 255, G: 220, B: 220, A: 255}
		}
	case KPipeLineAttention1:
		if reflect.DeepEqual(data.ColorBackground, color.RGBA{}) {
			data.ColorBackground = color.RGBA{R: 255, G: 180, B: 180, A: 255}
		}
	default:
		if reflect.DeepEqual(data.ColorBackground, color.RGBA{}) {
			data.ColorBackground = CategoryIconColor
		}
	}

	return data
}
