package eventInput

type EventInput int

func (el EventInput) String() string {
	return eventInputString[el]
}

var eventInputString = [...]string{
	"input",
}

const (
	// en: The event occurs when an element gets user input
	KInput EventInput = iota
)
