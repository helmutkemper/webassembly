package eventDrag

type EventDrag int

func (el EventDrag) String() string {
	return eventDragString[el]
}

var eventDragString = [...]string{
	"drag",
	"dragend",
	"dragenter",
	"dragleave",
	"dragover",
	"dragstart",
	"drop",
}

const (
	// KDrag
	// en: The event occurs when an element is being dragged
	KDrag EventDrag = iota

	// KDragEnd
	// en: The event occurs when the user has finished dragging an element
	KDragEnd

	// KDragEnter
	// en: The event occurs when the dragged element enters the drop target
	KDragEnter

	// KDragLeave
	// en: The event occurs when the dragged element leaves the drop target
	KDragLeave

	// KDragOver
	// en: The event occurs when the dragged element is over the drop target
	KDragOver

	// KDragStart
	// en: The event occurs when the user starts to drag an element
	KDragStart

	// KDrop
	// en: The event occurs when the dragged element is dropped on the drop target
	KDrop
)
