package html

type AdjacentPosition string

func (e AdjacentPosition) String() string {
	return string(e)
}

const (
	// KAdjacentPositionBeforeBegin
	//
	// English:
	//
	//  Before the element. Only valid if the element is in the DOM tree and has a parent element.
	//
	// Português:
	//
	//  Antes do elemento. Válido apenas se o elemento estiver na árvore DOM e tiver um elemento pai.
	KAdjacentPositionBeforeBegin AdjacentPosition = "beforebegin"

	// KAdjacentPositionAfterBegin
	//
	// English:
	//
	//  Just inside the element, before its first child.
	//
	// Português:
	//
	//  Apenas dentro do elemento, antes de seu primeiro filho.
	KAdjacentPositionAfterBegin AdjacentPosition = "afterbegin"

	// KAdjacentPositionBeforeEnd
	//
	// English:
	//
	//  Just inside the element, after its last child.
	//
	// Português:
	//
	//  Apenas dentro do elemento, após seu último filho.
	KAdjacentPositionBeforeEnd AdjacentPosition = "beforeend"

	// KAdjacentPositionAfterEnd
	//
	// english:
	//
	//  After the element. Only valid if the element is in the DOM tree and has a parent element.
	//
	// Português:
	//
	//  Depois do elemento. Válido apenas se o elemento estiver na árvore DOM e tiver um elemento pai.
	KAdjacentPositionAfterEnd AdjacentPosition = "afterend"
)
