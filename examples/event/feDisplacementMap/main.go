// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feDisplacementMap
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feDisplacementMap
//
//  <svg width="200" height="200" viewBox="0 0 220 220"
//       xmlns="http://www.w3.org/2000/svg">
//    <filter id="displacementFilter">
//      <feTurbulence type="turbulence" baseFrequency="0.05"
//          numOctaves="2" result="turbulence"/>
//      <feDisplacementMap in2="turbulence" in="SourceGraphic"
//          scale="50" xChannelSelector="R" yChannelSelector="G"/>
//    </filter>
//
//    <circle cx="100" cy="100" r="100"
//        style="filter: url(#displacementFilter)"/>
//  </svg>

//go:build js
// +build js

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/event/animation"
	"github.com/helmutkemper/iotmaker.webassembly/browser/event/mouse"
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
	"log"
	"time"
)

func main() {

	done := make(chan struct{}, 0)

	mouseEvent := make(chan mouse.Data)
	animationBeginEvent := make(chan animation.Data)
	animationRepeatEvent := make(chan animation.Data)
	animationEndEvent := make(chan animation.Data)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 880, 220}).Append(
		factoryBrowser.NewTagSvgFilter().Id("displacementFilter").Append(
			factoryBrowser.NewTagSvgFeTurbulence().Type(html.KSvgTypeTurbulenceTurbulence).BaseFrequency(0.05).NumOctaves(4).Result("turbulence").Append(
				factoryBrowser.NewTagSvgAnimate().AttributeName("baseFrequency").From(0.05).Values([]float64{0.05, 0.04, 0.05}).Dur(3*time.Second).RepeatCount(html.KSvgDurIndefinite),
			),
			factoryBrowser.NewTagSvgFeDisplacementMap().In2("turbulence").In(html.KSvgInSourceGraphic).Scale(50).XChannelSelector(html.KSvgChannelSelectorR).YChannelSelector(html.KSvgChannelSelectorG),
		),

		factoryBrowser.NewTagSvgCircle().AddListenerClick(&mouseEvent).Id("circle").Tabindex(0).Tabindex(1).Cx(100).Cy(100).R(100).Style("filter: url(#displacementFilter)").Append(
			factoryBrowser.NewTagSvgAnimate().AddListenerBegin(&animationBeginEvent).AddListenerRepeat(&animationRepeatEvent).AddListenerEnd(&animationEndEvent).AttributeName("cx").Dur(1*time.Second).From(100).To(300).RepeatCount(3),
		),
	)

	stage.Append(s1)

	go func() {
		for {
			select {
			case data := <-animationBeginEvent:
				log.Printf("begin: %v seconds", data.CurrentTime)
			case data := <-animationRepeatEvent:
				log.Printf("repeat: %v seconds", data.CurrentTime)
			case data := <-animationEndEvent:
				log.Printf("end: %v seconds", data.CurrentTime)
			case data := <-mouseEvent:
				log.Printf("click: (%v, %v)", data.ClientX, data.ClientY)
			}
		}
	}()

	<-done
}
