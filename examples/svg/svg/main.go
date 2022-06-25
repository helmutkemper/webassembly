//go:build js
// +build js

//
package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/platform/factoryColor"
)

func main() {

	done := make(chan struct{}, 0)

	// browser stage
	stage := factoryBrowser.NewStage()

	// https://developer.mozilla.org/pt-BR/docs/Web/SVG/Element/svg
	//
	//<svg viewBox="0 0 300 100" xmlns="http://www.w3.org/2000/svg" stroke="red" fill="grey">
	//  <circle cx="50" cy="50" r="40" />
	//  <circle cx="150" cy="50" r="4" />
	//
	//  <svg viewBox="0 0 10 10" x="200" width="100">
	//    <circle cx="5" cy="5" r="4" />
	//  </svg>
	//</svg>

	//<svg viewBox="0 0 300 100" xmlns="http://www.w3.org/2000/svg" stroke="red" fill="grey">
	s1 := factoryBrowser.NewTagSvg("svg").
		ViewBox([]float64{0, 0, 300, 100}).
		Stroke(factoryColor.NewRed()).
		Fill(factoryColor.NewGrey()).
		Append(

			//  <circle cx="50" cy="50" r="40" />
			factoryBrowser.NewTagSvgCircle("cir1").
				Cx(50).
				Cy(50).
				R(40),

			//  <circle cx="150" cy="50" r="4" />
			factoryBrowser.NewTagSvgCircle("cir2").
				Cx(150).
				Cy(50).
				R(4),

			//  <svg viewBox="0 0 10 10" x="200" width="100">
			factoryBrowser.NewTagSvg("view").
				ViewBox([]float64{0, 0, 10, 10}).
				X(200).
				Width(100).
				Append(

					//    <circle cx="5" cy="5" r="4" />
					factoryBrowser.NewTagSvgCircle("cir3").
						Cx(5).
						Cy(5).
						R(4),
				),
		)

	stage.Append(s1)

	<-done
}
