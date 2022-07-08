// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/operator
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/operator
//
//  <svg viewBox="0 0 120 70" xmlns="http://www.w3.org/2000/svg">
//    <filter id="erode">
//      <feMorphology operator="erode" radius="0.4"/>
//    </filter>
//    <filter id="dilate">
//      <feMorphology operator="dilate" radius="0.8"/>
//    </filter>
//
//    <text x="0" y="15">Normal text</text>
//    <text x="0" y="40" filter="url(#erode)">Thin text</text>
//    <text x="0" y="65" filter="url(#dilate)">Fat text</text>
//  </svg>

//go:build js
// +build js

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 120, 70}).Append(
		factoryBrowser.NewTagSvgFilter().Id("erode").Append(
			factoryBrowser.NewTagSvgFeMorphology().Operator(html.KKSvgOperatorFeCompositeErode).Radius(0.4),
		),
		factoryBrowser.NewTagSvgFilter().Id("dilate").Append(
			factoryBrowser.NewTagSvgFeMorphology().Operator(html.KKSvgOperatorFeCompositeDilate).Radius(0.8),
		),

		factoryBrowser.NewTagSvgText().X(0).Y(15).Text("Normal text"),
		factoryBrowser.NewTagSvgText().X(0).Y(40).Filter("url(#erode)").Text("Thin text"),
		factoryBrowser.NewTagSvgText().X(0).Y(65).Filter("url(#dilate)").Text("Fat text"),
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
