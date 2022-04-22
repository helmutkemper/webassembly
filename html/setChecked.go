package html

import "log"

// SetChecked
//
// English:
//
//  Valid for both radio and checkbox types, checked is a Boolean attribute. If present on a radio
//  type, it indicates that the radio button is the currently selected one in the group of same-named
//  radio buttons. If present on a checkbox type, it indicates that the checkbox is checked by default
//  (when the page loads).
//  It does not indicate whether this checkbox is currently checked: if the checkbox's state is
//  changed, this content attribute does not reflect the change.
//  (Only the HTMLInputElement's checked IDL attribute is updated.)
//
//   Note:
//     * Unlike other input controls, a checkboxes and radio buttons value are only included in the
//       submitted data if they are currently checked. If they are, the name and the value(s) of the
//       checked controls are submitted.
//       For example, if a checkbox whose name is fruit has a value of cherry, and the checkbox is
//       checked, the form data submitted will include fruit=cherry. If the checkbox isn't active,
//       it isn't listed in the form data at all. The default value for checkboxes and radio buttons
//       is on.
//
// Português:
//
//  Válido para os tipos de rádio e caixa de seleção, marcado é um atributo booleano. Se estiver
//  presente em um tipo de rádio, indica que o botão de opção é o selecionado atualmente no grupo de
//  botões de opção com o mesmo nome. Se estiver presente em um tipo de caixa de seleção, indica que
//  a caixa de seleção está marcada por padrão (quando a página é carregada). Não indica se esta caixa
//  de seleção está marcada no momento: se o estado da caixa de seleção for alterado, esse atributo
//  de conteúdo não reflete a alteração.
//  (Apenas o atributo IDL verificado do HTMLInputElement é atualizado.)
//
//   Nota:
//     * Ao contrário de outros controles de entrada, um valor de caixas de seleção e botões de opção
//       só são incluídos nos dados enviados se estiverem marcados no momento. Se estiverem, o nome e
//       o(s) valor(es) dos controles verificados são enviados.
//       Por exemplo, se uma caixa de seleção cujo nome é fruta tiver o valor cereja e a caixa de
//       seleção estiver marcada, os dados do formulário enviados incluirão fruta=cereja.
//       Se a caixa de seleção não estiver ativa, ela não está listada nos dados do formulário.
//       O valor padrão para caixas de seleção e botões de opção é ativado.
func (e *GlobalAttributes) SetChecked(checked bool) (ref *GlobalAttributes) {
	switch e.tag {
	case KTagInput:
	default:
		log.Printf("tag " + e.tag.String() + " does not support checked property")
	}

	e.selfElement.Set("checked", checked)
	return e
}
