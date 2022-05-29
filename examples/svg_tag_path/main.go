//go:build js
// +build js

//
package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
	"github.com/helmutkemper/iotmaker.webassembly/browser/stage"
)

func main() {

	// https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/kernelMatrix

	done := make(chan struct{}, 0)

	// browser stage
	var bs = stage.Stage{}
	bs.Init()

	factoryBrowser.NewTagSvg("svg1").
		ViewBox(0, 0, 100, 100).
		Width(100).
		Height(100).
		AppendToStage()

	path := new(html.SvgPath)
	path.MoveTo(10, 30).
		ArcCurve(20, 20, 0, 0, 1, 50, 30).
		ArcCurve(20, 20, 0, 0, 1, 90, 30).
		QuadraticBezierCurve(90, 60, 50, 90).
		QuadraticBezierCurve(10, 60, 10, 30).
		Close()

	factoryBrowser.NewTagSvgPath("path1").
		D(path).
		AppendById("svg1")

	factoryBrowser.NewTagSvg("svg2").
		ViewBox(0, 0, 100, 100).
		Width(100).
		Height(100).
		AppendToStage()

	path = new(html.SvgPath)
	path.M(10, 30).
		A(20, 20, 0, 0, 1, 50, 30).
		A(20, 20, 0, 0, 1, 90, 30).
		Q(90, 60, 50, 90).
		Q(10, 60, 10, 30).
		Z()

	factoryBrowser.NewTagSvgPath("path2").
		D(path).
		AppendById("svg2")

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
