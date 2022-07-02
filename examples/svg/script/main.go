// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Element/script
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Element/script
//
//  <svg viewBox="0 0 10 10" xmlns="http://www.w3.org/2000/svg">
//    <script>
//    // <![CDATA[
//    window.addEventListener('DOMContentLoaded', () => {
//      function getColor () {
//        const R = Math.round(Math.random() * 255).toString(16).padStart(2,'0')
//        const G = Math.round(Math.random() * 255).toString(16).padStart(2,'0')
//        const B = Math.round(Math.random() * 255).toString(16).padStart(2,'0')
//        return `#${R}${G}${B}`
//      }
//
//      document.querySelector('circle').addEventListener('click', (e) => {
//        e.target.style.fill = getColor()
//      })
//    })
//    // ]]>
//    </script>
//
//    <circle cx="5" cy="5" r="4" />
//  </svg>

//go:build js
// +build js

// fixme: bug
// document.addEventListener("DOMContentLoaded", (e) => {...});
// 'DOMContentLoaded' não funciona

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().
		ViewBox([]float64{0, 0, 10, 10}).
		Append(

			factoryBrowser.NewTagSvgScript().Script(
				"  // <![CDATA[\n"+
					"function getColor () {\n"+
					"const R = Math.round(Math.random() * 255).toString(16).padStart(2,'0')\n"+
					"const G = Math.round(Math.random() * 255).toString(16).padStart(2,'0')\n"+
					"const B = Math.round(Math.random() * 255).toString(16).padStart(2,'0')\n"+
					"return `#${R}${G}${B}`\n"+
					"}\n"+
					"document.querySelector('circle').addEventListener('click', (e) => {\n"+
					"e.target.style.fill = getColor()\n"+
					"})\n"+
					"// ]]>",
			),

			factoryBrowser.NewTagSvgCircle().Cx(5).Cy(5).R(4),
		)

	stage.Append(s1)

	<-done
}
