// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/visibility
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/visibility
//
//  <svg viewBox="0 0 220 120" xmlns="http://www.w3.org/2000/svg">
//    <rect x="10" y="10" width="200" height="100" stroke="black"
//        stroke-width="5" fill="transparent" />
//    <g stroke="seagreen" stroke-width="5" fill="skyblue">
//      <rect x="20" y="20" width="80" height="80" visibility="visible" />
//      <rect x="120" y="20" width="80" height="80" visibility="hidden"/>
//    </g>
//  </svg>

//go:build js

// bug: o original não apararece nada

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
	"github.com/helmutkemper/iotmaker.webassembly/platform/factoryColor"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()
	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 220, 120}).Append(
		factoryBrowser.NewTagSvgRect().X(10).Y(10).Width(200).Height(100).Stroke(factoryColor.NewBlack()).StrokeWidth(5).Fill("transparent"),
		factoryBrowser.NewTagSvgG().Stroke(factoryColor.NewSeagreen()).StrokeWidth(5).Fill(factoryColor.NewSkyblue()).Append(
			factoryBrowser.NewTagSvgRect().X(20).Y(20).Width(80).Height(80).Visibility(html.KSvgVisibilityVisible),
			factoryBrowser.NewTagSvgRect().X(120).Y(20).Width(80).Height(80).Visibility(html.KSvgVisibilityHidden),
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
