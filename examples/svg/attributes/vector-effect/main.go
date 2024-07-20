// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/vector-effect
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/vector-effect
//
//  <svg viewBox="0 0 500 240">
//    <!-- normal -->
//    <path d="M10,20 L40,100 L39,200 z" stroke="black" stroke-width="2px" fill="none"></path>
//
//    <!-- scaled -->
//    <path transform="translate(100,0) scale(4,1)" d="M10,20 L40,100 L39,200 z" stroke="black"
//        stroke-width="2px" fill="none"></path>
//
//    <!-- fixed-->
//    <path vector-effect="non-scaling-stroke" transform="translate(300,0) scale(4,1)" d="M10,20 L40,100 L39,200 z"
//        stroke="black" stroke-width="2px" fill="none"></path>
//  </svg>

//go:build js

// bug: o original não apararece nada

package main

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/platform/factoryColor"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()
	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 500, 240}).Append(
		// normal
		factoryBrowser.NewTagSvgPath().D(factoryBrowser.NewPath().M(10, 20).L(40, 100).L(39, 200).Z()).Stroke(factoryColor.NewBlack()).StrokeWidth("2px").Fill(nil),

		// scale
		factoryBrowser.NewTagSvgPath().Transform(factoryBrowser.NewTransform().Translate(100, 0).Scale(4, 1)).D(factoryBrowser.NewPath().M(10, 20).L(40, 100).L(39, 200).Z()).Stroke(factoryColor.NewBlack()).StrokeWidth("2px").Fill(nil),

		// fixed
		factoryBrowser.NewTagSvgPath().VectorEffect(html.KSvgVectorEffectNonScalingStroke).Transform(factoryBrowser.NewTransform().Translate(300, 0).Scale(4, 1)).D(factoryBrowser.NewPath().M(10, 20).L(40, 100).L(39, 200).Z()).Stroke(factoryColor.NewBlack()).StrokeWidth("2px").Fill(nil),
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
