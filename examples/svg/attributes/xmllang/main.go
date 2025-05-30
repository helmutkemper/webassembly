// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/xml:lang
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/xml:lang
//
//  <svg viewBox="0 0 200 100" xmlns="http://www.w3.org/2000/svg">
//    <text xml:lang="en-US">This is some English text</text>
//  </svg>

//go:build js

package main

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 200, 100}).Append(
		factoryBrowser.NewTagSvgText().Y(20).XmlLang("en-US").Text("This is some English text"),
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
