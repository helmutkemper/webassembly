// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/rx
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/rx
//
//  <svg viewBox="0 0 300 200" xmlns="http://www.w3.org/2000/svg">
//    <ellipse cx="50"  cy="50" rx="0"  ry="25" />
//    <ellipse cx="150" cy="50" rx="25" ry="25" />
//    <ellipse cx="250" cy="50" rx="50" ry="25" />
//
//    <rect x="20"  y="120" width="60" height="60" rx="0"   ry="15"/>
//    <rect x="120" y="120" width="60" height="60" rx="15"  ry="15"/>
//    <rect x="220" y="120" width="60" height="60" rx="150" ry="15"/>
//  </svg>

//go:build js
// +build js

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()
	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 300, 200}).Append(
		factoryBrowser.NewTagSvgEllipse().Cx(50).Cy(50).Rx(0).Ry(25),
		factoryBrowser.NewTagSvgEllipse().Cx(150).Cy(50).Rx(25).Ry(25),
		factoryBrowser.NewTagSvgEllipse().Cx(250).Cy(50).Rx(50).Ry(25),

		factoryBrowser.NewTagSvgRect().X(20).Y(120).Width(60).Height(60).Rx(0).Ry(15),
		factoryBrowser.NewTagSvgRect().X(120).Y(120).Width(60).Height(60).Rx(15).Ry(15),
		factoryBrowser.NewTagSvgRect().X(220).Y(120).Width(60).Height(60).Rx(150).Ry(15),
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
