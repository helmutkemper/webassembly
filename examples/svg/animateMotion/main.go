// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Element/animateMotion
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Element/animateMotion
//
//  <svg viewBox="0 0 200 100" xmlns="http://www.w3.org/2000/svg">
//    <path fill="none" stroke="lightgrey"
//      d="M20,50 C20,-50 180,150 180,50 C180-50 20,150 20,50 z" />
//
//    <circle r="5" fill="red">
//      <animateMotion dur="10s" repeatCount="indefinite"
//        path="M20,50 C20,-50 180,150 180,50 C180-50 20,150 20,50 z" />
//    </circle>
//  </svg>

//go:build js
// +build js

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

	s1 := factoryBrowser.NewTagSvg("svg1").
		ViewBox([]float64{0, 0, 200, 100}).
		Append(

			factoryBrowser.NewTagSvgPath("path1").
				Fill("none").
				Stroke(factoryColor.NewLightgrey()).
				D(

					factoryBrowser.NewPath().
						M(20, 50).
						C(20, -50, 180, 150, 180, 50).
						C(180, -50, 20, 150, 20, 50).
						Z(),
				),

			factoryBrowser.NewTagSvgCircle("cir1").
				R(5).
				Fill(factoryColor.NewRed()).
				Append(

					factoryBrowser.NewTagSvgAnimateMotion("mot1").
						Dur(10*time.Second).
						RepeatCount(html.KSvgDurIndefinite).
						Path(

							factoryBrowser.NewPath().
								M(20, 50).
								C(20, -50, 180, 150, 180, 50).
								C(180, -50, 20, 150, 20, 50).
								Z(),
						),
				),
		)

	stage.Append(s1)

	<-done
}
