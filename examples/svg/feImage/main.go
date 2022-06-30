// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feImage
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feImage
//
//  <svg viewBox="0 0 200 200" xmlns="http://www.w3.org/2000/svg"
//      xmlns:xlink="http://www.w3.org/1999/xlink">
//    <defs>
//      <filter id="image">
//        <feImage xlink:href="/files/6457/mdn_logo_only_color.png"/>
//      </filter>
//    </defs>
//
//    <rect x="10%" y="10%" width="80%" height="80%"
//        style="filter:url(#image);"/>
//  </svg>

//go:build js
// +build js

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().
		ViewBox([]float64{0, 0, 200, 200}).
		Append(

			factoryBrowser.NewTagSvgDefs().
				Append(

					factoryBrowser.NewTagSvgFilter().
						Id("image").
						Append(

							factoryBrowser.NewTagSvgFeImage().
								HRef("//developer.mozilla.org/files/6457/mdn_logo_only_color.png"),
						),
				),

			factoryBrowser.NewTagSvgRect().
				X(float32(0.1)).
				Y(float32(0.1)).
				Width(float32(0.8)).
				Height(float32(0.8)).
				Style("filter:url(#image);"),
		)

	stage.Append(s1)

	<-done
}
