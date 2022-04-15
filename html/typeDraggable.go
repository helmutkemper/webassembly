package html

// Draggable
//
// English:
//
//  Specifies whether an element is draggable or not
//
// Português:
//
//  Especifica se um elemento pode ser arrastado ou não
type Draggable int

func (e Draggable) String() (element string) {
	return draggableString[e]
}

const (
	// KDraggableYes
	//
	// English:
	//
	//  Specifies that the element is draggable.
	//
	// Português:
	//
	//  Especifica que o elemento pode ser arrastado.
	KDraggableYes = iota + 1

	// KDraggableNo
	//
	// English:
	//
	//  Specifies that the element is not draggable.
	//
	// Português:
	//
	//  Especifica que o elemento não pode ser arrastado.
	KDraggableNo

	// KDraggableAuto
	//
	// English:
	//
	//  Uses the default behavior of the browser.
	//
	// Português:
	//
	//  Usa o comportamento padrão do navegador.
	KDraggableAuto
)

var draggableString = [...]string{
	"",
	"true",
	"false",
	"auto",
}
