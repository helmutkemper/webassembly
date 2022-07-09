// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/spreadMethod
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/spreadMethod
//
//  <svg width="340" height="120" xmlns="http://www.w3.org/2000/svg">
//    <defs>
//     <radialGradient id="RadialPadGradient"
//                      cx="75%" cy="25%" r="33%"
//                      fx="64%" fy="18%" fr="17%">
//        <stop offset="0%"  stop-color="fuchsia"/>
//        <stop offset="100%" stop-color="orange"/>
//      </radialGradient>
//      <radialGradient id="RadialReflectGradient"
//                      spreadMethod="reflect"
//                      cx="75%" cy="25%" r="33%"
//                      fx="64%" fy="18%" fr="17%">
//        <stop offset="0%"  stop-color="fuchsia"/>
//        <stop offset="100%" stop-color="orange"/>
//      </radialGradient>
//      <radialGradient id="RadialRepeatGradient"
//                      spreadMethod="repeat"
//                      cx="75%" cy="25%" r="33%"
//                      fx="64%" fy="18%" fr="17%">
//        <stop offset="0%"  stop-color="fuchsia"/>
//        <stop offset="100%" stop-color="orange"/>
//      </radialGradient>
//    </defs>
//
//    <rect fill="url(#RadialPadGradient)"
//          x="10" y="10" width="100" height="100"/>
//    <rect fill="url(#RadialReflectGradient)"
//          x="120" y="10" width="100" height="100"/>
//    <rect fill="url(#RadialRepeatGradient)"
//          x="230" y="10" width="100" height="100"/>
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
	s1 := factoryBrowser.NewTagSvg().Width(340).Height(120).Append(
		factoryBrowser.NewTagSvgRadialGradient().Id("RadialPadGradient").Cx(float32(0.75)).Cy(0.25).R(float32(0.33)).Fx(float32(0.64)).Fy(float32(0.18)).Fr(float32(0.17)).Append(
			factoryBrowser.NewTagSvgStop().Offset(float32(0)).StopColor(factoryColor.NewFuchsia()),
			factoryBrowser.NewTagSvgStop().Offset(float32(1)).StopColor(factoryColor.NewOrange()),
		),
		factoryBrowser.NewTagSvgRadialGradient().Id("RadialReflectGradient").SpreadMethod(html.KSvgSpreadMethodReflect).Cx(float32(0.75)).Cy(0.25).R(float32(0.33)).Fx(float32(0.64)).Fy(float32(0.18)).Fr(float32(0.17)).Append(
			factoryBrowser.NewTagSvgStop().Offset(float32(0)).StopColor(factoryColor.NewFuchsia()),
			factoryBrowser.NewTagSvgStop().Offset(float32(1)).StopColor(factoryColor.NewOrange()),
		),
		factoryBrowser.NewTagSvgRadialGradient().Id("RadialRepeatGradient").SpreadMethod(html.KSvgSpreadMethodRepeat).Cx(float32(0.75)).Cy(0.25).R(float32(0.33)).Fx(float32(0.64)).Fy(float32(0.18)).Fr(float32(0.17)).Append(
			factoryBrowser.NewTagSvgStop().Offset(float32(0)).StopColor(factoryColor.NewFuchsia()),
			factoryBrowser.NewTagSvgStop().Offset(float32(1)).StopColor(factoryColor.NewOrange()),
		),

		factoryBrowser.NewTagSvgRect().Fill("url(#RadialPadGradient)").X(10).Y(10).Width(100).Height(100),
		factoryBrowser.NewTagSvgRect().Fill("url(#RadialReflectGradient)").X(120).Y(10).Width(100).Height(100),
		factoryBrowser.NewTagSvgRect().Fill("url(#RadialRepeatGradient)").X(230).Y(10).Width(100).Height(100),
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
