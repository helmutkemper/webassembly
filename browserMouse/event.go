package browserMouse

// Event
//
// English:
//
//  Type referring to mouse events, such as click, move, up, down and others.
//
// Português:
//
//  Tipo referente aos eventos do mouse, como click, move, up, down, etc.
type Event int

func (el Event) String() string {
	return eventString[el]
}

var eventString = [...]string{
	"click",
	"contextmenu",
	"dblclick",
	"mousedown",
	"mouseenter",
	"mouseleave",
	"mousemove",
	"mouseover",
	"mouseout",
	"mouseup",
}

const (
	// KClick
	// en: The event occurs when the user clicks on an element
	KClick Event = iota

	// KContextMenu
	// en: The event occurs when the user right-clicks on an element to open a context
	// menu
	KContextMenu

	// KDblClick
	// en: The event occurs when the user double-clicks on an element
	KDblClick

	// KMouseDown
	// en: The event occurs when the user presses a mouse button over an element
	KMouseDown

	// KMouseEnter
	// en: The event occurs when the pointer is moved onto an element
	KMouseEnter

	// KMouseLeave
	// en: The event occurs when the pointer is moved out of an element
	KMouseLeave

	// KMouseMove
	// en: The event occurs when the pointer is moving while it is over an element
	KMouseMove

	// KMouseOver
	// en: The event occurs when the pointer is moved onto an element, or onto one of
	// its children
	KMouseOver

	// KMouseOut
	// en: The event occurs when a user moves the mouse pointer out of an element, or
	// out of one of its children
	KMouseOut

	// KMouseUp
	// en: The event occurs when a user releases a mouse button over an element
	KMouseUp
)
