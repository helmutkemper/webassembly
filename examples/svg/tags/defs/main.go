// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Element/defs
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Element/defs
//
//  <svg viewBox="0 0 10 10" xmlns="http://www.w3.org/2000/svg">
//    <!-- Some graphical objects to use -->
//    <defs>
//      <circle id="myCircle" cx="0" cy="0" r="5" />
//
//      <linearGradient id="myGradient" gradientTransform="rotate(90)">
//        <stop offset="20%" stop-color="gold" />
//        <stop offset="90%" stop-color="red" />
//      </linearGradient>
//    </defs>
//
//    <!-- using my graphical objects -->
//    <use x="5" y="5" href="#myCircle" fill="url('#myGradient')" />
//  </svg>

//go:build js

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/platform/factoryColor"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 10, 10}).Append(
		factoryBrowser.NewTagSvgDefs().Append(
			factoryBrowser.NewTagSvgCircle().Id("myCircle").Cx(0).Cy(0).R(5),
			factoryBrowser.NewTagSvgLinearGradient().Id("myGradient").GradientTransform("rotate(90)").Append(
				factoryBrowser.NewTagSvgStop().Offset(float32(0.2)).StopColor(factoryColor.NewGold()),
				factoryBrowser.NewTagSvgStop().Offset(float32(0.9)).StopColor(factoryColor.NewRed()),
			),
		),
		factoryBrowser.NewTagSvgUse().X(5).Y(5).HRef("#myCircle").Fill("url('#myGradient')"),
	)

	stage.Append(s1)

	<-done
}
