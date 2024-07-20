// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/font-style
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/font-style
//
//  <svg viewBox="0 0 250 30" xmlns="http://www.w3.org/2000/svg">
//    <text y="20" font-style="normal">Normal font style</text>
//    <text x="150" y="20" font-style="italic">Italic font style</text>
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

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 250, 30}).Append(
		factoryBrowser.NewTagSvgText().Y(20).FontStyle(html.KFontStyleRuleNormal).Text("Normal font style"),
		factoryBrowser.NewTagSvgText().X(150).Y(20).FontStyle(html.KFontStyleRuleItalic).Text("Normal font style"),
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
