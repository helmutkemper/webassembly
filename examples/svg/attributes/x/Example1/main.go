// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/x
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/x
//
//  <svg viewBox="0 0 300 100" xmlns="http://www.w3.org/2000/svg">
//    <rect x="20"  y="20" width="60" height="60" />
//    <rect x="120" y="20" width="60" height="60" />
//    <rect x="220" y="20" width="60" height="60" />
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

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 300, 100}).Append(
		factoryBrowser.NewTagSvgRect().X(20).Y(20).Width(60).Height(60),
		factoryBrowser.NewTagSvgRect().X(120).Y(20).Width(60).Height(60),
		factoryBrowser.NewTagSvgRect().X(220).Y(20).Width(60).Height(60),
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
