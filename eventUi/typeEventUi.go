package eventUi

type EventUi int

func (el EventUi) String() string {
	return eventUiString[el]
}

var eventUiString = [...]string{
	"abort",
	"beforeunload",
	"error",
	"load",
	"resize",
	"scroll",
	"select",
	"unload",
}

const (
	// KAbort
	// en: The event occurs when the loading of a media is aborted
	KAbort EventUi = iota

	// KBeforeUnload
	// en: The event occurs before the document is about to be unloaded
	KBeforeUnload

	// KError
	// en: The event occurs when an error occurs while loading an external file
	KError

	// KLoad
	// en: The event occurs when an object has loaded
	KLoad

	// KResize
	// en: The event occurs when the document view is resized
	KResize

	// KScroll
	// en: The event occurs when an element's scrollbar is being scrolled
	KScroll

	// KSelect
	// en: The event occurs after the user selects some text (for <input> and <textarea>)
	KSelect

	// KUnload
	// en: The event occurs once a page has unloaded (for <body>)
	KUnload
)
