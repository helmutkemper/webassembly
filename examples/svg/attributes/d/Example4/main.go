// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/d
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/d
//
//  <svg viewBox="0 0 200 100" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink">
//
//    <!-- Cubic Bézier curve with absolute coordinates -->
//    <path fill="none" stroke="red"
//          d="M 10,90
//             C 30,90 25,10 50,10
//             S 70,90 90,90" />
//
//    <!-- Cubic Bézier curve with relative coordinates -->
//    <path fill="none" stroke="red"
//          d="M 110,90
//             c 20,0 15,-80 40,-80
//             s 20,80 40,80" />
//
//    <!-- Highlight the curve vertex and control points -->
//    <g id="ControlPoints">
//
//      <!-- First cubic command control points -->
//      <line x1="10" y1="90" x2="30" y2="90" stroke="lightgrey" />
//      <circle cx="30" cy="90" r="1.5"/>
//
//      <line x1="50" y1="10" x2="25" y2="10" stroke="lightgrey" />
//      <circle cx="25" cy="10" r="1.5"/>
//
//      <!-- Second smooth command control points (the first one is implicit) -->
//      <line x1="50" y1="10" x2="75" y2="10" stroke="lightgrey" stroke-dasharray="2" />
//      <circle cx="75" cy="10" r="1.5" fill="lightgrey"/>
//
//      <line x1="90" y1="90" x2="70" y2="90" stroke="lightgrey" />
//      <circle cx="70" cy="90" r="1.5" />
//
//      <!-- curve vertex points -->
//      <circle cx="10" cy="90" r="1.5"/>
//      <circle cx="50" cy="10" r="1.5"/>
//      <circle cx="90" cy="90" r="1.5"/>
//    </g>
//    <use xlink:href="#ControlPoints" x="100" />
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

	s1 := factoryBrowser.NewTagSvg().
		ViewBox([]float64{0, 0, 200, 100}).XmlnsXLink("http://www.w3.org/1999/xlink").Append(

		// Cubic Bézier curve with absolute coordinates
		factoryBrowser.NewTagSvgPath().Fill(nil).Stroke(factoryColor.NewRed()).D(factoryBrowser.NewPath().M(10, 90).C(30, 90, 25, 10, 50, 10).S(70, 90, 90, 90)),

		// Cubic Bézier curve with relative coordinates
		factoryBrowser.NewTagSvgPath().Fill(nil).Stroke(factoryColor.NewRed()).D(factoryBrowser.NewPath().M(110, 90).Cd(20, 0, 15, -80, 40, -80).Sd(20, 80, 40, 80)),

		// Highlight the curve vertex and control points
		factoryBrowser.NewTagSvgG().Id("ControlPoints").Append(

			// First cubic command control points
			factoryBrowser.NewTagSvgLine().X1(10).Y1(90).X2(30).Y2(90).Stroke(factoryColor.NewLightgray()),
			factoryBrowser.NewTagSvgCircle().Cx(30).Cy(90).R(1.5),

			factoryBrowser.NewTagSvgLine().X1(50).Y1(10).X2(25).Y2(10).Stroke(factoryColor.NewLightgray()),
			factoryBrowser.NewTagSvgCircle().Cx(25).Cy(10).R(1.5),

			// Second smooth command control points (the first one is implicit)
			factoryBrowser.NewTagSvgLine().X1(50).Y1(10).X2(75).Y2(10).Stroke(factoryColor.NewLightgray()).StrokeDasharray(2),
			factoryBrowser.NewTagSvgCircle().Cx(75).Cy(10).R(1.5).Fill(factoryColor.NewLightgray()),

			factoryBrowser.NewTagSvgLine().X1(90).Y1(90).X2(70).Y2(90).Stroke(factoryColor.NewLightgray()),
			factoryBrowser.NewTagSvgCircle().Cx(70).Cy(90).R(1.5),

			// curve vertex points
			factoryBrowser.NewTagSvgCircle().Cx(10).Cy(90).R(1.5),
			factoryBrowser.NewTagSvgCircle().Cx(50).Cy(10).R(1.5),
			factoryBrowser.NewTagSvgCircle().Cx(90).Cy(90).R(1.5),
		),

		factoryBrowser.NewTagSvgUse().HRef("#ControlPoints").X(100),
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
