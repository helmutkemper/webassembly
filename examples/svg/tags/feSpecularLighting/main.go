// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feSpecularLighting
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feSpecularLighting
//
//  <svg height="200" width="200" viewBox="0 0 220 220"
//      xmlns="http://www.w3.org/2000/svg">
//    <filter id = "filter">
//      <feSpecularLighting result="specOut"
//          specularExponent="20" lighting-color="#bbbbbb">
//        <fePointLight x="50" y="75" z="200"/>
//      </feSpecularLighting>
//      <feComposite in="SourceGraphic" in2="specOut"
//          operator="arithmetic" k1="0" k2="1" k3="1" k4="0"/>
//    </filter>
//    <circle cx="110" cy="110" r="100" style="filter:url(#filter)"/>
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

	s1 := factoryBrowser.NewTagSvg().
		ViewBox([]float64{0, 0, 220, 220}).XmlnsXLink("http://www.w3.org/1999/xlink").Append(
		factoryBrowser.NewTagSvgFilter().Id("filter").Append(
			factoryBrowser.NewTagSvgFeSpecularLighting().Result("specOut").SpecularExponent(20).LightingColor("#bbbbbb").Append(
				factoryBrowser.NewTagSvgFePointLight().X(50).Y(75).Z(200),
			),

			factoryBrowser.NewTagSvgFeComposite().In(html.KSvgInSourceGraphic).In2("specOut").Operator(html.KSvgOperatorFeCompositeArithmetic).K1(0).K2(1).K3(1).K4(0),
		),

		factoryBrowser.NewTagSvgCircle().Cx(110).Cy(110).R(100).Style("filter:url(#filter)"),
	)

	stage.Append(s1)

	<-done
}
