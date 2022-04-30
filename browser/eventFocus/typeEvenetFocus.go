package eventFocus

type EventFocus int

func (el EventFocus) String() string {
	return eventFocusString[el]
}

var eventFocusString = [...]string{
	"blur",
	"focus",
	"focusin",
	"focusout",
}

const (
	// KBlur
	// en: The event occurs when an element loses focus
	KBlur EventFocus = iota

	// KFocus
	// en: The event occurs when an element gets focus
	KFocus

	// KFocusIn
	// en: The event occurs when an element is about to get focus
	KFocusIn

	// KFocusOut
	// en: The event occurs when an element is about to lose focus
	KFocusOut
)
