package html

import "log"

// SetFormEncType
//
// English:
//
//  If the button is a submit button (it's inside/associated with a <form> and doesn't have
//  type="button"), specifies how to encode the form data that is submitted. Possible values:
//
//   Input:
//     formenctype: specifies how to encode the form data
//
//       application/x-www-form-urlencoded: The default if the attribute is not used.
//       multipart/form-data: Use to submit <input> elements with their type attributes set to file.
//       text/plain: Specified as a debugging aid; shouldn't be used for real form submission.
//
//   Note:
//     * If this attribute is specified, it overrides the enctype attribute of the button's form
//       owner.
//
// Português:
//
//  Se o botão for um botão de envio (está associado a um <form> e não possui type="button"),
//  especifica como codificar os dados do formulário que são enviados. Valores possíveis:
//
//   Entrada:
//     formenctype: especifica como codificar os dados do formulário
//
//       KFormEncTypeApplication: O padrão se o atributo não for usado.
//       KFormEncTypeMultiPart: Use para enviar elementos <input> com seus atributos de tipo definidos
//         para arquivo.
//       KFormEncTypeText: Especificado como auxiliar de depuração; não deve ser usado para envio de
//         formulário real.
//
//   Note:
//     * Se este atributo for especificado, ele substituirá o atributo enctype do proprietário do
//       formulário do botão.
func (e *GlobalAttributes) SetFormEncType(formenctype FormEncType) (ref *GlobalAttributes) {
	switch e.tag {
	case KTagButton:
	case KTagForm:
	default:
		log.Printf("tag " + e.tag.String() + " does not support formenctype property")
	}

	e.selfElement.Set("formenctype", formenctype.String())
	return e
}
