package mouse

type CursorType int

func (el CursorType) String() string {
	return CursorTypes[el]
}

var CursorTypes = [...]string{
	"cursor:auto",
	"cursor:alias",
	"cursor:all-scroll",
	"cursor:cell",
	"cursor:context-menu",
	"cursor:col-resize",
	"cursor:copy",
	"cursor:crosshair",
	"cursor:default",
	"cursor:e-resize",
	"cursor:ew-resize",
	"cursor:help",
	"cursor:move",
	"cursor:n-resize",
	"cursor:ne-resize",
	"cursor:nesw-resize",
	"cursor:ns-resize",
	"cursor:nw-resize",
	"cursor:nwse-resize",
	"cursor:no-drop",
	"cursor:none",
	"cursor:not-allowed",
	"cursor:pointer",
	"cursor:progress",
	"cursor:row-resize",
	"cursor:s-resize",
	"cursor:se-resize",
	"cursor:sw-resize",
	"cursor:text",
	"cursor:URL",
	"cursor:vertical-text",
	"cursor:w-resize",
	"cursor:wait",
	"cursor:zoom-in",
	"cursor:zoom-out",
	"cursor:initial",
	"cursor:inherit",
}

const (
	// en: Default. The browser sets a cursor
	KCursorAuto CursorType = iota

	// en: The cursor indicates an alias of something is to be created
	KCursorAlias

	// en: The cursor indicates that something can be scrolled in any direction
	KCursorAllScroll

	// en: The cursor indicates that a cell (or set of cells) may be selected
	KCursorCell

	// en: The cursor indicates that a context-menu is available
	KCursorContextMenu

	// en: The cursor indicates that the column can be resized horizontally
	KCursorColResize

	// en: The cursor indicates something is to be copied
	KCursorCopy

	// en: The cursor render as a crosshair
	KCursorCrossHair

	// en: The default cursor
	KCursorDefault

	// en: The cursor indicates that an edge of a box is to be moved right (east)
	KCursorEResize

	// en: Indicates a bidirectional resize cursor
	KCursorEwResize

	// en: The cursor indicates that help is available
	KCursorHelp

	// en: The cursor indicates something is to be moved
	KCursorMove

	// en: The cursor indicates that an edge of a box is to be moved up (north)
	KCursorNResize

	// en: The cursor indicates that an edge of a box is to be moved up and right (north/east)
	KCursorNeResize

	// en: Indicates a bidirectional resize cursor
	KCursorNeswResize

	// en: Indicates a bidirectional resize cursor
	KCursorNsResize

	// en: The cursor indicates that an edge of a box is to be moved up and left (north/west)
	KCursorNwResize

	// en: Indicates a bidirectional resize cursor
	KCursorNwseResize

	// en: The cursor indicates that the dragged item cannot be dropped here
	KCursorNoDrop

	// en: No cursor is rendered for the element
	KCursorNone

	// en: The cursor indicates that the requested action will not be executed
	KCursorNotAllowed

	// en: The cursor is a pointer and indicates a link
	KCursorPointer

	// en: The cursor indicates that the program is busy (in progress)
	KCursorProgress

	// en: The cursor indicates that the row can be resized vertically
	KCursorRowResize

	// en: The cursor indicates that an edge of a box is to be moved down (south)
	KCursorSResize

	// en: The cursor indicates that an edge of a box is to be moved down and right (south/east)
	KCursorSeResize

	// en: The cursor indicates that an edge of a box is to be moved down and left (south/west)
	KCursorSwResize

	// en: The cursor indicates text that may be selected
	KCursorText

	// en: A comma separated list of URLs to custom cursors. Note: Always specify a generic cursor at the end of the list, in case none of the URL-defined cursors can be used
	KCursorURL

	// en: The cursor indicates vertical-text that may be selected
	KCursorVerticalText

	// en: The cursor indicates that an edge of a box is to be moved left (west)
	KCursorWResize

	// en: The cursor indicates that the program is busy
	KCursorWait

	// en: The cursor indicates that something can be zoomed in
	KCursorZoomIn

	// en: The cursor indicates that something can be zoomed out
	KCursorZoomOut

	// en: Sets this property to its default value. Read about initial
	KCursorInitial

	// en: Inherits this property from its parent element. Read about inherit
	KCursorInherit
)
