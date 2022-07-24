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
	"fmt"
	"github.com/helmutkemper/iotmaker.webassembly/browser/event/animation"
	"github.com/helmutkemper/iotmaker.webassembly/browser/event/mouse"
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
	"strconv"
	"time"
)

func main() {

	done := make(chan struct{}, 0)

	// Um channel para cada evento desejado, por organização do código.
	// Pode-se usar o mesmo channel para vários eventos do mesmo tipo
	mouseEvent := make(chan mouse.Data)
	animationEvent := make(chan animation.Data)

	tagText := &html.TagSvgText{}
	text := ""

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 880, 500}).Append(
		factoryBrowser.NewTagSvgFilter().Id("displacementFilter").Append(
			factoryBrowser.NewTagSvgFeTurbulence().Type(html.KSvgTypeTurbulenceTurbulence).BaseFrequency(0.07).NumOctaves(16).Result("turbulence").Append(
				factoryBrowser.NewTagSvgAnimate().AttributeName("baseFrequency").From(0.05).To(0.04).Dur(3*time.Second).RepeatCount(html.KSvgDurIndefinite),
				factoryBrowser.NewTagSvgAnimate().AttributeName("seed").From(0).To(60).Dur(2*time.Second).RepeatCount(html.KSvgDurIndefinite),
			),
			factoryBrowser.NewTagSvgFeDisplacementMap().In2("turbulence").In(html.KSvgInSourceGraphic).Scale(50).XChannelSelector(html.KSvgChannelSelectorR).YChannelSelector(html.KSvgChannelSelectorG),
		),

		factoryBrowser.NewTagSvgCircle().AddListenerClick(&mouseEvent).Id("circle").Tabindex(0).Tabindex(1).Cx(100).Cy(100).R(100).Style("filter: url(#displacementFilter)").Append(
			factoryBrowser.NewTagSvgAnimate().AddListenerBegin(&animationEvent).AddListenerRepeat(&animationEvent).AddListenerEnd(&animationEvent).AttributeName("cx").Dur(1*time.Second).From(100).To(300).RepeatCount(3),
			factoryBrowser.NewTagSvgAnimate().AttributeName("opacity").Dur(2*time.Second).From(1.0).To(0.0),
		),
		factoryBrowser.NewTagSvgText().X(5).Y(250).FontSize(24).Text("FPS: "+strconv.FormatInt(int64(stage.GetFPS()), 10)),
		factoryBrowser.NewTagSvgText().X(5).Y(275).FontSize(24).Reference(&tagText),
	)

	stage.Append(s1)

	go func() {
		var y = 275
		for {
			select {
			case data := <-animationEvent:
				switch data.EventName {
				case animation.KEventBegin:
					text += fmt.Sprintf("<tspan x='5' y='%v'>begin: %v seconds</tspan>", y, data.CurrentTime)
				case animation.KEventRepeat:
					text += fmt.Sprintf("<tspan x='5' y='%v'>repeat: %v seconds</tspan>", y, data.CurrentTime)
				case animation.KEventEnd:
					text += fmt.Sprintf("<tspan x='5' y='%v'>end: %v seconds</tspan>", y, data.CurrentTime)
				}
				tagText.Html(text)
			case data := <-mouseEvent:
				text += fmt.Sprintf("<tspan x='5' y='%v'>click: (%v, %v)</tspan>", y, data.ClientX, data.ClientY)
				tagText.Html(text)
			}

			y += 25
		}
	}()

	<-done
}
