package block

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/examples/ide/ornament"
	"github.com/helmutkemper/webassembly/examples/ide/rulesSequentialId"
	"github.com/helmutkemper/webassembly/examples/ide/utils"
	"github.com/helmutkemper/webassembly/platform/components"
	"github.com/helmutkemper/webassembly/platform/factoryColor"
	"image/color"
	"strings"
	"syscall/js"
)

type Block struct {
	id     string
	autoId string
	name   string

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

	ideStage *html.TagSvg
	block    *html.TagSvg

	resizerTopLeft     *html.TagSvgRect
	resizerTopRight    *html.TagSvgRect
	resizerBottomLeft  *html.TagSvgRect
	resizerBottomRight *html.TagSvgRect

	resizerTopMiddle    *html.TagSvgRect
	resizerBottomMiddle *html.TagSvgRect
	resizerLeftMiddle   *html.TagSvgRect
	resizerRightMiddle  *html.TagSvgRect

	selectDiv *html.TagSvgRect

	ornament ornament.Draw

	onResizeFunc func(element js.Value, width, height int)
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
func (e *Block) createBlock(x, y, width, height int) {
	e.block = factoryBrowser.NewTagSvg().
		Id(e.id).
		//Class(e.classListName).
		//AddStyle("position", "absolute").
		//AddStyle("touchAction", "none").
		X(x).
		Y(y).
		Width(width).
		Height(height)
	e.ideStage.Append(e.block)

	e.selectDiv = factoryBrowser.NewTagSvgRect().
		X(x).
		Y(y).
		Width(width).
		Height(height).
		Fill("none").Stroke("red").
		StrokeDasharray([]int{4}).
		StrokeWidth(1).
		AddStyle("display", "none")
	e.ideStage.Append(e.selectDiv)

	e.resizerTopLeft = factoryBrowser.NewTagSvgRect().
		DataKey("name", "top-left").
		X(x + e.resizerDistance).
		Y(y + e.resizerDistance).
		Width(e.resizerWidth).
		Height(e.resizerHeight).
		Fill(e.resizerColor).
		Cursor("nwse-resize")
	e.ideStage.Append(e.resizerTopLeft)

	e.resizerTopRight = factoryBrowser.NewTagSvgRect().
		DataKey("name", "top-right").
		X(x + width + e.resizerDistance).
		Y(y + e.resizerDistance).
		Width(e.resizerWidth).
		Height(e.resizerHeight).
		Fill(e.resizerColor).
		Cursor("nesw-resize")
	//AddStyle("border-radius", fmt.Sprintf("%dpx", e.resizerRadius)).
	e.ideStage.Append(e.resizerTopRight)

	e.resizerBottomLeft = factoryBrowser.NewTagSvgRect().
		DataKey("name", "bottom-left").
		X(x + e.resizerDistance).
		Y(y + height + e.resizerDistance).
		Width(e.resizerWidth).
		Height(e.resizerHeight).
		Fill(e.resizerColor).
		Cursor("nesw-resize")
	//AddStyle("border-radius", fmt.Sprintf("%dpx", e.resizerRadius)).
	e.ideStage.Append(e.resizerBottomLeft)

	e.resizerBottomRight = factoryBrowser.NewTagSvgRect().
		DataKey("name", "bottom-right").
		X(x + width + e.resizerDistance).
		Y(y + height + e.resizerDistance).
		Width(e.resizerWidth).
		Height(e.resizerHeight).
		Fill(e.resizerColor).
		Cursor("nwse-resize")
	//AddStyle("position", "absolute").
	//AddStyle("background-color", e.resizerColor).
	//AddStyle("border-radius", fmt.Sprintf("%dpx", e.resizerRadius)).
	//AddStyle("cursor", "nwse-resize")
	e.ideStage.Append(e.resizerBottomRight)

	//----------------------------------------------------

	e.resizerTopMiddle = factoryBrowser.NewTagSvgRect().
		DataKey("name", "top-middle").
		X(x + width/2 + e.resizerDistance).
		Y(y + e.resizerDistance).
		Width(e.resizerWidth).
		Height(e.resizerHeight).
		Fill(e.resizerColor).
		Cursor("ns-resize")
	e.ideStage.Append(e.resizerTopMiddle)

	e.resizerBottomMiddle = factoryBrowser.NewTagSvgRect().
		DataKey("name", "bottom-middle").
		X(x + width/2 + e.resizerDistance).
		Y(y + height + e.resizerDistance).
		Width(e.resizerWidth).
		Height(e.resizerHeight).
		Fill(e.resizerColor).
		Cursor("ns-resize")
	e.ideStage.Append(e.resizerBottomMiddle)

	e.resizerLeftMiddle = factoryBrowser.NewTagSvgRect().
		DataKey("name", "left-middle").
		X(x + e.resizerDistance).
		Y(y + height/2 + e.resizerDistance).
		Width(e.resizerWidth).
		Height(e.resizerHeight).
		Fill(e.resizerColor).
		Cursor("ew-resize")
	e.ideStage.Append(e.resizerLeftMiddle)

	e.resizerRightMiddle = factoryBrowser.NewTagSvgRect().
		DataKey("name", "right-middle").
		X(x + width + e.resizerDistance).
		Y(y + height/2 + e.resizerDistance).
		Width(e.resizerWidth).
		Height(e.resizerHeight).
		Fill(e.resizerColor).
		Cursor("ew-resize")
	e.ideStage.Append(e.resizerRightMiddle)
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
func (e *Block) GetHeight() (height int) {
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
func (e *Block) GetWidth() (width int) {
	return e.width
}

// GetX Returns to coordinate X of the browser screen
func (e *Block) GetX() (x int) {
	return e.x
}

// GetY Returns to coordinate Y of the browser screen
func (e *Block) GetY() (y int) {
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

	e.block.X(e.x)
	e.block.Y(e.y)
	e.block.Width(e.width)
	e.block.Height(e.height)

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

	moveResizersX := func(x, width int) {
		e.resizerTopLeft.X(x + e.resizerDistance)
		e.resizerTopRight.X(x + width + e.resizerDistance)
		e.resizerBottomLeft.X(x + e.resizerDistance)
		e.resizerBottomRight.X(x + width + e.resizerDistance)

		e.resizerTopMiddle.X(x + width/2 + e.resizerDistance)
		e.resizerBottomMiddle.X(x + width/2 + e.resizerDistance)
		e.resizerLeftMiddle.X(x + e.resizerDistance)
		e.resizerRightMiddle.X(x + width + e.resizerDistance)
	}

	moveResizersY := func(y, height int) {
		e.resizerTopLeft.Y(y + e.resizerDistance)
		e.resizerTopRight.Y(y + e.resizerDistance)
		e.resizerBottomLeft.Y(y + height + e.resizerDistance)
		e.resizerBottomRight.Y(y + height + e.resizerDistance)

		e.resizerTopMiddle.Y(y + e.resizerDistance)
		e.resizerBottomMiddle.Y(y + height + e.resizerDistance)
		e.resizerLeftMiddle.Y(y + height/2 + e.resizerDistance)
		e.resizerRightMiddle.Y(y + height/2 + e.resizerDistance)
	}

	// Calculates the X position of the drag
	dragX := func(element js.Value) {
		dx := element.Get("screenX").Int() - startX
		newLeft := e.min(e.max(0, startLeft+dx), e.ideStage.GetClientWidth()-e.block.GetOffsetWidth())
		e.x = newLeft
		e.block.X(newLeft)
		e.selectDiv.X(newLeft)

		moveResizersX(newLeft, e.width)
	}

	// Calculates the Y position of the drag
	dragY := func(element js.Value) {
		dy := element.Get("screenY").Int() - startY
		newTop := e.min(e.max(0, startTop+dy), e.ideStage.GetClientHeight()-e.block.GetOffsetHeight())
		e.y = newTop
		e.block.Y(newTop)
		e.selectDiv.Y(newTop)

		moveResizersY(newTop, e.height)
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

		_ = e.ornament.Update(e.x, e.y, e.width, e.height)
		return nil
	})

	// Removes events when the drag ends
	stopDrag = js.FuncOf(func(this js.Value, args []js.Value) interface{} { // feito
		isDragging = false
		e.block.AddStyle("cursor", "grab")

		js.Global().Call("removeEventListener", "mousemove", drag)
		js.Global().Call("removeEventListener", "mouseup", stopDrag)
		return nil
	})

	// Adds the device drag event when the mouse pointer is pressed
	dragFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
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
		js.Global().Call("addEventListener", "touchmove", drag, false)

		js.Global().Call("addEventListener", "mouseup", stopDrag)
		js.Global().Call("addEventListener", "touchend", stopDrag, false)

		return nil
	})
	e.block.Get().Call("addEventListener", "mousedown", dragFunc)
	e.block.Get().Call("addEventListener", "touchstart", dragFunc, false)

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
		if newWidth < e.blockMinimumWidth {
			return
		}

		newWidth = e.max(e.blockMinimumWidth, newWidth)

		e.x = newLeft
		e.width = newWidth

		e.block.X(newLeft)
		e.block.Width(newWidth)
		e.selectDiv.X(newLeft)
		e.selectDiv.Width(newWidth)

		moveResizersX(newLeft, newWidth)
	}

	resizeVertical := func(element js.Value, name string) {
		dy := element.Get("screenY").Int() - startY
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
		if newHeight < e.blockMinimumHeight {
			return
		}

		newHeight = e.max(e.blockMinimumHeight, newHeight)

		e.y = newTop
		e.height = newHeight

		e.block.Y(newTop)
		e.block.Height(newHeight)
		e.selectDiv.Y(newTop)
		e.selectDiv.Height(newHeight)

		moveResizersY(newTop, newHeight)
	}

	resizeMouseMove = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if !isResizing {
			return nil
		}

		element := args[0]
		resizerName := e.block.Get().Get("dataset").Get("resizeName").String()
		resizeHorizontal(element, resizerName)
		resizeVertical(element, resizerName)
		_ = e.ornament.Update(e.x, e.y, e.width, e.height)

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

	e.resizerTopMiddle.Get().Call("addEventListener", "mousedown", resizeFunc)
	e.resizerBottomMiddle.Get().Call("addEventListener", "mousedown", resizeFunc)
	e.resizerLeftMiddle.Get().Call("addEventListener", "mousedown", resizeFunc)
	e.resizerRightMiddle.Get().Call("addEventListener", "mousedown", resizeFunc)
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

// OnResize cannot be shadowed by the main instance, so the function in SetOnResize
func (e *Block) OnResize(element js.Value, width, height int) {
	if e.onResizeFunc != nil {
		e.onResizeFunc(element, width, height)
	}
}

// resizeEnabledDraw Show/hide the resizing blocks on the screen
func (e *Block) resizeEnabledDraw() {
	if !e.initialized {
		return
	}

	e.resizerTopLeft.AddStyleConditional(e.resizeEnabled, "display", "block", "none")
	e.resizerTopRight.AddStyleConditional(e.resizeEnabled, "display", "block", "none")
	e.resizerBottomLeft.AddStyleConditional(e.resizeEnabled, "display", "block", "none")
	e.resizerBottomRight.AddStyleConditional(e.resizeEnabled, "display", "block", "none")

	e.resizerTopMiddle.AddStyleConditional(e.resizeEnabled, "display", "block", "none")
	e.resizerBottomMiddle.AddStyleConditional(e.resizeEnabled, "display", "block", "none")
	e.resizerLeftMiddle.AddStyleConditional(e.resizeEnabled, "display", "block", "none")
	e.resizerRightMiddle.AddStyleConditional(e.resizeEnabled, "display", "block", "none")
}

// SetFatherId Receives the div ID used as a stage for the IDE and puts it to occupy the entire browser area
func (e *Block) SetFatherId(fatherId string) {
	e.ideStage = factoryBrowser.NewTagSvg().
		Import(fatherId) //.
	//AddStyle("position", "relative").
	//AddStyle("width", "100vw"). // todo: transformar isto em algo chamado uma única vez e em outro lugar
	//AddStyle("height", "100vh") // todo: transformar isto em algo chamado uma única vez e em outro lugar
}

// SetHeight Defines the height property of the device
func (e *Block) SetHeight(height int) {
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
func (e *Block) SetMinimumHeight(height int) {
	e.blockMinimumHeight = height
}

// SetMinimumWidth Defines the minimum width of the device
func (e *Block) SetMinimumWidth(width int) {
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
func (e *Block) SetOnResize(f func(element js.Value, width, height int)) {
	e.onResizeFunc = f
}

// SetOrnament Sets the ornament draw object
//
//	ornament draw object is the instance in charge of making the SVG design of the device
func (e *Block) SetOrnament(ornament ornament.Draw) {
	e.ornament = ornament
}

// SetPosition Defines the coordinates (x, y) of the device
func (e *Block) SetPosition(x, y int) {
	e.SetX(x)
	e.SetY(y)
}

// SetSize Defines the height and width of the device
func (e *Block) SetSize(width, height int) {
	e.SetWidth(width)
	e.SetHeight(height)
}

// SetWidth Defines the width property of the device
func (e *Block) SetWidth(width int) {
	if !e.initialized {
		e.width = width
		return
	}

	e.block.Width(width)
}

// SetX Define a coordenada x da tela do navegador
func (e *Block) SetX(x int) {
	if !e.initialized {
		e.x = x
		return
	}

	e.block.X(x)
}

// SetY Set the coordinate X of the browser screen
func (e *Block) SetY(y int) {
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
