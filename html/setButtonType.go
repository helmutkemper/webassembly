package html

import "log"

// SetButtonType
//
// English:
//
//  The default behavior of the button.
//
//   Input:
//     value: default behavior of the button.
//       KButtonTypeSubmit: The button submits the form data to the server. This is the default if
//         the attribute is not specified for buttons associated with a <form>, or if the attribute
//         is an empty or invalid value.
//       KButtonTypeReset:  The button resets all the controls to their initial values, like
//         <input type="reset">. (This behavior tends to annoy users.)
//       KButtonTypeButton: The button has no default behavior, and does nothing when pressed by
//         default. It can have client-side scripts listen to the element's events, which are
//         triggered when the events occur.
//
// Português:
//
//  O comportamento padrão do botão. Os valores possíveis são:
//
//   Entrada:
//     value: comportamento padrão do botão
//       KButtonTypeSubmit: O botão envia os dados do formulário para o servidor. Este é o padrão se
//         o atributo não for especificado para botões associados a um <form> ou se o atributo for um
//         valor vazio ou inválido.
//       KButtonTypeReset:  O botão redefine todos os controles para seus valores iniciais, como
//         <input type="reset">. (Esse comportamento tende a incomodar os usuários.)
//       KButtonTypeButton: O botão não tem comportamento padrão e não faz nada quando pressionado por
//         padrão. Ele pode fazer com que os scripts do lado do cliente escutem os eventos do
//         elemento, que são acionados quando os eventos ocorrem.
func (e *GlobalAttributes) SetButtonType(value ButtonType) (ref *GlobalAttributes) {
	switch e.tag {
	case KTagButton:
	default:
		log.Printf("tag " + e.tag.String() + " does not support buttonType property")
	}

	e.selfElement.Set("type", value.String())
	return e
}
