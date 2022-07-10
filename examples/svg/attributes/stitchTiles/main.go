// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/stitchTiles
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/stitchTiles
//
//  <svg viewBox="0 0 420 200" xmlns="http://www.w3.org/2000/svg">
//    <filter id="noise1" x="0" y="0" width="100%" height="100%">
//      <feTurbulence baseFrequency="0.025" stitchTiles="noStitch" />
//    </filter>
//    <filter id="noise2" x="0" y="0" width="100%" height="100%">
//      <feTurbulence baseFrequency="0.025" stitchTiles="stitch" />
//    </filter>
//
//    <rect x="0" y="0" width="100" height="100" style="filter: url(#noise1);" />
//    <rect x="0" y="0" width="100" height="100" style="filter: url(#noise1); transform: translate(100px, 0);" />
//    <rect x="0" y="0" width="100" height="100" style="filter: url(#noise1); transform: translate(0, 100px);" />
//    <rect x="0" y="0" width="100" height="100" style="filter: url(#noise1); transform: translate(100px, 100px);" />
//
//    <rect x="0" y="0" width="100" height="100" style="filter: url(#noise2); transform: translate(220px, 0);" />
//    <rect x="0" y="0" width="100" height="100" style="filter: url(#noise2); transform: translate(320px, 0);" />
//    <rect x="0" y="0" width="100" height="100" style="filter: url(#noise2); transform: translate(220px, 100px);" />
//    <rect x="0" y="0" width="100" height="100" style="filter: url(#noise2); transform: translate(320px, 100px);" />
//  </svg>

//go:build js
// +build js

// bug: o original não apararece nada

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()
	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 420, 200}).Append(
		factoryBrowser.NewTagSvgFilter().Id("noise1").X(0).Y(0).Width(float32(1)).Height(float32(1)).Append(
			factoryBrowser.NewTagSvgFeTurbulence().BaseFrequency(0.025).StitchTiles(html.KSvgStitchTilesNoStitch),
		),
		factoryBrowser.NewTagSvgFilter().Id("noise2").X(0).Y(0).Width(float32(1)).Height(float32(1)).Append(
			factoryBrowser.NewTagSvgFeTurbulence().BaseFrequency(0.025).StitchTiles(html.KSvgStitchTilesStitch),
		),

		factoryBrowser.NewTagSvgRect().X(0).Y(0).Width(100).Height(100).Style("filter: url(#noise1);"),
		factoryBrowser.NewTagSvgRect().X(0).Y(0).Width(100).Height(100).Style("filter: url(#noise1); transform: translate(100px, 0);"),
		factoryBrowser.NewTagSvgRect().X(0).Y(0).Width(100).Height(100).Style("filter: url(#noise1); transform: translate(0, 100px);"),
		factoryBrowser.NewTagSvgRect().X(0).Y(0).Width(100).Height(100).Style("filter: url(#noise1); transform: translate(100px, 100px);"),

		factoryBrowser.NewTagSvgRect().X(0).Y(0).Width(100).Height(100).Style("filter: url(#noise2); transform: translate(220px, 0);"),
		factoryBrowser.NewTagSvgRect().X(0).Y(0).Width(100).Height(100).Style("filter: url(#noise2); transform: translate(320px, 0);"),
		factoryBrowser.NewTagSvgRect().X(0).Y(0).Width(100).Height(100).Style("filter: url(#noise2); transform: translate(220px, 100px);"),
		factoryBrowser.NewTagSvgRect().X(0).Y(0).Width(100).Height(100).Style("filter: url(#noise2); transform: translate(320px, 100px);"),
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
