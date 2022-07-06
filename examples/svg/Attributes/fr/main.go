// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/fr
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/fr
//
//  <svg viewBox="0 0 480 200" width="420" height="160" xmlns="http://www.w3.org/2000/svg">
//    <defs>
//      <radialGradient id="gradient1" cx="0.5" cy="0.5" r="0.5"
//          fx="0.35" fy="0.35" fr="5%">
//        <stop offset="0%" stop-color="white"/>
//        <stop offset="100%" stop-color="darkseagreen"/>
//      </radialGradient>
//      <radialGradient id="gradient2" cx="0.5" cy="0.5" r="0.5"
//          fx="0.35" fy="0.35" fr="25%">
//        <stop offset="0%" stop-color="white"/>
//        <stop offset="100%" stop-color="darkseagreen"/>
//      </radialGradient>
//    </defs>
//
//    <circle cx="100" cy="100" r="100" fill="url(#gradient1)" />
//    <circle cx="100" cy="100" r="100" fill="url(#gradient2)" style="transform: translateX(240px);" />
//  </svg>

//go:build js
// +build js

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/platform/factoryColor"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 480, 200}).Width(420).Height(160).Append(
		factoryBrowser.NewTagSvgDefs().Append(
			factoryBrowser.NewTagSvgRadialGradient().Id("gradient1").Cx(0.5).Cy(0.5).R(0.5).Fx(0.35).Fy(0.35).Fr(float32(0.05)).Append(
				factoryBrowser.NewTagSvgStop().Offset(float32(0.0)).StopColor(factoryColor.NewWhite()),
				factoryBrowser.NewTagSvgStop().Offset(float32(1.0)).StopColor(factoryColor.NewDarkseagreen()),
			),

			factoryBrowser.NewTagSvgRadialGradient().Id("gradient2").Cx(0.5).Cy(0.5).R(0.5).Fx(0.35).Fy(0.35).Fr(float32(0.25)).Append(
				factoryBrowser.NewTagSvgStop().Offset(float32(0.0)).StopColor(factoryColor.NewWhite()),
				factoryBrowser.NewTagSvgStop().Offset(float32(1.0)).StopColor(factoryColor.NewDarkseagreen()),
			),
		),

		factoryBrowser.NewTagSvgCircle().Cx(100).Cy(100).R(100).Fill("url(#gradient1)"),
		factoryBrowser.NewTagSvgCircle().Cx(100).Cy(100).R(100).Fill("url(#gradient2)").Style("transform: translateX(240px);"),
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
