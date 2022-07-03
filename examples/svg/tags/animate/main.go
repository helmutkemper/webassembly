// English:
//
// This example was taken from https://developer.mozilla.org/pt-BR/docs/Web/SVG/Element/animate
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/pt-BR/docs/Web/SVG/Element/animate
//
//  <svg viewBox="0 0 10 10" xmlns="http://www.w3.org/2000/svg">
//    <rect width="10" height="10">
//      <animate attributeName="rx" values="0;5;0" dur="10s" repeatCount="indefinite" />
//    </rect>
//  </svg>

//go:build js
// +build js

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
	"github.com/helmutkemper/iotmaker.webassembly/platform/factoryColor"
	"image/color"
	"time"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().
		ViewBox([]float64{0, 0, 10, 10}).
		Append(

			factoryBrowser.NewTagSvgRect().
				Width(10).
				Height(10).Append(

				factoryBrowser.NewTagSvgAnimate().
					AttributeName("rx").
					Values([]float64{0, 5, 0}).
					Dur(10*time.Second).
					RepeatCount(html.KSvgDurIndefinite),

				factoryBrowser.NewTagSvgAnimate().
					AttributeName("fill").
					Values(
						[]color.RGBA{
							factoryColor.NewBlack(),
							factoryColor.NewRed(),
							factoryColor.NewBlack(),
						},
					).
					Dur(10*time.Second).
					RepeatCount(html.KSvgDurIndefinite),
			),
		)

	stage.Append(s1)

	<-done
}
