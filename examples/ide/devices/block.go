package devices

import (
	"fmt"
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/examples/ide/ornament"
	"github.com/helmutkemper/webassembly/examples/ide/utils"
	"github.com/helmutkemper/webassembly/platform/factoryColor"
	"image/color"
	"strings"
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

	blockHorizontalMinimumSize int
	blockVerticalMinimumSize   int

	classListName string

	isResizing      bool
	resizeEnabled   bool
	resizerFlashing bool
	selectFlashing  bool
	selected        bool
	dragEnabled     bool

	resizerColor      color.RGBA
	resizerFlashColor color.RGBA

	container *html.TagDiv
	block     *html.TagDiv

	resizerTopLeft     *html.TagDiv
	resizerTopRight    *html.TagDiv
	resizerBottomLeft  *html.TagDiv
	resizerBottomRight *html.TagDiv

	selectDiv *html.TagDiv

	resizers []*html.TagDiv

	ornament ornament.Draw
}

func (e *Block) SetFatherId(fatherId string) {
	e.fatherId = fatherId

	e.container = factoryBrowser.NewTagDiv().
		AddStyle("position", "relative"). // todo: mudar e documentar
		AddStyle("width", "100vw").
		AddStyle("height", "100vh")

	e.container.AppendById(fatherId)
}

// SetOrnament Sets the ornament draw object
func (e *Block) SetOrnament(ornament ornament.Draw) {
	e.ornament = ornament
}

func (e *Block) SetID(id string) (err error) {
	e.id, err = utils.VerifyUniqueId(id)
	return
}

func (e *Block) GetID() (id string) {
	return e.id
}

func (e *Block) SetX(x int) {
	if !e.initialized {
		e.x = x
		return
	}
	e.block.AddStyle("left", fmt.Sprintf("%dpx", x))
}

func (e *Block) GetX() (x int) {
	return e.x
}

func (e *Block) SetY(y int) {
	if !e.initialized {
		e.y = y
		return
	}

	e.block.AddStyle("top", fmt.Sprintf("%dpx", y))
}

func (e *Block) GetY() (y int) {
	return e.y
}

func (e *Block) SetWidth(width int) {
	if !e.initialized {
		e.width = width
		return
	}

	e.block.AddStyle("width", fmt.Sprintf("%dpx", width))
}

func (e *Block) SetHeight(height int) {
	if !e.initialized {
		e.height = height
		return
	}

	e.block.AddStyle("height", fmt.Sprintf("%dpx", height))
}

//func (e *Block) SetFatherID(fatherID string) {
//	e.fatherId = fatherID
//}

func (e *Block) SetName(name string) (err error) {
	e.name, err = utils.VerifyName(name)
	return
}

func (e *Block) GetName() (name string) {
	return e.name
}

func (e *Block) SetPosition(x, y int) {
	e.SetX(x)
	e.SetY(y)
}

func (e *Block) SetSize(width, height int) {
	e.SetWidth(width)
	e.SetHeight(height)
}

func (e *Block) SetHorizontalMinimumSize(size int) {
	e.blockHorizontalMinimumSize = size
}

func (e *Block) SetVerticalMinimumSize(size int) {
	e.blockVerticalMinimumSize = size
}

func (e *Block) Init() (err error) {
	var base string
	if base, err = e.SequentialId.GetId(e.name); err != nil {
		return
	}
	e.id, err = utils.VerifyUniqueId(base)

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

	e.dragEnabledSupport()
	e.resizeEnabledSupport()
	//e.SetSelectEnabled(e.selected) // todo: descomentar

	return
}

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
	e.container.Append(e.block)

	e.selectDiv = factoryBrowser.NewTagDiv().
		AddStyle("position", "absolute").
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

func (e *Block) updateOrnament() (err error) {
	width := e.block.GetOffsetWidth()
	height := e.block.GetOffsetHeight()
	_ = e.ornament.Update(width, height)
	return
}

func (e *Block) min(a, b int) (min int) {
	if a < b {
		return a
	}
	return b
}

func (e *Block) max(a, b int) (max int) {
	if a > b {
		return a
	}
	return b
}

func (e *Block) SetSelectEnabled(enabled bool) {
	e.selected = enabled
	e.selectDiv.AddStyleConditional(enabled, "display", "block", "none")

	if enabled && e.resizeEnabled {
		e.SetResizeEnabled(false)
	}
}

func (e *Block) SetResizeEnabled(enabled bool) {
	e.resizeEnabled = enabled
	e.resizeEnabledSupport()
	if enabled && e.selected {
		e.SetSelectEnabled(false)
	}
}

func (e *Block) resizeEnabledSupport() {
	if !e.initialized {
		return
	}

	e.resizerTopLeft.AddStyleConditional(e.resizeEnabled, "display", "block", "none")
	e.resizerTopRight.AddStyleConditional(e.resizeEnabled, "display", "block", "none")
	e.resizerBottomLeft.AddStyleConditional(e.resizeEnabled, "display", "block", "none")
	e.resizerBottomRight.AddStyleConditional(e.resizeEnabled, "display", "block", "none")
}

func (e *Block) SetDragEnabled(enabled bool) {
	e.dragEnabled = enabled
	e.dragEnabledSupport()

	if enabled && e.dragEnabled {
		e.SetResizeEnabled(true)
	}

	if enabled && e.selected {
		e.SetSelectEnabled(false)
	}
}

func (e *Block) dragEnabledSupport() {
	if !e.initialized {
		return
	}

	e.block.AddStyleConditional(e.dragEnabled, "cursor", "grab", "default")
}

func (e *Block) initEvents() {
	var isDragging, isResizing bool
	var startX, startY, startWidth, startHeight, startLeft, startTop int

	var resizeMouseMove, stopResize js.Func
	var drag, stopDrag js.Func

	dragX := func(element js.Value) {
		dx := element.Get("screenX").Int() - startX
		newLeft := e.min(e.max(0, startLeft+dx), e.container.GetClientWidth()-e.block.GetOffsetWidth())
		e.block.AddStyle("left", fmt.Sprintf("%dpx", newLeft))
	}

	dragY := func(element js.Value) {
		dy := element.Get("screenY").Int() - startY
		newTop := e.min(e.max(0, startTop+dy), e.container.GetClientHeight()-e.block.GetOffsetHeight())
		e.block.AddStyle("top", fmt.Sprintf("%dpx", newTop))
	}

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

	stopDrag = js.FuncOf(func(this js.Value, args []js.Value) interface{} { // feito
		isDragging = false
		e.block.AddStyle("cursor", "grab")

		js.Global().Call("removeEventListener", "mousemove", drag)
		js.Global().Call("removeEventListener", "mouseup", stopDrag)
		return nil
	})

	// drag main block
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
		js.Global().Call("addEventListener", "mousemove", drag)
		js.Global().Call("addEventListener", "mouseup", stopDrag)

		return nil
	}))

	resizeHorizontal := func(element js.Value, name string) {
		/*
		   bug:
		   [tl]--------------[tr]
		     |                |
		     |                |
		     |                |
		   [bl]--------------[br]

		   If I drag TR or BR left, and the size is below minimum, the block is dragged left.
		*/

		dx := element.Get("screenX").Int() - startX
		newLeft := startLeft
		newWidth := startWidth

		if name == "bottom-right" {
			newWidth = e.min(startWidth+dx, e.container.Get().Get("clientWidth").Int()-startLeft)
		} else if name == "bottom-left" {
			newWidth = e.min(startWidth-dx, startLeft+startWidth)
			newLeft = e.max(0, startLeft+dx)
		} else if name == "top-right" {
			newWidth = e.min(startWidth+dx, e.container.Get().Get("clientWidth").Int()-startLeft)
		} else if name == "top-left" {
			newWidth = e.min(startWidth-dx, startLeft+startWidth)
			newLeft = e.max(0, startLeft+dx)
		}

		e.block.AddStyle("left", fmt.Sprintf("%dpx", newLeft))
		e.block.AddStyle("width", fmt.Sprintf("%dpx", e.max(e.blockHorizontalMinimumSize, newWidth)))
		e.selectDiv.AddStyle("width", fmt.Sprintf("%dpx", e.max(e.blockHorizontalMinimumSize, newWidth)))
	}

	resizeVertical := func(element js.Value, name string) {
		/*
		   bug:
		   [tl]--------------[tr]
		     |                |
		     |                |
		     |                |
		   [bl]--------------[br]

		   If I drag TL or TR down, and the size is below minimum, the block is dragged down.
		*/

		dy := element.Get("screenY").Int() - startY
		newTop := startTop
		newHeight := startHeight

		if name == "bottom-right" {
			newHeight = e.min(startHeight+dy, e.container.Get().Get("clientHeight").Int()-startTop)
		} else if name == "bottom-left" {
			newHeight = e.min(startHeight+dy, e.container.Get().Get("clientHeight").Int()-newTop)
		} else if name == "top-right" {
			newHeight = e.min(startHeight-dy, startTop+startHeight)
			newTop = e.max(0, startTop+dy)
		} else if name == "top-left" {
			newHeight = e.min(startHeight-dy, startTop+startHeight)
			newTop = e.max(0, startTop+dy)
		}

		e.block.AddStyle("top", fmt.Sprintf("%dpx", newTop))
		e.block.AddStyle("height", fmt.Sprintf("%dpx", e.max(e.blockVerticalMinimumSize, newHeight)))
		e.selectDiv.AddStyle("height", fmt.Sprintf("%dpx", e.max(e.blockVerticalMinimumSize, newHeight)))
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
		e.OnResize(element, width, height) // todo: virar ponteiro de func

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

func (e *Block) OnResize(element js.Value, width, height int) {}
