// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/mask
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/mask
//
//  <svg viewBox="0 0 200 100" xmlns="http://www.w3.org/2000/svg">
//    <mask id="myMask" maskContentUnits="objectBoundingBox">
//      <rect    fill="white" x="0" y="0" width="100%" height="100%" />
//      <polygon fill="black" points="0.5,0.2 0.68,0.74 0.21,0.41 0.79,0.41 0.32,0.74" />
//    </mask>
//
//    <!--
//    Punch a hole in a shape of a star inside the red circle,
//    revealing the yellow circle underneath
//    -->
//    <circle cx="50" cy="50" r="20" fill="yellow" />
//    <circle cx="50" cy="50" r="45" fill="red"
//            mask="url(#myMask)"/>
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

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 200, 100}).Append(
		factoryBrowser.NewTagSvgMask().Id("myMask").MaskContentUnits(html.KSvgUnitsObjectBoundingBox).Append(
			factoryBrowser.NewTagSvgRect().Fill(factoryColor.NewWhite()).X(0).Y(0).Width(float32(1.0)).Height(float32(1.0)),
			factoryBrowser.NewTagSvgPolygon().Fill(factoryColor.NewBlack()).Points([][]float64{{0.5, 0.2}, {0.68, 0.74}, {0.21, 0.41}, {0.79, 0.41}, {0.32, 0.74}}),
		),

		// Punch a hole in a shape of a star inside the red circle,
		// revealing the yellow circle underneath

		factoryBrowser.NewTagSvgCircle().Cx(50).Cy(50).R(20).Fill(factoryColor.NewYellow()),
		factoryBrowser.NewTagSvgCircle().Cx(50).Cy(50).R(45).Fill(factoryColor.NewRed()).Mask("url(#myMask)"),
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
