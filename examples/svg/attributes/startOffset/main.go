// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/startOffset
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/startOffset
//
//  <svg viewBox="0 0 220 100" xmlns="http://www.w3.org/2000/svg">
//    <path id="path1" fill="none" stroke="red"
//          d="M10,90 Q90,90 90,45 Q90,10 50,10 Q10,10 10,40 Q10,70 45,70 Q70,70 75,50" />
//    <path id="path2" fill="none" stroke="red"
//          d="M130,90 Q210,90 210,45 Q210,10 170,10 Q130,10 130,40 Q130,70 165,70 Q190,70 195,50" />
//
//    <text>
//      <textPath href="#path1" startOffset="0">
//        Quick brown fox jumps over the lazy dog.
//      </textPath>
//    </text>
//
//    <text>
//      <textPath href="#path2" startOffset="40">
//        Quick brown fox jumps over the lazy dog.
//      </textPath>
//    </text>
//  </svg>

//go:build js
// +build js

// bug: o original não apararece nada

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/platform/factoryColor"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()
	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 220, 100}).Append(
		factoryBrowser.NewTagSvgPath().Id("path1").Fill(nil).Stroke(factoryColor.NewRed()).D(factoryBrowser.NewPath().M(10, 90).Q(90, 90, 90, 45).Q(90, 10, 50, 10).Q(10, 10, 10, 40).Q(10, 70, 45, 70).Q(70, 70, 75, 50)),
		factoryBrowser.NewTagSvgPath().Id("path2").Fill(nil).Stroke(factoryColor.NewRed()).D(factoryBrowser.NewPath().M(130, 90).Q(210, 90, 210, 45).Q(210, 10, 170, 10).Q(130, 10, 130, 40).Q(130, 70, 165, 70).Q(190, 70, 195, 50)),

		factoryBrowser.NewTagSvgText().Append(
			factoryBrowser.NewTagSvgTextPath().HRef("#path1").StartOffset(0).Text("Quick brown fox jumps over the lazy dog."),
		),
		factoryBrowser.NewTagSvgText().Append(
			factoryBrowser.NewTagSvgTextPath().HRef("#path2").StartOffset(40).Text("Quick brown fox jumps over the lazy dog."),
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
