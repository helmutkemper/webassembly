// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Element/stop
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Element/stop
//
//  <svg viewBox="0 0 10 10" xmlns="http://www.w3.org/2000/svg"
//       xmlns:xlink="http://www.w3.org/1999/xlink">
//    <defs>
//      <linearGradient id="myGradient" gradientTransform="rotate(90)">
//        <stop offset="5%"  stop-color="gold" />
//        <stop offset="95%" stop-color="red" />
//      </linearGradient>
//    </defs>
//
//    <!-- using my linear gradient -->
//    <circle cx="5" cy="5" r="4" fill="url('#myGradient')" />
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

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 10, 10}).XmlnsXLink("http://www.w3.org/1999/xlink").Append(
		factoryBrowser.NewTagSvgDefs().Append(
			factoryBrowser.NewTagSvgLinearGradient().Id("myGradient").GradientTransform(factoryBrowser.NewTransform().RotateAngle(90)).Append(
				factoryBrowser.NewTagSvgStop().Offset(float32(0.05)).StopColor(factoryColor.NewGold()),
				factoryBrowser.NewTagSvgStop().Offset(float32(0.95)).StopColor(factoryColor.NewRed()),
			),
		),

		//using my linear gradient
		factoryBrowser.NewTagSvgCircle().Cx(5).Cy(5).R(4).Fill("url('#myGradient')"),
	)

	stage.Append(s1)

	<-done
}
