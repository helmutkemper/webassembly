package textUtil

import (
	"fmt"
	"log"
	"syscall/js"
)

func GetTextSize(text, fontFamily string, bold, italic bool, fontSize int) (width, height int) {
	doc := js.Global().Get("document")
	svgNS := "http://www.w3.org/2000/svg"

	// Cria um elemento SVG invisível temporário
	svg := doc.Call("createElementNS", svgNS, "svg")
	svg.Call("setAttribute", "width", "0")
	svg.Call("setAttribute", "height", "0")
	svg.Call("setAttribute", "style", "position:absolute; left:-9999px; top:-9999px;")

	// Cria o elemento <text>
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

	// Adiciona o texto ao SVG e o SVG ao documento
	svg.Call("appendChild", svgText)
	doc.Get("body").Call("appendChild", svg)

	// Mede o texto
	bbox := svgText.Call("getBBox")
	width = bbox.Get("width").Int()
	height = bbox.Get("height").Int()

	log.Printf("width: %v, height: %v", width, height)

	// Remove o SVG temporário do DOM
	doc.Get("body").Call("removeChild", svg)

	return
}
