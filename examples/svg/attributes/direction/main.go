// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/direction
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/direction
//
//  <svg viewBox="0 0 600 72" xmlns="http://www.w3.org/2000/svg"
//      direction="rtl" lang="fa">
//    <text x="300" y="50" text-anchor="middle"
//        font-size="36">داستان SVG 1.1 SE طولا ني است.</text>
//  </svg>

//go:build js

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 600, 72}).Direction(html.KSvgDirectionRtl).Lang(html.KLanguagePersianFarsi).Append(
		factoryBrowser.NewTagSvgText().X(300).Y(50).TextAnchor(html.KSvgTextAnchorMiddle).FontSize(36).Text(" SVG 1.1 SE طولا ني است."),
	)

	stage.Append(s1)

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
