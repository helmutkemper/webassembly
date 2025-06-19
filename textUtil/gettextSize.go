package textUtil

import (
	"fmt"
	"syscall/js"
)

// GetTextSize
//
// English:
//
//	Returns space in pixels, occupied by a text.
//
//	  Note:
//	    * To measure fontawesome.com icon size, use the constant textUtil.KFontAwesomeRegular or
//	      textUtil.KFontAwesomeSolid to use font-family="FARegular" or font-family="FASolid"
//
// Português:
//
//	Retorna o espaço, em pixels, ocupado por um texto
//
//	  Nota:
//	    * para medir o tamanho de ícones da fontawesome.com, use as constantes textUtil.KFontAwesomeRegular ou
//	      textUtil.KFontAwesomeSolid para usar font-family="FARegular" ou font-family="FASolid"
func GetTextSize(text, fontFamily, fontWeight, fontStyle, fontSize string) (width, height int) {
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
	svgText.Call("setAttribute", "font-weight", fontWeight)
	svgText.Call("setAttribute", "font-style", fontStyle)

	svg.Call("appendChild", svgText)
	doc.Get("body").Call("appendChild", svg)

	bbox := svgText.Call("getBBox")
	width = bbox.Get("width").Int()
	height = bbox.Get("height").Int()

	doc.Get("body").Call("removeChild", svg)
	return
}
