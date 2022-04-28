package _global

import "log"

// SetStep
//
// English:
//
//  Valid for the numeric input types, including number, date/time input types, and range, the step
//  attribute is a number that specifies the granularity that the value must adhere to.
//
//   If not explicitly included:
//     * step defaults to 1 for number and range;
//     * For the date/time input types, step is expressed in seconds, with the default step being 60
//       seconds. The step scale factor is 1000 (which converts the seconds to milliseconds, as used
//       in other algorithms);
//     * The value must be a positive number—integer or float—or the special value any, which means
//       no stepping is implied, and any value is allowed (barring other constraints, such as min and
//       max).
//
// If any is not explicitly set, valid values for the number, date/time input types, and range input
// types are equal to the basis for stepping — the min value and increments of the step value, up to
// the max value, if specified.
//
// For example, if you have <input type="number" min="10" step="2">, then any even integer, 10 or
// greater, is valid. If omitted, <input type="number">, any integer is valid, but floats (like 4.2)
// are not valid, because step defaults to 1. For 4.2 to be valid, step would have had to be set to
// any, 0.1, 0.2, or any the min value would have had to be a number ending in .2, such as
// <input type="number" min="-5.2">
//
//   Note:
//     * When the data entered by the user doesn't adhere to the stepping configuration, the value is
//       considered invalid in constraint validation and will match the :invalid pseudoclass.
//
// Português:
//
//  Válido para os tipos de entrada numérica, incluindo número, tipos de entrada de data e hora e
//  intervalo, o atributo step é um número que especifica a granularidade à qual o valor deve aderir.
//
//   Se não estiver explicitamente incluído:
//     * step padroniza para 1 para número e intervalo.
//     * Para os tipos de entrada de data e hora, a etapa é expressa em segundos, com a etapa padrão
//       sendo 60 segundos. O fator de escala de passo é 1000 (que converte os segundos em
//       milissegundos, conforme usado em outros algoritmos).
//     * O valor deve ser um número positivo — inteiro ou flutuante — ou o valor especial any, o que
//       significa que nenhuma depuração está implícita e qualquer valor é permitido (exceto outras
//       restrições, como min e max).
//
// Se algum não for definido explicitamente, os valores válidos para o número, tipos de entrada de
// data e hora e tipos de entrada de intervalo são iguais à base para a depuração — o valor mínimo e
// os incrementos do valor da etapa, até o valor máximo, se especificado.
//
// Por exemplo, se você tiver <input type="number" min="10" step="2">, qualquer número inteiro par,
// 10 ou maior, é válido. Se omitido, <input type="number">, qualquer inteiro é válido, mas floats
// (como 4.2) não são válidos, porque step é padronizado como 1. Para 4.2 ser válido, step teria que
// ser definido como any, 0.1 , 0.2 ou qualquer valor mínimo teria que ser um número que terminasse
// em .2, como <input type="number" min="-5.2">
//
//   Nota:
//     * Quando os dados inseridos pelo usuário não estão de acordo com a configuração de stepping,
//       o valor é considerado inválido na validação da restrição e corresponderá à
//       :invalid pseudoclass.
func (e *GlobalAttributes) SetStep(step int) (ref *GlobalAttributes) {
	switch e.tag {
	case KTagInput:
	default:
		log.Printf("tag " + e.tag.String() + " does not support step property")
	}

	e.selfElement.Set("step", step)
	return e
}
