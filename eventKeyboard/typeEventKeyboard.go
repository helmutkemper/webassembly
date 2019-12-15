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
	// en: The event occurs when the user is pressing a key
	KKeyDown EventKeyboard = iota

	// en: The event occurs when the user presses a key
	KKeyPress

	// en: The event occurs when the user releases a key
	KKeyUp
)
