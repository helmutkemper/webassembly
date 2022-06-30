// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feSpotLight
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feSpotLight
//
//  <svg width="200" height="200" xmlns="http://www.w3.org/2000/svg"
//      xmlns:xlink="http://www.w3.org/1999/xlink">
//    <defs>
//      <filter id="spotlight">
//        <feSpecularLighting result="spotlight" specularConstant="1.5"
//            specularExponent="4" lighting-color="#FFF">
//          <feSpotLight x="600" y="600" z="400" limitingConeAngle="5.5" />
//        </feSpecularLighting>
//        <feComposite in="SourceGraphic" in2="spotlight" operator="out"
//            k1="0" k2="1" k3="1" k4="0"/>
//      </filter>
//    </defs>
//
//    <image xlink:href="/files/6457/mdn_logo_only_color.png" x="10%" y="10%"
//        width="80%" height="80%" style="filter:url(#spotlight);"/>
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
		Width(200).
		Height(200).
		XmlnsXLink("http://www.w3.org/1999/xlink").
		Append(

			factoryBrowser.NewTagSvgDefs().
				Append(

					factoryBrowser.NewTagSvgFilter().
						Id("spotlight").
						Append(

							factoryBrowser.NewTagSvgFeSpecularLighting().
								Result("spotlight").
								SpecularConstant(1.5).
								SpecularExponent(4).
								LightingColor("#FFF").
								Append(

									factoryBrowser.NewTagSvgFeSpotLight().
										X(600).
										Y(600).
										Z(400).
										LimitingConeAngle(5.5),
								),

							factoryBrowser.NewTagSvgFeComposite().
								In(html.KSvgInSourceGraphic).
								In2("spotlight").
								Operator(html.KSvgOperatorFeCompositeOut).
								K1(0).
								K2(1).
								K3(1).
								K4(0),
						),
				),

			factoryBrowser.NewTagSvgImage().
				HRef("//media.prod.mdn.mozit.cloud/attachments/2013/11/15/6457/5e0f6aa96fb8e4593f143aa803576698/mdn_logo_only_color.png").
				X(float32(0.1)).
				Y(float32(0.1)).
				Width(float32(0.8)).
				Height(float32(0.8)).
				Style("filter:url(#spotlight);"),
		)

	stage.Append(s1)

	<-done
}
