package block

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/examples/ide/ornament"
	"github.com/helmutkemper/webassembly/examples/ide/rulesDensity"
	"github.com/helmutkemper/webassembly/examples/ide/rulesSequentialId"
	"github.com/helmutkemper/webassembly/examples/ide/rulesStage"
	"github.com/helmutkemper/webassembly/examples/ide/utils"
	"github.com/helmutkemper/webassembly/platform/components"
	"github.com/helmutkemper/webassembly/platform/factoryColor"
	"image/color"
	"log"
	"syscall/js"
)

type Block struct {
	id     string
	autoId string
	name   string

	x      rulesDensity.Density
	y      rulesDensity.Density
	width  rulesDensity.Density
	height rulesDensity.Density

	initialized bool

	blockMinimumWidth  rulesDensity.Density
	blockMinimumHeight rulesDensity.Density

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

	ideStage *html.TagSvg
	block    *html.TagSvg
	main     *html.TagSvg

	resizeButton ResizeButton

	resizerTopLeft     ResizeButton
	resizerTopRight    ResizeButton
	resizerBottomLeft  ResizeButton
	resizerBottomRight ResizeButton

	resizerTopMiddle    ResizeButton
	resizerBottomMiddle ResizeButton
	resizerLeftMiddle   ResizeButton
	resizerRightMiddle  ResizeButton

	selectDiv *html.TagSvgRect

	ornament ornament.Draw

	onResizeFunc func(args []js.Value, width, height rulesDensity.Density)

	gridAdjust rulesStage.GridAdjust
}

func (e *Block) SetMainSvg(svg *html.TagSvg) {
	e.main = svg
}

func (e *Block) SetResizeButton(resizeButton ResizeButton) {
	e.resizeButton = resizeButton
}

func (e *Block) SetGridAdjust(gridAdjust rulesStage.GridAdjust) {
	e.gridAdjust = gridAdjust
}

func (e *Block) adjustXYToGrid(x, y int) (cx, cy int) {
	return e.gridAdjust.AdjustCenter(x, y)
}

// GetInitialized Returns if the instance is ready for use
func (e *Block) GetInitialized() bool {
	return e.initialized
}

// SetWarning sets the visibility of the warning mark
func (e *Block) SetWarning(warning bool) {
	e.ornament.SetWarning(warning)
}

// SetDragBlocked Disables the use of the drag tool
func (e *Block) SetDragBlocked(blocked bool) {
	e.dragBlocked = blocked
}

// DragBlockedInvert Invert the drag tool enable status. Note: Used in the menu
func (e *Block) DragBlockedInvert() {
	e.dragBlocked = !e.dragBlocked
}

// GetDragBlocked Return the drag tool enable status
func (e *Block) GetDragBlocked() (blocked bool) {
	return e.dragBlocked
}

// GetDrag Return the drag tool status
func (e *Block) GetDrag() (enabled bool) {
	return e.dragEnabled
}

// SetDragInvert Invert the drag tool status. Note: Used in the menu
func (e *Block) SetDragInvert() {
	e.dragEnabled = !e.dragEnabled
}

// SetDrag Enables the device's drag tool
func (e *Block) SetDrag(enabled bool) {
	e.dragEnabled = enabled

	if !e.initialized {
		return
	}

	if e.dragBlocked {
		e.dragEnabled = false
	}

	e.dragCursorChange()

	if e.dragEnabled {
		e.SetResize(true)
	}

	if e.dragEnabled && e.selected {
		e.SetSelected(false)
	}
}

// ResizeInverter Invert the resize tool status
func (e *Block) ResizeInverter() {
	e.resizeEnabled = !e.resizeEnabled
}

// GetResize Return the resize tool status
func (e *Block) GetResize() (enabled bool) {
	return e.resizeEnabled
}

// SetResize Defines the resize tool status
func (e *Block) SetResize(enabled bool) {
	e.resizeEnabled = enabled

	if !e.initialized {
		return
	}

	if e.resizeBlocked {
		e.resizeEnabled = false
	}

	e.resizeEnabledDraw()
	if enabled && e.selected {
		e.SetSelected(false)
	}
}

// ResizeBlockedInvert Invert the status from disables resize tool. Note: Used in the menu
func (e *Block) ResizeBlockedInvert() {
	e.resizeBlocked = !e.resizeBlocked
}

// GetResizeBlocked Return the status from disables resize tool
func (e *Block) GetResizeBlocked() (blocked bool) {
	return e.resizeBlocked
}

// SetResizeBlocked Disables the use of the resize tool
func (e *Block) SetResizeBlocked(blocked bool) {
	e.resizeBlocked = blocked
}

// SelectBlockedInvert Invert the status of the selection tool lock. Note: Used in the menu
func (e *Block) SelectBlockedInvert() {
	e.selectBlocked = !e.selectBlocked
}

// GetSelectBlocked Returns the status of the selection tool lock
func (e *Block) GetSelectBlocked() (blocked bool) {
	return e.selectBlocked
}

// SetSelectBlocked Lock the use of the selection tool
func (e *Block) SetSelectBlocked(blocked bool) {
	e.selectBlocked = blocked
}

// SelectedInvert Invert the status of the selection tool. Note: Used in the menu
func (e *Block) SelectedInvert() {
	e.SetSelected(!e.selected)
}

// SetSelected Defines if the device selection tool is active
func (e *Block) SetSelected(selected bool) {
	e.selected = selected

	if !e.initialized {
		return
	}

	if e.selectBlocked {
		e.selected = false
	}

	e.selectDiv.AddStyleConditional(e.selected, "display", "block", "none")
	e.SetResize(false)
}

// GetSelected Return the select tool status
func (e *Block) GetSelected() (selected bool) {
	return e.selected
}

// createBlock Prepare all divs and CSS
func (e *Block) createBlock(x, y, width, height rulesDensity.Density) {

	xInt, yInt := e.adjustXYToGrid(x.GetInt(), y.GetInt())
	widthInt, heightInt := e.adjustXYToGrid(width.GetInt(), height.GetInt())

	x, y = rulesDensity.Convert(xInt), rulesDensity.Convert(yInt)
	width, height = rulesDensity.Convert(widthInt), rulesDensity.Convert(heightInt)

	e.block = factoryBrowser.NewTagSvg().
		Id(e.id).
		X(x.GetInt()).
		Y(y.GetInt()).
		Width(width.GetInt()).
		Height(height.GetInt())
	e.ideStage.Append(e.block)

	e.selectDiv = factoryBrowser.NewTagSvgRect().
		X(x.GetInt()).
		Y(y.GetInt()).
		Width(width.GetInt()).
		Height(height.GetInt()).
		Fill("none").Stroke("red").
		StrokeDasharray([]int{4}).
		StrokeWidth(1)
	e.ideStage.Append(e.selectDiv)

	e.resizerTopLeft = e.resizeButton.GetNew()
	e.resizerTopLeft.SetName("top-left")
	e.resizerTopLeft.SetCursor("nwse-resize")
	e.resizerTopLeft.SetCX(x)
	e.resizerTopLeft.SetCY(y)
	e.ideStage.Append(e.resizerTopLeft.GetSvg())

	e.resizerTopRight = e.resizeButton.GetNew()
	e.resizerTopRight.SetName("top-right")
	e.resizerTopRight.SetCursor("nesw-resize")
	e.resizerTopRight.SetCX(x + width)
	e.resizerTopRight.SetCY(y)
	e.ideStage.Append(e.resizerTopRight.GetSvg())

	log.Printf("y: %v height: %v", y.GetInt(), height.GetInt())
	e.resizerBottomLeft = e.resizeButton.GetNew()
	e.resizerBottomLeft.SetName("bottom-left")
	e.resizerBottomLeft.SetCursor("nesw-resize")
	e.resizerBottomLeft.SetCX(x)
	e.resizerBottomLeft.SetCY(y + height)
	e.ideStage.Append(e.resizerBottomLeft.GetSvg())

	e.resizerBottomRight = e.resizeButton.GetNew()
	e.resizerBottomRight.SetName("bottom-right")
	e.resizerBottomRight.SetCursor("nwse-resize")
	e.resizerBottomRight.SetCX(x + width)
	e.resizerBottomRight.SetCY(y + height)
	e.ideStage.Append(e.resizerBottomRight.GetSvg())

	//----------------------------------------------------

	e.resizerTopMiddle = e.resizeButton.GetNew()
	e.resizerTopMiddle.SetName("top-middle")
	e.resizerTopMiddle.SetCursor("ns-resize")
	e.resizerTopMiddle.SetCX(x + width/2)
	e.resizerTopMiddle.SetCY(y)
	e.ideStage.Append(e.resizerTopMiddle.GetSvg())

	e.resizerBottomMiddle = e.resizeButton.GetNew()
	e.resizerBottomMiddle.SetName("bottom-middle")
	e.resizerBottomMiddle.SetCursor("ns-resize")
	e.resizerBottomMiddle.SetCX(x + width/2)
	e.resizerBottomMiddle.SetCY(y + height)
	e.ideStage.Append(e.resizerBottomMiddle.GetSvg())

	e.resizerLeftMiddle = e.resizeButton.GetNew()
	e.resizerLeftMiddle.SetName("left-middle")
	e.resizerLeftMiddle.SetCursor("ew-resize")
	e.resizerLeftMiddle.SetCX(x)
	e.resizerLeftMiddle.SetCY(y + height/2)
	e.ideStage.Append(e.resizerLeftMiddle.GetSvg())

	e.resizerRightMiddle = e.resizeButton.GetNew()
	e.resizerRightMiddle.SetName("right-middle")
	e.resizerRightMiddle.SetCursor("ew-resize")
	e.resizerRightMiddle.SetCX(x + width)
	e.resizerRightMiddle.SetCY(y + height/2)
	e.ideStage.Append(e.resizerRightMiddle.GetSvg())
}

// dragCursorChange Change the cursor when the device is being dragged
func (e *Block) dragCursorChange() {
	if !e.initialized {
		return
	}

	e.block.AddStyleConditional(e.dragEnabled, "cursor", "grab", "")
}

// GetDeviceDiv Returns the div from device
func (e *Block) GetDeviceDiv() (element *html.TagSvg) {
	return e.block
}

// GetHeight returns the current height of the device.
func (e *Block) GetHeight() (height rulesDensity.Density) {
	return e.height
}

// GetID Returns the device's div ID
func (e *Block) GetID() (id string) {
	return e.id
}

// GetIdeStage Returns to Div where IDE is drawn
func (e *Block) GetIdeStage() (ideStage *html.TagSvg) {
	return e.ideStage
}

// GetName Returns the single name of the device
func (e *Block) GetName() (name string) {
	return e.name
}

// GetWidth returns the current width of the device.
func (e *Block) GetWidth() (width rulesDensity.Density) {
	return e.width
}

// GetX Returns to coordinate X of the browser screen
func (e *Block) GetX() (x rulesDensity.Density) {
	return e.x
}

// GetY Returns to coordinate Y of the browser screen
func (e *Block) GetY() (y rulesDensity.Density) {
	return e.y
}

// Init Initializes the generic functions of the device
func (e *Block) Init() (err error) {
	var id string
	id = rulesSequentialId.GetIdFromBase(e.name)
	if e.id, err = utils.VerifyUniqueId(id); err != nil {
		return
	}

	e.autoId = utils.GetRandomId()

	e.classListName = "block"

	e.resizerFlashing = true
	e.selectFlashing = true

	e.resizerFlashColor = factoryColor.NewYellow()
	e.resizerColor = factoryColor.NewRed()

	e.createBlock(e.x, e.y, e.width, e.height)
	e.initEvents()

	e.initialized = true

	if e.ornament != nil {
		svg := e.ornament.GetSvg()
		e.block.Append(svg)
		_ = e.ornament.Update(e.x, e.y, e.width, e.height)
	}

	e.dragCursorChange()
	e.resizeEnabledDraw()
	e.SetSelected(e.selected)

	return
}

// initEvents initialize mouse events
func (e *Block) initEvents() {
	var isDragging, isResizing bool
	var startX, startY, startWidth, startHeight, startLeft, startTop int

	// add / remove event listener requires pointers, so the variable should be initialized in this way
	var drag, stopDrag, resizeMouseMove, stopResize js.Func

	moveResizersX := func() {
		e.resizerTopLeft.SetCX(e.x)
		e.resizerTopRight.SetCX(e.x + e.width)
		e.resizerBottomLeft.SetCX(e.x)
		e.resizerBottomRight.SetCX(e.x + e.width)

		e.resizerTopMiddle.SetCX(e.x + e.width/2)
		e.resizerBottomMiddle.SetCX(e.x + e.width/2)
		e.resizerLeftMiddle.SetCX(e.x)
		e.resizerRightMiddle.SetCX(e.x + e.width)
	}

	moveResizersY := func() {
		e.resizerTopLeft.SetCY(e.y)
		e.resizerTopRight.SetCY(e.y)
		e.resizerBottomLeft.SetCY(e.y + e.height)
		e.resizerBottomRight.SetCY(e.y + e.height)

		e.resizerTopMiddle.SetCY(e.y)
		e.resizerBottomMiddle.SetCY(e.y + e.height)
		e.resizerLeftMiddle.SetCY(e.y + e.height/2)
		e.resizerRightMiddle.SetCY(e.y + e.height/2)
	}

	// Calculates the X position of the drag
	dragX := func(args []js.Value) {

		dx, dy := e.block.GetPointerPosition(args, e.main)

		dx -= startX
		dy -= startY

		dx, dy = e.adjustXYToGrid(dx, dy)

		newLeft := e.min(e.max(0, startLeft+dx), e.ideStage.GetClientWidth()-e.block.GetOffsetWidth())
		e.x = rulesDensity.Convert(newLeft)
		e.block.X(newLeft)
		e.selectDiv.X(newLeft)

		moveResizersX()
		return
	}

	// Calculates the Y position of the drag
	dragY := func(args []js.Value) {

		dx, dy := e.block.GetPointerPosition(args, e.main)

		dx -= startX
		dy -= startY

		dx, dy = e.adjustXYToGrid(dx, dy)

		newTop := e.min(e.max(0, startTop+dy), e.ideStage.GetClientHeight()-e.block.GetOffsetHeight())
		e.y = rulesDensity.Convert(newTop)
		e.block.Y(newTop)
		e.selectDiv.Y(newTop)

		moveResizersY()
		return
	}

	// Joins the calculations of X and Y of the drag
	drag = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if !isDragging {
			return nil
		}

		e.block.AddStyle("cursor", "grabbing")

		dragX(args)
		dragY(args)

		_ = e.ornament.Update(e.x, e.y, e.width, e.height)
		return nil
	})

	var pFunc func()
	// Removes events when the drag ends
	stopDrag = js.FuncOf(func(this js.Value, args []js.Value) interface{} { // feito
		pFunc()
		return nil
	})
	pFunc = func() {
		isDragging = false
		e.block.AddStyle("cursor", "grab")

		js.Global().Call("removeEventListener", "mousemove", drag)
		js.Global().Call("removeEventListener", "touchmove", drag, false)

		js.Global().Call("removeEventListener", "mouseup", stopDrag)
		js.Global().Call("removeEventListener", "touchend", stopDrag, false)
		js.Global().Call("removeEventListener", "touchcancel", stopDrag, false)
	}

	// Adds the device drag event when the mouse pointer is pressed
	dragFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if !e.dragEnabled {
			return nil
		}

		startX, startY = e.block.GetPointerPosition(args, e.main)

		isDragging = true
		startLeft = e.block.GetOffsetLeft()
		startTop = e.block.GetOffsetTop()

		// The movement of the mouse must be captured from the document and not the dragged element, or when the mouse moves
		// very fast, the drag to
		js.Global().Call("addEventListener", "mousemove", drag, map[string]any{"passive": true})
		js.Global().Call("addEventListener", "touchmove", drag, map[string]any{"passive": true})

		js.Global().Call("addEventListener", "mouseup", stopDrag, map[string]any{"passive": true})
		js.Global().Call("addEventListener", "touchend", stopDrag, map[string]any{"passive": true})
		js.Global().Call("addEventListener", "touchcancel", stopDrag, map[string]any{"passive": true})

		return nil
	})
	e.block.Get().Call("addEventListener", "mousedown", dragFunc, map[string]any{"passive": true})
	e.block.Get().Call("addEventListener", "touchstart", dragFunc, map[string]any{"passive": true})

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

	resizeHorizontal := func(args []js.Value, name string) {
		dx, dy := e.block.GetPointerPosition(args, e.main)

		dx -= startX
		dy -= startY

		dx, dy = e.adjustXYToGrid(dx, dy)

		newLeft := startLeft
		newWidth := startWidth

		if name == "bottom-right" {
			newWidth = e.min(startWidth+dx, e.ideStage.GetClientWidth()-startLeft)
		} else if name == "bottom-left" {
			newWidth = e.min(startWidth-dx, startLeft+startWidth)
			newLeft = e.max(0, startLeft+dx)
		} else if name == "top-right" {
			newWidth = e.min(startWidth+dx, e.ideStage.GetClientWidth()-startLeft)
		} else if name == "top-left" {
			newWidth = e.min(startWidth-dx, startLeft+startWidth)
			newLeft = e.max(0, startLeft+dx)
		} else if name == "top-middle" {
			return
		} else if name == "bottom-middle" {
			return
		} else if name == "left-middle" {
			newWidth = e.min(startWidth-dx, startLeft+startWidth)
			newLeft = e.max(0, startLeft+dx)
		} else if name == "right-middle" {
			newWidth = e.min(startWidth+dx, e.ideStage.GetClientWidth()-startLeft)
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
		if newWidth < e.blockMinimumWidth.GetInt() {
			return
		}

		newWidth = e.max(e.blockMinimumWidth.GetInt(), newWidth)

		//newLeft = int(float64(newLeft) / rulesDensity.GetDensity())
		//newWidth = int(float64(newWidth) / rulesDensity.GetDensity())

		e.x = rulesDensity.Convert(newLeft)
		e.width = rulesDensity.Convert(newWidth)

		e.block.X(e.x.GetInt())
		e.block.Width(e.width.GetInt())
		e.selectDiv.X(e.x.GetInt())
		e.selectDiv.Width(e.width.GetInt())

		moveResizersX()
	}

	resizeVertical := func(args []js.Value, name string) {
		dx, dy := e.block.GetPointerPosition(args, e.main)

		dx -= startX
		dy -= startY

		dx, dy = e.adjustXYToGrid(dx, dy)

		newTop := startTop
		newHeight := startHeight

		if name == "bottom-right" {
			newHeight = e.min(startHeight+dy, e.ideStage.GetClientHeight()-startTop)
		} else if name == "bottom-left" {
			newHeight = e.min(startHeight+dy, e.ideStage.GetClientHeight()-newTop)
		} else if name == "top-right" {
			newHeight = e.min(startHeight-dy, startTop+startHeight)
			newTop = e.max(0, startTop+dy)
		} else if name == "top-left" {
			newHeight = e.min(startHeight-dy, startTop+startHeight)
			newTop = e.max(0, startTop+dy)
		} else if name == "top-middle" {
			newHeight = e.min(startHeight-dy, startTop+startHeight)
			newTop = e.max(0, startTop+dy)
		} else if name == "bottom-middle" {
			newHeight = e.min(startHeight+dy, e.ideStage.GetClientHeight()-newTop)
		} else if name == "left-middle" {
			return
		} else if name == "right-middle" {
			return
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
		if newHeight < e.blockMinimumHeight.GetInt() {
			return
		}

		newHeight = e.max(e.blockMinimumHeight.GetInt(), newHeight)

		e.y = rulesDensity.Convert(newTop)
		e.height = rulesDensity.Convert(newHeight)

		e.block.Y(e.y.GetInt())
		e.block.Height(e.height.GetInt())
		e.selectDiv.Y(e.y.GetInt())
		e.selectDiv.Height(e.height.GetInt())

		moveResizersY()
	}

	resizeMouseMove = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if !isResizing {
			return nil
		}

		resizerName := e.block.Get().Get("dataset").Get("resizeName").String()
		resizeHorizontal(args, resizerName)
		resizeVertical(args, resizerName)
		_ = e.ornament.Update(e.x, e.y, e.width, e.height)

		width := rulesDensity.Convert(e.block.GetOffsetWidth())
		height := rulesDensity.Convert(e.block.GetOffsetHeight())
		e.OnResize(args, width, height)

		return nil
	})

	stopResize = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		isResizing = false
		js.Global().Call("removeEventListener", "mousemove", resizeMouseMove)
		js.Global().Call("removeEventListener", "mouseup", stopResize)

		js.Global().Call("removeEventListener", "touchmove", resizeMouseMove, false)
		js.Global().Call("removeEventListener", "touchend", stopResize, false)
		js.Global().Call("removeEventListener", "touchcancel", stopResize, false)
		return nil
	})

	resizeFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if !e.resizeEnabled {
			return nil
		}

		resizerName := this.Get("dataset").Get("name").String()
		e.block.DataKey("resizeName", resizerName)

		isResizing = true

		startX, startY = e.block.GetPointerPosition(args, e.main)

		startWidth = e.block.GetOffsetWidth()
		startHeight = e.block.GetOffsetHeight()
		startLeft = e.block.GetOffsetLeft()
		startTop = e.block.GetOffsetTop()

		js.Global().Call("addEventListener", "mousemove", resizeMouseMove, map[string]any{"passive": true})
		js.Global().Call("addEventListener", "mouseup", stopResize, map[string]any{"passive": true})

		js.Global().Call("addEventListener", "touchmove", resizeMouseMove, map[string]any{"passive": true})
		js.Global().Call("addEventListener", "touchend", stopResize, map[string]any{"passive": true})
		js.Global().Call("addEventListener", "touchcancel", stopResize, map[string]any{"passive": true})
		return nil
	})

	e.resizerTopLeft.GetSvg().Get().Call("addEventListener", "mousedown", resizeFunc, map[string]any{"passive": true})
	e.resizerTopRight.GetSvg().Get().Call("addEventListener", "mousedown", resizeFunc, map[string]any{"passive": true})
	e.resizerBottomLeft.GetSvg().Get().Call("addEventListener", "mousedown", resizeFunc, map[string]any{"passive": true})
	e.resizerBottomRight.GetSvg().Get().Call("addEventListener", "mousedown", resizeFunc, map[string]any{"passive": true})

	e.resizerTopMiddle.GetSvg().Get().Call("addEventListener", "mousedown", resizeFunc, map[string]any{"passive": true})
	e.resizerBottomMiddle.GetSvg().Get().Call("addEventListener", "mousedown", resizeFunc, map[string]any{"passive": true})
	e.resizerLeftMiddle.GetSvg().Get().Call("addEventListener", "mousedown", resizeFunc, map[string]any{"passive": true})
	e.resizerRightMiddle.GetSvg().Get().Call("addEventListener", "mousedown", resizeFunc, map[string]any{"passive": true})

	e.resizerTopLeft.GetSvg().Get().Call("addEventListener", "touchstart", resizeFunc, map[string]any{"passive": true})
	e.resizerTopRight.GetSvg().Get().Call("addEventListener", "touchstart", resizeFunc, map[string]any{"passive": true})
	e.resizerBottomLeft.GetSvg().Get().Call("addEventListener", "touchstart", resizeFunc, map[string]any{"passive": true})
	e.resizerBottomRight.GetSvg().Get().Call("addEventListener", "touchstart", resizeFunc, map[string]any{"passive": true})

	e.resizerTopMiddle.GetSvg().Get().Call("addEventListener", "touchstart", resizeFunc, map[string]any{"passive": true})
	e.resizerBottomMiddle.GetSvg().Get().Call("addEventListener", "touchstart", resizeFunc, map[string]any{"passive": true})
	e.resizerLeftMiddle.GetSvg().Get().Call("addEventListener", "touchstart", resizeFunc, map[string]any{"passive": true})
	e.resizerRightMiddle.GetSvg().Get().Call("addEventListener", "touchstart", resizeFunc, map[string]any{"passive": true})
}

// max Returns the maximum value
func (e *Block) maxD(a, b rulesDensity.Density) (max rulesDensity.Density) {
	if a > b {
		return a
	}
	return b
}

// max Returns the maximum value
func (e *Block) max(a, b int) (max int) {
	if a > b {
		return a
	}
	return b
}

// min Returns the minimum value
func (e *Block) min(a, b int) (min int) {
	if a < b {
		return a
	}
	return b
}

// min Returns the minimum value
func (e *Block) minD(a, b rulesDensity.Density) (min rulesDensity.Density) {
	if a < b {
		return a
	}
	return b
}

// OnResize cannot be shadowed by the main instance, so the function in SetOnResize
func (e *Block) OnResize(args []js.Value, width, height rulesDensity.Density) {
	if e.onResizeFunc != nil {
		e.onResizeFunc(args, width, height)
	}
}

// resizeEnabledDraw Show/hide the resizing blocks on the screen
func (e *Block) resizeEnabledDraw() {
	if !e.initialized {
		return
	}

	e.resizerTopLeft.SetVisible(e.resizeEnabled)
	e.resizerTopRight.SetVisible(e.resizeEnabled)
	e.resizerBottomLeft.SetVisible(e.resizeEnabled)
	e.resizerBottomRight.SetVisible(e.resizeEnabled)

	e.resizerTopMiddle.SetVisible(e.resizeEnabled)
	e.resizerBottomMiddle.SetVisible(e.resizeEnabled)
	e.resizerLeftMiddle.SetVisible(e.resizeEnabled)
	e.resizerRightMiddle.SetVisible(e.resizeEnabled)
}

// SetFatherId Receives the div ID used as a stage for the IDE and puts it to occupy the entire browser area
func (e *Block) SetFatherId(fatherId string) {
	e.ideStage = factoryBrowser.NewTagSvg().
		Import(fatherId)
}

// SetHeight Defines the height property of the device
func (e *Block) SetHeight(height rulesDensity.Density) {
	if !e.initialized {
		e.height = height
		return
	}

	e.block.Height(height)
}

// SetID Define the device's div ID
func (e *Block) SetID(id string) (err error) {
	e.id, err = utils.VerifyUniqueId(id)
	return
}

// SetMinimumHeight Defines the minimum height of the device
func (e *Block) SetMinimumHeight(height rulesDensity.Density) {
	e.blockMinimumHeight = height
}

// SetMinimumWidth Defines the minimum width of the device
func (e *Block) SetMinimumWidth(width rulesDensity.Density) {
	e.blockMinimumWidth = width
}

// SetName Defines a unique name for the device [compulsory]
func (e *Block) SetName(name string) {
	e.name = rulesSequentialId.GetIdFromBase(name)
	return
}

// SetOnResize Receives the pointer to a function to be invoked during resizing
//
//	This function is due to the fact that the OnResize function cannot be shadowed by the main instance
func (e *Block) SetOnResize(f func(args []js.Value, width, height rulesDensity.Density)) {
	e.onResizeFunc = f
}

// SetOrnament Sets the ornament draw object
//
//	ornament draw object is the instance in charge of making the SVG design of the device
func (e *Block) SetOrnament(ornament ornament.Draw) {
	e.ornament = ornament
}

// SetPosition Defines the coordinates (x, y) of the device
func (e *Block) SetPosition(x, y rulesDensity.Density) {
	e.SetX(x)
	e.SetY(y)
}

// SetSize Defines the height and width of the device
func (e *Block) SetSize(width, height rulesDensity.Density) {
	e.SetWidth(width)
	e.SetHeight(height)
}

// SetWidth Defines the width property of the device
func (e *Block) SetWidth(width rulesDensity.Density) {
	if !e.initialized {
		e.width = width
		return
	}

	e.block.Width(width)
}

// SetX Define a coordenada x da tela do navegador
func (e *Block) SetX(x rulesDensity.Density) {
	if !e.initialized {
		e.x = x
		return
	}

	e.block.X(x)
}

// SetY Set the coordinate X of the browser screen
func (e *Block) SetY(y rulesDensity.Density) {
	if !e.initialized {
		e.y = y
		return
	}

	e.block.Y(y)
}

func (e *Block) getMenuLabel(condition bool, labelTrue, labelFalse string) (label string) {
	if condition {
		return labelTrue
	}

	return labelFalse
}

func (e *Block) GetMenuDebug() (options []components.MenuOptions) {
	// mover para o topo
	// mover para cima
	// mover para baixo
	// mover para o fim
	options = []components.MenuOptions{
		{
			Label: "Debug",
			Submenu: []components.MenuOptions{
				{
					Label: e.getMenuLabel(e.GetSelected(), "Unselect", "Select"),
					Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} {
						e.SetSelected(!e.GetSelected())
						return nil
					}),
				},
				{
					Label: e.getMenuLabel(e.GetSelectBlocked(), "Select lock disable", "Select lock enable"),
					Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} {
						e.SetSelectBlocked(!e.GetSelectBlocked())
						return nil
					}),
				},
				{
					Label: e.getMenuLabel(e.GetResize(), "Resize disable", "Resize enable"),
					Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} {
						e.SetResize(!e.GetResize())
						return nil
					}),
				},
				{
					Label: e.getMenuLabel(e.GetResizeBlocked(), "Resize lock disable", "Resize lock enable"),
					Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} {
						e.SetResizeBlocked(!e.GetResizeBlocked())
						return nil
					}),
				},
				{
					Label: e.getMenuLabel(e.GetDrag(), "Drag disable", "Drag enable"),
					Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} {
						e.SetDrag(!e.GetDrag())
						return nil
					}),
				},
				{
					Label: e.getMenuLabel(e.GetDragBlocked(), "Drag lock disable", "Drag lock enable"),
					Action: js.FuncOf(func(this js.Value, args []js.Value) interface{} {
						e.SetDragBlocked(!e.GetDragBlocked())
						return nil
					}),
				},
			},
		},
	}

	return
}
