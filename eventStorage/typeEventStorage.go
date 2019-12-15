package eventStorage

type EventStorage int

func (el EventStorage) String() string {
	return eventStorageString[el]
}

var eventStorageString = [...]string{
	"storage",
}

const (
	// en: The event occurs when a Web Storage area is updated
	KStorage EventStorage = iota
)
