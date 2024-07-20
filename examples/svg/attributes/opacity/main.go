// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/opacity
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/opacity
//
//  <svg viewBox="0 0 200 100" xmlns="http://www.w3.org/2000/svg">
//    <defs>
//      <linearGradient id="gradient" x1="0%" y1="0%" x2="0" y2="100%">
//        <stop offset="0%" style="stop-color:skyblue;" />
//        <stop offset="100%" style="stop-color:seagreen;" />
//      </linearGradient>
//    </defs>
//    <rect x="0" y="0" width="100%" height="100%" fill="url(#gradient)" />
//    <circle cx="50" cy="50" r="40" fill="black" />
//    <circle cx="150" cy="50" r="40" fill="black" opacity="0.3" />
//  </svg>

//go:build js

package main

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/platform/factoryColor"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 200, 100}).Append(
		factoryBrowser.NewTagSvgDefs().Append(
			factoryBrowser.NewTagSvgLinearGradient().Id("gradient").X1(float32(0)).Y1(float32(0)).X2(float32(0)).Y2(float32(1)).Append(
				factoryBrowser.NewTagSvgStop().Offset(float32(0)).Style("stop-color:skyblue;"),
				factoryBrowser.NewTagSvgStop().Offset(float32(1)).Style("stop-color:seagreen;"),
			),
		),

		factoryBrowser.NewTagSvgRect().X(0).Y(0).Width(float32(1)).Height(float32(1)).Fill("url(#gradient)"),
		factoryBrowser.NewTagSvgCircle().Cx(50).Cy(50).R(40).Fill(factoryColor.NewBlack()),
		factoryBrowser.NewTagSvgCircle().Cx(150).Cy(50).R(40).Fill(factoryColor.NewBlack()).Opacity(0.3),
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
