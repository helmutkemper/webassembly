package document

import (
	"github.com/helmutkemper/iotmaker.platform.webbrowser/event"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/eventAnimation"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/eventClipBoard"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/eventDrag"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/eventFocus"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/eventHashChange"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/eventInput"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/eventKeyboard"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/eventMouse"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/eventPageTransition"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/eventUi"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/eventWheel"
	"log"
	"syscall/js"
)

type Document struct {
	hasInitialized bool
	SelfDocument   js.Value
}

func (el *Document) Initialize() {
	el.hasInitialized = true
	el.SelfDocument = js.Global().Get("document")
}

func (el *Document) Get() js.Value {

	if el.hasInitialized == false {
		el.Initialize()
	}

	return el.SelfDocument
}

func (el *Document) HideMousePointer() {
	if el.hasInitialized == false {
		el.Initialize()
	}

	el.SelfDocument.Get("body").Set("style", "cursor: none")
}

func (el *Document) AppendChildToDocumentBody(value interface{}) {

	if el.hasInitialized == false {
		el.Initialize()
	}

	el.SelfDocument.Get("body").Call("appendChild", value)
}

func (el *Document) RemoveChildFromDocumentBody(value interface{}) {

	if el.hasInitialized == false {
		el.Initialize()
	}

	el.SelfDocument.Get("body").Call("removeChild", value)
}

func (el Document) GetDocumentWidth() int {
	return el.SelfDocument.Get("body").Get("clientWidth").Int()
}

func (el Document) GetDocumentHeight() int {
	return el.SelfDocument.Get("body").Get("clientHeight").Int()
}

func (el Document) ResizeToScreen() {
	el.SelfDocument.Get("body").Set("width", el.SelfDocument.Get("body").Get("clientWidth").Int())
	el.SelfDocument.Get("body").Set("height", el.SelfDocument.Get("body").Get("clientHeight").Int())
}

func (el *Document) AppendChild(element string, value interface{}) {

	if el.hasInitialized == false {
		el.Initialize()
	}

	el.SelfDocument.Get(element).Call("appendChild", value)
}

func (el *Document) AddEventListener(eventType interface{}, mouseMoveEvt js.Func) {
	switch converted := eventType.(type) {
	case event.Event:
		el.SelfDocument.Call("addEventListener", converted.String(), mouseMoveEvt)

	case eventAnimation.EventAnimation:
		el.SelfDocument.Call("addEventListener", converted.String(), mouseMoveEvt)

	case eventClipBoard.EventClipBoard:
		el.SelfDocument.Call("addEventListener", converted.String(), mouseMoveEvt)

	case eventDrag.EventDrag:
		el.SelfDocument.Call("addEventListener", converted.String(), mouseMoveEvt)

	case eventFocus.EventFocus:
		el.SelfDocument.Call("addEventListener", converted.String(), mouseMoveEvt)

	case eventHashChange.EventHashChange:
		el.SelfDocument.Call("addEventListener", converted.String(), mouseMoveEvt)

	case eventInput.EventInput:
		el.SelfDocument.Call("addEventListener", converted.String(), mouseMoveEvt)

	case eventKeyboard.EventKeyboard:
		el.SelfDocument.Call("addEventListener", converted.String(), mouseMoveEvt)

	case eventMouse.EventMouse:
		el.SelfDocument.Call("addEventListener", converted.String(), mouseMoveEvt)

	case eventPageTransition.EventPageTransition:
		el.SelfDocument.Call("addEventListener", converted.String(), mouseMoveEvt)

	case eventUi.EventUi:
		el.SelfDocument.Call("addEventListener", converted.String(), mouseMoveEvt)

	case eventWheel.EventWheel:
		el.SelfDocument.Call("addEventListener", converted.String(), mouseMoveEvt)

	default:
		log.Fatalf("event must be a event type")
	}
}
