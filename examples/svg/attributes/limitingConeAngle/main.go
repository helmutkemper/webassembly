// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/limitingConeAngle
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/limitingConeAngle
//
//  <svg viewBox="0 0 480 200" xmlns="http://www.w3.org/2000/svg">
//    <filter id="spotLight1" x="0" y="0" width="100%" height="100%">
//      <feDiffuseLighting diffuseConstant="2">
//        <feSpotLight x="10" y="10" z="50" pointsAtX="100" pointsAtY="100" limitingConeAngle="10" />
//      </feDiffuseLighting>
//    </filter>
//    <filter id="spotLight2" x="0" y="0" width="100%" height="100%">
//      <feDiffuseLighting diffuseConstant="2">
//        <feSpotLight x="10" y="10" z="50" pointsAtX="100" pointsAtY="100" limitingConeAngle="40" />
//      </feDiffuseLighting>
//    </filter>
//
//    <rect x="0" y="0" width="200" height="200" style="filter: url(#spotLight1);" />
//    <rect x="0" y="0" width="200" height="200" style="filter: url(#spotLight2); transform: translateX(220px);" />
//  </svg>

//go:build js

package main

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 480, 200}).Append(
		factoryBrowser.NewTagSvgFilter().Id("spotLight1").X(0).Y(0).Width(float32(1.0)).Height(float32(1.0)).Append(
			factoryBrowser.NewTagSvgFeDiffuseLighting().DiffuseConstant(2).Append(
				factoryBrowser.NewTagSvgFeSpotLight().X(10).Y(10).Z(50).PointsAtX(100).PointsAtY(100).LimitingConeAngle(10),
			),
		),

		factoryBrowser.NewTagSvgFilter().Id("spotLight2").X(0).Y(0).Width(float32(1.0)).Height(float32(1.0)).Append(
			factoryBrowser.NewTagSvgFeDiffuseLighting().DiffuseConstant(2).Append(
				factoryBrowser.NewTagSvgFeSpotLight().X(10).Y(10).Z(50).PointsAtX(100).PointsAtY(100).LimitingConeAngle(40),
			),
		),

		factoryBrowser.NewTagSvgRect().X(0).Y(0).Width(200).Height(200).Style("filter: url(#spotLight1);"),
		factoryBrowser.NewTagSvgRect().X(0).Y(0).Width(200).Height(200).Style("filter: url(#spotLight2); transform: translateX(220px);"),
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
