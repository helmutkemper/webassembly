package _global

import "log"

// SetMultiple
//
// English:
//
//  This Boolean attribute indicates that multiple options can be selected in the list. If it is not
//  specified, then only one option can be selected at a time. When multiple is specified, most
//  browsers will show a scrolling list box instead of a single line dropdown.
//
// Português:
//
//  Este atributo booleano indica que várias opções podem ser selecionadas na lista. Se não for
//  especificado, apenas uma opção pode ser selecionada por vez. Quando vários são especificados, a
//  maioria dos navegadores mostrará uma caixa de listagem de rolagem em vez de uma lista suspensa
//  de uma única linha.
func (e *GlobalAttributes) SetMultiple(multiple bool) (ref *GlobalAttributes) {
	switch e.tag {
	case KTagSelect:
	case KTagInput:
	default:
		log.Printf("tag " + e.tag.String() + " does not support multiple property")
	}

	e.selfElement.Set("multiple", multiple)
	return e
}
