package devices

import (
	"errors"
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

type DeviceOpAmp struct {
	sequentialId utils.SequentialId
	// block - start
	id     string
	autoId string

	resizerWidth      string
	resizerHeight     string
	resizerDistance   string
	resizerColor      color.RGBA
	resizerRadius     string
	classListName     string
	resizerFlashing   bool
	selectFlashing    bool
	selected          bool
	resizerFlashColor string
	initialized       bool
	draw              ornament.Draw
	resizers          []*html.TagDiv
	container         *html.TagDiv
	block             *html.TagDiv
	selectDiv         *html.TagDiv
	utils.SequentialId
	fatherId      string
	isResizing    bool
	resizeEnabled bool
	dragEnabled   bool
	ornament      ornament.Draw

	blockHorizontalMinimumSize int
	blockVerticalMinimumSize   int
	// block - end

	x                     int
	y                     int
	width                 int
	height                int
	defaultWidth          int
	defaultHeight         int
	horizontalMinimumSize int
	verticalMinimumSize   int

	name     string
	deviceId string
}

// SetHorizontalMinimumSize sets the minimum horizontal size of the block.
func (e *DeviceOpAmp) SetHorizontalMinimumSize(size int) {
	e.blockHorizontalMinimumSize = size
}

// SetVerticalMinimumSize sets the minimum vertical size of the block.
func (e *DeviceOpAmp) SetVerticalMinimumSize(size int) {
	e.blockVerticalMinimumSize = size
}

// SetPosition sets the position of the block.
func (e *DeviceOpAmp) SetPosition(x, y int) {
	e.SetX(x)
	e.SetY(y)
}

// SetSize sets the size of the block.
func (e *DeviceOpAmp) SetSize(width, height int) {
	e.SetWidth(width)
	e.SetHeight(height)
}

// SetZIndex sets the z-index of the block.
func (e *DeviceOpAmp) SetZIndex(zIndex int) {
	if !e.initialized {
		return
	}

	e.block.AddStyle("zIndex", fmt.Sprintf("%d", zIndex))
	for _, resizer := range e.resizers {
		resizer.AddStyle("zIndex", fmt.Sprintf("%d", zIndex))
	}
}

// SetResizerSize sets the width and height of the resizers.
func (e *DeviceOpAmp) SetResizerSize(width, height string) {
	e.resizerWidth = width
	e.resizerHeight = height

	if !e.initialized {
		return
	}

	for _, resizer := range e.resizers {
		resizer.AddStyle("width", width)
		resizer.AddStyle("height", height)
	}
}

// SetResizerDistance sets the distances between the resizers and the component.
func (e *DeviceOpAmp) SetResizerDistance(distance string) {
	e.resizerDistance = distance

	if !e.initialized {
		return
	}

	for _, resizer := range e.resizers {
		if strings.Contains(resizer.Get().Get("classList").String(), "top-left") {
			resizer.AddStyle("top", distance)
			resizer.AddStyle("left", distance)
		} else if strings.Contains(resizer.Get().Get("classList").String(), "top-right") {
			resizer.AddStyle("top", distance)
			resizer.AddStyle("right", distance)
		} else if strings.Contains(resizer.Get().Get("classList").String(), "bottom-left") {
			resizer.AddStyle("bottom", distance)
			resizer.AddStyle("left", distance)
		} else if strings.Contains(resizer.Get().Get("classList").String(), "bottom-right") {
			resizer.AddStyle("bottom", distance)
			resizer.AddStyle("right", distance)
		}
	}
}

func (e *DeviceOpAmp) SetX(x int) {
	e.x = x
	if !e.initialized {
		return
	}

	e.block.AddStyle("left", fmt.Sprintf("%dpx", x))
}

func (e *DeviceOpAmp) SetY(y int) {
	e.y = y
	if !e.initialized {
		return
	}

	e.block.AddStyle("top", fmt.Sprintf("%dpx", y))
}

func (e *DeviceOpAmp) SetWidth(width int) {
	e.width = width
	if !e.initialized {
		return
	}

	e.block.AddStyle("width", fmt.Sprintf("%dpx", width))
}

func (e *DeviceOpAmp) SetHeight(height int) {
	e.height = height
	if !e.initialized {
		return
	}

	e.block.AddStyle("height", fmt.Sprintf("%dpx", height))
}

func (e *DeviceOpAmp) SetName(name string) (err error) {
	e.name, err = utils.VerifyName(name)
	return
}

func (e *DeviceOpAmp) GetName() (name string) {
	return e.name
}

func (e *DeviceOpAmp) Init(id string, x, y int) (err error) {
	e.block = factoryBrowser.NewTagDiv()

	// block - start
	for {
		e.id, err = e.sequentialId.GetId(e.name)
		if errors.Is(err, errors.New("base is required")) {
			return
		}
		e.id, err = utils.VerifyUniqueId(e.id)
		if err == nil {
			break
		}
	}

	e.autoId = utils.GetRandomId()

	e.resizerWidth = "10px"
	e.resizerHeight = "10px"
	e.resizerDistance = "-5px"
	e.resizerColor = factoryColor.NewRed()
	e.resizerRadius = "2px"
	e.classListName = "block"
	e.resizerFlashColor = "yellow"
	// block - end

	e.x = x
	e.y = y
	e.deviceId = id

	e.defaultWidth = 50
	e.defaultHeight = 50
	e.horizontalMinimumSize = 100
	e.verticalMinimumSize = 100

	if e.width == 0 {
		e.width = e.defaultWidth
	}

	if e.height == 0 {
		e.height = e.defaultHeight
	}

	return
}

func (e *DeviceOpAmp) createBlock(x, y, width, height int) {
	e.resizers = make([]*html.TagDiv, 0)

	e.block = factoryBrowser.NewTagDiv().
		Id(e.id).
		Class(e.classListName).
		AddStyle("position", "absolute").
		AddStyle("top", "0px").
		AddStyle("left", "0px").
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

	resizerPositions := []string{"top-left", "top-right", "bottom-left", "bottom-right"}
	for _, pos := range resizerPositions {
		resizer := factoryBrowser.NewTagDiv().
			Class("resizer", pos).
			DataKey("name", pos).
			AddStyle("position", "absolute").
			AddStyle("width", fmt.Sprintf("%dpx", e.resizerWidth)).
			AddStyle("height", fmt.Sprintf("%dpx", e.resizerHeight)).
			AddStyle("background-color", html.RGBAToJs(e.resizerColor)).
			AddStyle("border-radius", fmt.Sprintf("%dpx", e.resizerRadius))
		e.resizers = append(e.resizers, resizer)

		resizer.Get().Call("addEventListener", "mousedown", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			if !e.resizeEnabled {
				return nil
			}

			args[0].Call("stopPropagation")
			e.isResizing = true

			return nil
		}))

		switch pos {
		case "top-left":
			resizer.AddStyle("top", fmt.Sprintf("%dpx", e.resizerDistance)).
				AddStyle("left", fmt.Sprintf("%dpx", e.resizerDistance)).
				AddStyle("cursor", "nwse-resize")

		case "top-right":
			resizer.AddStyle("top", fmt.Sprintf("%dpx", e.resizerDistance)).
				AddStyle("right", fmt.Sprintf("%dpx", e.resizerDistance)).
				AddStyle("cursor", "nesw-resize")

		case "bottom-left":
			resizer.AddStyle("bottom", fmt.Sprintf("%dpx", e.resizerDistance)).
				AddStyle("left", fmt.Sprintf("%dpx", e.resizerDistance)).
				AddStyle("cursor", "nesw-resize")

		case "bottom-right":
			resizer.AddStyle("bottom", fmt.Sprintf("%dpx", e.resizerDistance)).
				AddStyle("right", fmt.Sprintf("%dpx", e.resizerDistance)).
				AddStyle("cursor", "nwse-resize")
		}

		e.block.Append(resizer)
		e.updateOrnament()
	}
}

func (e *DeviceOpAmp) updateOrnament() (err error) {
	width := e.block.GetOffsetWidth()
	height := e.block.GetOffsetHeight()
	err = e.ornament.Update(width, height)
	return
}

func (e *DeviceOpAmp) min(a, b int) (min int) {
	if a < b {
		return a
	}
	return b
}

func (e *DeviceOpAmp) max(a, b int) (min int) {
	if a > b {
		return a
	}
	return b
}

func (e *DeviceOpAmp) initEvents() {
	var isDragging, isResizing bool
	var startX, startY, startWidth, startHeight, startLeft, startTop int
	var currentResizer js.Value

	dragX := func(element js.Value) {
		dx := element.Get("clientX").Int() - startX
		newLeft := e.min(0, startTop+dx)
		e.block.AddStyle("left", fmt.Sprintf("%dpx", newLeft))
	}

	dragY := func(element js.Value) {
		dy := element.Get("clientY").Int() - startY
		newTop := e.min(0, startTop+dy)
		e.block.AddStyle("top", fmt.Sprintf("%dpx", newTop))
	}

	drag := func(this js.Value, args []js.Value) interface{} {
		element := args[0]

		if !isDragging {
			return nil
		}

		e.block.AddStyle("cursor", "grabbing")
		dragX(element)
		dragY(element)
		return nil
	}

	resizeHorizontal := func(this js.Value, args []js.Value) interface{} {
		element := args[0]

		/*
		   bug:
		   [tl]--------------[tr]
		     |                |
		     |                |
		     |                |
		   [bl]--------------[br]

		   If I drag TR or BR left, and the size is below minimum, the block is dragged left.
		*/

		dx := element.Get("clientX").Int() - startX
		newLeft := startLeft
		newWidth := startWidth

		if strings.Contains(currentResizer.Get("classList").String(), "bottom-right") {
			newWidth = e.min(startWidth+dx, e.container.Get().Get("clientWidth").Int()-startLeft)
		} else if strings.Contains(currentResizer.Get("classList").String(), "bottom-left") {
			newWidth = e.min(startWidth-dx, startLeft+startWidth)
			newLeft = e.max(0, startLeft+dx)
		} else if strings.Contains(currentResizer.Get("classList").String(), "top-right") {
			newWidth = e.min(startWidth+dx, e.container.Get().Get("clientWidth").Int()-startLeft)
		} else if strings.Contains(currentResizer.Get("classList").String(), "top-left") {
			newWidth = e.min(startWidth-dx, startLeft+startWidth)
			newLeft = e.max(0, startLeft+dx)
		}

		e.block.AddStyle("left", fmt.Sprintf("%dpx", newLeft))
		e.block.AddStyle("width", fmt.Sprintf("%dpx", e.max(e.blockHorizontalMinimumSize, newWidth)))

		e.selectDiv.AddStyle("width", fmt.Sprintf("%dpx", e.max(e.blockHorizontalMinimumSize, newWidth)))
		return nil
	}

	resizeVertical := func(this js.Value, args []js.Value) interface{} {
		element := args[0]

		/*
		   bug:
		   [tl]--------------[tr]
		     |                |
		     |                |
		     |                |
		   [bl]--------------[br]

		   If I drag TL or TR down, and the size is below minimum, the block is dragged down.
		*/

		dy := element.Get("clientY").Int() - startY
		newTop := startTop
		newHeight := startHeight

		if strings.Contains(currentResizer.Get("classList").String(), "bottom-right") {
			newHeight = e.min(startHeight+dy, e.container.Get().Get("clientHeight").Int()-startTop)
		} else if strings.Contains(currentResizer.Get("classList").String(), "bottom-left") {
			newHeight = e.min(startHeight+dy, e.container.Get().Get("clientHeight").Int()-newTop)
		} else if strings.Contains(currentResizer.Get("classList").String(), "top-right") {
			newHeight = e.min(startHeight-dy, startTop+startHeight)
			newTop = e.max(0, startTop+dy)
		} else if strings.Contains(currentResizer.Get("classList").String(), "top-left") {
			newHeight = e.min(startHeight-dy, startTop+startHeight)
			newTop = e.max(0, startTop+dy)
		}

		e.block.AddStyle("top", fmt.Sprintf("%dpx", newTop))
		e.block.AddStyle("height", fmt.Sprintf("%dpx", e.max(e.blockVerticalMinimumSize, newHeight)))

		e.selectDiv.AddStyle("height", fmt.Sprintf("%dpx", e.max(e.blockVerticalMinimumSize, newHeight)))
		return nil
	}

	resize := func(this js.Value, args []js.Value) interface{} {
		if !isResizing {
			return nil
		}

		element := args[0]

		resizeHorizontal(this, args)
		resizeVertical(this, args)
		_ = e.updateOrnament() // todo: erro aqui, tem lógica?

		width := e.block.Get().Get("offsetWidth").Int()
		height := e.block.Get().Get("offsetHeight").Int()

		e.onResize(element, width, height)
		return nil
	}

	stopResize := func(this js.Value, args []js.Value) interface{} {
		//element := args[0]
		isResizing = false
		//document.removeEventListener("mousemove", resize);
		//document.removeEventListener("mouseup", stopResize);
		return nil
	}

	stopDrag := func(this js.Value, args []js.Value) interface{} {
		isDragging = false
		e.block.AddStyle("cursor", "grab")

		js.Global().Call("removeEventListener", "mousemove", drag)
		//js.Global().Call("removeEventListener","mouseup", stopDrag) // todo: resolver
		return nil
	}

	e.block.Get().Call("mousedown", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		element := args[0]
		if !e.dragEnabled || strings.Contains(element.Get("className").String(), "resizer") {
			return nil
		}

		isDragging = true
		startX = element.Get("clientX").Int()
		startY = element.Get("clientY").Int()
		startLeft = e.block.GetOffsetLeft()
		startTop = e.block.GetOffsetTop()
		js.Global().Call("addEventListener", "mousemove", drag)
		js.Global().Call("addEventListener", "mouseup", stopDrag)

		return nil
	}))

	for k := range e.resizers {
		e.resizers[k].Get().Call("addEventListener", "mousedown", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			if !e.resizeEnabled {
				return nil
			}

			element := args[0]
			element.Call("stopPropagation") // preventDefault

			isResizing = true
			startX = element.Get("clientX").Int()
			startY = element.Get("clientY").Int()
			startWidth = e.block.Get().Get("offsetWidth").Int()
			startHeight = e.block.Get().Get("offsetHeight").Int()
			startLeft = e.block.Get().Get("offsetLeft").Int()
			startTop = e.block.Get().Get("offsetTop").Int()

			currentResizer = element.Get("target")
			currentResizer.Call("addEventListener", "mousemove", resize)
			currentResizer.Call("addEventListener", "mouseup", stopResize)

			return nil
		}))
	}

}

func (e *DeviceOpAmp) onResize(element js.Value, width, height int) {
	//this.connStop.setX(width-50-4);
	//this.connStop.setY(height-40-2);
	//console.log("onResize", e, width, height);
}
