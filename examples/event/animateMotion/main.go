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
	"github.com/helmutkemper/iotmaker.webassembly/browser/event/document"
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
	"github.com/helmutkemper/iotmaker.webassembly/platform/factoryColor"
	"log"
	"math"
	"syscall/js"
	"time"
)

func main() {
	js.Global().Get("window").Set("name", "test")

	var container *html.TagSvg
	var circle *html.TagSvgCircle
	var svgG *html.TagSvgG
	var animateMotion *html.TagSvgAnimateMotion

	var factor = 1.0
	var width = 400.0

	animationEvent := make(chan animation.Data)
	animationResize := make(chan document.Data)
	windowResize := make(chan document.Data)
	newWindow := make(chan document.Data)
	closeWindow := make(chan document.Data)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().Reference(&container).ViewBox([]float64{0, 0, width, 200}).Append(
		// caminho da bola vermelha
		factoryBrowser.NewTagSvg().Append(
			factoryBrowser.NewTagSvgPath().Fill(nil).Stroke(factoryColor.NewLightgrey()).D(factoryBrowser.NewPath().M(20, 50).C(20, -50, 180, 150, 180, 50).C(180, -50, 20, 150, 20, 50).Z()),
			factoryBrowser.NewTagSvgCircle().Reference(&circle).R(5).Fill(factoryColor.NewRed()).Append(
				factoryBrowser.NewTagSvgAnimateMotion().Reference(&animateMotion).AddListenerMotion(&animationEvent).Dur(3*time.Second).RepeatCount(html.KSvgDurIndefinite).Path(factoryBrowser.NewPath().M(20, 50).C(20, -50, 180, 150, 180, 50).C(180, -50, 20, 150, 20, 50).Z()),
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
	)

	stage.Append(s1)
	nw := stage.NewWindow("http://localhost:3000/documentation/")
	nw.AddListenerLoad(&newWindow)
	nw.AddListenerResize(&windowResize)

	stage.AddListenerResize(&animationResize)

	timeOut := time.NewTicker(10 * time.Second)

	go func() {
		for {
			select {
			case <-windowResize:
				log.Printf("is closed: %v", nw.GetIsClosed())

			case <-closeWindow:
				tm := time.NewTimer(time.Second)
				go func() {
					select {
					case <-tm.C:
						log.Printf("is closed: %v", nw.GetIsClosed())
					}
				}()
				log.Printf("is closed: %v", nw.GetIsClosed())

			case <-newWindow:
				nw.Scroll(0, 1000)
				nw.MoveTo(100, 100)
				nw.ResizeTo(200, 500)
				nw.AddListenerOnUnload(&closeWindow)

			case <-timeOut.C:
				animateMotion.RemoveListenerMotion()

			case data := <-animationResize:
				log.Printf("%+v", data)
				//data.Blur()

			case <-animationEvent:
				factor = (container.GetRight() - container.GetX()) / width
				angle := math.Atan2(120-circle.GetY()/factor, 95-circle.GetX()/factor)
				svgG.Transform(factoryBrowser.NewTransform().Rotate(angle*180/math.Pi-90, 25, 25))
			}
		}
	}()

	done := make(chan struct{}, 0)
	<-done
}
