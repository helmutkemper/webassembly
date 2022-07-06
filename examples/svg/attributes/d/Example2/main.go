// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/d
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/d
//
//  <svg viewBox="0 0 100 100" xmlns="http://www.w3.org/2000/svg">
//    <path fill="none" stroke="red"
//      d="M 10,10 h 10
//         m  0,10 h 10
//         m  0,10 h 10
//         M 40,20 h 10
//         m  0,10 h 10
//         m  0,10 h 10
//         m  0,10 h 10
//         M 50,50 h 10
//         m-20,10 h 10
//         m-20,10 h 10
//         m-20,10 h 10" />
//  </svg>

//go:build js
// +build js

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/platform/factoryColor"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 100, 100}).Append(

		factoryBrowser.NewTagSvgPath().Fill(nil).Stroke(factoryColor.NewRed()).D(factoryBrowser.NewPath().M(10, 10).Hd(10).Md(0, 10).Hd(10).Md(0, 10).Hd(10).M(40, 20).Hd(10).Md(0, 10).Hd(10).Md(0, 10).Hd(10).Md(0, 10).Hd(10).M(50, 50).Hd(10).Md(-20, 10).Hd(10).Md(-20, 10).Hd(10).Md(-20, 10).Hd(10)),
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
