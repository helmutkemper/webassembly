// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/stdDeviation
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/stdDeviation
//
//  <svg viewBox="0 0 480 200" xmlns="http://www.w3.org/2000/svg">
//    <filter id="gaussianBlur1">
//      <feGaussianBlur stdDeviation="1" />
//    </filter>
//    <filter id="gaussianBlur2">
//      <feGaussianBlur stdDeviation="5" />
//    </filter>
//    <filter id="gaussianBlur3" x="-30%" y="-30%" width="160%" height="160%">
//      <feGaussianBlur stdDeviation="10" />
//    </filter>
//
//    <circle cx="100" cy="100" r="50" style="filter: url(#gaussianBlur1);" />
//    <circle cx="100" cy="100" r="50" style="filter: url(#gaussianBlur2); transform: translateX(140px);" />
//    <circle cx="100" cy="100" r="50" style="filter: url(#gaussianBlur3); transform: translateX(280px);" />
//  </svg>

//go:build js

// bug: o original não apararece nada

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()
	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 480, 200}).Append(
		factoryBrowser.NewTagSvgFilter().Id("gaussianBlur1").Append(
			factoryBrowser.NewTagSvgFeGaussianBlur().StdDeviation(1),
		),
		factoryBrowser.NewTagSvgFilter().Id("gaussianBlur2").Append(
			factoryBrowser.NewTagSvgFeGaussianBlur().StdDeviation(5),
		),
		factoryBrowser.NewTagSvgFilter().Id("gaussianBlur3").X(float32(-0.3)).Y(float32(-0.3)).Width(float32(1.6)).Height(float32(1.6)).Append(
			factoryBrowser.NewTagSvgFeGaussianBlur().StdDeviation(10),
		),

		factoryBrowser.NewTagSvgCircle().Cx(100).Cy(100).R(50).Style("filter: url(#gaussianBlur1);"),
		factoryBrowser.NewTagSvgCircle().Cx(100).Cy(100).R(50).Style("filter: url(#gaussianBlur2); transform: translateX(140px);"),
		factoryBrowser.NewTagSvgCircle().Cx(100).Cy(100).R(50).Style("filter: url(#gaussianBlur3); transform: translateX(280px);"),
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
