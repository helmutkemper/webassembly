// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/in
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/in
//
//  <div style="width: 420px; height: 220px;">
//    <svg style="width:200px; height:200px; display: inline;" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink">
//      <defs>
//        <filter id="backgroundMultiply">
//          <!-- This will not work. -->
//          <feBlend in="BackgroundImage" in2="SourceGraphic" mode="multiply"/>
//        </filter>
//      </defs>
//      <image xlink:href="mdn_logo_only_color.png" x="10%" y="10%" width="80%" height="80%"/>
//      <circle cx="50%" cy="40%" r="40%" fill="#c00" style="filter:url(#backgroundMultiply);" />
//    </svg>
//
//    <svg style="width:200px; height:200px; display: inline;" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink">
//      <defs>
//        <filter id="imageMultiply">
//          <!-- This is a workaround. -->
//          <feImage xlink:href="mdn_logo_only_color.png" x="10%" y="10%" width="80%" height="80%"/>
//          <feBlend in2="SourceGraphic" mode="multiply"/>
//        </filter>
//      </defs>
//      <circle cx="50%" cy="40%" r="40%" fill="#c00" style="filter:url(#imageMultiply);"/>
//    </svg>
//  </div>

//go:build js

package main

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/html"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagDiv().Style("width: 420px; height: 220px;").Append(
		factoryBrowser.NewTagSvg().Style("width:200px; height:200px; display: inline;").XmlnsXLink("http://www.w3.org/1999/xlink").Append(
			factoryBrowser.NewTagSvgDefs().Append(
				factoryBrowser.NewTagSvgFilter().Id("backgroundMultiply").Append(
					//This will not work.
					factoryBrowser.NewTagSvgFeBlend().In(html.KSvgInBackgroundImage).In2(html.KSvgIn2SourceGraphic).Mode(html.KSvgModeMultiply),
				),
			),

			factoryBrowser.NewTagSvgImage().HRef("//developer.mozilla.org/files/6457/mdn_logo_only_color.png").X(float32(0.1)).Y(float32(0.1)).Width(float32(0.8)).Width(float32(0.8)),
			factoryBrowser.NewTagSvgCircle().Cx(float32(0.5)).Cy(float32(0.4)).R(float32(0.4)).Fill("#C00").Style("filter:url(#backgroundMultiply);"),
		),

		factoryBrowser.NewTagSvg().Style("width:200px; height:200px; display: inline;").XmlnsXLink("http://www.w3.org/1999/xlink").Append(
			factoryBrowser.NewTagSvgDefs().Append(
				factoryBrowser.NewTagSvgFilter().Id("imageMultiply").Append(
					// This is a workaround.
					factoryBrowser.NewTagSvgFeImage().HRef("//developer.mozilla.org/files/6457/mdn_logo_only_color.png").X(float32(0.1)).Y(float32(0.1)).Width(float32(0.8)).Height(float32(0.8)),
					factoryBrowser.NewTagSvgFeBlend().In2(html.KSvgIn2SourceGraphic).Mode(html.KSvgModeMultiply),
				),
			),
			factoryBrowser.NewTagSvgCircle().Cx(float32(0.5)).Cy(float32(0.4)).R(float32(0.4)).Fill("#C00").Style("filter:url(#imageMultiply);"),
		),
	)

	stage.Append(s1)

	<-done
}
