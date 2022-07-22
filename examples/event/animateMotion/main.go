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
	"time"
)

func main() {
	var container *html.TagSvg
	//var circle *html.TagSvgCircle
	var test *html.TagSvgCircle
	var svgG *html.TagSvgG
	//var animateMotion *html.TagSvgAnimateMotion

	//var factor = 1.0
	var width = 400.0

	animationEvent := make(chan animation.Data)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().Reference(&container).ViewBox([]float64{0, 0, width, 200}).Append(
		// caminho da bola vermelha
		//factoryBrowser.NewTagSvg().Append(
		//	factoryBrowser.NewTagSvgPath().Fill(nil).Stroke(factoryColor.NewLightgrey()).D(factoryBrowser.NewPath().M(20, 50).C(20, -50, 180, 150, 180, 50).C(180, -50, 20, 150, 20, 50).Z()),
		//	factoryBrowser.NewTagSvgCircle().Reference(&circle).Id("trinidad").R(5).Fill(factoryColor.NewRed()).Append(
		//		factoryBrowser.NewTagSvgAnimateMotion().Reference(&animateMotion).AddListenerMotion(&animationEvent).Dur(3*time.Second).RepeatCount(0).Path(factoryBrowser.NewPath().M(20, 50).C(20, -50, 180, 150, 180, 50).C(180, -50, 20, 150, 20, 50).Z()),
		//	),
		//),

		// Seta dinâmica
		factoryBrowser.NewTagSvg().X(100-25).Y(100).Append(
			factoryBrowser.NewTagSvgG().Reference(&svgG).Append(
				factoryBrowser.NewTagSvgCircle().Cx(25).Cy(25).R(10).Fill(nil).StrokeWidth(1).Stroke(factoryColor.NewGray()),
				factoryBrowser.NewTagSvgLine().X1(25).X2(25).Y1(0).Y2(50).StrokeWidth(1).Stroke(factoryColor.NewGray()),
				factoryBrowser.NewTagSvgLine().X1(25).Y1(0).X2(20).Y2(5).StrokeWidth(1).Stroke(factoryColor.NewGray()),
				factoryBrowser.NewTagSvgLine().X1(25).Y1(0).X2(30).Y2(5).StrokeWidth(1).Stroke(factoryColor.NewGray()),
			),
		),

		factoryBrowser.NewTagSvgCircle().Reference(&test).R(5).Fill(factoryColor.NewRed()).Append(
			factoryBrowser.NewTagSvgAnimate().Dur(5*time.Second).AttributeName("cx").To(2*w),
			factoryBrowser.NewTagSvgAnimate().Dur(5*time.Second).AttributeName("cy").To(0*h),
		),
	)

	stage.Append(s1)

	timeOut := time.NewTicker(5 * time.Second)

	go func() {
		var i = 0
		for {
			select {
			case <-timeOut.C:
				//circle.Cx(20).Cy(50)
				//animateMotion.RemoveListenerMotion()
				switch i {
				case 0:
					//test.ClearContent()
					test.Append(
						st[0]...,
					)
					i += 1
				case 1:
					//test.ClearContent()
					test.Append(
						st[1]...,
					)
					i += 1
				case 2:
					//test.ClearContent()
					test.Append(
						st[2]...,
					)
					i += 1
				case 3:
					//test.ClearContent()
					test.Append(
						st[3]...,
					)
					i = 0
				}

			case <-animationEvent:
				//factor = (container.GetRight() - container.GetX()) / width
				//angle := math.Atan2(120-circle.GetY()/factor, 95-circle.GetX()/factor)
				//svgG.Transform(factoryBrowser.NewTransform().Rotate(angle*180/math.Pi-90, 25, 25))
			}
		}
	}()

	done := make(chan struct{}, 0)
	<-done
}

//          (x,y)           (x,y)
//          (0,0)           (2,0)
//            +-------+-------+
//            |               |
//            |               |
//            +               +
//            |               |
//            |               |
//            +-------+-------+
//          (0,2)           (2,2)
//          (x,y)           (x,y)
type stateMachine [][]html.Compatible

var w = 50
var h = 50
var st = stateMachine{
	{
		//factoryBrowser.NewTagSvgAnimate().AttributeName("cx").Values(w * 2),
		//factoryBrowser.NewTagSvgAnimate().AttributeName("cy").Values(h * 0),

		//factoryBrowser.NewTagSvgAnimate().Dur(5 * time.Second).AttributeName("cx").From(w * 2).To(w * 2),
		factoryBrowser.NewTagSvgAnimate().AttributeName("cy").From(h * 0).To(h * 2),
	},
	{
		//factoryBrowser.NewTagSvgAnimate().AttributeName("cx").Values(w * 2),
		//factoryBrowser.NewTagSvgAnimate().AttributeName("cy").Values(h * 2),

		factoryBrowser.NewTagSvgAnimate().AttributeName("cx").To(w * 0),
		factoryBrowser.NewTagSvgAnimate().AttributeName("cy").To(h * 2),
	},
	{
		factoryBrowser.NewTagSvgAnimate().Dur(5 * time.Second).AttributeName("cx").Values(w * 0),
		//factoryBrowser.NewTagSvgAnimate().AttributeName("cy").Values(h * 2),

		factoryBrowser.NewTagSvgAnimate().AttributeName("cx").To(w * 0),
		//factoryBrowser.NewTagSvgAnimate().Dur(5 * time.Second).AttributeName("cy").From(h * 2).To(h * 0),
	},
	{
		factoryBrowser.NewTagSvgAnimate().Dur(5 * time.Second).AttributeName("cx").Values(w * 0),
		//factoryBrowser.NewTagSvgAnimate().AttributeName("cy").To(h * 0),

		factoryBrowser.NewTagSvgAnimate().AttributeName("cx").To(w * 2),
		//factoryBrowser.NewTagSvgAnimate().Dur(5 * time.Second).AttributeName("cy").From(h * 0).To(h * 0),
	},
}
