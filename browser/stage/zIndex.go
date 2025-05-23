package stage

import (
	"strconv"
	"syscall/js"
)

// GetNextZIndex
//
// English:
//
//	Looking for all graphic elements in the document and then calculates the next Zindex
//
// Português:
//
//	Procura todos os elementos gráficos no documento e em seguida calcula o próximo zIndex
func GetNextZIndex() int {

	maxZIndex := 0
	elements := js.Global().Get("document").Call("getElementsByTagName", "*")
	length := elements.Length()

	for i := 0; i < length; i++ {
		element := elements.Index(i)
		style := js.Global().Get("window").Call("getComputedStyle", element)
		zIndex := style.Get("zIndex").String()
		if zIndex != "auto" {
			if parsedZIndex, err := strconv.Atoi(zIndex); err == nil {
				if parsedZIndex > maxZIndex {
					maxZIndex = parsedZIndex
				}
			}
		}
	}

	return maxZIndex + 1
}
