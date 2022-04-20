package html

import "log"

// SetNewOption
//
// English:
//
//  The <option> HTML element is used to define an item contained in a <select>, an <optgroup>, or
//  a <datalist> element. As such, <option> can represent menu items in popups and other lists of
//  items in an HTML document.
//
//   Input:
//     id: a unique id for an element;
//     label: This attribute is text for the label indicating the meaning of the option. If the label
//       attribute isn't defined, its value is that of the element text content;
//     value: The content of this attribute represents the value to be submitted with the form, should
//       this option be selected. If this attribute is omitted, the value is taken from the text
//       content of the option element;
//     disabled: If this Boolean attribute is set, this option is not checkable. Often browsers grey
//       out such control and it won't receive any browsing event, like mouse clicks or focus-related
//       ones. If this attribute is not set, the element can still be disabled if one of its ancestors
//       is a disabled <optgroup> element;
//     selected: If present, this Boolean attribute indicates that the option is initially selected.
//       If the <option> element is the descendant of a <select> element whose multiple attribute is
//       not set, only one single <option> of this <select> element may have the selected attribute.
//
// Português:
//
//  O elemento HTML <option> é usado para definir um item contido em um elemento <select>, <optgroup>
//  ou <datalist>. Como tal, <option> pode representar itens de menu em pop-ups e outras listas de
//  itens em um documento HTML.
//
//   Entrada:
//     id: um id exclusivo para um elemento;
//     label: Este atributo é um texto para o rótulo que indica o significado da opção. Se o atributo
//       label não estiver definido, seu valor será o do conteúdo do texto do elemento;
//     value: O conteúdo deste atributo representa o valor a ser enviado com o formulário, caso esta
//       opção seja selecionada. Se este atributo for omitido, o valor será obtido do conteúdo de
//       texto do elemento de opção;
//     disabled: Se este atributo booleano estiver definido, esta opção não poderá ser marcada.
//       Muitas vezes, os navegadores desativam esse controle e não recebem nenhum evento de
//       navegação, como cliques do mouse ou relacionados ao foco. Se este atributo não for definido,
//       o elemento ainda poderá ser desabilitado se um de seus ancestrais for um elemento <optgroup>
//       desabilitado;
//     selected: Se presente, este atributo booleano indica que a opção foi selecionada inicialmente.
//       Se o elemento <option> é descendente de um elemento <select> cujo atributo múltiplo não está
//       definido, apenas um único <option> deste elemento <select> pode ter o atributo selecionado.
func (e *GlobalAttributes) SetNewOption(id, label, value string, disabled, selected bool) (ref *GlobalAttributes) {
	switch e.tag {
	case KTagDatalist:
	default:
		log.Printf("tag " + e.tag.String() + " does not support option property")
	}

	ref = &GlobalAttributes{}
	ref.CreateElement(KTagOption)
	ref.SetId(id)
	ref.SetValue(value)
	ref.SetTextContent(label)

	if disabled == true {
		ref.SetDisabled(disabled)
	}

	if selected == true {
		ref.SetSelected(selected)
	}

	e.Append(ref.selfElement)

	return e
}
