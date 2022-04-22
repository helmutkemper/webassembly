package html

import "log"

// SetSize
//
// English:
//
//  If the control is presented as a scrolling list box (e.g. when multiple is specified), this
//  attribute represents the number of rows in the list that should be visible at one time.
//  Browsers are not required to present a select element as a scrolled list box. The default value
//  is 0.
//
//   Note:
//     * According to the HTML5 specification, the default value for size should be 1; however, in
//       practice, this has been found to break some web sites, and no other browser currently does
//       that, so Mozilla has opted to continue to return 0 for the time being with Firefox.
//
// Português:
//
//  Se o controle for apresentado como uma caixa de listagem de rolagem (por exemplo, quando múltiplo
//  é especificado), esse atributo representa o número de linhas na lista que devem estar visíveis ao
//  mesmo tempo. Os navegadores não precisam apresentar um elemento de seleção como uma caixa de
//  listagem rolada. O valor padrão é 0.
//
//   Nota:
//     * De acordo com a especificação HTML5, o valor padrão para tamanho deve ser 1; no entanto, na
//       prática, descobriu-se que isso quebra alguns sites, e nenhum outro navegador atualmente faz
//       isso, então a Mozilla optou por continuar retornando 0 por enquanto com o Firefox.
func (e *GlobalAttributes) SetSize(size int) (ref *GlobalAttributes) {
	switch e.tag {
	case KTagSelect:
	case KTagInput:
	default:
		log.Printf("tag " + e.tag.String() + " does not support size property")
	}

	e.selfElement.Set("size", size)
	return e
}
