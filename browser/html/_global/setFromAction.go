package _global

import "log"

// SetFromAction
//
// English:
//
//  The URL that processes the information submitted by the button. Overrides the action attribute of
//  the button's form owner. Does nothing if there is no form owner.
//
//   Input:
//     action: The URL that processes the form submission. This value can be overridden by a
//             formaction attribute on a <button>, <input type="submit">, or <input type="image">
//             element. This attribute is ignored when method="dialog" is set.
//
// Português:
//
//  A URL que processa as informações enviadas pelo botão. Substitui o atributo de ação do
//  proprietário do formulário do botão. Não faz nada se não houver um proprietário de formulário.
//
//   Entrada:
//     action: A URL que processa o envio do formulário. Esse valor pode ser substituído por um
//             atributo formaction em um elemento <button>, <input type="submit"> ou
//             <input type="image">. Este atributo é ignorado quando method="dialog" é definido.
func (e *GlobalAttributes) SetFromAction(action string) (ref *GlobalAttributes) {
	switch e.tag {
	case KTagButton:
	default:
		log.Printf("tag " + e.tag.String() + " does not support formaction property")
	}

	e.selfElement.Set("formaction", action)
	return e
}
