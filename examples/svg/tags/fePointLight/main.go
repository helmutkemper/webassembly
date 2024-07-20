// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Element/fePointLight
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Element/fePointLight
//
//  <svg width="200" height="200" xmlns="http://www.w3.org/2000/svg"
//      xmlns:xlink="http://www.w3.org/1999/xlink">
//    <defs>
//      <filter id="spotlight">
//        <feSpecularLighting result="spotlight" specularConstant="1.5"
//            specularExponent="80" lighting-color="#FFF">
//          <fePointLight x="50" y="50" z="220"/>
//        </feSpecularLighting>
//        <feComposite in="SourceGraphic" in2="spotlight" operator="arithmetic"
//            k1="0" k2="1" k3="1" k4="0"/>
//      </filter>
//    </defs>
//
//    <image xlink:href="/files/6457/mdn_logo_only_color.png" x="10%" y="10%"
//        width="80%" height="80%" style="filter:url(#spotlight);"/>
//  </svg>

//go:build js

package main

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/html"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().Width(200).Height(200).XmlnsXLink("http://www.w3.org/1999/xlink").Append(
		factoryBrowser.NewTagSvgDefs().Append(
			factoryBrowser.NewTagSvgFilter().Id("spotlight").Append(
				factoryBrowser.NewTagSvgFeSpecularLighting().Result("spotlight").SpecularConstant(1.5).SpecularExponent(80).LightingColor("#FFF").Append(
					factoryBrowser.NewTagSvgFePointLight().X(50).Y(50).Z(220),
				),
				factoryBrowser.NewTagSvgFeComposite().In(html.KSvgInSourceGraphic).In2("spotlight").Operator(html.KSvgOperatorFeCompositeArithmetic).K1(0).K2(1).K3(1).K4(0),
			),
		),
		factoryBrowser.NewTagSvgImage().HRef("//developer.mozilla.org/files/6457/mdn_logo_only_color.png").X(float32(0.1)).Y(float32(0.1)).Width(float32(0.8)).Height(float32(0.8)).Style("filter:url(#spotlight);"),
	)

	stage.Append(s1)

	<-done
}
