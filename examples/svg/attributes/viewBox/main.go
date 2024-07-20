// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/viewBox
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/viewBox
//
//  <svg viewBox="0 0 100 100" xmlns="http://www.w3.org/2000/svg">
//    <!--
//    with relative unit such as percentage, the visual size
//    of the square looks unchanged regardless of the viewBox
//    -->
//    <rect x="0" y="0" width="100%" height="100%"/>
//
//    <!--
//    with a large viewBox the circle looks small
//    as it is using user units for the r attribute:
//    4 resolved against 100 as set in the viewBox
//    -->
//    <circle cx="50%" cy="50%" r="4" fill="white"/>
//  </svg>
//
//  <svg viewBox="0 0 10 10" xmlns="http://www.w3.org/2000/svg">
//    <!--
//    with relative unit such as percentage, the visual size
//    of the square looks unchanged regardless of the viewBox
//    -->
//    <rect x="0" y="0" width="100%" height="100%"/>
//
//    <!--
//    with a small viewBox the circle looks large
//    as it is using user units for the r attribute:
//    4 resolved against 10 as set in the viewBox
//    -->
//    <circle cx="50%" cy="50%" r="4" fill="white"/>
//  </svg>
//
//  <svg viewBox="-5 -5 10 10" xmlns="http://www.w3.org/2000/svg">
//    <!--
//    The point of coordinate 0,0 is now in the center of the viewport,
//    and 100% is still resolve to a width or height of 10 user units so
//    the rectangle looks shifted to the bottom/right corner of the viewport
//    -->
//    <rect x="0" y="0" width="100%" height="100%"/>
//
//    <!--
//    With the point of coordinate 0,0 in the center of the viewport the
//    value 50% is resolve to 5 which means the center of the circle is
//    in the bottom/right corner of the viewport.
//    -->
//    <circle cx="50%" cy="50%" r="4" fill="white"/>
//  </svg>

//go:build js

// bug: o original não apararece nada

package main

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/platform/factoryColor"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()
	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 100, 100}).Append(
		// with relative unit such as percentage, the visual size
		// of the square looks unchanged regardless of the viewBox
		factoryBrowser.NewTagSvgRect().X(0).Y(0).Width(float32(1)).Height(float32(1)),

		// with a large viewBox the circle looks small
		// as it is using user units for the r attribute:
		// 4 resolved against 100 as set in the viewBox
		factoryBrowser.NewTagSvgCircle().Cx(float32(0.5)).Cy(float32(0.5)).R(4).Fill(factoryColor.NewWhite()),
	)

	s2 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 10, 10}).Append(
		// with relative unit such as percentage, the visual size
		// of the square looks unchanged regardless of the viewBox
		factoryBrowser.NewTagSvgRect().X(0).Y(0).Width(float32(1)).Height(float32(1)),

		// with a small viewBox the circle looks large
		// as it is using user units for the r attribute:
		// 4 resolved against 10 as set in the viewBox
		factoryBrowser.NewTagSvgCircle().Cx(float32(0.5)).Cy(float32(0.5)).R(4).Fill(factoryColor.NewWhite()),
	)

	s3 := factoryBrowser.NewTagSvg().ViewBox([]float64{-5, -5, 10, 10}).Append(
		// The point of coordinate 0,0 is now in the center of the viewport,
		// and 100% is still resolve to a width or height of 10 user units so
		// the rectangle looks shifted to the bottom/right corner of the viewport
		factoryBrowser.NewTagSvgRect().X(0).Y(0).Width(float32(1)).Height(float32(1)),

		// With the point of coordinate 0,0 in the center of the viewport the
		// value 50% is resolve to 5 which means the center of the circle is
		// in the bottom/right corner of the viewport.
		factoryBrowser.NewTagSvgCircle().Cx(float32(0.5)).Cy(float32(0.5)).R(4).Fill(factoryColor.NewWhite()),
	)

	stage.Append(s1)
	stage.Append(s2)
	stage.Append(s3)

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
