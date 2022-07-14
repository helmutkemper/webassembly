// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Element/animateMotion
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Element/animateMotion
//
//  <svg viewBox="0 0 200 100" xmlns="http://www.w3.org/2000/svg">
//    <path fill="none" stroke="lightgrey"
//      d="M20,50 C20,-50 180,150 180,50 C180-50 20,150 20,50 z" />
//
//    <circle r="5" fill="red">
//      <animateMotion dur="10s" repeatCount="indefinite"
//        path="M20,50 C20,-50 180,150 180,50 C180-50 20,150 20,50 z" />
//    </circle>
//  </svg>

//go:build js
// +build js

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
	"github.com/helmutkemper/iotmaker.webassembly/platform/factoryColor"
	"math"
	"time"
)

var circle *html.TagSvgCircle
var svgG *html.TagSvgG

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	circle = factoryBrowser.NewTagSvgCircle()
	svgG = factoryBrowser.NewTagSvgG()

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 200, 300}).Width(300).Append(

		factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 50, 50}).X(100-25).Y(200-50).Width(50).Height(50).Append(
			svgG.Append(
				factoryBrowser.NewTagSvgCircle().Cx(25).Cy(25).R(10).Fill(nil).StrokeWidth(1).Stroke(factoryColor.NewGray()),
				factoryBrowser.NewTagSvgLine().X1(25).X2(25).Y1(0).Y2(50).StrokeWidth(1).Stroke(factoryColor.NewGray()),
				factoryBrowser.NewTagSvgLine().X1(25).Y1(0).X2(20).Y2(5).StrokeWidth(1).Stroke(factoryColor.NewGray()),
				factoryBrowser.NewTagSvgLine().X1(25).Y1(0).X2(30).Y2(5).StrokeWidth(1).Stroke(factoryColor.NewGray()),
			),
		),

		factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 200, 100}).Height(100).Append(
			factoryBrowser.NewTagSvgPath().Fill("none").Stroke(factoryColor.NewLightgrey()).D(factoryBrowser.NewPath().M(20, 50).C(20, -50, 180, 150, 180, 50).C(180, -50, 20, 150, 20, 50).Z()),
			circle.Id("circle").R(5).Fill(factoryColor.NewRed()).Append(
				factoryBrowser.NewTagSvgAnimateMotion().Id("test").Dur(10*time.Second).RepeatCount(html.KSvgDurIndefinite).Path(factoryBrowser.NewPath().M(20, 50).C(20, -50, 180, 150, 180, 50).C(180, -50, 20, 150, 20, 50).Z()).
					AddListener(),
			),
		),
	)

	stage.Append(s1)

	stage.AddDrawFunctions(func() {
		angle := math.Atan2(svgG.GetY()-circle.GetY(), svgG.GetX()-circle.GetX())
		svgG.Transform(factoryBrowser.NewTransform().Rotate(angle*180/math.Pi-90, 25, 25))
	})

	<-done
}
