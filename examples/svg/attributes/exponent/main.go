// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/exponent
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/exponent
//
//  <svg viewBox="0 0 420 200" xmlns="http://www.w3.org/2000/svg">
//    <defs>
//      <linearGradient id="gradient" gradientUnits="userSpaceOnUse"
//          x1="0" y1="0" x2="200" y2="0">
//        <stop offset="0" stop-color="#ff0000" />
//        <stop offset="0.5" stop-color="#00ff00" />
//        <stop offset="1" stop-color="#0000ff" />
//      </linearGradient>
//    </defs>
//
//    <filter id="componentTransfer1" x="0" y="0" width="100%" height="100%">
//      <feComponentTransfer>
//        <feFuncR type="gamma" exponent="1"/>
//        <feFuncG type="gamma" exponent="1"/>
//        <feFuncB type="gamma" exponent="1"/>
//      </feComponentTransfer>
//    </filter>
//    <filter id="componentTransfer2" x="0" y="0" width="100%" height="100%">
//      <feComponentTransfer>
//        <feFuncR type="gamma" exponent="5"/>
//        <feFuncG type="gamma" exponent="5"/>
//        <feFuncB type="gamma" exponent="5"/>
//      </feComponentTransfer>
//    </filter>
//
//    <rect x="0" y="0" width="200" height="200" fill="url(#gradient)"
//        style="filter: url(#componentTransfer1);" />
//    <rect x="0" y="0" width="200" height="200" fill="url(#gradient)"
//        style="filter: url(#componentTransfer2); transform: translateX(220px);" />
//  </svg>

//go:build js

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 420, 200}).Append(
		factoryBrowser.NewTagSvgDefs().Append(
			factoryBrowser.NewTagSvgLinearGradient().Id("gradient").GradientUnits(html.KSvgGradientUnitsUserSpaceOnUse).X1(0).Y1(0).X2(200).Y2(0).Append(
				factoryBrowser.NewTagSvgStop().Offset(0.0).StopColor("#FF0000"),
				factoryBrowser.NewTagSvgStop().Offset(0.5).StopColor("#00FF00"),
				factoryBrowser.NewTagSvgStop().Offset(1.0).StopColor("#0000FF"),
			),
		),

		factoryBrowser.NewTagSvgFilter().Id("componentTransfer1").X(0).Y(0).Width(float32(1.0)).Height(float32(1.0)).Append(
			factoryBrowser.NewTagSvgFeComponentTransfer().Append(
				factoryBrowser.NewTagSvgFeFuncR().Type(html.KSvgTypeFeFuncGamma).Exponent(1),
				factoryBrowser.NewTagSvgFeFuncG().Type(html.KSvgTypeFeFuncGamma).Exponent(1),
				factoryBrowser.NewTagSvgFeFuncB().Type(html.KSvgTypeFeFuncGamma).Exponent(1),
			),
		),

		factoryBrowser.NewTagSvgFilter().Id("componentTransfer2").X(0).Y(0).Width(float32(1.0)).Height(float32(1.0)).Append(
			factoryBrowser.NewTagSvgFeComponentTransfer().Append(
				factoryBrowser.NewTagSvgFeFuncR().Type(html.KSvgTypeFeFuncGamma).Exponent(5),
				factoryBrowser.NewTagSvgFeFuncG().Type(html.KSvgTypeFeFuncGamma).Exponent(5),
				factoryBrowser.NewTagSvgFeFuncB().Type(html.KSvgTypeFeFuncGamma).Exponent(5),
			),
		),

		factoryBrowser.NewTagSvgRect().X(0).Y(0).Width(200).Height(200).Fill("url(#gradient)").Style("filter: url(#componentTransfer1);"),
		factoryBrowser.NewTagSvgRect().X(0).Y(0).Width(200).Height(200).Fill("url(#gradient)").Style("filter: url(#componentTransfer2); transform: translateX(220px);"),
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
