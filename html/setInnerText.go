package html

// SetInnerText
//
// English:
//
//  The innerText property of the HTMLElement interface represents the rendered text content of a node
//  and its descendants.
//
// As a getter, it approximates the text the user would get if they highlighted the contents of the
// element with the cursor and then copied it to the clipboard. As a setter this will replace the
// element's children with the given value, converting any line breaks into <br> elements.
//
//   Note:
//     * SetInnerText() is easily confused with SetTextContent(), but there are important differences
//       between the two:
//       Basically, SetInnerText() is aware of the rendered appearance of text, while SetTextContent()
//       is not.
//
// Value: A DOMString representing the rendered text content of an element.
//
// If the element itself is not being rendered (for example, is detached from the document or is
// hidden from view), the returned value is the same as the GetTextContent() property.
//
// Português:
//
//  A propriedade innerText da interface HTMLElement representa o conteúdo de texto renderizado de um
//  nó e seus descendentes.
//
// Como um getter, ele aproxima o texto que o usuário obteria se destacasse o conteúdo do elemento com
// o cursor e o copiasse para a área de transferência. Como um setter, isso substituirá os filhos do
// elemento pelo valor fornecido, convertendo qualquer quebra de linha em elementos <br>.
//
//   Nota:
//     * SetInnerText() é facilmente confundido com SetTextContent(), mas existem diferenças
//       importantes entre os dois:
//       Basicamente, SetInnerText() está ciente da aparência renderizada do texto, enquanto
//       SetTextContent() não.
//
// Valor: Um DOMString que representa o conteúdo de texto renderizado de um elemento.
//
// Se o próprio elemento não estiver sendo renderizado (por exemplo, estiver desanexado do documento
// ou estiver oculto da exibição), o valor retornado será o mesmo que a propriedade GetTextContent().
func (e *GlobalAttributes) SetInnerText(text string) (ref *GlobalAttributes) {
	e.selfElement.Set("innerText", text)
	return e
}
