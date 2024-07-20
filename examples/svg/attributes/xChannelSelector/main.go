// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/xChannelSelector
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/xChannelSelector
//
//  <svg viewBox="0 0 440 160" xmlns="http://www.w3.org/2000/svg">
//    <filter id="displacementFilter">
//      <feImage xlink:href="https://yari-demos.prod.mdn.mozit.cloud/en-US/docs/Web/SVG/Attribute/xChannelSelector/mdn.svg"
//           x="0" y="0" width="100%" height="100%" result="abc"></feImage>
//      <feDisplacementMap in2="abc" in="SourceGraphic" scale="30" xChannelSelector="R"></feDisplacementMap>
//    </filter>
//    <filter id="displacementFilter2">
//      <feImage xlink:href="https://yari-demos.prod.mdn.mozit.cloud/en-US/docs/Web/SVG/Attribute/xChannelSelector/mdn.svg"
//           x="0" y="0" width="100%" height="100%" result="abc"></feImage>
//      <feDisplacementMap in2="abc" in="SourceGraphic" scale="30" xChannelSelector="B"></feDisplacementMap>
//    </filter>
//
//    <text x="10" y="60" font-size="50" filter="url(#displacementFilter)">Some displaced text</text>
//    <text x="10" y="120" font-size="50" filter="url(#displacementFilter2)">Some displaced text</text>
//  </svg>

//go:build js

package main

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/html"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 440, 160}).Append(
		factoryBrowser.NewTagSvgFilter().Id("displacementFilter").Append(
			factoryBrowser.NewTagSvgFeImage().HRef("https://yari-demos.prod.mdn.mozit.cloud/en-US/docs/Web/SVG/Attribute/xChannelSelector/mdn.svg").
				X(0).Y(0).Width(float32(1)).Height(float32(1)).Result("abc"),
			factoryBrowser.NewTagSvgFeDisplacementMap().In2("abc").In(html.KSvgInSourceGraphic).
				Scale(30).XChannelSelector(html.KSvgChannelSelectorR),
		),
		factoryBrowser.NewTagSvgFilter().Id("displacementFilter2").Append(
			factoryBrowser.NewTagSvgFeImage().HRef("https://yari-demos.prod.mdn.mozit.cloud/en-US/docs/Web/SVG/Attribute/xChannelSelector/mdn.svg").
				X(0).Y(0).Width(float32(1)).Height(float32(1)).Result("abc"),
			factoryBrowser.NewTagSvgFeDisplacementMap().In2("abc").In(html.KSvgInSourceGraphic).
				Scale(30).XChannelSelector(html.KSvgChannelSelectorB),
		),

		factoryBrowser.NewTagSvgText().X(10).Y(60).FontSize(50).
			Filter("url(#displacementFilter)").Text("Some displaced text"),
		factoryBrowser.NewTagSvgText().X(10).Y(120).FontSize(50).
			Filter("url(#displacementFilter2)").Text("Some displaced text"),
	)

	stage.Append(s1)

	<-done
}

//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
