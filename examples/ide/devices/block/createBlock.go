package block

import (
	"fmt"
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/html"
)

// createBlock Prepare all divs and CSS
func (e *Block) createBlock(x, y, width, height int) {
	e.resizers = make([]*html.TagDiv, 0)

	e.block = factoryBrowser.NewTagDiv().
		Id(e.id).
		//Class(e.classListName).
		AddStyle("position", "absolute").
		AddStyle("top", fmt.Sprintf("%dpx", x)).
		AddStyle("left", fmt.Sprintf("%dpx", y)).
		AddStyle("width", fmt.Sprintf("%dpx", width)).
		AddStyle("height", fmt.Sprintf("%dpx", height))
	e.ideStage.Append(e.block)

	e.selectDiv = factoryBrowser.NewTagDiv().
		AddStyle("position", "absolute").
		AddStyle("display", "none").
		AddStyle("top", "0px").
		AddStyle("left", "0px").
		AddStyle("width", fmt.Sprintf("%dpx", width)).
		AddStyle("height", fmt.Sprintf("%dpx", height)).
		AddStyle("border", "1px dashed red").
		AddStyle("background", "transparent")
	e.block.Append(e.selectDiv)

	e.resizerTopLeft = factoryBrowser.NewTagDiv().
		DataKey("name", "top-left").
		AddStyle("position", "absolute").
		AddStyle("width", fmt.Sprintf("%dpx", e.resizerWidth)).
		AddStyle("height", fmt.Sprintf("%dpx", e.resizerHeight)).
		AddStyle("background-color", e.resizerColor).
		AddStyle("border-radius", fmt.Sprintf("%dpx", e.resizerRadius)).
		AddStyle("top", fmt.Sprintf("%dpx", e.resizerDistance)).
		AddStyle("left", fmt.Sprintf("%dpx", e.resizerDistance)).
		AddStyle("cursor", "nwse-resize")
	e.block.Append(e.resizerTopLeft)

	e.resizerTopRight = factoryBrowser.NewTagDiv().
		DataKey("name", "top-right").
		AddStyle("position", "absolute").
		AddStyle("width", fmt.Sprintf("%dpx", e.resizerWidth)).
		AddStyle("height", fmt.Sprintf("%dpx", e.resizerHeight)).
		AddStyle("background-color", e.resizerColor).
		AddStyle("border-radius", fmt.Sprintf("%dpx", e.resizerRadius)).
		AddStyle("top", fmt.Sprintf("%dpx", e.resizerDistance)).
		AddStyle("right", fmt.Sprintf("%dpx", e.resizerDistance)).
		AddStyle("cursor", "nesw-resize")
	e.block.Append(e.resizerTopRight)

	e.resizerBottomLeft = factoryBrowser.NewTagDiv().
		DataKey("name", "bottom-left").
		AddStyle("position", "absolute").
		AddStyle("width", fmt.Sprintf("%dpx", e.resizerWidth)).
		AddStyle("height", fmt.Sprintf("%dpx", e.resizerHeight)).
		AddStyle("background-color", e.resizerColor).
		AddStyle("border-radius", fmt.Sprintf("%dpx", e.resizerRadius)).
		AddStyle("bottom", fmt.Sprintf("%dpx", e.resizerDistance)).
		AddStyle("left", fmt.Sprintf("%dpx", e.resizerDistance)).
		AddStyle("cursor", "nesw-resize")
	e.block.Append(e.resizerBottomLeft)

	e.resizerBottomRight = factoryBrowser.NewTagDiv().
		DataKey("name", "bottom-right").
		AddStyle("position", "absolute").
		AddStyle("width", fmt.Sprintf("%dpx", e.resizerWidth)).
		AddStyle("height", fmt.Sprintf("%dpx", e.resizerHeight)).
		AddStyle("background-color", e.resizerColor).
		AddStyle("border-radius", fmt.Sprintf("%dpx", e.resizerRadius)).
		AddStyle("bottom", fmt.Sprintf("%dpx", e.resizerDistance)).
		AddStyle("right", fmt.Sprintf("%dpx", e.resizerDistance)).
		AddStyle("cursor", "nwse-resize")
	e.block.Append(e.resizerBottomRight)

	_ = e.updateOrnament()
}
