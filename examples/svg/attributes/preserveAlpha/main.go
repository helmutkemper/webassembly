// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/preserveAlpha
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/preserveAlpha
//
//  <svg viewBox="0 0 420 200" xmlns="http://www.w3.org/2000/svg">
//    <filter id="convolveMatrix1" x="0" y="0" width="100%" height="100%">
//      <feConvolveMatrix kernelMatrix="1 -1 2 0 0 0 0 0 -2" preserveAlpha="true"/>
//    </filter>
//    <filter id="convolveMatrix2" x="0" y="0" width="100%" height="100%">
//      <feConvolveMatrix kernelMatrix="1 -1 2 0 0 0 0 0 -2" preserveAlpha="false"/>
//    </filter>
//
//    <image xlink:href="//developer.mozilla.org/files/6457/mdn_logo_only_color.png" width="200" height="200"
//        style="filter:url(#convolveMatrix1);"/>
//    <image xlink:href="//developer.mozilla.org/files/6457/mdn_logo_only_color.png" width="200" height="200"
//        style="filter:url(#convolveMatrix2); transform:translateX(220px);"/>
//  </svg>

//go:build js

package main

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()
	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 420, 200}).Append(
		factoryBrowser.NewTagSvgFilter().Id("convolveMatrix1").X(0).Y(0).Width(float32(1.0)).Height(float32(1.0)).Append(
			factoryBrowser.NewTagSvgFeConvolveMatrix().KernelMatrix([]float64{1, -1, 2, 0, 0, 0, 0, 0, -2}).PreserveAlpha(true),
		),
		factoryBrowser.NewTagSvgFilter().Id("convolveMatrix2").X(0).Y(0).Width(float32(1.0)).Height(float32(1.0)).Append(
			factoryBrowser.NewTagSvgFeConvolveMatrix().KernelMatrix([]float64{1, -1, 2, 0, 0, 0, 0, 0, -2}).PreserveAlpha(false),
		),

		factoryBrowser.NewTagSvgImage().HRef("//developer.mozilla.org/files/6457/mdn_logo_only_color.png").Width(200).Height(200).Style("filter:url(#convolveMatrix1);"),
		factoryBrowser.NewTagSvgImage().HRef("//developer.mozilla.org/files/6457/mdn_logo_only_color.png").Width(200).Height(200).Style("filter:url(#convolveMatrix2); transform:translateX(220px);"),
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
