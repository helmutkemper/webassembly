package eventMouse

import "strings"

type EventMouse int

func (el EventMouse) String() string {
	return eventMouseString[el]
}

func (el EventMouse) StringToType(value string) EventMouse {
	return eventFromStringMap[strings.ToLower(value)]
}

var eventFromStringMap = map[string]EventMouse{
	"click":       KClick,
	"contextMenu": KContextMenu,
	"doubleClick": KDoubleClick,
	"mouseDown":   KMouseDown,
	"mouseEnter":  KMouseEnter,
	"mouseLeave":  KMouseLeave,
	"mouseMove":   KMouseMove,
	"mouseOver":   KMouseOver,
	"mouseOut":    KMouseOut,
	"mouseUp":     KMouseUp,
}

var eventMouseString = [...]string{
	"click",
	"contextMenu",
	"doubleClick",
	"mouseDown",
	"mouseEnter",
	"mouseLeave",
	"mouseMove",
	"mouseOver",
	"mouseOut",
	"mouseUp",
}

const (
	// en: The event occurs when the user clicks on an element
	KClick EventMouse = iota

	// en: The event occurs when the user right-clicks on an element to open a context
	// menu
	KContextMenu

	// en: The event occurs when the user double-clicks on an element
	KDoubleClick

	// en: The event occurs when the user presses a mouse button over an element
	KMouseDown

	// en: The event occurs when the pointer is moved onto an element
	KMouseEnter

	// en: The event occurs when the pointer is moved out of an element
	KMouseLeave

	// en: The event occurs when the pointer is moving while it is over an element
	KMouseMove

	// en: The event occurs when the pointer is moved onto an element, or onto one of
	// its children
	KMouseOver

	// en: The event occurs when a user moves the mouse pointer out of an element, or
	// out of one of its children
	KMouseOut

	// en: The event occurs when a user releases a mouse button over an element
	KMouseUp
)
