// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/To
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/To
//
//  <svg viewBox="0 0 200 200" xmlns="http://www.w3.org/2000/svg">
//    <rect x="10" y="10" width="100" height="100">
//      <animate attributeType="XML" attributeName="width" fill="freeze"
//          from="100" to="150" dur="3s"/>
//    </rect>
//  </svg>

//go:build js

// bug: o original não apararece nada

package main

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"time"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()
	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 200, 200}).Append(
		factoryBrowser.NewTagSvgRect().X(10).Y(10).Width(100).Height(100).Append(
			factoryBrowser.NewTagSvgAnimate().AttributeName("width").Fill("freeze").From(100).To(150).Dur(3 * time.Second),
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
