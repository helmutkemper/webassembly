// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/mode
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/mode
//
//  <svg viewBox="0 0 480 200" xmlns="http://www.w3.org/2000/svg">
//    <filter id="blending1" x="0" y="0" width="100%" height="100%">
//      <feFlood result="floodFill" x="0" y="0" width="100%" height="100%"
//          flood-color="seagreen" flood-opacity="1"/>
//      <feBlend in="SourceGraphic" in2="floodFill" mode="multiply"/>
//    </filter>
//    <filter id="blending2" x="0" y="0" width="100%" height="100%">
//      <feFlood result="floodFill" x="0" y="0" width="100%" height="100%"
//          flood-color="seagreen" flood-opacity="1"/>
//      <feBlend in="SourceGraphic" in2="floodFill" mode="color-dodge"/>
//    </filter>
//
//    <image xlink:href="//developer.mozilla.org/files/6457/mdn_logo_only_color.png" width="200" height="200"
//        style="filter:url(#blending1);"/>
//    <image xlink:href="//developer.mozilla.org/files/6457/mdn_logo_only_color.png" width="200" height="200"
//        style="filter:url(#blending2); transform:translateX(220px);"/>
//  </svg>

//go:build js

package main

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/platform/factoryColor"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 480, 200}).Append(
		factoryBrowser.NewTagSvgFilter().Id("blending1").X(0).Y(0).Width(float32(1.0)).Height(float32(1.0)).Append(
			factoryBrowser.NewTagSvgFeFlood().Result("floodFill").X(0).Y(0).Width(float32(1.0)).Height(float32(1.0)).FloodColor(factoryColor.NewSeagreen()).FloodOpacity(1),
			factoryBrowser.NewTagSvgFeBlend().In(html.KSvgInSourceGraphic).In2("floodFill").Mode(html.KSvgModeMultiply),
		),
		factoryBrowser.NewTagSvgFilter().Id("blending2").X(0).Y(0).Width(float32(1.0)).Height(float32(1.0)).Append(
			factoryBrowser.NewTagSvgFeFlood().Result("floodFill").X(0).Y(0).Width(float32(1.0)).Height(float32(1.0)).FloodColor(factoryColor.NewSeagreen()).FloodOpacity(1),
			factoryBrowser.NewTagSvgFeBlend().In(html.KSvgInSourceGraphic).In2("floodFill").Mode(html.KSvgModeColorDodge),
		),

		factoryBrowser.NewTagSvgImage().HRef("//developer.mozilla.org/files/6457/mdn_logo_only_color.png").Width(200).Style("filter:url(#blending1);"),
		factoryBrowser.NewTagSvgImage().HRef("//developer.mozilla.org/files/6457/mdn_logo_only_color.png").Width(200).Style("filter:url(#blending2); transform:translateX(220px);"),
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
