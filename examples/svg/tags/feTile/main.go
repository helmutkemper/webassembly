// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feTile
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feTile
//
//  <svg width="200" height="200" xmlns="http://www.w3.org/2000/svg"
//      xmlns:xlink="http://www.w3.org/1999/xlink">
//    <defs>
//      <filter id="tile" x="0" y="0" width="100%" height="100%">
//        <feTile in="SourceGraphic" x="50" y="50"
//            width="100" height="100" />
//        <feTile/>
//      </filter>
//    </defs>
//
//    <image xlink:href="/files/6457/mdn_logo_only_color.png"
//        x="10%" y="10%" width="80%" height="80%"
//        style="filter:url(#tile);"/>
//  </svg>

//fixme: bug. Filter doesn't seem to work as expected

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

	s1 := factoryBrowser.NewTagSvg().Width(200).Height(200).XmlnsXLink("http://www.w3.org/1999/xlink").Append(
		factoryBrowser.NewTagSvgDefs().Append(
			factoryBrowser.NewTagSvgFilter().Id("tile").X(0).Y(0).Width(float32(1.0)).Height(float32(1.0)).Append(
				factoryBrowser.NewTagSvgFeTile().In(html.KSvgInSourceGraphic).X(50).Y(50).Width(100).Height(100),
			),
		),
		factoryBrowser.NewTagSvgImage().HRef("//developer.mozilla.org/files/6457/mdn_logo_only_color.png").X(float32(0.1)).Y(float32(0.1)).Width(float32(0.8)).Height(float32(0.8)).Style("filter:url(#tile);"),
	)

	stage.Append(s1)

	<-done
}
