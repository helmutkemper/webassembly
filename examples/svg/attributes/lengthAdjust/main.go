// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/lengthAdjust
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/lengthAdjust
//
//  <svg width="300" height="150" xmlns="http://www.w3.org/2000/svg">
//    <g font-face="sans-serif">
//      <text x="0" y="20" textLength="300" lengthAdjust="spacing">
//        Stretched using spacing only.
//      </text>
//      <text x="0" y="50" textLength="300" lengthAdjust="spacingAndGlyphs">
//        Stretched using spacing and glyphs.
//      </text>
//      <text x="0" y="80" textLength="100" lengthAdjust="spacing">
//        Shrunk using spacing only.
//      </text>
//      <text x="0" y="110" textLength="100" lengthAdjust="spacingAndGlyphs">
//        Shrunk using spacing and glyphs.
//      </text>
//    </g>
//  </svg>

//go:build js

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().Width(300).Height(150).Append(
		factoryBrowser.NewTagSvgG().FontFamily("sans-serif").Append(
			factoryBrowser.NewTagSvgText().X(0).Y(20).TextLength(300).LengthAdjust(html.KSvgLengthAdjustSpacing).Text("Stretched using spacing only."),
			factoryBrowser.NewTagSvgText().X(0).Y(50).TextLength(300).LengthAdjust(html.KSvgLengthAdjustSpacingAndGlyphs).Text("Stretched using spacing and glyphs."),
			factoryBrowser.NewTagSvgText().X(0).Y(80).TextLength(100).LengthAdjust(html.KSvgLengthAdjustSpacing).Text("Shrunk using spacing only."),
			factoryBrowser.NewTagSvgText().X(0).Y(110).TextLength(100).LengthAdjust(html.KSvgLengthAdjustSpacingAndGlyphs).Text("Shrunk using spacing and glyphs."),
		),
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
