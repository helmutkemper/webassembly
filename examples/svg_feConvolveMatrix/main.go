//go:build js
// +build js

//
package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/stage"
)

func main() {

	// https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/kernelMatrix

	done := make(chan struct{}, 0)

	// browser stage
	var bs = stage.Stage{}
	bs.Init()

	factoryBrowser.NewTagSvg("svg").
		ViewBox(0, 0, 420, 200).
		AppendToStage()

	factoryBrowser.NewTagSvgFilter("convolveMatrix1").
		X(0).
		Y(0).
		Width(1.0).
		Height(1.0).
		AppendById("svg")

	factoryBrowser.NewTagSvgFeConvolveMatrix("feMatrix1").
		KernelMatrix([]float64{1, 1, 0, 0, 0, 0, 0, 0, -1}).
		AppendById("convolveMatrix1")

	factoryBrowser.NewTagSvgFilter("convolveMatrix2").
		X(0).
		Y(0).
		Width(1.0).
		Height(1.0).
		AppendById("svg")

	factoryBrowser.NewTagSvgFeConvolveMatrix("feMatrix2").
		KernelMatrix([]float64{-1, 0, 0, 0, 0, 0, 0, 0, 1}).
		AppendById("convolveMatrix2")

	factoryBrowser.NewTagSvgImage("image1").
		XLinkHRef("//developer.mozilla.org/files/6457/mdn_logo_only_color.png").
		Width(200).
		Height(200).
		Style("filter:url(#convolveMatrix1);").
		AppendById("svg")

	factoryBrowser.NewTagSvgImage("image2").
		XLinkHRef("//developer.mozilla.org/files/6457/mdn_logo_only_color.png").
		Width(200).
		Height(200).
		Style("filter:url(#convolveMatrix2); transform:translateX(220px);").
		AppendById("svg")

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
