//go:build js
// +build js

//
package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/stage"
	"github.com/helmutkemper/iotmaker.webassembly/platform/factoryColor"
)

func main() {

	done := make(chan struct{}, 0)

	// browser stage
	var browserStage = stage.Stage{}
	browserStage.Init()

	//<svg viewBox="0 0 300 100" xmlns="http://www.w3.org/2000/svg" stroke="red" fill="grey">
	//  <circle cx="50" cy="50" r="40" />
	//  <circle cx="150" cy="50" r="4" />
	//
	//  <svg viewBox="0 0 10 10" x="200" width="100">
	//    <circle cx="5" cy="5" r="4" />
	//  </svg>
	//</svg>

	factoryBrowser.NewTagSvg("svg").
		ViewBox([]float64{0, 0, 300, 100}).
		Stroke(factoryColor.NewRed()).
		Fill(factoryColor.NewGrey()).
		AppendToStage()

	factoryBrowser.NewTagSvgCircle("cir1").
		Cx(50).
		Cy(51).
		R(40).
		AppendById("svg")

	factoryBrowser.NewTagSvgCircle("cir2").
		Cx(150).
		Cy(50).
		R(4).
		AppendById("svg")

	factoryBrowser.NewTagSvg("view").
		ViewBox([]float64{0, 0, 10, 10}).
		X(200).
		Width(100).
		AppendById("svg")

	factoryBrowser.NewTagSvgCircle("cir3").
		Cx(5).
		Cy(5).
		R(4).
		AppendById("view")

	<-done
}
