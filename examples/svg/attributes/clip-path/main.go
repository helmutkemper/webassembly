// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/clip-path
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/clip-path
//
//  <svg viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
//    <clipPath id="myClip" clipPathUnits="objectBoundingBox">
//      <circle cx=".5" cy=".5" r=".5" />
//    </clipPath>
//
//    <!-- Top-left: Apply a custom defined clipping path -->
//    <rect x="1" y="1" width="8" height="8" stroke="green"
//          clip-path="url(#myClip)" />
//
//    <!-- Top-right: Apply a CSS basic shape on a fill-box
//         geometry. This is the same as having a custom clipping
//         path with a clipPathUnits set to objectBoundingBox -->
//    <rect x="11" y="1" width="8" height="8" stroke="green"
//          clip-path="circle() fill-box" />
//
//    <!-- Bottom-left -->
//    <rect x="1" y="11" width="8" height="8" stroke="green"
//          clip-path="circle() stroke-box" />
//
//    <!-- Bottom-right: Apply a CSS basic shape on a view-box
//         geometry. This is the same as having a custom clipping
//         path with a clipPathUnits set to userSpaceOnUse -->
//    <rect x="11" y="11" width="8" height="8" stroke="green"
//          clip-path="circle() view-box" />
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

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 20, 20}).Append(

		factoryBrowser.NewTagSvgClipPath().Id("myClip").ClipPathUnits(html.KSvgClipPathUnitsObjectBoundingBox).Append(
			factoryBrowser.NewTagSvgCircle().Cx(0.5).Cy(0.5).R(0.5),
		),

		// Top-left: Apply a custom defined clipping path
		factoryBrowser.NewTagSvgRect().X(1).Y(1).Width(8).Height(8).Stroke(factoryColor.NewGreen()).ClipPath("url(#myClip)"),

		// Top-right: Apply a CSS basic shape on a fill-box
		// geometry. This is the same as having a custom clipping
		// path with a clipPathUnits set to objectBoundingBox
		factoryBrowser.NewTagSvgRect().X(11).Y(1).Width(8).Height(8).Stroke(factoryColor.NewGreen()).ClipPath("circle() fill-box"),

		// Bottom-left
		factoryBrowser.NewTagSvgRect().X(1).Y(11).Width(8).Height(8).Stroke(factoryColor.NewGreen()).ClipPath("circle() stroke-box"),

		// Bottom-right: Apply a CSS basic shape on a view-box
		// geometry. This is the same as having a custom clipping
		// path with a clipPathUnits set to userSpaceOnUse
		factoryBrowser.NewTagSvgRect().X(11).Y(11).Width(8).Height(8).Stroke(factoryColor.NewGreen()).ClipPath("circle() view-box"),
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
