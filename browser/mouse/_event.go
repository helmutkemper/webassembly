package mouse

// Event
//
// English:
//
//  Type referring to mouse events, such as click, move, up, down and others.
//
// Português:
//
//  Tipo referente aos eventos do mouse, como click, move, up, down, etc.
type Event string

func (e Event) String() string {
	return string(e)
}

const (
	// KEventClick
	//
	// English:
	//
	//  The event occurs when the user clicks on an element.
	//
	// Português:
	//
	//  O evento acontece quando o usuário clica em um elemento.
	KEventClick Event = "click"

	// KEventContextMenu
	//
	// English:
	//
	//  The event occurs when the user right-clicks on an element to open a context menu.
	//
	// Português:
	//
	//  O evento acontece quando o usuário usa o clique direito para abrir o menu de contexto.
	KEventContextMenu Event = "contextmenu"

	// KEventDoubleClick
	//
	// English:
	//
	//  The event occurs when the user double-clicks on an element.
	//
	// Português:
	//
	//  O evento acontece quando o usuário dá dois cliques no elemento.
	KEventDoubleClick Event = "dblclick"

	// KEventMouseDown
	//
	// English:
	//
	//  The event occurs when the user presses a mouse button over an element.
	//
	// Português:
	//
	//  O evento acontece quando o usuário pressiona o botão do mouse em cima do elemento.
	KEventMouseDown Event = "mousedown"

	// KEventMouseEnter
	//
	// English:
	//
	//  The event occurs when the pointer is moved onto an element.
	//
	// Português:
	//
	//  O evento acontece quando o ponteiro do mouse é movido para um elemento
	KEventMouseEnter Event = "mouseenter"

	// KEventMouseLeave
	//
	// English:
	//
	//  The event occurs when the pointer is moved out of an element.
	//
	// Português:
	//
	//  O evento acontece quando o ponteiro do mouse é movido para fora do elemento.
	KEventMouseLeave Event = "mouseleave"

	// KEventMouseMove
	//
	// English:
	//
	//  The event occurs when the pointer is moving while it is over an element.
	//
	// Português:
	//
	//  O evento acontece durante o movimento do ponteiro do mouse em cima do elemento.
	KEventMouseMove Event = "mousemove"

	// KEventMouseOver
	//
	// English:
	//
	//  The event occurs when the pointer is moved onto an element, or onto one of its children.
	//
	// Português:
	//
	//  O evento acontece quando quando o ponteiro do mouse é movido para cima do elemento ou um de seus
	//  elementos filhos.
	KEventMouseOver Event = "mouseover"

	// KEventMouseOut
	//
	// English:
	//
	//  The event occurs when a user moves the mouse pointer out of an element, or out of one of its
	//  children.
	//
	// Português:
	//
	//  O evento acontece quando quando o ponteiro do mouse é movido para fora do elemento ou um de seus
	//  elementos filhos.
	KEventMouseOut Event = "mouseout"

	// KEventMouseUp
	//
	// English:
	//
	//  The event occurs when a user releases a mouse button over an element.
	//
	// Português:
	//
	//  O evento acontece quando o usuário solta o botão do mouse em cima do elemento.
	KEventMouseUp Event = "mouseup"
)
