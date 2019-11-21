package iotmaker_platform_webbrowser

import (
	"syscall/js"
)

const (
	kCanvasNotSet int = iota
	kCanvas2DContext
	kCanvas3DContext
)

// todo: SelfContextType deve ser um enum

// en: The Canvas API provides a means for drawing graphics via JavaScript and the HTML <canvas> element. Among other
// things, it can be used for animation, game graphics, data visualization, photo manipulation, and real-time video
// processing.
//
// The Canvas API largely focuses on 2D graphics. The WebGL API, which also uses the <canvas> element, draws
// hardware-accelerated 2D and 3D graphics.
type Canvas struct {
	SelfContext     js.Value
	SelfContextType int
	Element
}

func (el *Canvas) GetCanvas() js.Value {
	return el.SelfElement
}

func (el *Canvas) GetContext() js.Value {
	return el.SelfContext
}

func (el *Canvas) Call(jsFunction string, value interface{}) js.Value {
	return el.selfDocument.Call(jsFunction, value)
}

func (el *Canvas) Set(jsParam string, value ...interface{}) {
	el.selfDocument.Set(jsParam, value)
}

/*func (el *Canvas) CreateNewWith3DContext(width, height iotmaker_types.Pixel) {
	el.selfCanvas = el.Call("getElementsById", "myCanvas")
	el.Set("width", width)
	el.selfCanvas.Set("height", height)
	el.selfCanvas.Call("getContext", "3d")
}*/

// todo: tem que saber que id é um canvas
func (el *Canvas) InitializeContext2DById(id string) {
	el.Document.Initialize()
	el.SelfElement = el.Element.NewCanvas(id)
	el.SelfContextType = 1
	el.SelfContext = el.SelfElement.Call("getContext", "2d")
}

// todo: tem que saber que id é um canvas
func (el *Canvas) InitializeContext3DById(id string) {
	el.Element.NewCanvas(id)
	el.SelfContextType = 2
	el.SelfContext = el.SelfElement.Call("getContext", "3d")
}

func (el *Canvas) AppendToDocumentBody() {
	el.selfDocument.Get("body").Call("appendChild", el.SelfElement)
}

// en: Saves the state of the current context
func (el *Canvas) Save() {
	el.selfDocument.Call("save")
}

// en: Returns previously saved path state and attributes
func (el *Canvas) Restore() {
	el.selfDocument.Call("restore")
}

func (el *Canvas) CreateEvent() {
	el.selfDocument.Call("createEvent")
}

// todo: toDataURL()
