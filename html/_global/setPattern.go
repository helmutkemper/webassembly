package _global

import "log"

// SetPattern
//
// English:
//
//  The pattern attribute, when specified, is a regular expression that the input's value must match
//  in order for the value to pass constraint validation. It must be a valid JavaScript regular
//  expression, as used by the RegExp type, and as documented in our guide on regular expressions;
//  the 'u' flag is specified when compiling the regular expression, so that the pattern is treated
//  as a sequence of Unicode code points, instead of as ASCII. No forward slashes should be specified
//  around the pattern text.
//
// If the pattern attribute is present but is not specified or is invalid, no regular expression is
// applied and this attribute is ignored completely. If the pattern attribute is valid and a non-empty
// value does not match the pattern, constraint validation will prevent form submission.
//
//   Note:
//     * If using the pattern attribute, inform the user about the expected format by including
//       explanatory text nearby. You can also include a title attribute to explain what the
//       requirements are to match the pattern; most browsers will display this title as a tooltip.
//       The visible explanation is required for accessibility. The tooltip is an enhancement.
//
// Português:
//
//  O atributo pattern, quando especificado, é uma expressão regular que o valor da entrada deve
//  corresponder para que o valor passe na validação de restrição. Deve ser uma expressão regular
//  JavaScript válida, conforme usada pelo tipo RegExp e conforme documentado em nosso guia sobre
//  expressões regulares; o sinalizador 'u' é especificado ao compilar a expressão regular, para que
//  o padrão seja tratado como uma sequência de pontos de código Unicode, em vez de como ASCII.
//  Nenhuma barra deve ser especificada ao redor do texto do padrão.
//
// Se o atributo pattern estiver presente, mas não for especificado ou for inválido, nenhuma
// expressão regular será aplicada e esse atributo será completamente ignorado. Se o atributo de
// padrão for válido e um valor não vazio não corresponder ao padrão, a validação de restrição
// impedirá o envio do formulário.
//
//   Nota:
//     * Se estiver usando o atributo pattern, informe o usuário sobre o formato esperado incluindo
//       um texto explicativo próximo. Você também pode incluir um atributo title para explicar quais
//       são os requisitos para corresponder ao padrão; a maioria dos navegadores exibirá este título
//       como uma dica de ferramenta. A explicação visível é necessária para acessibilidade. A dica
//       de ferramenta é um aprimoramento.
func (e *GlobalAttributes) SetPattern(pattern string) (ref *GlobalAttributes) {
	switch e.tag {
	case KTagInput:
	default:
		log.Printf("tag " + e.tag.String() + " does not support pattern property")
	}

	e.selfElement.Set("pattern", pattern)
	return e
}
