package eventPageTransition

type EventPageTransition int

func (el EventPageTransition) String() string {
	return eventPageTransitionString[el]
}

var eventPageTransitionString = [...]string{
	"pagehide",
	"pageshow",
}

const (
	// KPageHide
	// en: The event occurs when the user navigates away from a webpage
	KPageHide EventPageTransition = iota

	// KPageShow
	// en: The event occurs when the user navigates to a webpage
	KPageShow
)
