// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/order
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/order
//
//  <svg viewBox="0 0 420 200" xmlns="http://www.w3.org/2000/svg">
//    <filter id="emboss1" x="0" y="0" width="100%" height="100%">
//      <feTurbulence baseFrequency="0.025" seed="0" />
//      <feConvolveMatrix kernelMatrix="3 0 0 -4" order="2"/>
//    </filter>
//    <filter id="emboss2" x="0" y="0" width="100%" height="100%">
//      <feTurbulence baseFrequency="0.025" seed="0" />
//      <feConvolveMatrix kernelMatrix="3 0 0 0 0 0 0 0 -4" order="3"/>
//    </filter>
//
//    <rect x="0" y="0" width="200" height="200" style="filter:url(#emboss1);" />
//    <rect x="0" y="0" width="200" height="200" style="filter:url(#emboss2); transform: translateX(220px);" />
//  </svg>

//go:build js

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 420, 200}).Append(
		factoryBrowser.NewTagSvgFilter().Id("emboss1").X(0).Y(0).Width(float32(1)).Height(float32(1)).Append(
			factoryBrowser.NewTagSvgFeTurbulence().BaseFrequency(0.025).Seed(0),
			factoryBrowser.NewTagSvgFeConvolveMatrix().KernelMatrix([]float64{3, 0, 0, -4}).Order(2),
		),
		factoryBrowser.NewTagSvgFilter().Id("emboss2").X(0).Y(0).Width(float32(1)).Height(float32(1)).Append(
			factoryBrowser.NewTagSvgFeTurbulence().BaseFrequency(0.025).Seed(0),
			factoryBrowser.NewTagSvgFeConvolveMatrix().KernelMatrix([]float64{3, 0, 0, 0, 0, 0, 0, 0, -4}).Order(3),
		),

		factoryBrowser.NewTagSvgRect().X(0).Y(0).Width(200).Height(200).Style("filter:url(#emboss1);"),
		factoryBrowser.NewTagSvgRect().X(0).Y(0).Width(200).Height(200).Style("filter:url(#emboss2); transform: translateX(220px);"),
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
