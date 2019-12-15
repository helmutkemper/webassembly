package eventClipBoard

type EventClipBoard int

func (el EventClipBoard) String() string {
	return eventClipBoardString[el]
}

var eventClipBoardString = [...]string{
	"copy",
	"cut",
	"paste",
}

const (
	// en: The event occurs when the user copies the content of an element
	KCopy EventClipBoard = iota

	// en: The event occurs when the user cuts the content of an element
	KCut

	// en: The event occurs when the user pastes some content in an element
	KPaste
)
