package _global

import "log"

// SetName
//
// English:
//
//  The name of the button, submitted as a pair with the button's value as part of the form data,
//  when that button is used to submit the form.
//
// Português:
//
//  O nome do botão, enviado como um par com o valor do botão como parte dos dados do formulário,
//  quando esse botão é usado para enviar o formulário.
func (e *GlobalAttributes) SetName(name string) (ref *GlobalAttributes) {
	switch e.tag {
	case KTagButton:
	case KTagFieldset:
	case KTagInput:
	default:
		log.Printf("tag " + e.tag.String() + " does not support name property")
	}

	e.selfElement.Set("name", name)
	return e
}
