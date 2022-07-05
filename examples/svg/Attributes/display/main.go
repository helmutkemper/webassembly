// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/display
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/display
//
//  <svg viewBox="0 0 220 100" xmlns="http://www.w3.org/2000/svg">
//    <!-- Here the yellow rectangle is displayed -->
//    <rect x="0" y="0" width="100" height="100" fill="skyblue"></rect>
//    <rect x="20" y="20" width="60" height="60" fill="yellow"></rect>
//
//    <!-- Here the yellow rectangle is not displayed -->
//    <rect x="120" y="0" width="100" height="100" fill="skyblue"></rect>
//    <rect x="140" y="20" width="60" height="60" fill="yellow" display="none"></rect>
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

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 220, 100}).Append(

		// Here the yellow rectangle is displayed
		factoryBrowser.NewTagSvgRect().X(0).Y(0).Width(100).Height(100).Fill(factoryColor.NewSkyblue()),
		factoryBrowser.NewTagSvgRect().X(20).Y(20).Width(60).Height(60).Fill(factoryColor.NewYellow()),

		// Here the yellow rectangle is not displayed
		factoryBrowser.NewTagSvgRect().X(120).Y(0).Width(100).Height(100).Fill(factoryColor.NewSkyblue()),
		factoryBrowser.NewTagSvgRect().X(140).Y(20).Width(60).Height(60).Fill(factoryColor.NewYellow()).Display(nil),
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
