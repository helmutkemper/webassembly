// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/repeatDur
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/repeatDur
//
//  <svg viewBox="0 0 220 150" xmlns="http://www.w3.org/2000/svg">
//    <rect x="0" y="0" width="100" height="100">
//      <animate attributeType="XML" attributeName="y" from="0" to="50"
//          dur="1s" repeatDur="5s"/>
//    </rect>
//    <rect x="120" y="0" width="100" height="100">
//      <animate attributeType="XML" attributeName="y" from="0" to="50"
//          dur="1s" repeatDur="indefinite"/>
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
	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 220, 150}).Append(
		factoryBrowser.NewTagSvgRect().X(0).Y(0).Width(100).Height(100).Append(
			factoryBrowser.NewTagSvgAnimate().AttributeName("y").From(0).To(50).Dur(1*time.Second).RepeatDur(5*time.Second),
		),
		factoryBrowser.NewTagSvgRect().X(120).Y(0).Width(100).Height(100).Append(
			factoryBrowser.NewTagSvgAnimate().AttributeName("y").From(0).To(50).Dur(1*time.Second).RepeatDur(html.KSvgDurIndefinite),
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
