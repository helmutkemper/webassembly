// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/surfaceScale
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/surfaceScale
//
//  <svg viewBox="0 0 260 260" xmlns="http://www.w3.org/2000/svg">
//    <circle cx="60" cy="60" r="15" tabindex="1" />
//    <circle cx="60" cy="160" r="30" tabindex="3" />
//    <circle cx="160" cy="60" r="30" tabindex="2" />
//    <circle cx="160" cy="160" r="60" tabindex="4" />
//  </svg>

//go:build js
// +build js

// bug: o original não apararece nada

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()
	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 260, 260}).Append(
		factoryBrowser.NewTagSvgCircle().Cx(60).Cy(60).R(15).Tabindex(1),
		factoryBrowser.NewTagSvgCircle().Cx(60).Cy(160).R(30).Tabindex(3),
		factoryBrowser.NewTagSvgCircle().Cx(160).Cy(60).R(30).Tabindex(2),
		factoryBrowser.NewTagSvgCircle().Cx(160).Cy(160).R(60).Tabindex(4),
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
