package _global

import "log"

// SetPlaceholder
//
// English:
//
//  The placeholder attribute is a string that provides a brief hint to the user as to what kind of
//  information is expected in the field. It should be a word or short phrase that provides a hint
//  as to the expected type of data, rather than an explanation or prompt. The text must not include
//  carriage returns or line feeds. So for example if a field is expected to capture a user's first
//  name, and its label is "First Name", a suitable placeholder might be "e.g. Mustafa".
//
//   Note:
//     * The placeholder attribute is not as semantically useful as other ways to explain your form,
//       and can cause unexpected technical issues with your content. See Labels for more information.
//
// Português:
//
//  O atributo placeholder é uma string que fornece uma breve dica ao usuário sobre que tipo de
//  informação é esperada no campo. Deve ser uma palavra ou frase curta que forneça uma dica sobre o
//  tipo de dados esperado, em vez de uma explicação ou prompt. O texto não deve incluir retornos de
//  carro ou feeds de linha. Assim, por exemplo, se espera-se que um campo capture o primeiro nome de
//  um usuário e seu rótulo for "Nome", um espaço reservado adequado pode ser "por exemplo, Mustafa".
//
//   Nota:
//     * O atributo placeholder não é tão semanticamente útil quanto outras formas de explicar seu
//       formulário e pode causar problemas técnicos inesperados com seu conteúdo. Consulte Rótulos
//       para obter mais informações.
func (e *GlobalAttributes) SetPlaceholder(placeholder string) (ref *GlobalAttributes) {
	switch e.tag {
	case KTagInput:
	default:
		log.Printf("tag " + e.tag.String() + " does not support placeholder property")
	}

	e.selfElement.Set("placeholder", placeholder)
	return e
}
