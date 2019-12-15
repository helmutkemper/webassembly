package eventWheel

type EventWheel int

func (el EventWheel) String() string {
	return eventWheelString[el]
}

var eventWheelString = [...]string{
	"mousewheel",
	"wheel",
}

const (
	// en: Deprecated. Use the wheel event instead
	KMouseWheel EventWheel = iota

	// en: The event occurs when the mouse wheel rolls up or down over an element
	KWheel
)
