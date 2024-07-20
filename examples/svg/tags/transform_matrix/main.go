// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/transform
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/transform
//
//  <svg viewBox="0 0 200 200" xmlns="http://www.w3.org/2000/svg">
//    <rect x="10" y="10" width="30" height="20" fill="green" />
//
//    <!--
//    In the following example we are applying the matrix:
//    [a c e]    [3 -1 30]
//    [b d f] => [1  3 40]
//    [0 0 1]    [0  0  1]
//
//    which transform the rectangle as such:
//
//    top left corner: oldX=10 oldY=10
//    newX = a * oldX + c * oldY + e = 3 * 10 - 1 * 10 + 30 = 50
//    newY = b * oldX + d * oldY + f = 1 * 10 + 3 * 10 + 40 = 80
//
//    top right corner: oldX=40 oldY=10
//    newX = a * oldX + c * oldY + e = 3 * 40 - 1 * 10 + 30 = 140
//    newY = b * oldX + d * oldY + f = 1 * 40 + 3 * 10 + 40 = 110
//
//    bottom left corner: oldX=10 oldY=30
//    newX = a * oldX + c * oldY + e = 3 * 10 - 1 * 30 + 30 = 30
//    newY = b * oldX + d * oldY + f = 1 * 10 + 3 * 30 + 40 = 140
//
//    bottom right corner: oldX=40 oldY=30
//    newX = a * oldX + c * oldY + e = 3 * 40 - 1 * 30 + 30 = 120
//    newY = b * oldX + d * oldY + f = 1 * 40 + 3 * 30 + 40 = 170
//    -->
//    <rect x="10" y="10" width="30" height="20" fill="red"
//          transform="matrix(3 1 -1 3 30 40)" />
//  </svg>

//go:build js

package main

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/platform/factoryColor"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 200, 200}).XmlnsXLink("http://www.w3.org/1999/xlink").Append(
		factoryBrowser.NewTagSvgRect().X(10).Y(10).Width(30).Height(20).Fill(factoryColor.NewGreen()),

		//    In the following example we are applying the matrix:
		//    [a c e]    [3 -1 30]
		//    [b d f] => [1  3 40]
		//    [0 0 1]    [0  0  1]
		//
		//    which transform the rectangle as such:
		//
		//    top left corner: oldX=10 oldY=10
		//    newX = a * oldX + c * oldY + e = 3 * 10 - 1 * 10 + 30 = 50
		//    newY = b * oldX + d * oldY + f = 1 * 10 + 3 * 10 + 40 = 80
		//
		//    top right corner: oldX=40 oldY=10
		//    newX = a * oldX + c * oldY + e = 3 * 40 - 1 * 10 + 30 = 140
		//    newY = b * oldX + d * oldY + f = 1 * 40 + 3 * 10 + 40 = 110
		//
		//    bottom left corner: oldX=10 oldY=30
		//    newX = a * oldX + c * oldY + e = 3 * 10 - 1 * 30 + 30 = 30
		//    newY = b * oldX + d * oldY + f = 1 * 10 + 3 * 30 + 40 = 140
		//
		//    bottom right corner: oldX=40 oldY=30
		//    newX = a * oldX + c * oldY + e = 3 * 40 - 1 * 30 + 30 = 120
		//    newY = b * oldX + d * oldY + f = 1 * 40 + 3 * 30 + 40 = 170

		factoryBrowser.NewTagSvgRect().X(10).Y(10).Width(30).Height(20).Fill(factoryColor.NewRed()).Transform(
			factoryBrowser.NewTransform().Matrix(3, 1, -1, 3, 30, 40),
		),
	)

	stage.Append(s1)

	<-done
}
