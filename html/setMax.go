package html

import "log"

// SetMax
//
// English:
//
//  Valid for date, month, week, time, datetime-local, number, and range, it defines the greatest
//  value in the range of permitted values.
//  If the value entered into the element exceeds this, the element fails constraint validation.
//  If the value of the max attribute isn't a number, then the element has no maximum value.
//
// There is a special case: if the data type is periodic (such as for dates or times), the value of
// max may be lower than the value of min, which indicates that the range may wrap around;
// for example, this allows you to specify a time range from 10 PM to 4 AM.
//
// Português:
//
//  Válido para data, mês, semana, hora, datetime-local, número e intervalo, define o maior valor no
//  intervalo de valores permitidos. Se o valor inserido no elemento exceder isso, o elemento falhará
//  na validação de restrição. Se o valor do atributo max não for um número, o elemento não terá
//  valor máximo.
//
// Há um caso especial: se o tipo de dado for periódico (como para datas ou horas), o valor de max
// pode ser menor que o valor de min, o que indica que o intervalo pode ser contornado;
// por exemplo, isso permite que você especifique um intervalo de tempo das 22h às 4h.
func (e *GlobalAttributes) SetMax(max int) (ref *GlobalAttributes) {
	switch e.tag {
	case KTagInput:
	default:
		log.Printf("tag " + e.tag.String() + " does not support max property")
	}

	e.selfElement.Set("max", max)
	return e
}
