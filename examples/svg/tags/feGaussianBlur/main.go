// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feGaussianBlur
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feGaussianBlur
//
//  <svg width="230" height="120"
//          xmlns="http://www.w3.org/2000/svg"
//          xmlns:xlink="http://www.w3.org/1999/xlink">
//
//    <filter id="blurMe">
//      <feGaussianBlur in="SourceGraphic" stdDeviation="5" />
//    </filter>
//
//    <circle cx="60"  cy="60" r="50" fill="green" />
//
//    <circle cx="170" cy="60" r="50" fill="green"
//            filter="url(#blurMe)" />
//  </svg>

//go:build js

package main

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/platform/factoryColor"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().Width(230).Height(120).XmlnsXLink("http://www.w3.org/1999/xlink").Append(
		factoryBrowser.NewTagSvgFilter().Id("blurMe").Append(
			factoryBrowser.NewTagSvgFeGaussianBlur().In(html.KSvgInSourceGraphic).StdDeviation(5),
		),

		factoryBrowser.NewTagSvgCircle().Cx(60).Cy(50).R(50).Fill(factoryColor.NewGreen()),
		factoryBrowser.NewTagSvgCircle().Cx(170).Cy(50).R(50).Fill(factoryColor.NewGreen()).Filter("url(#blurMe)"),
	)

	stage.Append(s1)

	<-done
}
