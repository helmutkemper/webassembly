package canvas

import (
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/browserMouse"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/event"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/eventAnimation"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/eventClipBoard"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/eventDrag"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/eventFocus"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/eventHashChange"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/eventInput"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/eventKeyboard"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/eventPageTransition"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/eventUi"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/eventWheel"
	"log"
	"syscall/js"
)

func (el *Canvas) AddEventListener(eventType interface{}, mouseMoveEvt interface{}) {
	switch converted := eventType.(type) {
	case event.Event:
		el.SelfElement.Call("addEventListener", converted.String(), mouseMoveEvt.(js.Func))

	case eventAnimation.EventAnimation:
		el.SelfElement.Call("addEventListener", converted.String(), mouseMoveEvt.(js.Func))

	case eventClipBoard.EventClipBoard:
		el.SelfElement.Call("addEventListener", converted.String(), mouseMoveEvt.(js.Func))

	case eventDrag.EventDrag:
		el.SelfElement.Call("addEventListener", converted.String(), mouseMoveEvt.(js.Func))

	case eventFocus.EventFocus:
		el.SelfElement.Call("addEventListener", converted.String(), mouseMoveEvt.(js.Func))

	case eventHashChange.EventHashChange:
		el.SelfElement.Call("addEventListener", converted.String(), mouseMoveEvt.(js.Func))

	case eventInput.EventInput:
		el.SelfElement.Call("addEventListener", converted.String(), mouseMoveEvt.(js.Func))

	case eventKeyboard.EventKeyboard:
		el.SelfElement.Call("addEventListener", converted.String(), mouseMoveEvt.(js.Func))

	case browserMouse.Event:
		el.SelfElement.Call("addEventListener", converted.String(), mouseMoveEvt.(js.Func))

	case eventPageTransition.EventPageTransition:
		el.SelfElement.Call("addEventListener", converted.String(), mouseMoveEvt.(js.Func))

	case eventUi.EventUi:
		el.SelfElement.Call("addEventListener", converted.String(), mouseMoveEvt.(js.Func))

	case eventWheel.EventWheel:
		el.SelfElement.Call("addEventListener", converted.String(), mouseMoveEvt.(js.Func))

	default:
		log.Fatalf("event must be a event type")
	}
}
