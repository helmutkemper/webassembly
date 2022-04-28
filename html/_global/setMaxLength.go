package _global

import "log"

// SetMaxLength
//
// English:
//
//  Valid for text, search, url, tel, email, and password, it defines the maximum number of characters
//  (as UTF-16 code units) the user can enter into the field. This must be an integer value 0 or
//  higher. If no maxlength is specified, or an invalid value is specified, the field has no maximum
//  length. This value must also be greater than or equal to the value of minlength.
//
// The input will fail constraint validation if the length of the text entered into the field is
// greater than maxlength UTF-16 code units long. By default, browsers prevent users from entering
// more characters than allowed by the maxlength attribute.
//
// Português:
//
//  Válido para texto, pesquisa, url, tel, email e senha, define o número máximo de caracteres
//  (como unidades de código UTF-16) que o usuário pode inserir no campo.
//
// Este deve ser um valor inteiro 0 ou superior. Se nenhum comprimento máximo for especificado ou um
// valor inválido for especificado, o campo não terá comprimento máximo. Esse valor também deve ser
// maior ou igual ao valor de minlength.
//
// A entrada falhará na validação de restrição se o comprimento do texto inserido no campo for maior
// que o comprimento máximo das unidades de código UTF-16. Por padrão, os navegadores impedem que os
// usuários insiram mais caracteres do que o permitido pelo atributo maxlength.
func (e *GlobalAttributes) SetMaxLength(maxlength int) (ref *GlobalAttributes) {
	switch e.tag {
	case KTagInput:
	default:
		log.Printf("tag " + e.tag.String() + " does not support maxlength property")
	}

	e.selfElement.Set("maxlength", maxlength)
	return e
}
