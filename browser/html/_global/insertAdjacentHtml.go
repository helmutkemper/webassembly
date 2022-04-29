package _global

// InsertAdjacentHtml
//
// English:
//
//  The InsertAdjacentHtml() method of the Element interface parses the specified text as HTML or XML
//  and inserts the resulting nodes into the DOM tree at a specified position.
//
//   Input:
//     position: a representing the position relative to the element.
//       KAdjacentPositionBeforeBegin: Before the element. Only valid if the element is in the DOM
//         tree and has a parent element.
//       KAdjacentPositionAfterBegin: Just inside the element, before its first child.
//       KAdjacentPositionBeforeEnd: Just inside the element, after its last child.
//       KAdjacentPositionAfterEnd: After the element. Only valid if the element is in the DOM tree
//         and has a parent element.
//     text: The string to be parsed as HTML or XML and inserted into the tree.
//
// Português:
//
//  O método InsertAdjacentHtml() da interface Element analisa o texto especificado como HTML ou XML e
//  insere os nós resultantes na árvore DOM em uma posição especificada.
//
//   Entrada:
//     position: a representando a posição relativa ao elemento.
//       KAdjacentPositionBeforeBegin: Antes do elemento. Válido apenas se o elemento estiver na
//         árvore DOM e tiver um elemento pai.
//       KAdjacentPositionAfterBegin: Apenas dentro do elemento, antes de seu primeiro filho.
//       KAdjacentPositionBeforeEnd: Apenas dentro do elemento, após seu último filho.
//       KAdjacentPositionAfterEnd: Depois do elemento. Válido apenas se o elemento estiver na árvore
//         DOM e tiver um elemento pai.
//     text: A string a ser analisada como HTML ou XML e inserida na árvore.
func (e *GlobalAttributes) InsertAdjacentHtml(position AdjacentPosition, text string) (ref *GlobalAttributes) {
	e.selfElement.Call("insertAdjacentHTML", position.String(), text)
	return e
}
