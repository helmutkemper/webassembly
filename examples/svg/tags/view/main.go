// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Element/view
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Element/view
//
//  <svg viewBox="0 0 300 100" width="300" height="100"
//      xmlns="http://www.w3.org/2000/svg">
//
//    <view id="one" viewBox="0 0 100 100" />
//    <circle cx="50" cy="50" r="40" fill="red" />
//
//    <view id="two" viewBox="100 0 100 100" />
//    <circle cx="150" cy="50" r="40" fill="green" />
//
//    <view id="three" viewBox="200 0 100 100" />
//    <circle cx="250" cy="50" r="40" fill="blue" />
//  </svg>

//go:build js
// +build js

// fixme: bug. este exemplo procura um arquivo svg?

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/platform/factoryColor"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 300, 100}).Id("svgImg").Width(300).Height(100).Append(

		factoryBrowser.NewTagSvgView().Id("one").ViewBox([]float64{0, 0, 100, 100}),
		factoryBrowser.NewTagSvgCircle().Cx(50).Cy(50).R(40).Fill(factoryColor.NewRed()),

		factoryBrowser.NewTagSvgView().Id("one").ViewBox([]float64{100, 0, 100, 100}),
		factoryBrowser.NewTagSvgCircle().Cx(150).Cy(50).R(40).Fill(factoryColor.NewGreen()),

		factoryBrowser.NewTagSvgView().Id("one").ViewBox([]float64{200, 0, 100, 100}),
		factoryBrowser.NewTagSvgCircle().Cx(250).Cy(50).R(40).Fill(factoryColor.NewGreen()),
	)

	stage.Append(s1).Append("<img src=\"url(#svgImg)\" alt=\"three circles\" width=\"300\" height=\"100\" />\n<br />\n<img src=\"#three\" alt=\"blue circle\" width=\"100\" height=\"100\" />")

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
//
//
//
//
//
//
//
//
