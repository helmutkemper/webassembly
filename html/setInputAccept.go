package html

import "log"

// SetInputAccept
//
// English:
//
//  Valid for the file input type only, the accept attribute defines which file types are selectable
//  in a file upload control. See the file input type.
//
// Português:
//
//  Válido apenas para o tipo de entrada de arquivo, o atributo accept define quais tipos de arquivo
//  são selecionáveis em um controle de upload de arquivo. Consulte o tipo de entrada do arquivo.
func (e *GlobalAttributes) SetInputAccept(accept string) (ref *GlobalAttributes) {
	switch e.tag {
	case KTagInput:
	default:
		log.Printf("tag " + e.tag.String() + " does not support accept property")
	}

	e.selfElement.Set("accept", accept)
	return e
}
