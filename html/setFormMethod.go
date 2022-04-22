package html

import "log"

// SetFormMethod
//
// English:
//
//  If the button is a submit button (it's inside/associated with a <form> and doesn't have
//  type="button"), this attribute specifies the HTTP method used to submit the form.
//
//   Input:
//     method: specifies the HTTP method used to submit
//       KFormMethodPost: The data from the form are included in the body of the HTTP request when
//         sent to the server. Use when the form contains information that shouldn't be public, like
//         login credentials;
//       KFormMethodGet: The form data are appended to the form's action URL, with a ? as a separator,
//         and the resulting URL is sent to the server. Use this method when the form has no side
//         effects, like search forms;
//       KFormMethodDialog: When the form is inside a <dialog>, closes the dialog and throws a submit
//         event on submission without submitting data or clearing the form.
//
//   Note:
//     * If specified, this attribute overrides the method attribute of the button's form owner.
//
// Português:
//
//  Se o botão for um botão de envio (está associado a um <form> e não possui type="button"),
//  este atributo especifica o método HTTP usado para enviar o formulário.
//
//   Input:
//     method: especifica o método HTTP usado para enviar
//       KFormMethodPost: Os dados do formulário são incluídos no corpo da solicitação HTTP quando
//         enviados ao servidor. Use quando o formulário contém informações que não devem ser
//         públicas, como credenciais de login;
//       KFormMethodGet: Os dados do formulário são anexados à URL de ação do formulário, com um ?
//         como separador e a URL resultante é enviada ao servidor. Use este método quando o
//         formulário não tiver efeitos colaterais, como formulários de pesquisa;
//       KFormMethodDialog: Quando o formulário está dentro de um <dialog>, fecha o diálogo e lança um
//         evento submit no envio sem enviar dados ou limpar o formulário.
//
//   Nota:
//     * Se especificado, este atributo substitui o atributo method do proprietário do formulário do
//       botão.
func (e *GlobalAttributes) SetFormMethod(method FormMethod) (ref *GlobalAttributes) {
	switch e.tag {
	case KTagButton:
	case KTagForm:
	default:
		log.Printf("tag " + e.tag.String() + " does not support formmethod property")
	}

	e.selfElement.Set("formmethod", method.String())
	return e
}
