// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Element/clipPath
//
//   Notes:
//     * The CSS is inside the example HTML file.
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Element/clipPath
//
//   Notas:
//     * O CSS está dentro do arquivo HTML de exemplo.
//
// CSS:
//
//  /* With a touch of CSS for browsers who *
//   * implemented the r Geometry Property. */
//
//  @keyframes openYourHeart {from {r: 0} to {r: 60px}}
//
//  #myClip circle {
//    animation: openYourHeart 15s infinite;
//  }
//
// HTML:
//
//  <svg viewBox="0 0 100 100">
//    <clipPath id="myClip">
//      <!--
//        Everything outside the circle will be
//        clipped and therefore invisible.
//      -->
//      <circle cx="40" cy="35" r="35" />
//    </clipPath>
//
//    <!-- The original black heart, for reference -->
//    <path id="heart" d="M10,30 A20,20,0,0,1,50,30 A20,20,0,0,1,90,30 Q90,60,50,90 Q10,60,10,30 Z" />
//
//    <!--
//      Only the portion of the red heart
//      inside the clip circle is visible.
//    -->
//    <use clip-path="url(#myClip)" href="#heart" fill="red" />
//  </svg>

//go:build js
// +build js

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/platform/factoryColor"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 100, 100}).Append(

		factoryBrowser.NewTagSvgClipPath().Id("myClip").Append(

			// Everything outside the circle will be
			// clipped and therefore invisible.
			factoryBrowser.NewTagSvgCircle().Cx(40).Cy(35).R(35),
		),

		// The original black heart, for reference
		factoryBrowser.NewTagSvgPath().Id("heart").D(factoryBrowser.NewPath().M(10, 30).A(20, 20, 0, 0, 1, 50, 30).A(20, 20, 0, 0, 1, 90, 30).Q(90, 60, 50, 90).Q(10, 60, 10, 30).Z()),

		// Only the portion of the red heart
		// inside the clip circle is visible.
		factoryBrowser.NewTagSvgUse().ClipPath("url(#myClip)").HRef("#heart").Fill(factoryColor.NewRed()),
	)

	stage.Append(s1)

	<-done
}
