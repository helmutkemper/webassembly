// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/attributeName
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/attributeName
//
//  <svg viewBox="0 0 250 250" xmlns="http://www.w3.org/2000/svg">
//    <rect x="50" y="50" width="100" height="100">
//      <animate attributeType="XML" attributeName="y" from="0" to="50"
//          dur="5s" repeatCount="indefinite"/>
//    </rect>
//  </svg>

//go:build js

package main

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/html"
	"time"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 250, 250}).Append(
		factoryBrowser.NewTagSvgRect().X(50).Y(50).Width(100).Height(100).Append(
			factoryBrowser.NewTagSvgAnimate().AttributeName("y").From(0).To(50).Dur(5 * time.Second).RepeatCount(html.KSvgDurIndefinite),
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
