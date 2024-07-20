// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feConvolveMatrix
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feConvolveMatrix
//
//  <svg width="200" height="200" viewBox="0 0 200 200" xmlns="http://www.w3.org/2000/svg"
//      xmlns:xlink="http://www.w3.org/1999/xlink">
//    <defs>
//      <filter id="emboss">
//        <feConvolveMatrix
//            kernelMatrix="3 0 0
//                          0 0 0
//                          0 0 -3"/>
//      </filter>
//    </defs>
//
//    <image xlink:href="mdn.svg" x="0" y="0"
//        height="200" width="200" style="filter:url(#emboss);" />
//  </svg>

//go:build js

package main

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().
		ViewBox([]float64{0, 0, 200, 200}).Width(200).Height(200).XmlnsXLink("http://www.w3.org/1999/xlink").Append(
		factoryBrowser.NewTagSvgDefs().Append(
			factoryBrowser.NewTagSvgFilter().Id("emboss").Append(
				factoryBrowser.NewTagSvgFeConvolveMatrix().KernelMatrix(
					[]float64{
						3, 0, 0,
						0, 0, 0,
						0, 0, -3,
					},
				),
			),
		),

		factoryBrowser.NewTagSvgImage().HRef("//yari-demos.prod.mdn.mozit.cloud/en-US/docs/Web/SVG/Element/feConvolveMatrix/mdn.svg").X(0).Y(0).Height(200).Width(200).Style("filter:url(#emboss);"),
	)

	stage.Append(s1)

	<-done
}
