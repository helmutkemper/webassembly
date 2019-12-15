package eventHashChange

type EventHashChange int

func (el EventHashChange) String() string {
	return eventHashChangeString[el]
}

var eventHashChangeString = [...]string{
	"hashchange",
}

const (
	// en: The event occurs when there has been changes to the anchor part of a URL
	KHashChange EventHashChange = iota
)
