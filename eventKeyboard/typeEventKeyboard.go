package eventKeyboard

type EventKeyboard int

func (el EventKeyboard) String() string {
	return eventKeyboardString[el]
}

var eventKeyboardString = [...]string{
	"keydown",
	"keypress",
	"keyup",
}

const (
	// KKeyDown
	// en: The event occurs when the user is pressing a key
	KKeyDown EventKeyboard = iota

	// KKeyPress
	// en: The event occurs when the user presses a key
	KKeyPress

	// KKeyUp
	// en: The event occurs when the user releases a key
	KKeyUp
)
