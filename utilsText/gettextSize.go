package utilsText

import (
	"strings"
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

// WrapText quebra `text` em múltiplas linhas, cada uma com largura ≤ maxWidth.
// Usa utilsText.GetTextSize(line) para medir a largura em pixels.
func WrapText(text string, maxWidth float64, fontFamily, fontWeight, fontStyle string, fontSize int) []string {
	var lines []string
	words := strings.Fields(text)
	var currentLine string

	for _, w := range words {
		// Se a palavra isolada já ultrapassa maxWidth, quebramos ela em pedaços:
		textWidth, _ := GetTextSize(w, fontFamily, fontWeight, fontStyle, fontSize)
		if float64(textWidth) > maxWidth {
			// Primeiro, finaliza a linha atual (se existir)
			if currentLine != "" {
				lines = append(lines, currentLine)
				currentLine = ""
			}
			// Quebra a palavra longa
			for _, part := range breakLongWord(w, maxWidth, fontFamily, fontWeight, fontStyle, fontSize) {
				lines = append(lines, part)
			}
			continue
		}

		// Tenta encaixar no final da linha atual
		if currentLine == "" {
			currentLine = w
		} else {
			candidate := currentLine + " " + w
			textWidth, _ = GetTextSize(candidate, fontFamily, fontWeight, fontStyle, fontSize)
			if float64(textWidth) <= maxWidth {
				currentLine = candidate
			} else {
				// Não cabe: fecha a linha atual e inicia uma nova com a palavra
				lines = append(lines, currentLine)
				currentLine = w
			}
		}
	}

	// Adiciona a última linha remanescente
	if currentLine != "" {
		lines = append(lines, currentLine)
	}
	return lines
}

// breakLongWord fragmenta uma palavra cuja largura excede maxWidth.
// Retorna um slice de pedaços que cabem em ≤ maxWidth pixels.
func breakLongWord(word string, maxWidth float64, fontFamily, fontWeight, fontStyle string, fontSize int) []string {
	var parts []string
	var buf []rune

	for _, r := range word {
		buf = append(buf, r)
		textWidth, _ := GetTextSize(string(buf), fontFamily, fontWeight, fontStyle, fontSize)
		if float64(textWidth) > maxWidth {
			// Se estourou, retira o último rune e fecha o pedaço
			if len(buf) > 1 {
				last := buf[len(buf)-1]
				part := string(buf[:len(buf)-1])
				parts = append(parts, part)
				// reinicia buf com o rune que sobrou
				buf = []rune{last}
			} else {
				// rune único muito largo (caso extremo): empurra mesmo assim
				parts = append(parts, string(buf))
				buf = []rune{}
			}
		}
	}
	// resto do buffer vira último pedaço
	if len(buf) > 0 {
		parts = append(parts, string(buf))
	}
	return parts
}
