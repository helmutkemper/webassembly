package html

import "log"

// SetFormValidate
//
// English:
//
//  If the button is a submit button, this Boolean attribute specifies that the form is not to be
//  validated when it is submitted.
//
// If this attribute is specified, it overrides the novalidate attribute of the button's form owner.
//
// Português:
//
//  Se o botão for um botão de envio, este atributo booleano especifica que o formulário não deve ser
//  validado quando for enviado.
//
// Se este atributo for especificado, ele substituirá o atributo novalidate do proprietário do
// formulário do botão.
func (e *GlobalAttributes) SetFormValidate(validate bool) (ref *GlobalAttributes) {
	switch e.tag {
	case KTagButton:
	case KTagForm:
	default:
		log.Printf("tag " + e.tag.String() + " does not support formvalidade property")
	}

	e.selfElement.Set("formnovalidate", validate)
	return e
}
