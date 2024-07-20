// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/font-family
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/font-family
//
//  <svg viewBox="0 0 200 30" xmlns="http://www.w3.org/2000/svg">
//    <text y="20" font-family="Arial, Helvetica, sans-serif">Sans serif</text>
//    <text x="100" y="20" font-family="monospace">Monospace</text>
//  </svg>

//go:build js

package main

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/factoryFontFamily"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 200, 30}).Append(
		factoryBrowser.NewTagSvgText().Y(20).FontFamily("Arial, Helvetica, sans-serif").Text("Sans serif"),
		factoryBrowser.NewTagSvgText().X(100).Y(20).FontFamily(factoryFontFamily.NewMonoSpace()).Text("Monospace"),
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
