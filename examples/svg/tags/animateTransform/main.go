// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Element/animateTransform
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Element/animateTransform
//
//  <svg width="120" height="120" viewBox="0 0 120 120"
//       xmlns="http://www.w3.org/2000/svg">
//
//      <polygon points="60,30 90,90 30,90">
//          <animateTransform attributeName="transform"
//                            attributeType="XML"
//                            type="rotate"
//                            from="0 60 70"
//                            to="360 60 70"
//                            dur="10s"
//                            repeatCount="indefinite"/>
//      </polygon>
//  </svg>

//go:build js
// +build js

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
	"time"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().Width(120).Height(120).ViewBox([]float64{0, 0, 120, 120}).Append(
		factoryBrowser.NewTagSvgPolygon().Points(html.Points{{60, 30}, {90, 90}, {30, 90}}).Append(
			factoryBrowser.NewTagSvgAnimateTransform().AttributeName("transform").Type(html.KSvgTypeTransformRotate).From([]float64{0, 60, 70}).To([]float64{360, 60, 70}).Dur(10 * time.Second).RepeatCount(html.KSvgDurIndefinite),
		),
	)

	stage.Append(s1)

	<-done
}
