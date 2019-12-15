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
	// en: The event occurs when an element loses focus
	KBlur EventFocus = iota

	// en: The event occurs when an element gets focus
	KFocus

	// en: The event occurs when an element is about to get focus
	KFocusIn

	// en: The event occurs when an element is about to lose focus
	KFocusOut
)
