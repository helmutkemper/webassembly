package block

import (
	"fmt"
	"github.com/helmutkemper/webassembly/examples/ide/utils"
	"github.com/helmutkemper/webassembly/platform/factoryColor"
)

// Init Initializes the generic functions of the device
func (e *Block) Init() (err error) {
	var base string
	if base, err = e.SequentialId.GetId(e.name); err != nil {
		return
	}
	if e.id, err = utils.VerifyUniqueId(base); err != nil {
		return
	}

	e.autoId = utils.GetRandomId()

	e.resizerWidth = 10
	e.resizerHeight = 10
	e.resizerDistance = -5
	e.resizerRadius = 2

	e.classListName = "block"

	e.resizerFlashing = true
	e.selectFlashing = true

	e.resizerFlashColor = factoryColor.NewYellow()
	e.resizerColor = factoryColor.NewRed()

	e.createBlock(e.x, e.y, e.width, e.height)
	e.initEvents()

	e.initialized = true

	e.block.AddStyle("left", fmt.Sprintf("%dpx", e.x))
	e.block.AddStyle("top", fmt.Sprintf("%dpx", e.y))
	e.block.AddStyle("width", fmt.Sprintf("%dpx", e.width))
	e.block.AddStyle("height", fmt.Sprintf("%dpx", e.height))

	if e.ornament != nil {
		svg := e.ornament.GetSvg()
		e.block.Append(svg)
		_ = e.updateOrnament()
	}

	e.dragCursorChange()
	e.resizeEnabledDraw()
	e.SetSelected(e.selected)

	return
}
