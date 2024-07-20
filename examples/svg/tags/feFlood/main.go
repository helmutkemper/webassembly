// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feFlood
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feFlood
//
//  <svg xmlns="http://www.w3.org/2000/svg" width="200" height="200">
//    <defs>
//      <filter id="floodFilter" filterUnits="userSpaceOnUse">
//        <feFlood x="50" y="50" width="100" height="100"
//            flood-color="green" flood-opacity="0.5"/>
//      </filter>
//    </defs>
//
//    <use style="filter: url(#floodFilter);"/>
//  </svg>

//go:build js

//fixme: bug. it is necessary to resize the screen to show the content
//When a graphic element, type, rectangle, is added to the screen, the bug does not happen.

package main

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/platform/factoryColor"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().
		ViewBox([]float64{0, 0, 200, 200}).Width(200).Height(200).Append(
		factoryBrowser.NewTagSvgDefs().Append(
			factoryBrowser.NewTagSvgFilter().Id("floodFilter").FilterUnits("userSpaceOnUse").Append(
				factoryBrowser.NewTagSvgFeFlood().X(50).Y(50).Width(100).Height(100).FloodColor(factoryColor.NewGreen()).FloodOpacity(0.5),
			),
		),

		factoryBrowser.NewTagSvgUse().Style("filter: url(#floodFilter);"),
	)

	stage.Append(s1)

	<-done
}
