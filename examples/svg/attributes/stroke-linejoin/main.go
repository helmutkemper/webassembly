// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/stroke-linejoin
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/stroke-linejoin
//
//  <svg viewBox="0 0 18 12" xmlns="http://www.w3.org/2000/svg">
//    <!--
//    Upper left path:
//    Effect of the "miter" value
//    -->
//    <path d="M1,5 a2,2 0,0,0 2,-3 a3,3 0 0 1 2,3.5" stroke="black" fill="none"
//          stroke-linejoin="miter" />
//
//    <!--
//    Center path:
//    Effect of the "round" value
//    -->
//    <path d="M7,5 a2,2 0,0,0 2,-3 a3,3 0 0 1 2,3.5" stroke="black" fill="none"
//          stroke-linejoin="round" />
//
//    <!--
//    Upper right path:
//    Effect of the "bevel" value
//    -->
//    <path d="M13,5 a2,2 0,0,0 2,-3 a3,3 0 0 1 2,3.5" stroke="black" fill="none"
//          stroke-linejoin="bevel" />
//
//    <!--
//    Bottom left path:
//    Effect of the "miter-clip" value
//    with fallback to "miter" if not supported.
//    -->
//    <path d="M3,11 a2,2 0,0,0 2,-3 a3,3 0 0 1 2,3.5" stroke="black" fill="none"
//          stroke-linejoin="miter-clip" />
//
//    <!--
//    Bottom right path:
//    Effect of the "arcs" value
//    with fallback to "miter" if not supported.
//    -->
//    <path d="M9,11 a2,2 0,0,0 2,-3 a3,3 0 0 1 2,3.5" stroke="black" fill="none"
//          stroke-linejoin="arcs" />
//
//    <!--
//    the following pink lines highlight the
//    position of the path for each stroke
//    -->
//    <g id="highlight">
//      <path d="M1,5 a2,2 0,0,0 2,-3 a3,3 0 0 1 2,3.5"
//            stroke="pink" fill="none" stroke-width="0.025" />
//      <circle cx="1" cy="5"   r="0.05" fill="pink" />
//      <circle cx="3" cy="2"   r="0.05" fill="pink" />
//      <circle cx="5" cy="5.5" r="0.05" fill="pink" />
//    </g>
//    <use xlink:href="#highlight" x="6" />
//    <use xlink:href="#highlight" x="12" />
//    <use xlink:href="#highlight" x="2" y="6" />
//    <use xlink:href="#highlight" x="8" y="6" />
//  </svg>

//go:build js
// +build js

// bug: o original não apararece nada

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
	"github.com/helmutkemper/iotmaker.webassembly/platform/factoryColor"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()
	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 18, 12}).Append(
		// Upper left path:
		// Effect of the "miter" value
		factoryBrowser.NewTagSvgPath().D(factoryBrowser.NewPath().M(1, 5).Ad(2, 2, 0, 0, 0, 2, -3).Ad(3, 3, 0, 0, 1, 2, 3.5)).Stroke(factoryColor.NewBlack()).Fill(nil).StrokeLineJoin(html.KSvgStrokeLinejoinMiter),

		// Center path:
		// Effect of the "round" value
		factoryBrowser.NewTagSvgPath().D(factoryBrowser.NewPath().M(7, 5).Ad(2, 2, 0, 0, 0, 2, -3).Ad(3, 3, 0, 0, 1, 2, 3.5)).Stroke(factoryColor.NewBlack()).Fill(nil).StrokeLineJoin(html.KSvgStrokeLinejoinRound),

		// Upper right path:
		// Effect of the "bevel" value
		factoryBrowser.NewTagSvgPath().D(factoryBrowser.NewPath().M(13, 5).Ad(2, 2, 0, 0, 0, 2, -3).Ad(3, 3, 0, 0, 1, 2, 3.5)).Stroke(factoryColor.NewBlack()).Fill(nil).StrokeLineJoin(html.KSvgStrokeLinejoinBevel),

		// Bottom left path:
		// Effect of the "miter-clip" value
		// with fallback to "miter" if not supported.
		factoryBrowser.NewTagSvgPath().D(factoryBrowser.NewPath().M(3, 11).Ad(2, 2, 0, 0, 0, 2, -3).Ad(3, 3, 0, 0, 1, 2, 3.5)).Stroke(factoryColor.NewBlack()).Fill(nil).StrokeLineJoin(html.KSvgStrokeLinejoinMiterClip),

		// Bottom right path:
		// Effect of the "arcs" value
		// with fallback to "miter" if not supported.
		factoryBrowser.NewTagSvgPath().D(factoryBrowser.NewPath().M(9, 11).Ad(2, 2, 0, 0, 0, 2, -3).Ad(3, 3, 0, 0, 1, 2, 3.5)).Stroke(factoryColor.NewBlack()).Fill(nil).StrokeLineJoin(html.KSvgStrokeLinejoinArcs),

		// the following pink lines highlight the
		// position of the path for each stroke
		factoryBrowser.NewTagSvgG().Id("highlight").Append(
			factoryBrowser.NewTagSvgPath().D(factoryBrowser.NewPath().M(1, 5).Ad(2, 2, 0, 0, 0, 2, -3).Ad(3, 3, 0, 0, 1, 2, 3.5)).Stroke(factoryColor.NewPink()).Fill(nil).StrokeWidth(0.025),
			factoryBrowser.NewTagSvgCircle().Cx(1).Cy(5).R(0.05).Fill(factoryColor.NewPink()),
			factoryBrowser.NewTagSvgCircle().Cx(3).Cy(2).R(0.05).Fill(factoryColor.NewPink()),
			factoryBrowser.NewTagSvgCircle().Cx(5).Cy(5.5).R(0.05).Fill(factoryColor.NewPink()),
		),
		factoryBrowser.NewTagSvgUse().HRef("#highlight").X(6),
		factoryBrowser.NewTagSvgUse().HRef("#highlight").X(12),
		factoryBrowser.NewTagSvgUse().HRef("#highlight").X(2).Y(6),
		factoryBrowser.NewTagSvgUse().HRef("#highlight").X(8).Y(6),
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
