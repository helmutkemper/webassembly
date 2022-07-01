// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Element/image
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Element/image
//
//  <svg width="200" height="200"
//    xmlns="http://www.w3.org/2000/svg">
//    <image href="mdn_logo_only_color.png" height="200" width="200"/>
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
		Width(200).
		Height(200).
		Append(

			factoryBrowser.NewTagSvgImage().
				HRef("//developer.mozilla.org/files/6457/mdn_logo_only_color.png").
				Height(200).
				Width(200),
		)

	stage.Append(s1)

	<-done
}
