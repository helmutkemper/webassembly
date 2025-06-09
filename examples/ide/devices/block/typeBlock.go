package block

import (
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/examples/ide/ornament"
	"github.com/helmutkemper/webassembly/examples/ide/utils"
	"image/color"
	"syscall/js"
)

type Block struct {
	utils.SequentialId

	id       string
	autoId   string
	fatherId string
	name     string

	x      int
	y      int
	width  int
	height int

	initialized bool

	resizerWidth    int
	resizerHeight   int
	resizerDistance int
	resizerRadius   int

	blockMinimumWidth  int
	blockMinimumHeight int

	classListName string

	isResizing      bool
	resizeEnabled   bool
	resizeBlocked   bool
	resizerFlashing bool
	selectFlashing  bool
	selected        bool
	selectBlocked   bool
	dragEnabled     bool
	dragBlocked     bool

	resizerColor      color.RGBA
	resizerFlashColor color.RGBA

	ideStage *html.TagDiv
	block    *html.TagDiv

	resizerTopLeft     *html.TagDiv
	resizerTopRight    *html.TagDiv
	resizerBottomLeft  *html.TagDiv
	resizerBottomRight *html.TagDiv

	selectDiv *html.TagDiv

	resizers []*html.TagDiv

	ornament ornament.Draw

	onResizeFunc func(element js.Value, width, height int)
}
