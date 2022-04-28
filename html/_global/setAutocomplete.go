package _global

import "log"

// SetAutocomplete
//
// English:
//
//  The HTML autocomplete attribute lets web developers specify what if any permission the user agent
//  has to provide automated assistance in filling out form field values, as well as guidance to the
//  browser as to the type of information expected in the field.
//
// It is available on <input> elements that take a text or numeric value as input, <textarea>
// elements, <select> elements, and <form> elements.
//
// The source of the suggested values is generally up to the browser; typically values come from past
// values entered by the user, but they may also come from pre-configured values. For instance, a
// browser might let the user save their name, address, phone number, and email addresses for
// autocomplete purposes. Perhaps the browser offers the ability to save encrypted credit card
// information, for autocompletion following an authentication procedure.
//
// If an <input>, <select> or <textarea> element has no autocomplete attribute, then browsers use the
// autocomplete attribute of the element's form owner, which is either the <form> element that the
// element is a descendant of, or the <form> whose id is specified by the form attribute of the
// element.
//
//   Note:
//     * In order to provide autocompletion, user-agents might require <input>/<select>/<textarea>
//       elements to:
//         Have a name and/or id attribute;
//         Be descendants of a <form> element;
//         The form to have a submit button.
//
// Português:
//
//  O atributo autocomplete HTML permite que os desenvolvedores da Web especifiquem se existe alguma
//  permissão que o agente do usuário tenha para fornecer assistência automatizada no preenchimento
//  dos valores dos campos do formulário, bem como orientação ao navegador quanto ao tipo de
//  informação esperado no campo.
//
// Ele está disponível em elementos <input> que recebem um texto ou valor numérico como entrada,
// elementos <textarea>, elementos <select> e elementos <form>.
//
// A origem dos valores sugeridos geralmente depende do navegador; normalmente os valores vêm de
// valores passados inseridos pelo usuário, mas também podem vir de valores pré-configurados.
// Por exemplo, um navegador pode permitir que o usuário salve seu nome, endereço, número de telefone
// e endereços de e-mail para fins de preenchimento automático. Talvez o navegador ofereça a
// capacidade de salvar informações de cartão de crédito criptografadas, para preenchimento automático
// após um procedimento de autenticação.
//
// Se um elemento <input>, <select> ou <textarea> não tiver um atributo autocomplete, os navegadores
// usarão o atributo autocomplete do proprietário do formulário do elemento, que é o elemento <form>
// do qual o elemento é descendente ou o < form> cujo id é especificado pelo atributo form do
// elemento.
//
//   Nota:
//     * Para fornecer preenchimento automático, os agentes do usuário podem exigir elementos
//       <input> / <select> / <textarea> para:
//         Ter um atributo name e ou id;
//         Ser descendentes de um elemento <form>;
//         O formulário para ter um botão de envio.
func (e *GlobalAttributes) SetAutocomplete(autocomplete Autocomplete) (ref *GlobalAttributes) {
	switch e.tag {
	case KTagInput:
	case KTagTextarea:
	case KTagSelect:
	default:
		log.Printf("tag " + e.tag.String() + " does not support autocomplete property")
	}

	e.selfElement.Set("autocomplete", autocomplete.String())
	return e
}
