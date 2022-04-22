package html

import "log"

// SetHeight
//
// English:
//
//  The height is the height of the image file to display to represent the graphical submit button.
//
// Português:
//
//  A altura é a altura do arquivo de imagem a ser exibido para representar o botão de envio gráfico.
func (e *GlobalAttributes) SetHeight(height int) (ref *GlobalAttributes) {
	switch e.tag {
	case KTagInput:
	default:
		log.Printf("tag " + e.tag.String() + " does not support height property")
	}

	e.selfElement.Set("height", height)
	return e
}
