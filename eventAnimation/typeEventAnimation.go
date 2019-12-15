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
	// en: The event occurs when a CSS animation has completed
	KAnimationEnd EventAnimation = iota

	// en: The event occurs when a CSS animation is repeated
	KAnimationIteration

	// en: The event occurs when a CSS animation has started
	KAnimationStart
)
