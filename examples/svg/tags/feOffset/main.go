// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feOffset
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feOffset
//
//  <svg width="200" height="200" xmlns="http://www.w3.org/2000/svg">
//    <defs>
//      <filter id="offset" width="180" height="180">
//        <feOffset in="SourceGraphic" dx="60" dy="60" />
//      </filter>
//    </defs>
//
//    <rect x="0" y="0" width="100" height="100" stroke="black" fill="green"/>
//    <rect x="0" y="0" width="100" height="100" stroke="black" fill="green" filter="url(#offset)"/>
//  </svg>

//go:build js
// +build js

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
			factoryBrowser.NewTagSvgFilter().Id("offset").Width(180).Height(180).Append(
				factoryBrowser.NewTagSvgFeOffset().In(html.KSvgInSourceGraphic).Dx(60).Dy(60),
			),
		),

		factoryBrowser.NewTagSvgRect().X(0).Y(0).Width(100).Height(100).Stroke(factoryColor.NewBlack()).Fill(factoryColor.NewGreen()),
		factoryBrowser.NewTagSvgRect().X(0).Y(0).Width(100).Height(100).Stroke(factoryColor.NewBlack()).Fill(factoryColor.NewGreen()).Filter("url(#offset)"),
	)

	stage.Append(s1)

	<-done
}
