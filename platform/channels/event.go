package mouse

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
	// English: The event occurs when the user clicks on an element
	KClick Event = iota

	// KContextMenu
	// English: The event occurs when the user right-clicks on an element to open a context
	// menu
	KContextMenu

	// KDblClick
	// English: The event occurs when the user double-clicks on an element
	KDblClick

	// KMouseDown
	// English: The event occurs when the user presses a mouse button over an element
	KMouseDown

	// KMouseEnter
	// English: The event occurs when the pointer is moved onto an element
	KMouseEnter

	// KMouseLeave
	// English: The event occurs when the pointer is moved out of an element
	KMouseLeave

	// KMouseMove
	// English: The event occurs when the pointer is moving while it is over an element
	KMouseMove

	// KMouseOver
	// English: The event occurs when the pointer is moved onto an element, or onto one of
	// its children
	KMouseOver

	// KMouseOut
	// English: The event occurs when a user moves the mouse pointer out of an element, or
	// out of one of its children
	KMouseOut

	// KMouseUp
	// English: The event occurs when a user releases a mouse button over an element
	KMouseUp
)
