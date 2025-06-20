package textUtil

import (
	"syscall/js"
)

// GetTextSize
//
// English:
//
//		Returns space in pixels, occupied by a text.
//
//		  Note:
//		    * To measure fontawesome.com icon size, use the constant textUtil.KFontAwesomeRegular or
//		      textUtil.KFontAwesomeSolid to use font-family="FARegular" or font-family="FASolid"
//	     * Some icons have a contour box incompatible with the cartoon contained in the icon and need manual correction.
//
// Português:
//
//		Retorna o espaço, em pixels, ocupado por um texto
//
//		  Nota:
//		    * Para medir o tamanho de ícones da fontawesome.com, use as constantes textUtil.KFontAwesomeRegular ou
//		      textUtil.KFontAwesomeSolid para usar font-family="FARegular" ou font-family="FASolid";
//	     * Alguns ícones apresentam uma caixa de contorno incompatível com o desenho contido no ícone e necessitam de
//	       correção manual.
func GetTextSize(text, fontFamily, fontWeight, fontStyle string, fontSize int) (width, height int) {
	doc := js.Global().Get("document")
	svgNS := "http://www.w3.org/2000/svg"

	svg := doc.Call("createElementNS", svgNS, "svg")
	svg.Call("setAttribute", "width", "0")
	svg.Call("setAttribute", "height", "0")
	svg.Call("setAttribute", "style", "position:absolute; left:-9999px; top:-9999px;")

	svgText := doc.Call("createElementNS", svgNS, "text")
	svgText.Call("setAttribute", "xmlns", "http://www.w3.org/2000/svg")
	svgText.Call("setAttribute", "x", "0")
	svgText.Call("setAttribute", "y", "0")
	svgText.Call("setAttribute", "font-size", fontSize)
	svgText.Call("setAttribute", "font-family", fontFamily)
	svgText.Call("setAttribute", "font-weight", fontWeight)
	svgText.Call("setAttribute", "font-style", fontStyle)
	svgText.Set("textContent", text)

	svg.Call("appendChild", svgText)
	doc.Get("body").Call("appendChild", svg)

	bbox := svgText.Call("getBBox")
	width = bbox.Get("width").Int()
	height = bbox.Get("height").Int()

	doc.Get("body").Call("removeChild", svg)
	return
}
