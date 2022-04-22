package html

import "log"

// SetMin
//
// English:
//
//  Valid for date, month, week, time, datetime-local, number, and range, it defines the most negative
//  value in the range of permitted values.
//
// If the value entered into the element is less than this, the element fails constraint validation.
// If the value of the min attribute isn't a number, then the element has no minimum value.
//
// This value must be less than or equal to the value of the max attribute. If the min attribute is
// present but is not specified or is invalid, no min value is applied. If the min attribute is valid
// and a non-empty value is less than the minimum allowed by the min attribute, constraint validation
// will prevent form submission. See Client-side validation for more information.
//
// There is a special case: if the data type is periodic (such as for dates or times), the value of
// max may be lower than the value of min, which indicates that the range may wrap around; for
// example, this allows you to specify a time range from 10 PM to 4 AM.
//
// Português:
//
//  Válido para data, mês, semana, hora, data e hora local, número e intervalo, define o valor mais
//  negativo no intervalo de valores permitidos.
//
// Se o valor inserido no elemento for menor que isso, o elemento falhará na validação de restrição.
// Se o valor do atributo min não for um número, o elemento não terá valor mínimo.
//
// Esse valor deve ser menor ou igual ao valor do atributo max. Se o atributo min estiver presente,
// mas não for especificado ou for inválido, nenhum valor min será aplicado. Se o atributo min for
// válido e um valor não vazio for menor que o mínimo permitido pelo atributo min, a validação de
// restrição impedirá o envio do formulário. Consulte Validação do lado do cliente para obter mais
// informações.
//
// Há um caso especial: se o tipo de dado for periódico (como para datas ou horas), o valor de max
// pode ser menor que o valor de min, o que indica que o intervalo pode ser contornado; por exemplo,
// isso permite que você especifique um intervalo de tempo das 22h às 4h.
func (e *GlobalAttributes) SetMin(min int) (ref *GlobalAttributes) {
	switch e.tag {
	case KTagInput:
	default:
		log.Printf("tag " + e.tag.String() + " does not support min property")
	}

	e.selfElement.Set("max", min)
	return e
}
