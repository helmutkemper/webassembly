package textUtil

import (
	"fmt"
	"syscall/js"
)

func GetTextSize(text, fontFamily string, bold, italic bool, fontSize int) (width, height int) {
	doc := js.Global().Get("document")
	svgNS := "http://www.w3.org/2000/svg"

	svg := doc.Call("createElementNS", svgNS, "svg")
	svg.Call("setAttribute", "width", "0")
	svg.Call("setAttribute", "height", "0")
	svg.Call("setAttribute", "style", "position:absolute; left:-9999px; top:-9999px;")

	svgText := doc.Call("createElementNS", svgNS, "text")
	svgText.Set("textContent", text)
	svgText.Call("setAttribute", "x", "0")
	svgText.Call("setAttribute", "y", "0")
	svgText.Call("setAttribute", "font-size", fmt.Sprintf("%dpx", fontSize))
	svgText.Call("setAttribute", "font-family", fontFamily)

	if bold {
		svgText.Call("setAttribute", "font-weight", "bold")
	}
	if italic {
		svgText.Call("setAttribute", "font-style", "italic")
	}

	svg.Call("appendChild", svgText)
	doc.Get("body").Call("appendChild", svg)

	bbox := svgText.Call("getBBox")
	width = bbox.Get("width").Int()
	height = bbox.Get("height").Int()

	doc.Get("body").Call("removeChild", svg)
	return
}
