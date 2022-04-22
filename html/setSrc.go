package html

import "log"

// SetSrc
//
// English:
//
//  Valid for the image input button only, the src is string specifying the URL of the image file to
//  display to represent the graphical submit button.
//
// Português:
//
//  Válido apenas para o botão de entrada de imagem, o src é uma string que especifica a URL do
//  arquivo de imagem a ser exibido para representar o botão de envio gráfico.
func (e *GlobalAttributes) SetSrc(src string) (ref *GlobalAttributes) {
	switch e.tag {
	case KTagInput:
	default:
		log.Printf("tag " + e.tag.String() + " does not support src property")
	}

	e.selfElement.Set("src", src)
	return e
}
