// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/maskUnits
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/maskUnits
//
//  <svg viewBox="0 0 100 100" xmlns="http://www.w3.org/2000/svg">
//    <mask id="myMask1" maskUnits="userSpaceOnUse"
//          x="20%" y="20%" width="60%" height="60%">
//      <rect   fill="black" x="0" y="0" width="100%" height="100%" />
//      <circle fill="white" cx="50" cy="50" r="35" />
//    </mask>
//
//    <mask id="myMask2" maskUnits="objectBoundingBox"
//          x="20%" y="20%" width="60%" height="60%">
//      <rect   fill="black" x="0" y="0" width="100%" height="100%" />
//      <circle fill="white" cx="50" cy="50" r="35" />
//    </mask>
//
//    <!-- Some reference rect to materialized the mask -->
//    <rect id="r1" x="0"  y="0"  width="45" height="45" />
//    <rect id="r2" x="0"  y="55" width="45" height="45" />
//    <rect id="r3" x="55" y="55" width="45" height="45" />
//    <rect id="r4" x="55" y="0"  width="45" height="45" />
//
//    <!-- The first 3 rect are masked with useSpaceOnUse units -->
//    <use mask="url(#myMask1)" xlink:href="#r1" fill="red" />
//    <use mask="url(#myMask1)" xlink:href="#r2" fill="red" />
//    <use mask="url(#myMask1)" xlink:href="#r3" fill="red" />
//
//    <!-- The last rect is masked with objectBoundingBox units -->
//    <use mask="url(#myMask2)" xlink:href="#r4" fill="red" />
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

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 100, 100}).Append(
		factoryBrowser.NewTagSvgMask().Id("myMask1").MaskUnits(html.KSvgUnitsUserSpaceOnUse).X(float32(0.2)).Y(float32(0.2)).Width(float32(0.6)).Height(float32(0.6)).Append(
			factoryBrowser.NewTagSvgRect().Fill(factoryColor.NewBlack()).X(0).Y(0).Width(float32(1.0)).Height(float32(1.0)),
			factoryBrowser.NewTagSvgCircle().Fill(factoryColor.NewWhite()).Cx(50).Cy(50).R(35),
		),

		factoryBrowser.NewTagSvgMask().Id("myMask2").MaskUnits(html.KSvgUnitsObjectBoundingBox).X(float32(0.2)).Y(float32(0.2)).Width(float32(0.6)).Height(float32(0.6)).Append(
			factoryBrowser.NewTagSvgRect().Fill(factoryColor.NewBlack()).X(0).Y(0).Width(float32(1.0)).Height(float32(1.0)),
			factoryBrowser.NewTagSvgCircle().Fill(factoryColor.NewWhite()).Cx(50).Cy(50).R(35),
		),

		// Some reference rect to materialized the mask
		factoryBrowser.NewTagSvgRect().Id("r1").X(0).Y(0).Width(45).Height(45),
		factoryBrowser.NewTagSvgRect().Id("r2").X(0).Y(55).Width(45).Height(45),
		factoryBrowser.NewTagSvgRect().Id("r3").X(55).Y(55).Width(45).Height(45),
		factoryBrowser.NewTagSvgRect().Id("r4").X(55).Y(0).Width(45).Height(45),

		// The first 3 rect are masked with useSpaceOnUse units
		factoryBrowser.NewTagSvgUse().Mask("url(#myMask1)").HRef("#r1").Fill(factoryColor.NewRed()),
		factoryBrowser.NewTagSvgUse().Mask("url(#myMask1)").HRef("#r2").Fill(factoryColor.NewRed()),
		factoryBrowser.NewTagSvgUse().Mask("url(#myMask1)").HRef("#r3").Fill(factoryColor.NewRed()),

		// The last rect is masked with objectBoundingBox units
		factoryBrowser.NewTagSvgUse().Mask("url(#myMask2)").HRef("#r4").Fill(factoryColor.NewRed()),
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
