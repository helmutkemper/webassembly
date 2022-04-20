package html

import "log"

// SetSelected
//
// English:
//
//  If present, this Boolean attribute indicates that the option is initially selected. If the
//  <option> element is the descendant of a <select> element whose multiple attribute is not set,
//  only one single <option> of this <select> element may have the selected attribute.
//
// Português:
//
//  Se presente, este atributo booleano indica que a opção foi selecionada inicialmente. Se o elemento
//  <option> é descendente de um elemento <select> cujo atributo múltiplo não está definido, apenas um
//  único <option> deste elemento <select> pode ter o atributo selecionado.
func (e *GlobalAttributes) SetSelected(selected bool) (ref *GlobalAttributes) {
	switch e.tag {
	case KTagOption:
	default:
		log.Printf("tag " + e.tag.String() + " does not support selected property")
	}

	e.selfElement.Set("selected", selected)
	return e
}
