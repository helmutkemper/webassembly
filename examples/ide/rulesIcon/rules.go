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
var BorderWidth = rulesDensity.Convert(4)
var TextY = rulesDensity.Convert(160)
var FontFamily = "Helvetica"
var FontSize = rulesDensity.Convert(20)
var Width = rulesDensity.Convert(200)
var Height = rulesDensity.Convert(200)
var SizeRatio = 0.5
var FilterIcon *html.TagSvgFilter
var FilterText *html.TagSvgFilter

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
