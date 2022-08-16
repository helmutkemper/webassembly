// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/flood-color
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/flood-color
//
//  <svg viewBox="0 0 420 200" xmlns="http://www.w3.org/2000/svg">
//    <filter id="flood1">
//      <feFlood flood-color="skyblue" x="0" y="0" width="200" height="200"/>
//    </filter>
//    <filter id="flood2">
//      <feFlood flood-color="seagreen" x="0" y="0" width="200" height="200"/>
//    </filter>
//
//    <rect x="0" y="0" width="200" height="200" style="filter: url(#flood1);" />
//    <rect x="0" y="0" width="200" height="200" style="filter: url(#flood2); transform: translateX(220px);" />
//  </svg>

//go:build js

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/platform/factoryColor"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 420, 200}).Append(
		factoryBrowser.NewTagSvgFilter().Id("flood1").Append(
			factoryBrowser.NewTagSvgFeFlood().FloodColor(factoryColor.NewSkyblue()).X(0).Y(0).Width(200).Height(200),
		),

		factoryBrowser.NewTagSvgFilter().Id("flood2").Append(
			factoryBrowser.NewTagSvgFeFlood().FloodColor(factoryColor.NewSeagreen()).X(0).Y(0).Width(200).Height(200),
		),

		factoryBrowser.NewTagSvgRect().X(0).Y(0).Width(200).Height(200).Style("filter: url(#flood1);"),
		factoryBrowser.NewTagSvgRect().X(0).Y(0).Width(200).Height(200).Style("filter: url(#flood2); transform: translateX(220px);"),
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
