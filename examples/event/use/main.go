// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Element/use
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Element/use
//
//  <svg viewBox="0 0 30 10" xmlns="http://www.w3.org/2000/svg">
//    <circle id="myCircle" cx="5" cy="5" r="4" stroke="blue"/>
//    <use href="#myCircle" x="10" fill="blue"/>
//    <use href="#myCircle" x="20" fill="white" stroke="red"/>
//    <!--
//      stroke="red" will be ignored here, as stroke was already set on myCircle.
//      Most attributes (except for x, y, width, height and (xlink:)href)
//      do not override those set in the ancestor.
//      That's why the circles have different x positions, but the same stroke value.
//    -->
//  </svg>

//go:build js

package main

import (
	"github.com/helmutkemper/webassembly/browser/event/mouse"
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/platform/factoryColor"
)

func main() {

	done := make(chan struct{}, 0)

	tagDiv := &html.TagDiv{}
	tagUse := &html.TagSvgUse{}
	mouseEvent := make(chan mouse.Data)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
		factoryBrowser.NewTagSvgCircle().AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
		factoryBrowser.NewTagSvgUse().Reference(&tagUse).HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
		factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
	)
	div1 := factoryBrowser.NewTagDiv().Reference(&tagDiv)

	go func() {
		text := ""
		for {
			select {
			case <-mouseEvent:
				text += "click<br>"
				tagDiv.Html(text)
				// English: addEventListener('click') was created on the <circle> element, so the reference is invalid and the command does not work.
				// Português: addEventListener('click') foi criado no elemento <circle>, por isto, a refereência é inválida e o comando não funciona.
				tagUse.RemoveListenerClick()
			}
		}
	}()

	stage.Append(s1, div1)

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
//
//
//
//
//
//
//
//
