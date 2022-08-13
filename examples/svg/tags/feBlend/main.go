// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feBlend
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feBlend
//
//  <svg width="200" height="200" xmlns="http://www.w3.org/2000/svg"
//      xmlns:xlink="http://www.w3.org/1999/xlink">
//    <defs>
//      <filter id="spotlight">
//        <feFlood result="floodFill" x="0" y="0" width="100%" height="100%"
//            flood-color="green" flood-opacity="1"/>
//        <feBlend in="SourceGraphic" in2="floodFill" mode="multiply"/>
//      </filter>
//    </defs>
//
//    <image xlink:href="//developer.mozilla.org/files/6457/mdn_logo_only_color.png"
//        x="10%" y="10%" width="80%" height="80%"
//        style="filter:url(#spotlight);"/>
//  </svg>

//go:build js

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
	"github.com/helmutkemper/iotmaker.webassembly/platform/factoryColor"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().Width(200).Height(200).XmlnsXLink("http://www.w3.org/1999/xlink").Append(
		factoryBrowser.NewTagSvgDefs().Append(
			factoryBrowser.NewTagSvgFilter().Id("spotlight").Append(
				factoryBrowser.NewTagSvgFeFlood().Result("floodFill").X(0).Y(0).Width(float32(1)).Height(float32(1)).FloodColor(factoryColor.NewGreen()).FloodOpacity(1),
				factoryBrowser.NewTagSvgFeBlend().In(html.KSvgInSourceGraphic).In2("floodFill").Mode(html.KSvgModeMultiply),
			),
		),

		factoryBrowser.NewTagSvgImage().HRef("//developer.mozilla.org/files/6457/mdn_logo_only_color.png").X(float32(0.1)).Y(float32(0.1)).Width(float32(0.8)).Height(float32(0.8)).Style("filter:url(#spotlight);"),
	)

	stage.Append(s1)

	<-done
}
