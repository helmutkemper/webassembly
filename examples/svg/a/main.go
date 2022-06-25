//go:build js
// +build js

//
package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
)

func main() {

	done := make(chan struct{}, 0)

	// browser stage
	stage := factoryBrowser.NewStage()

	// https://developer.mozilla.org/pt-BR/docs/Web/SVG/Element/a
	//
	//  /* As SVG does not provide a default visual style for links,
	//     it's considered best practice to add some */
	//
	//  @namespace svg url(http://www.w3.org/2000/svg);
	//  /* Necessary to select only SVG <a> elements, and not also HTML’s.
	//     See warning below */
	//
	//  svg|a:link, svg|a:visited {
	//    cursor: pointer;
	//  }
	//
	//  svg|a text,
	//  text svg|a {
	//    fill: blue; /* Even for text, SVG uses fill over color */
	//    text-decoration: underline;
	//  }
	//
	//  svg|a:hover, svg|a:active {
	//    outline: dotted 1px blue;
	//  }
	//
	//  <svg viewBox="0 0 100 100" xmlns="http://www.w3.org/2000/svg">
	//    <!-- A link around a shape -->
	//    <a href="/docs/Web/SVG/Element/circle">
	//      <circle cx="50" cy="40" r="35"/>
	//    </a>
	//
	//    <!-- A link around a text -->
	//    <a href="/docs/Web/SVG/Element/text">
	//      <text x="50" y="90" text-anchor="middle">
	//        &lt;circle&gt;
	//      </text>
	//    </a>
	//  </svg>
	//
	//  Notes:
	//    * The CSS is inside the example HTML file.

	s1 := factoryBrowser.NewTagSvg("svg1").
		ViewBox([]float64{0, 0, 100, 100}).
		Append(

			factoryBrowser.NewTagSvgA("a1").HRef("/docs/Web/SVG/Element/circle").
				Append(

					factoryBrowser.NewTagSvgCircle("cir1").
						Cx(50).
						Cy(40).
						R(35),
				),

			factoryBrowser.NewTagSvgA("a2").HRef("/docs/Web/SVG/Element/text").
				Append(

					factoryBrowser.NewTagSvgText("txt1").
						X(50).
						Y(90).
						TextAnchor(html.KSvgTextAnchorMiddle).
						Text("<circle>"),
				),
		)

	stage.Append(s1)

	<-done
}
