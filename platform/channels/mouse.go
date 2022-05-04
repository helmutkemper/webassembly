package mouse

// cursor name
type Cursor struct {
	Name string
}

type Coordinate struct {
	X int
	Y int
}

type Move Coordinate
type Click Coordinate
type DoubleClick Coordinate
type Press Coordinate
type Release Coordinate

var BrowserMouseToPlatformMouseCoordinate chan Move
var BrowserMouseClickToPlatformMouseClickEvent chan Click
var BrowserMouseDoubleClickToPlatformMouseDoubleClickEvent chan DoubleClick
var BrowserMouseDownToPlatformMouseDownEvent chan Press
var BrowserMouseUpToPlatformMouseUpEvent chan Release

func init() {
	BrowserMouseToPlatformMouseCoordinate = make(chan Move, 1)
	BrowserMouseClickToPlatformMouseClickEvent = make(chan Click, 1)
	BrowserMouseDoubleClickToPlatformMouseDoubleClickEvent = make(chan DoubleClick, 1)
	BrowserMouseDownToPlatformMouseDownEvent = make(chan Press, 1)
	BrowserMouseUpToPlatformMouseUpEvent = make(chan Release, 1)
}
