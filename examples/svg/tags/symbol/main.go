// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Element/symbol
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Element/symbol
//
//  <svg viewBox="0 0 80 20" xmlns="http://www.w3.org/2000/svg">
//    <!-- Our symbol in its own coordinate system -->
//    <symbol id="myDot" width="10" height="10" viewBox="0 0 2 2">
//      <circle cx="1" cy="1" r="1" />
//    </symbol>
//
//    <!-- A grid to materialize our symbol positioning -->
//    <path d="M0,10 h80 M10,0 v20 M25,0 v20 M40,0 v20 M55,0 v20 M70,0 v20" fill="none" stroke="pink" />
//
//    <!-- All instances of our symbol -->
//    <use href="#myDot" x="5"  y="5" style="opacity:1.0" />
//    <use href="#myDot" x="20" y="5" style="opacity:0.8" />
//    <use href="#myDot" x="35" y="5" style="opacity:0.6" />
//    <use href="#myDot" x="50" y="5" style="opacity:0.4" />
//    <use href="#myDot" x="65" y="5" style="opacity:0.2" />
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

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 80, 20}).Append(

		// Our symbol in its own coordinate system
		factoryBrowser.NewTagSvgSymbol().Id("myDot").Width(10).Height(10).ViewBox([]float64{0, 0, 2, 2}).Append(
			factoryBrowser.NewTagSvgCircle().Cx(1).Cy(1).R(1),
		),

		// A grid to materialize our symbol positioning
		factoryBrowser.NewTagSvgPath().D(factoryBrowser.NewPath().M(0, 10).Hd(80).M(10, 0).Vd(20).M(25, 0).Vd(20).M(40, 0).Vd(20).M(55, 0).Vd(20).M(70, 0).Vd(20)).Fill(nil).Stroke(factoryColor.NewPink()),

		// All instances of our symbol
		factoryBrowser.NewTagSvgUse().HRef("#myDot").X(5).Y(5).Style("opacity:1.0"),
		factoryBrowser.NewTagSvgUse().HRef("#myDot").X(20).Y(5).Style("opacity:0.8"),
		factoryBrowser.NewTagSvgUse().HRef("#myDot").X(35).Y(5).Style("opacity:0.6"),
		factoryBrowser.NewTagSvgUse().HRef("#myDot").X(50).Y(5).Style("opacity:0.4"),
		factoryBrowser.NewTagSvgUse().HRef("#myDot").X(65).Y(5).Style("opacity:0.2"),
	)

	stage.Append(s1)

	<-done
}
