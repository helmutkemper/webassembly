// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/pointsAtY
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/pointsAtY
//
//  <svg viewBox="0 0 420 200" xmlns="http://www.w3.org/2000/svg">
//
//    <filter id="lighting1" x="0" y="0" width="100%" height="100%">
//      <feDiffuseLighting in="SourceGraphic">
//        <feSpotLight x="60" y="60" z="50" pointsAtY="0" />
//      </feDiffuseLighting>
//    </filter>
//    <filter id="lighting2" x="0" y="0" width="100%" height="100%">
//      <feDiffuseLighting in="SourceGraphic">
//        <feSpotLight x="60" y="60" z="50" pointsAtY="400" />
//      </feDiffuseLighting>
//    </filter>
//
//    <rect x="0" y="0" width="200" height="200" style="filter: url(#lighting1);" />
//    <rect x="0" y="0" width="200" height="200" style="filter: url(#lighting2); transform: translateX(220px);" />
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
	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 420, 200}).Append(
		factoryBrowser.NewTagSvgFilter().Id("lighting1").X(0).Y(0).Width(float32(1.0)).Height(float32(1.0)).Append(
			factoryBrowser.NewTagSvgFeDiffuseLighting().In(html.KSvgInSourceGraphic).Append(
				factoryBrowser.NewTagSvgFeSpotLight().X(60).Y(60).Z(50).PointsAtY(0),
			),
		),

		factoryBrowser.NewTagSvgFilter().Id("lighting2").X(0).Y(0).Width(float32(1.0)).Height(float32(1.0)).Append(
			factoryBrowser.NewTagSvgFeDiffuseLighting().In(html.KSvgInSourceGraphic).Append(
				factoryBrowser.NewTagSvgFeSpotLight().X(60).Y(60).Z(50).PointsAtY(400),
			),
		),

		factoryBrowser.NewTagSvgRect().X(0).Y(0).Width(200).Height(200).Style("filter: url(#lighting1);"),
		factoryBrowser.NewTagSvgRect().X(0).Y(0).Width(200).Height(200).Style("filter: url(#lighting2); transform: translateX(220px);"),
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
