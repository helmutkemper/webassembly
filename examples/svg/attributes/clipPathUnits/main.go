// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/clipPathUnits
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/clipPathUnits
//
//  <svg viewBox="0 0 100 100">
//    <clipPath id="myClip1" clipPathUnits="userSpaceOnUse">
//      <circle cx="50" cy="50" r="35" />
//    </clipPath>
//
//    <clipPath id="myClip2" clipPathUnits="objectBoundingBox">
//      <circle cx=".5" cy=".5" r=".35" />
//    </clipPath>
//
//    <!-- Some reference rect to materialized to clip path -->
//    <rect id="r1" x="0"  y="0"  width="45" height="45" />
//    <rect id="r2" x="0"  y="55" width="45" height="45" />
//    <rect id="r3" x="55" y="55" width="45" height="45" />
//    <rect id="r4" x="55" y="0"  width="45" height="45" />
//
//    <!-- The first 3 rect are clipped with useSpaceOnUse units -->
//    <use clip-path="url(#myClip1)" xlink:href="#r1" fill="red" />
//    <use clip-path="url(#myClip1)" xlink:href="#r2" fill="red" />
//    <use clip-path="url(#myClip1)" xlink:href="#r3" fill="red" />
//
//    <!-- The last rect is clipped with objectBoundingBox units -->
//    <use clip-path="url(#myClip2)" xlink:href="#r4" fill="red" />
//  </svg>

//go:build js

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
	"github.com/helmutkemper/iotmaker.webassembly/platform/factoryColor"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().Width(120).Height(220).ViewBox([]float64{0, 0, 100, 100}).Append(

		factoryBrowser.NewTagSvgClipPath().Id("myClip1").ClipPathUnits(html.KSvgClipPathUnitsUserSpaceOnUse).Append(
			factoryBrowser.NewTagSvgCircle().Cx(50).Cy(50).R(35),
		),

		factoryBrowser.NewTagSvgClipPath().Id("myClip2").ClipPathUnits(html.KSvgClipPathUnitsObjectBoundingBox).Append(
			factoryBrowser.NewTagSvgCircle().Cx(0.5).Cy(0.5).R(0.35),
		),

		// Some reference rect to materialized to clip path
		factoryBrowser.NewTagSvgRect().Id("r1").X(0).Y(0).Width(45).Height(45),
		factoryBrowser.NewTagSvgRect().Id("r2").X(0).Y(55).Width(45).Height(45),
		factoryBrowser.NewTagSvgRect().Id("r3").X(55).Y(55).Width(45).Height(45),
		factoryBrowser.NewTagSvgRect().Id("r4").X(55).Y(0).Width(45).Height(45),

		// The first 3 rect are clipped with useSpaceOnUse units
		factoryBrowser.NewTagSvgUse().ClipPath("url(#myClip1)").HRef("#r1").Fill(factoryColor.NewRed()),
		factoryBrowser.NewTagSvgUse().ClipPath("url(#myClip1)").HRef("#r2").Fill(factoryColor.NewRed()),
		factoryBrowser.NewTagSvgUse().ClipPath("url(#myClip1)").HRef("#r3").Fill(factoryColor.NewRed()),

		// The last rect is clipped with objectBoundingBox units
		factoryBrowser.NewTagSvgUse().ClipPath("url(#myClip2)").HRef("#r4").Fill(factoryColor.NewRed()),
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
