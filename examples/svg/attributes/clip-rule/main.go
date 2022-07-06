// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/clip-rule
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/clip-rule
//
//  <svg width="100" viewBox="0 0 100 90" xmlns="http://www.w3.org/2000/svg" version="1.1">
//    <!-- Define star path -->
//    <defs>
//      <path d="M50,0 21,90 98,35 2,35 79,90z" id="star" />
//    </defs>
//
//    <!-- Left: evenodd -->
//    <clipPath id="emptyStar">
//      <use xlink:href="#star" clip-rule="evenodd" />
//    </clipPath>
//    <rect clip-path="url(#emptyStar)" width="50" height="90" fill="blue" />
//
//    <!-- Right: nonzero -->
//    <clipPath id="filledStar">
//      <use xlink:href="#star" clip-rule="nonzero" />
//    </clipPath>
//    <rect clip-path="url(#filledStar)" width="50" height="90" x="50" fill="red" />
//  </svg>

//go:build js
// +build js

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
	"github.com/helmutkemper/iotmaker.webassembly/platform/factoryColor"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().
		Width(100).ViewBox([]float64{0, 0, 100, 90}).Append(

		// Define star path
		factoryBrowser.NewTagSvgDefs().Append(
			//factoryBrowser.NewTagSvgPath().D("M50,0 21,90 98,35 2,35 79,90z").Id("star"),
			factoryBrowser.NewTagSvgPath().D(factoryBrowser.NewPath().M(50, 0).L(21, 90).L(98, 35).L(2, 35).L(79, 90).Z()).Id("star"),
		),

		// Left: evenodd
		factoryBrowser.NewTagSvgClipPath().Id("emptyStar").Append(
			factoryBrowser.NewTagSvgUse().HRef("#star").ClipRule(html.KSvgClipRuleEvenOdd),
		),
		factoryBrowser.NewTagSvgRect().ClipPath("url(#emptyStar)").Width(50).Height(90).Fill(factoryColor.NewBlue()),

		// Right: nonzero
		factoryBrowser.NewTagSvgClipPath().Id("filledStar").Append(
			factoryBrowser.NewTagSvgUse().HRef("#star").ClipRule(html.KSvgClipRuleNonzero),
		),
		factoryBrowser.NewTagSvgRect().ClipPath("url(#filledStar)").Width(50).Height(90).X(50).Fill(factoryColor.NewRed()),
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
