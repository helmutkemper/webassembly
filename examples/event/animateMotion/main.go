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
	"github.com/helmutkemper/iotmaker.webassembly/browser/event/animation"
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
	"github.com/helmutkemper/iotmaker.webassembly/platform/factoryColor"
	"math"
	"time"
)

func main() {
	var container *html.TagSvg
	var circle *html.TagSvgCircle
	var svgG *html.TagSvgG
	var line *html.TagSvgLine
	var cruz *html.TagSvg
	var animateMotion *html.TagSvgAnimateMotion

	var factor = 1.0
	var width = 400.0

	animationEvent := make(chan animation.Data)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().Reference(&container).ViewBox([]float64{0, 0, width, 200}).Append(
		// caminho da bola vermelha
		factoryBrowser.NewTagSvg().Append(
			factoryBrowser.NewTagSvgG().Append(
				factoryBrowser.NewTagSvgPath().Fill(nil).Stroke(factoryColor.NewLightgrey()).D(factoryBrowser.NewPath().M(20, 50).C(20, -50, 180, 150, 180, 50).C(180, -50, 20, 150, 20, 50).Z()),
				factoryBrowser.NewTagSvgCircle().Reference(&circle).Id("trinidad").R(5).Fill(factoryColor.NewRed()).Append(
					factoryBrowser.NewTagSvgAnimateMotion().Reference(&animateMotion).AddListenerMotion(&animationEvent).Dur(10*time.Second).RepeatCount(html.KSvgDurIndefinite).Path(factoryBrowser.NewPath().M(20, 50).C(20, -50, 180, 150, 180, 50).C(180, -50, 20, 150, 20, 50).Z()),
				),
			),
		),

		// Seta dinâmica
		factoryBrowser.NewTagSvg().X(100-25).Y(100).Append(
			factoryBrowser.NewTagSvgG().Reference(&svgG).Append(
				factoryBrowser.NewTagSvgCircle().Cx(25).Cy(25).R(10).Fill(nil).StrokeWidth(1).Stroke(factoryColor.NewGray()),
				factoryBrowser.NewTagSvgLine().X1(25).X2(25).Y1(0).Y2(50).StrokeWidth(1).Stroke(factoryColor.NewGray()),
				factoryBrowser.NewTagSvgLine().X1(25).Y1(0).X2(20).Y2(5).StrokeWidth(1).Stroke(factoryColor.NewGray()),
				factoryBrowser.NewTagSvgLine().X1(25).Y1(0).X2(30).Y2(5).StrokeWidth(1).Stroke(factoryColor.NewGray()),
			),
		),

		// cruz vermelha no centro da seta
		factoryBrowser.NewTagSvg().Reference(&cruz).X(95).Y(120).Stroke(factoryColor.NewRed()).StrokeWidth(0.5).Fill(nil).Append(
			factoryBrowser.NewTagSvgLine().X1(5).Y1(0).X2(5).Y2(10).StrokeWidth(1).Stroke(factoryColor.NewRed()),
			factoryBrowser.NewTagSvgLine().X1(0).Y1(5).X2(10).Y2(5).StrokeWidth(1).Stroke(factoryColor.NewRed()),
		),

		// linha azul ligando o centro da seta e a bola vermelha
		factoryBrowser.NewTagSvgLine().Reference(&line).Stroke(factoryColor.NewBlue()).StrokeWidth(0.1).Fill(nil),
	)

	stage.Append(s1)

	timeOut := time.NewTimer(10 * time.Second)

	go func() {
		for {
			select {
			case <-timeOut.C:
				svgG.Opacity(0.0)
				cruz.Opacity(0.0)
				line.Opacity(0.0)
				animateMotion.RemoveListenerMotion()

			case <-animationEvent:
				factor = (container.GetRight() - container.GetX()) / width
				angle := math.Atan2(120-circle.GetY()/factor, 95-circle.GetX()/factor)
				svgG.Transform(factoryBrowser.NewTransform().Rotate(angle*180/math.Pi-90, 25, 25))
				line.X1(100).Y1(125).X2(circle.GetX()/factor + 5).Y2(circle.GetY()/factor + 5)
			}
		}
	}()

	done := make(chan struct{}, 0)
	<-done
}
