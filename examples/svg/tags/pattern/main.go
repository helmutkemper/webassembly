// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Element/pattern
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Element/pattern
//
//  <svg viewBox="0 0 230 100" xmlns="http://www.w3.org/2000/svg">
//    <defs>
//      <pattern id="star" viewBox="0,0,10,10" width="10%" height="10%">
//        <polygon points="0,0 2,5 0,10 5,8 10,10 8,5 10,0 5,2"/>
//      </pattern>
//    </defs>
//
//    <circle cx="50"  cy="50" r="50" fill="url(#star)"/>
//    <circle cx="180" cy="50" r="40" fill="none" stroke-width="20" stroke="url(#star)"/>
//  </svg>

//go:build js

package main

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 230, 100}).Append(
		factoryBrowser.NewTagSvgPattern().Id("star").ViewBox([]float64{0, 0, 10, 10}).Width(float32(0.1)).Height(float32(0.1)).Append(
			factoryBrowser.NewTagSvgPolygon().Points([][]float64{{0, 0}, {2, 5}, {0, 10}, {5, 8}, {10, 10}, {8, 5}, {10, 0}, {5, 2}}),
		),

		factoryBrowser.NewTagSvgCircle().Cx(50).Cy(50).R(50).Fill("url(#star)"),
		factoryBrowser.NewTagSvgCircle().Cx(180).Cy(50).R(40).Fill(nil).StrokeWidth(20).Stroke("url(#star)"),
	)

	stage.Append(s1)

	<-done
}
