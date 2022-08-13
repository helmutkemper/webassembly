// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feMorphology
//
//   Notes:
//     * The CSS is inside the example HTML file.
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feMorphology
//
//   Notas:
//     * O CSS está dentro do arquivo HTML de exemplo.
//
// CSS:
//
//  text {
//    font-family: Arial, Helvetica, sans-serif;
//    font-size: 3em;
//  }
//
//  #thin {
//    filter: url(#erode);
//  }
//
//  #thick {
//    filter: url(#dilate);
//  }
//
// HTML:
//
//  <svg xmlns="http://www.w3.org/2000/svg" width="300" height="180">
//    <filter id="erode">
//      <feMorphology operator="erode" radius="1"/>
//    </filter>
//    <filter id="dilate">
//      <feMorphology operator="dilate" radius="2"/>
//    </filter>
//    <text y="1em">Normal text</text>
//    <text id="thin" y="2em">Thinned text</text>
//    <text id="thick" y="3em">Fattened text</text>
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

	s1 := factoryBrowser.NewTagSvg().Width(300).Height(180).XmlnsXLink("http://www.w3.org/1999/xlink").Append(

		factoryBrowser.NewTagSvgFilter().Id("erode").Append(
			factoryBrowser.NewTagSvgFeMorphology().Operator("erode").Radius(1),
		),

		factoryBrowser.NewTagSvgFilter().Id("dilate").Append(
			factoryBrowser.NewTagSvgFeMorphology().Operator(html.KKSvgOperatorFeCompositeDilate).Radius(2),
		),

		factoryBrowser.NewTagSvgText().Y("1em").Text("Normal text"),
		factoryBrowser.NewTagSvgText().Id("thin").Y("2em").Text("Thinned text"),
		factoryBrowser.NewTagSvgText().Id("thick").Y("3em").Text("Thinned text"),
	)

	stage.Append(s1)

	<-done
}
