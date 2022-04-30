package eventTouch

type EventTouch int

func (el EventTouch) String() string {
	return eventTouchString[el]
}

var eventTouchString = [...]string{
	"touchcancel",
	"touchend",
	"touchmove",
	"touchstart",
}

const (
	// KTouchCancel
	// en: The event occurs when the touch is interrupted
	KTouchCancel EventTouch = iota

	// KTouchEnd
	// en: The event occurs when a finger is removed from a touch screen
	KTouchEnd

	// KTouchMove
	// en: The event occurs when a finger is dragged across the screen
	KTouchMove

	// KTouchStart
	// en: The event occurs when a finger is placed on a touch screen
	KTouchStart
)
