// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/transform
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/transform
//
//  <svg viewBox="-12 -2 34 14" xmlns="http://www.w3.org/2000/svg">
//    <rect x="0" y="0" width="10" height="10" />
//
//    <!-- rotation is done around the point 0,0 -->
//    <rect x="0" y="0" width="10" height="10" fill="red"
//          transform="rotate(100)" />
//
//    <!-- rotation is done around the point 10,10 -->
//    <rect x="0" y="0" width="10" height="10" fill="green"
//          transform="rotate(100,10,10)" />
//  </svg>

//go:build js

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/platform/factoryColor"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{-12, -2, 34, 14}).Append(
		factoryBrowser.NewTagSvgRect().X(0).Y(0).Width(10).Height(10),
		factoryBrowser.NewTagSvgRect().X(0).Y(0).Width(10).Height(10).Fill(factoryColor.NewRed()).Transform(factoryBrowser.NewTransform().RotateAngle(100)),
		factoryBrowser.NewTagSvgRect().X(0).Y(0).Width(10).Height(10).Fill(factoryColor.NewGreen()).Transform(factoryBrowser.NewTransform().Rotate(100, 10, 10)),
	)

	stage.Append(s1)

	<-done
}
