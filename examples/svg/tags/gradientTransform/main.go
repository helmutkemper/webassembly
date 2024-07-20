// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/gradientTransform
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/gradientTransform
//
//  <svg viewBox="0 0 420 200" xmlns="http://www.w3.org/2000/svg">
//    <radialGradient id="gradient1" gradientUnits="userSpaceOnUse"
//        cx="100" cy="100" r="100" fx="100" fy="100">
//      <stop offset="0%" stop-color="darkblue" />
//      <stop offset="50%" stop-color="skyblue" />
//      <stop offset="100%" stop-color="darkblue" />
//    </radialGradient>
//    <radialGradient id="gradient2" gradientUnits="userSpaceOnUse"
//        cx="100" cy="100" r="100" fx="100" fy="100"
//        gradientTransform="skewX(20) translate(-35, 0)">
//      <stop offset="0%" stop-color="darkblue" />
//      <stop offset="50%" stop-color="skyblue" />
//      <stop offset="100%" stop-color="darkblue" />
//    </radialGradient>
//
//    <rect x="0" y="0" width="200" height="200" fill="url(#gradient1)" />
//    <rect x="0" y="0" width="200" height="200" fill="url(#gradient2)" style="transform: translateX(220px);" />
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

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 420, 200}).Append(
		factoryBrowser.NewTagSvgRadialGradient().Id("gradient1").GradientUnits(html.KSvgGradientUnitsUserSpaceOnUse).Cx(100).Cy(100).R(100).Fx(100).Fy(100).Append(
			factoryBrowser.NewTagSvgStop().Offset(float32(0.0)).StopColor(factoryColor.NewDarkblue()),
			factoryBrowser.NewTagSvgStop().Offset(float32(0.5)).StopColor(factoryColor.NewSkyblue()),
			factoryBrowser.NewTagSvgStop().Offset(float32(1.0)).StopColor(factoryColor.NewDarkblue()),
		),

		factoryBrowser.NewTagSvgRadialGradient().Id("gradient2").GradientUnits(html.KSvgGradientUnitsUserSpaceOnUse).Cx(100).Cy(100).R(100).Fx(100).GradientTransform(factoryBrowser.NewTransform().SkewX(20).Translate(-35, 0)).Fy(100).Append(
			factoryBrowser.NewTagSvgStop().Offset(float32(0.0)).StopColor(factoryColor.NewDarkblue()),
			factoryBrowser.NewTagSvgStop().Offset(float32(0.5)).StopColor(factoryColor.NewSkyblue()),
			factoryBrowser.NewTagSvgStop().Offset(float32(1.0)).StopColor(factoryColor.NewDarkblue()),
		),

		factoryBrowser.NewTagSvgRect().X(0).Y(0).Width(200).Height(200).Fill("url(#gradient1)"),
		factoryBrowser.NewTagSvgRect().X(0).Y(0).Width(200).Height(200).Fill("url(#gradient2)").Style("transform: translateX(220px);"),
	)

	stage.Append(s1)

	<-done
}
