// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Element/textPath
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Element/textPath
//
//  <svg viewBox="0 0 100 100" xmlns="http://www.w3.org/2000/svg">
//
//    <!-- to hide the path, it is usually wrapped in a <defs> element -->
//    <!-- <defs> -->
//    <path id="MyPath" fill="none" stroke="red"
//          d="M10,90 Q90,90 90,45 Q90,10 50,10 Q10,10 10,40 Q10,70 45,70 Q70,70 75,50" />
//    <!-- </defs> -->
//
//    <text>
//      <textPath href="#MyPath">
//        Quick brown fox jumps over the lazy dog.
//      </textPath>
//    </text>
//
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

	s1 := factoryBrowser.NewTagSvg().
		ViewBox([]float64{0, 0, 100, 100}).
		Append(

			// to hide the path, it is usually wrapped in a <defs> element
			// <defs>
			factoryBrowser.NewTagSvgPath().Id("MyPath").Fill(nil).Stroke(factoryColor.NewRed()).D(factoryBrowser.NewPath().M(10, 90).Q(90, 90, 90, 45).Q(90, 10, 50, 10).Q(10, 10, 10, 40).Q(10, 70, 45, 70).Q(70, 70, 75, 50)),
			// <defs>

			factoryBrowser.NewTagSvgText().Append(

				factoryBrowser.NewTagSvgTextPath().HRef("#MyPath").Html("Quick brown fox jumps over the lazy dog."),
			),
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
//
//
//
//
//
//
//
//
