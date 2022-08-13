// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Element/text
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Element/text
//
//  <svg viewBox="0 0 240 80" xmlns="http://www.w3.org/2000/svg">
//    <style>
//      .small { font: italic 13px sans-serif; }
//      .heavy { font: bold 30px sans-serif; }
//
//      /* Note that the color of the text is set with the    *
//       * fill property, the color property is for HTML only */
//      .Rrrrr { font: italic 40px serif; fill: red; }
//    </style>
//
//    <text x="20" y="35" class="small">My</text>
//    <text x="40" y="35" class="heavy">cat</text>
//    <text x="55" y="55" class="small">is</text>
//    <text x="65" y="55" class="Rrrrr">Grumpy!</text>
//  </svg>

//go:build js

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 240, 80}).Append(
		factoryBrowser.NewTagSvgStyle().Style(
			".small { font: italic 13px sans-serif; }\n"+
				".heavy { font: bold 30px sans-serif; }\n\n"+
				"/* Note that the color of the text is set with the    *\n"+
				"* fill property, the color property is for HTML only */\n"+
				".Rrrrr { font: italic 40px serif; fill: red; }\n",
		),

		factoryBrowser.NewTagSvgText().X(20).Y(35).Class("small").Text("My"),
		factoryBrowser.NewTagSvgText().X(40).Y(35).Class("heavy").Text("cat"),
		factoryBrowser.NewTagSvgText().X(55).Y(55).Class("small").Text("is"),
		factoryBrowser.NewTagSvgText().X(65).Y(55).Class("Rrrrr").Text("Grumpy!"),
	)

	stage.Append(s1)

	<-done
}
