// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Element/foreignObject
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Element/foreignObject
//
//  <svg viewBox="0 0 200 200" xmlns="http://www.w3.org/2000/svg">
//    <style>
//      div {
//        color: white;
//        font: 18px serif;
//        height: 100%;
//        overflow: auto;
//      }
//    </style>
//
//    <polygon points="5,5 195,10 185,185 10,195" />
//
//    <!-- Common use case: embed HTML text into SVG -->
//    <foreignObject x="20" y="20" width="160" height="160">
//      <!--
//        In the context of SVG embedded in an HTML document, the XHTML
//        namespace could be omitted, but it is mandatory in the
//        context of an SVG document
//      -->
//      <div xmlns="http://www.w3.org/1999/xhtml">
//        Lorem ipsum dolor sit amet, consectetur adipiscing elit.
//        Sed mollis mollis mi ut ultricies. Nullam magna ipsum,
//        porta vel dui convallis, rutrum imperdiet eros. Aliquam
//        erat volutpat.
//      </div>
//    </foreignObject>
//  </svg>

//go:build js

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 200, 200}).XmlnsXLink("http://www.w3.org/1999/xlink").Append(
		factoryBrowser.NewTagSvgStyle().Style(
			"div {"+
				"color: white;"+
				"font: 18px serif;"+
				"height: 100%;"+
				"overflow: auto;"+
				"}"),

		factoryBrowser.NewTagSvgPolygon().Points([][]float64{{5, 5}, {195, 10}, {185, 185}, {10, 195}}),

		// Common use case: embed HTML text into SVG
		factoryBrowser.NewTagSvgForeignObject().X(20).Y(20).Width(160).Height(160).Html(
			" <!--"+
				"In the context of SVG embedded in an HTML document, the XHTML"+
				"namespace could be omitted, but it is mandatory in the"+
				"context of an SVG document"+
				"-->"+
				"<div xmlns=\"http://www.w3.org/1999/xhtml\">"+
				"  Lorem ipsum dolor sit amet, consectetur adipiscing elit."+
				"  Sed mollis mollis mi ut ultricies. Nullam magna ipsum,"+
				"  porta vel dui convallis, rutrum imperdiet eros. Aliquam"+
				"  erat volutpat."+
				"</div>",
		),
	)

	stage.Append(s1)

	<-done
}
