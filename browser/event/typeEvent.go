package event

type Event int

func (el Event) String() string {
	return eventString[el]
}

var eventString = [...]string{
	"abort",
	"afterprint",
	"beforeprint",
	"beforeunload",
	"canplay",
	"canplaythrough",
	"change",
	"durationchange",
	"ended",
	"error",
	"fullscreenchange",
	"fullscreenerror",
	"input",
	"invalid",
	"load",
	"loadeddata",
	"loadedmetadata",
	"message",
	"offline",
	"online",
	"open",
	"pause",
	"play",
	"playing",
	"progress",
	"ratechange",
	"resize",
	"reset",
	"scroll",
	"search",
	"seeked",
	"seeking",
	"select",
	"show",
	"stalled",
	"submit",
	"suspend",
	"timeupdate",
	"toggle",
	"unload",
	"volumechange",
	"waiting",
}

const (
	// en: The event occurs when the loading of a media is aborted
	KAbort Event = iota

	// en: The event occurs when a page has started printing, or if the print dialogue
	// box has been closed
	KAfterPrint

	// en: The event occurs when a page is about to be printed
	KBeforePrint

	// en: The event occurs before the document is about to be unloaded
	KBeforeUnload

	// en: The event occurs when the browser can start playing the media (when it has
	// buffered enough to begin)
	KCanPlay

	// en: The event occurs when the browser can play through the media without
	// stopping for buffering
	KCanPlayThrough

	// en: The event occurs when the content of a form element, the selection, or the
	// checked state have changed (for <input>, <select>, and <textarea>)
	KChange

	// en: The event occurs when the duration of the media is changed
	KDurationChange

	// en: The event occurs when the media has reach the end (useful for messages like
	// "thanks for listening")
	KEnded

	// en: The event occurs when an error occurs while loading an external file
	KError

	// en: The event occurs when an element is displayed in fullscreen mode
	KFullScreenChange

	// en: The event occurs when an element can not be displayed in fullscreen mode
	KFullScreenError

	// en: The event occurs when an element gets user input
	KInput

	// en: The event occurs when an element is invalid
	KInvalid

	// en: The event occurs when an object has loaded
	KLoad

	// en: The event occurs when media data is loaded
	KLoadedData

	// en: The event occurs when meta data (like dimensions and duration) are loaded
	KLoadedMetaData

	// en: The event occurs when a message is received through the event source
	KMessage

	// en: The event occurs when the browser starts to work offline
	KOffLine

	// en: The event occurs when the browser starts to work online
	KOnLine

	// en: The event occurs when a connection with the event source is opened
	KOpen

	// en: The event occurs when the media is paused either by the user or
	// programmatically
	KPause

	// en: The event occurs when the media has been started or is no longer paused
	KPlay

	// en: The event occurs when the media is playing after having been paused or
	// stopped for buffering
	KPlaying

	// en: The event occurs when the browser is in the process of getting the media
	// data (downloading the media)
	KProgress

	// en: The event occurs when the playing speed of the media is changed
	KRateChange

	// en: The event occurs when the document view is resized
	KResize

	// en: The event occurs when a form is reset
	KReset

	// en: The event occurs when an element's scrollbar is being scrolled
	KScroll

	// en: The event occurs when the user writes something in a search field
	// (for <input="search">)
	KSearch

	// en: The event occurs when the user is finished moving/skipping to a new position
	// in the media
	KSeeked

	// en: The event occurs when the user starts moving/skipping to a new position in
	// the media
	KSeeking

	// en: The event occurs after the user selects some text (for <input> and
	// <textarea>)
	KSelect

	// en: The event occurs when a <menu> element is shown as a context menu
	KShow

	// en: The event occurs when the browser is trying to get media data, but data is
	// not available
	KStalled

	// en: The event occurs when a form is submitted
	KSubmit

	// en: The event occurs when the browser is intentionally not getting media data
	KSuspend

	// en: The event occurs when the playing position has changed (like when the user
	// fast forwards to a different point in the media)
	KTimeUpdate

	// en: The event occurs when the user opens or closes the <details> element
	KToggle

	// en: The event occurs once a page has unloaded (for <body>)
	KUnload

	// en: The event occurs when the volume of the media has changed (includes setting
	// the volume to "mute")
	KVolumeChange

	// en: The event occurs when the media has paused but is expected to resume (like
	// when the media pauses to buffer more data)
	KWaiting
)
