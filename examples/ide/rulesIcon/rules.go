package rulesIcon

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/examples/ide/rulesDensity"
	"image/color"
)

var BorderColor = color.RGBA{R: 0x5F, G: 0x5F, B: 0x5F, A: 0xFF}
var FillColor = color.RGBA{R: 180, G: 180, B: 255, A: 255}
var TextColor = color.RGBA{R: 0, G: 0, B: 0, A: 255}
var TextColorSelected = color.RGBA{R: 128, G: 0, B: 0, A: 255}
var TextColorDisabled = color.RGBA{R: 0, G: 0, B: 0, A: 0x8f}
var CategoryIconColorSelected = color.RGBA{R: 128, G: 0, B: 0, A: 255}
var CategoryIconColor = color.RGBA{R: 0xf8, G: 0xf8, B: 0xef, A: 0xff}
var CategoryIconColorDisabled = color.RGBA{R: 0xb4, G: 0xb4, B: 0xff, A: 0x8f}
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
