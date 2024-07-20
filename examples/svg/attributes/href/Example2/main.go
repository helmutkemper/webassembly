// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/href
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/href
//
//  <svg viewBox="0 0 100 100" xmlns="http://www.w3.org/2000/svg">
//    <image href="https://yari-demos.prod.mdn.mozit.cloud/en-US/docs/Web/SVG/Attribute/href/fxlogo.png" x="0" y="0" height="100" width="100"/>
//  </svg>

//go:build js

package main

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().
		ViewBox([]float64{0, 0, 100, 100}).
		Append(
			factoryBrowser.NewTagSvgImage().HRef("https://yari-demos.prod.mdn.mozit.cloud/en-US/docs/Web/SVG/Attribute/href/fxlogo.png").X(0).Y(0).Height(100).Width(100),
		)

	stage.Append(s1)

	<-done
}
