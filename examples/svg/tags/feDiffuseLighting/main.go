// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feDiffuseLighting
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feDiffuseLighting
//
//  <svg width="440" height="140" xmlns="http://www.w3.org/2000/svg">
//
//    <!-- No light is applied -->
//    <text text-anchor="middle" x="60" y="22">No Light</text> ok
//    <circle cx="60" cy="80" r="50" fill="green" /> ok
//
//    <!-- the light source is a fePointLight element -->
//    <text text-anchor="middle" x="170" y="22">fePointLight</text>
//    <filter id="lightMe1">
//      <feDiffuseLighting in="SourceGraphic" result="light"
//          lighting-color="white">
//        <fePointLight x="150" y="60" z="20" />
//      </feDiffuseLighting>
//
//      <feComposite in="SourceGraphic" in2="light"
//                   operator="arithmetic" k1="1" k2="0" k3="0" k4="0"/>
//    </filter>
//
//    <circle cx="170" cy="80" r="50" fill="green"
//        filter="url(#lightMe1)" /> ok
//
//    <!-- the light source is a feDistantLight element -->
//    <text text-anchor="middle" x="280" y="22">feDistantLight</text>
//    <filter id="lightMe2">
//      <feDiffuseLighting in="SourceGraphic" result="light"
//          lighting-color="white">
//        <feDistantLight azimuth="240" elevation="20"/>
//      </feDiffuseLighting>
//
//      <feComposite in="SourceGraphic" in2="light"
//                   operator="arithmetic" k1="1" k2="0" k3="0" k4="0"/>
//    </filter>
//
//    <circle cx="280" cy="80" r="50" fill="green"
//        filter="url(#lightMe2)" />
//
//    <!-- the light source is a feSpotLight source -->
//    <text text-anchor="middle" x="390" y="22">feSpotLight</text>
//    <filter id="lightMe3">
//      <feDiffuseLighting in="SourceGraphic" result="light"
//          lighting-color="white">
//        <feSpotLight x="360" y="5" z="30" limitingConeAngle="20"
//                     pointsAtX="390" pointsAtY="80" pointsAtZ="0"/>
//      </feDiffuseLighting>
//
//      <feComposite in="SourceGraphic" in2="light"
//                   operator="arithmetic" k1="1" k2="0" k3="0" k4="0"/>
//    </filter>
//
//    <circle cx="390" cy="80" r="50" fill="green"
//        filter="url(#lightMe3)" />
//  </svg>

//go:build js
// +build js

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
	"github.com/helmutkemper/iotmaker.webassembly/platform/factoryColor"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().Width(440).Height(140).Append(

		// No light is applied
		factoryBrowser.NewTagSvgText().TextAnchor(html.KSvgTextAnchorMiddle).X(60).Y(22).Text("No Light"),
		factoryBrowser.NewTagSvgCircle().Cx(60).Cy(80).R(50).Fill(factoryColor.NewGreen()),

		// the light source is a fePointLight element
		factoryBrowser.NewTagSvgText().TextAnchor(html.KSvgTextAnchorMiddle).X(170).Y(22).Text("fePointLight"),
		factoryBrowser.NewTagSvgFilter().Id("lightMe1").Append(
			factoryBrowser.NewTagSvgFeDiffuseLighting().In(html.KSvgInSourceGraphic).Result("light").LightingColor("white").Append(
				factoryBrowser.NewTagSvgFePointLight().X(150).Y(60).Z(20),
			),
			factoryBrowser.NewTagSvgFeComposite().In(html.KSvgInSourceGraphic).In2("light").Operator(html.KSvgOperatorFeCompositeArithmetic).K1(1).K2(0).K3(0).K4(0),
		),
		factoryBrowser.NewTagSvgCircle().Cx(170).Cy(80).R(50).Fill(factoryColor.NewGreen()).Filter("url(#lightMe1)"),

		// the light source is a feDistantLight element
		factoryBrowser.NewTagSvgText().TextAnchor(html.KSvgTextAnchorMiddle).X(280).Y(22).Text("feDistantLight"),
		factoryBrowser.NewTagSvgFilter().Id("lightMe2").Append(
			factoryBrowser.NewTagSvgFeDiffuseLighting().In(html.KSvgInSourceGraphic).Result("light").LightingColor("white").Append(
				factoryBrowser.NewTagSvgFeDistantLight().Azimuth(240).Elevation(20),
			),
			factoryBrowser.NewTagSvgFeComposite().In(html.KSvgInSourceGraphic).In2("light").Operator(html.KSvgOperatorFeCompositeArithmetic).K1(1).K2(0).K3(0).K4(0),
		),
		factoryBrowser.NewTagSvgCircle().Cx(280).Cy(80).R(50).Fill(factoryColor.NewGreen()).Filter("url(#lightMe2)"),

		// the light source is a feSpotLight source
		factoryBrowser.NewTagSvgText().TextAnchor(html.KSvgTextAnchorMiddle).X(390).Y(22).Text("feSpotLight"),
		factoryBrowser.NewTagSvgFilter().Id("lightMe3").Append(
			factoryBrowser.NewTagSvgFeDiffuseLighting().In(html.KSvgInSourceGraphic).Result("light").LightingColor("white").Append(
				factoryBrowser.NewTagSvgFeSpotLight().X(360).Y(5).Z(30).LimitingConeAngle(20).PointsAtX(390).PointsAtY(80).PointsAtZ(0),
			),
			factoryBrowser.NewTagSvgFeComposite().In(html.KSvgInSourceGraphic).In2("light").Operator(html.KSvgOperatorFeCompositeArithmetic).K1(1).K2(0).K3(0).K4(0),
		),
		factoryBrowser.NewTagSvgCircle().Cx(390).Cy(80).R(50).Fill(factoryColor.NewGreen()).Filter("url(#lightMe3)"),
	)

	stage.Append(s1)

	<-done
}
