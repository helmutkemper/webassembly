package eventTransition

type EventTransition int

func (el EventTransition) String() string {
	return eventTransitionString[el]
}

var eventTransitionString = [...]string{
	"transitionend",
}

const (
	// en: The event occurs when a CSS transition has completed
	KTransitionend EventTransition = iota
)
