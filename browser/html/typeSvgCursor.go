package html

type SvgCursor string

func (e SvgCursor) String() string {
	return string(e)
}

const (
	KSvgCursorAuto      SvgCursor = "auto"
	KSvgCursorCrossHair SvgCursor = "crosshair"
	KSvgCursorDefault   SvgCursor = "default"
	KSvgCursorPointer   SvgCursor = "pointer"
	KSvgCursorMove      SvgCursor = "move"
	KSvgCursorEResize   SvgCursor = " e-resize"
	KSvgCursorNeResize  SvgCursor = "ne-resize"
	KSvgCursorNwResize  SvgCursor = "nw-resize"
	KSvgCursorNResize   SvgCursor = "n-resize"
	KSvgCursorSeResize  SvgCursor = "se-resize"
	KSvgCursorSwResize  SvgCursor = "sw-resize"
	KSvgCursorSResize   SvgCursor = "s-resize"
	KSvgCursorWResize   SvgCursor = "w-resize"
	KSvgCursorText      SvgCursor = "text"
	KSvgCursorWait      SvgCursor = "wait"
	KSvgCursorHelp      SvgCursor = "help"
)
