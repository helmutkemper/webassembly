// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Element/mpath
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Element/mpath
//
//  <svg width="100%" height="100%"  viewBox="0 0 500 300"
//       xmlns="http://www.w3.org/2000/svg"
//       xmlns:xlink="http://www.w3.org/1999/xlink" >
//
//    <rect x="1" y="1" width="498" height="298"
//          fill="none" stroke="blue" stroke-width="2" />
//
//    <!-- Draw the outline of the motion path in blue, along
//            with three small circles at the start, middle and end. -->
//    <path id="path1" d="M100,250 C 100,50 400,50 400,250"
//          fill="none" stroke="blue" stroke-width="7.06"  />
//    <circle cx="100" cy="250" r="17.64" fill="blue"  />
//    <circle cx="250" cy="100" r="17.64" fill="blue"  />
//    <circle cx="400" cy="250" r="17.64" fill="blue"  />
//
//    <!-- Here is a triangle which will be moved about the motion path.
//         It is defined with an upright orientation with the base of
//         the triangle centered horizontally just above the origin. -->
//    <path d="M-25,-12.5 L25,-12.5 L 0,-87.5 z"
//          fill="yellow" stroke="red" stroke-width="7.06"  >
//      <!-- Define the motion path animation -->
//      <animateMotion dur="6s" repeatCount="indefinite" rotate="auto" >
//         <mpath xlink:href="#path1"/>
//      </animateMotion>
//    </path>
//  </svg>

//go:build js

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
	"github.com/helmutkemper/iotmaker.webassembly/platform/factoryColor"
	"time"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().Width(float32(1.0)).Height(float32(1.0)).ViewBox([]float64{0, 0, 500, 500}).XmlnsXLink("http://www.w3.org/1999/xlink").Append(

		factoryBrowser.NewTagSvgRect().X(1).Y(1).Width(498).Height(298).Fill(nil).Stroke(factoryColor.NewBlue()).StrokeWidth(2),

		// Draw the outline of the motion path in blue, along
		// with three small circles at the start, middle and end.
		factoryBrowser.NewTagSvgPath().Id("path1").D(factoryBrowser.NewPath().M(100, 250).C(100, 50, 400, 50, 400, 250)).Fill(nil).Stroke(factoryColor.NewBlue()).StrokeWidth(7.06),
		factoryBrowser.NewTagSvgCircle().Cx(100).Cy(250).R(17.64).Fill(factoryColor.NewBlue()),
		factoryBrowser.NewTagSvgCircle().Cx(250).Cy(10).R(17.64).Fill(factoryColor.NewBlue()),
		factoryBrowser.NewTagSvgCircle().Cx(400).Cy(250).R(17.64).Fill(factoryColor.NewBlue()),

		// Here is a triangle which will be moved about the motion path.
		// It is defined with an upright orientation with the base of
		// the triangle centered horizontally just above the origin.
		factoryBrowser.NewTagSvgPath().D(factoryBrowser.NewPath().M(-25, -12.5).L(25, -12.5).L(0, -87.5).Z()).Fill(factoryColor.NewYellow()).Stroke(factoryColor.NewRed()).StrokeWidth(7.06).Append(

			// Define the motion path animation
			factoryBrowser.NewTagSvgAnimateMotion().Dur(6*time.Second).RepeatCount(html.KSvgDurIndefinite).Rotate(html.KSvgRotateAuto).Append(

				factoryBrowser.NewTagSvgMPath().HRef("#path1"),
			),
		),
	)

	stage.Append(s1)

	<-done
}
