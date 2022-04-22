package html

import "log"

// SetMinLength
//
// English:
//
//  Valid for text, search, url, tel, email, and password, it defines the minimum number of
//  characters (as UTF-16 code units) the user can enter into the entry field.
//
// This must be an non-negative integer value smaller than or equal to the value specified by
// maxlength. If no minlength is specified, or an invalid value is specified, the input has no
// minimum length.
//
// The input will fail constraint validation if the length of the text entered into the field is
// fewer than minlength UTF-16 code units long, preventing form submission.
//
// Português:
//
//  Válido para texto, pesquisa, url, tel, email e senha, define o número mínimo de caracteres
//  (como unidades de código UTF-16) que o usuário pode inserir no campo de entrada.
//
// Este deve ser um valor inteiro não negativo menor ou igual ao valor especificado por maxlength.
// Se nenhum comprimento mínimo for especificado ou um valor inválido for especificado, a entrada não
// terá comprimento mínimo.
//
// A entrada falhará na validação de restrição se o comprimento do texto inserido no campo for
// inferior a unidades de código UTF-16 de comprimento mínimo, impedindo o envio do formulário.
func (e *GlobalAttributes) SetMinLength(minlength int) (ref *GlobalAttributes) {
	switch e.tag {
	case KTagInput:
	default:
		log.Printf("tag " + e.tag.String() + " does not support minlength property")
	}

	e.selfElement.Set("minlength", minlength)
	return e
}
