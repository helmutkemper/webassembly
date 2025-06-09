package block

import (
	"fmt"
	"strings"
	"syscall/js"
)

// initEvents initialize mouse events
func (e *Block) initEvents() {
	var isDragging, isResizing bool
	var startX, startY, startWidth, startHeight, startLeft, startTop int

	// add / remove event listener requires pointers, so the variable should be initialized in this way
	var drag, stopDrag, resizeMouseMove, stopResize js.Func

	// Calculates the X position of the drag
	dragX := func(element js.Value) {
		dx := element.Get("screenX").Int() - startX
		newLeft := e.min(e.max(0, startLeft+dx), e.ideStage.GetClientWidth()-e.block.GetOffsetWidth())
		e.block.AddStyle("left", fmt.Sprintf("%dpx", newLeft))
	}

	// Calculates the Y position of the drag
	dragY := func(element js.Value) {
		dy := element.Get("screenY").Int() - startY
		newTop := e.min(e.max(0, startTop+dy), e.ideStage.GetClientHeight()-e.block.GetOffsetHeight())
		e.block.AddStyle("top", fmt.Sprintf("%dpx", newTop))
	}

	// Joins the calculations of X and Y of the drag
	drag = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if !isDragging {
			return nil
		}

		e.block.AddStyle("cursor", "grabbing")
		element := args[0]
		dragX(element)
		dragY(element)
		return nil
	})

	// Removes events when the drag ends
	stopDrag = js.FuncOf(func(this js.Value, args []js.Value) interface{} { // feito
		isDragging = false
		e.block.AddStyle("cursor", "")

		js.Global().Call("removeEventListener", "mousemove", drag)
		js.Global().Call("removeEventListener", "mouseup", stopDrag)
		return nil
	})

	// Adds the device drag event when the mouse pointer is pressed
	e.block.Get().Call("addEventListener", "mousedown", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		element := args[0]
		if !e.dragEnabled || strings.Contains(element.Get("classList").String(), "resizer") {
			return nil
		}

		isDragging = true
		startX = element.Get("screenX").Int()
		startY = element.Get("screenY").Int()
		startLeft = e.block.GetOffsetLeft()
		startTop = e.block.GetOffsetTop()

		// The movement of the mouse must be captured from the document and not the dragged element, or when the mouse moves
		// very fast, the drag to
		js.Global().Call("addEventListener", "mousemove", drag)
		js.Global().Call("addEventListener", "mouseup", stopDrag)

		return nil
	}))

	// When the resizing tool is activated, four rectangles are designed in the corners of the device.
	// These rectangles are called top-right e top-left, bottom-right, bottom-left.
	//
	// [tl]           [tr]
	//    +-----------+
	//    |           |
	//    |  device   |
	//    |           |
	//    +-----------+
	// [bl]           [br]

	resizeHorizontal := func(element js.Value, name string) {
		dx := element.Get("screenX").Int() - startX
		newLeft := startLeft
		newWidth := startWidth

		if name == "bottom-right" {
			newWidth = e.min(startWidth+dx, e.ideStage.Get().Get("clientWidth").Int()-startLeft)
		} else if name == "bottom-left" {
			newWidth = e.min(startWidth-dx, startLeft+startWidth)
			newLeft = e.max(0, startLeft+dx)
		} else if name == "top-right" {
			newWidth = e.min(startWidth+dx, e.ideStage.Get().Get("clientWidth").Int()-startLeft)
		} else if name == "top-left" {
			newWidth = e.min(startWidth-dx, startLeft+startWidth)
			newLeft = e.max(0, startLeft+dx)
		}

		// [tl]           [tr]
		//    +-----------+
		//    |           |
		//    |  device   |
		//    |           |
		//    +-----------+
		// [bl]           [br]
		//
		// Prevents the effect:
		//   When drag TR or BR left, and the size is below minimum, the block is dragged left.
		if newWidth < e.blockMinimumWidth {
			return
		}

		newWidth = e.max(e.blockMinimumWidth, newWidth)

		e.block.AddStyle("left", fmt.Sprintf("%dpx", newLeft))
		e.block.AddStyle("width", fmt.Sprintf("%dpx", newWidth))
		e.selectDiv.AddStyle("width", fmt.Sprintf("%dpx", newWidth))
	}

	resizeVertical := func(element js.Value, name string) {
		dy := element.Get("screenY").Int() - startY
		newTop := startTop
		newHeight := startHeight

		if name == "bottom-right" {
			newHeight = e.min(startHeight+dy, e.ideStage.Get().Get("clientHeight").Int()-startTop)
		} else if name == "bottom-left" {
			newHeight = e.min(startHeight+dy, e.ideStage.Get().Get("clientHeight").Int()-newTop)
		} else if name == "top-right" {
			newHeight = e.min(startHeight-dy, startTop+startHeight)
			newTop = e.max(0, startTop+dy)
		} else if name == "top-left" {
			newHeight = e.min(startHeight-dy, startTop+startHeight)
			newTop = e.max(0, startTop+dy)
		}

		// [tl]           [tr]
		//    +-----------+
		//    |           |
		//    |  device   |
		//    |           |
		//    +-----------+
		// [bl]           [br]
		//
		// Prevents the effect:
		//   When drag TL or TR down, and the size is below minimum, the block is dragged down.
		if newHeight < e.blockMinimumHeight {
			return
		}

		newHeight = e.max(e.blockMinimumHeight, newHeight)

		e.block.AddStyle("top", fmt.Sprintf("%dpx", newTop))
		e.block.AddStyle("height", fmt.Sprintf("%dpx", newHeight))
		e.selectDiv.AddStyle("height", fmt.Sprintf("%dpx", newHeight))
	}

	resizeMouseMove = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if !isResizing {
			return nil
		}

		element := args[0]
		resizerName := e.block.Get().Get("dataset").Get("resizeName").String()
		resizeHorizontal(element, resizerName)
		resizeVertical(element, resizerName)
		_ = e.updateOrnament()

		width := e.block.GetOffsetWidth()
		height := e.block.GetOffsetHeight()
		e.OnResize(element, width, height)

		return nil
	})

	stopResize = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		isResizing = false
		js.Global().Call("removeEventListener", "mousemove", resizeMouseMove)
		js.Global().Call("removeEventListener", "mouseup", stopResize)
		return nil
	})

	resizeFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if !e.resizeEnabled {
			return nil
		}

		resizerName := this.Get("dataset").Get("name").String()
		e.block.DataKey("resizeName", resizerName)

		element := args[0]
		element.Call("stopPropagation") // preventDefault
		isResizing = true
		startX = element.Get("screenX").Int()
		startY = element.Get("screenY").Int()
		startWidth = e.block.GetOffsetWidth()
		startHeight = e.block.GetOffsetHeight()
		startLeft = e.block.GetOffsetLeft()
		startTop = e.block.GetOffsetTop()

		js.Global().Call("addEventListener", "mousemove", resizeMouseMove)
		js.Global().Call("addEventListener", "mouseup", stopResize)
		return nil
	})

	e.resizerTopLeft.Get().Call("addEventListener", "mousedown", resizeFunc)
	e.resizerTopRight.Get().Call("addEventListener", "mousedown", resizeFunc)
	e.resizerBottomLeft.Get().Call("addEventListener", "mousedown", resizeFunc)
	e.resizerBottomRight.Get().Call("addEventListener", "mousedown", resizeFunc)
}
