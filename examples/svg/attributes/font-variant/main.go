// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/font-variant
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/font-variant
//
//  <svg viewBox="0 0 250 30" xmlns="http://www.w3.org/2000/svg">
//    <text y="20" font-variant="normal">Normal text</text>
//    <text x="100" y="20" font-variant="small-caps">Small-caps text</text>
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

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 250, 30}).Append(
		factoryBrowser.NewTagSvgText().Y(20).FontVariant(html.KFontVariantRuleNormal).Text("Normal text"),
		factoryBrowser.NewTagSvgText().X(150).Y(20).FontVariant(html.KFontVariantRuleSmallCaps).Text("Small-caps text"),
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
