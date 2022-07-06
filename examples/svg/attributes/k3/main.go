// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/k3
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/k3
//
//  <svg viewBox="0 0 420 200" xmlns="http://www.w3.org/2000/svg">
//    <filter id="composite1" x="0" y="0" width="100%" height="100%">
//      <feComposite in2="SourceGraphic" operator="arithmetic" k1="1" k2="0" k3="1" k4="0" />
//    </filter>
//    <filter id="composite2" x="0" y="0" width="100%" height="100%">
//      <feComposite in2="SourceGraphic" operator="arithmetic" k1="1" k2="0" k3="10" k4="0" />
//    </filter>
//
//    <image href="mdn.svg" x="0" y="0"
//        width="200" height="200" style="filter: url(#composite1);" />
//    <image href="mdn.svg" x="0" y="0"
//        width="200" height="200" style="filter: url(#composite2); transform: translateX(220px);" />
//  </svg>

//go:build js
// +build js

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 420, 200}).Append(
		factoryBrowser.NewTagSvgFilter().Id("composite1").X(0).Y(0).Width(float32(1.0)).Height(float32(1.0)).Append(
			factoryBrowser.NewTagSvgFeComposite().In2(html.KSvgIn2SourceGraphic).Operator(html.KSvgOperatorFeCompositeArithmetic).K1(1).K2(0).K3(1).K4(0),
		),

		factoryBrowser.NewTagSvgFilter().Id("composite2").X(0).Y(0).Width(float32(1.0)).Height(float32(1.0)).Append(
			factoryBrowser.NewTagSvgFeComposite().In2(html.KSvgIn2SourceGraphic).Operator(html.KSvgOperatorFeCompositeArithmetic).K1(1).K2(0).K3(10).K4(0),
		),

		factoryBrowser.NewTagSvgImage().HRef("//yari-demos.prod.mdn.mozit.cloud/en-US/docs/Web/SVG/Element/feConvolveMatrix/mdn.svg").X(0).Y(0).Width(200).Height(200).Style("filter: url(#composite1);"),
		factoryBrowser.NewTagSvgImage().HRef("//yari-demos.prod.mdn.mozit.cloud/en-US/docs/Web/SVG/Element/feConvolveMatrix/mdn.svg").X(0).Y(0).Width(200).Height(200).Style("filter: url(#composite2); transform: translateX(220px);"),
	)

	stage.Append(s1)

	<-done
}
