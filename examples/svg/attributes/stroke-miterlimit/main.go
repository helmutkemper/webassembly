// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/stroke-miterlimit
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/stroke-miterlimit
//
//  <svg viewBox="0 0 38 30" xmlns="http://www.w3.org/2000/svg">
//    <!-- Impact of the default miter limit -->
//    <path stroke="black" fill="none" stroke-linejoin="miter" id="p1"
//          d="M1,9 l7   ,-3 l7   ,3
//             m2,0 l3.5 ,-3 l3.5 ,3
//             m2,0 l2   ,-3 l2   ,3
//             m2,0 l0.75,-3 l0.75,3
//             m2,0 l0.5 ,-3 l0.5 ,3" />
//
//    <!-- Impact of the smallest miter limit (1) -->
//    <path stroke="black" fill="none" stroke-linejoin="miter"
//          stroke-miterlimit="1" id="p2"
//          d="M1,19 l7   ,-3 l7   ,3
//             m2, 0 l3.5 ,-3 l3.5 ,3
//             m2, 0 l2   ,-3 l2   ,3
//             m2, 0 l0.75,-3 l0.75,3
//             m2, 0 l0.5 ,-3 l0.5 ,3" />
//
//    <!-- Impact of a large miter limit (8) -->
//    <path stroke="black" fill="none" stroke-linejoin="miter"
//          stroke-miterlimit="8" id="p3"
//          d="M1,29 l7   ,-3 l7   ,3
//             m2, 0 l3.5 ,-3 l3.5 ,3
//             m2, 0 l2   ,-3 l2   ,3
//             m2, 0 l0.75,-3 l0.75,3
//             m2, 0 l0.5 ,-3 l0.5 ,3" />
//
//    <!-- the following pink lines highlight the position of the path for each stroke -->
//    <path stroke="pink" fill="none" stroke-width="0.05"
//          d="M1, 9 l7,-3 l7,3 m2,0 l3.5,-3 l3.5,3 m2,0 l2,-3 l2,3 m2,0 l0.75,-3 l0.75,3 m2,0 l0.5,-3 l0.5,3
//             M1,19 l7,-3 l7,3 m2,0 l3.5,-3 l3.5,3 m2,0 l2,-3 l2,3 m2,0 l0.75,-3 l0.75,3 m2,0 l0.5,-3 l0.5,3
//             M1,29 l7,-3 l7,3 m2,0 l3.5,-3 l3.5,3 m2,0 l2,-3 l2,3 m2,0 l0.75,-3 l0.75,3 m2,0 l0.5,-3 l0.5,3" />
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
	s1 := factoryBrowser.NewTagSvg().Width(640).Height(380).Append(
		factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 30}).Append(
			// Impact of the default miter limit
			factoryBrowser.NewTagSvgPath().Stroke(factoryColor.NewBlack()).Fill(nil).
				StrokeLineJoin(html.KSvgStrokeLinejoinMiter).Id("p1").D(factoryBrowser.NewPath().
				M(1, 9).Ld(7, -3).Ld(7, 3).
				Md(2, 0).Ld(3.5, -3).Ld(3.5, 3).
				Md(2, 0).Ld(2, -3).Ld(2, 3).
				Md(2, 0).Ld(0.75, -3).Ld(0.75, 3).
				Md(2, 0).Ld(0.5, -3).Ld(0.5, 3)),

			// Impact of the smallest miter limit (1)
			factoryBrowser.NewTagSvgPath().Stroke(factoryColor.NewBlack()).Fill(nil).
				StrokeLineJoin(html.KSvgStrokeLinejoinMiter).StrokeMiterLimit(1).Id("p2").D(factoryBrowser.NewPath().
				M(1, 19).Ld(7, -3).Ld(7, 3).
				Md(2, 0).Ld(3.5, -3).Ld(3.5, 3).
				Md(2, 0).Ld(2, -3).Ld(2, 3).
				Md(2, 0).Ld(0.75, -3).Ld(0.75, 3).
				Md(2, 0).Ld(0.5, -3).Ld(0.5, 3)),

			// Impact of a large miter limit (8)
			factoryBrowser.NewTagSvgPath().Stroke(factoryColor.NewBlack()).Fill(nil).
				StrokeLineJoin(html.KSvgStrokeLinejoinMiter).StrokeMiterLimit(8).Id("p3").D(factoryBrowser.NewPath().
				M(1, 29).Ld(7, -3).Ld(7, 3).
				Md(2, 0).Ld(3.5, -3).Ld(3.5, 3).
				Md(2, 0).Ld(2, -3).Ld(2, 3).
				Md(2, 0).Ld(0.75, -3).Ld(0.75, 3).
				Md(2, 0).Ld(0.5, -3).Ld(0.5, 3)),

			// the following pink lines highlight the position of the path for each stroke
			factoryBrowser.NewTagSvgPath().Stroke(factoryColor.NewPink()).Fill(nil).StrokeWidth(0.05).D(factoryBrowser.NewPath().
				M(1, 9).Ld(7, -3).Ld(7, 3).Md(2, 0).Ld(3.5, -3).Ld(3.5, 3).Md(2, 0).Ld(2, -3).Ld(2, 3).Md(2, 0).Ld(0.75, -3).Ld(0.75, 3).Md(2, 0).Ld(0.5, -3).Ld(0.5, 3).
				M(1, 19).Ld(7, -3).Ld(7, 3).Md(2, 0).Ld(3.5, -3).Ld(3.5, 3).Md(2, 0).Ld(2, -3).Ld(2, 3).Md(2, 0).Ld(0.75, -3).Ld(0.75, 3).Md(2, 0).Ld(0.5, -3).Ld(0.5, 3).
				M(1, 29).Ld(7, -3).Ld(7, 3).Md(2, 0).Ld(3.5, -3).Ld(3.5, 3).Md(2, 0).Ld(2, -3).Ld(2, 3).Md(2, 0).Ld(0.75, -3).Ld(0.75, 3).Md(2, 0).Ld(0.5, -3).Ld(0.5, 3),
			),
		),
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
