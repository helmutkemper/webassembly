package eventAnimation

type EventAnimation int

func (el EventAnimation) String() string {
	return eventAnimationString[el]
}

var eventAnimationString = [...]string{
	"animationend",
	"animationiteration",
	"animationstart",
}

const (
	// KAnimationEnd
	// en: The event occurs when a CSS animation has completed
	KAnimationEnd EventAnimation = iota

	// KAnimationIteration
	// en: The event occurs when a CSS animation is repeated
	KAnimationIteration

	// KAnimationStart
	// en: The event occurs when a CSS animation has started
	KAnimationStart
)
