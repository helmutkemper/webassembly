// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/text-decoration
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/text-decoration
//
//  <svg viewBox="0 0 250 50" xmlns="http://www.w3.org/2000/svg">
//    <text y="20" text-decoration="underline">Underlined text</text>
//    <text x="0" y="40" text-decoration="line-through">Struck-through text</text>
//  </svg>

//go:build js
// +build js

// bug: o original não apararece nada

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()
	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 250, 50}).Append(
		factoryBrowser.NewTagSvgText().Y(20).TextDecoration(html.KSvgTextDecorationLineUnderline).Text("Underlined text"),
		factoryBrowser.NewTagSvgText().X(0).Y(40).TextDecoration(html.KSvgTextDecorationLineLineThrough).Text("Struck-through text"),
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
