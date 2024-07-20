// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/target
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/target
//
//  <svg viewBox="0 0 300 120" xmlns="http://www.w3.org/2000/svg">
//    <a href="https://developer.mozilla.org" target="_self">
//      <text x="0" y="20">Open link within iframe</text>
//    </a>
//    <a href="https://developer.mozilla.org" target="_blank">
//      <text x="0" y="60">Open link in new tab or window</text>
//    </a>
//    <a href="https://developer.mozilla.org" target="_top">
//      <text x="0" y="100">Open link in this tab or window</text>
//    </a>
//  </svg>

//go:build js

// bug: o original não apararece nada

package main

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/html"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()
	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 300, 120}).Append(
		factoryBrowser.NewTagSvgA().HRef("https://developer.mozilla.org").Target(html.KTargetSelf).Append(
			factoryBrowser.NewTagSvgText().X(0).Y(20).Text("Open link within iframe"),
		),

		factoryBrowser.NewTagSvgA().HRef("https://developer.mozilla.org").Target(html.KTargetBlank).Append(
			factoryBrowser.NewTagSvgText().X(0).Y(60).Text("Open link in new tab or window"),
		),

		factoryBrowser.NewTagSvgA().HRef("https://developer.mozilla.org").Target(html.KTargetTop).Append(
			factoryBrowser.NewTagSvgText().X(0).Y(100).Text("Open link in this tab or window"),
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
