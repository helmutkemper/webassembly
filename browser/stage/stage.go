package stage

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/event"
	"github.com/helmutkemper/iotmaker.webassembly/browser/eventAnimation"
	"github.com/helmutkemper/iotmaker.webassembly/browser/eventClipBoard"
	"github.com/helmutkemper/iotmaker.webassembly/browser/eventDrag"
	"github.com/helmutkemper/iotmaker.webassembly/browser/eventFocus"
	"github.com/helmutkemper/iotmaker.webassembly/browser/eventHashChange"
	"github.com/helmutkemper/iotmaker.webassembly/browser/eventInput"
	eventKeyboard "github.com/helmutkemper/iotmaker.webassembly/browser/eventKeyboard"
	"github.com/helmutkemper/iotmaker.webassembly/browser/eventPageTransition"
	"github.com/helmutkemper/iotmaker.webassembly/browser/eventUi"
	"github.com/helmutkemper/iotmaker.webassembly/browser/eventWheel"
	"github.com/helmutkemper/iotmaker.webassembly/browser/mouse"
	"log"
	"sync"
	"syscall/js"
)

type Compatible interface {
	Get() js.Value
}

type Stage struct {
	selfDocument js.Value

	// listener
	//
	// English:
	//
	//  The javascript function removeEventListener needs to receive the function passed in addEventListener
	//
	// Português:
	//
	//  A função javascript removeEventListener necessitam receber a função passada em addEventListener
	listener *sync.Map
}

// Init
//
// English:
//
//  Initializes the document with the browser's main document.
//
// Português:
//
//  Inicializa o documento com o documento principal do navegador.
func (e *Stage) Init() {
	e.selfDocument = js.Global().Get("document")
	e.listener = new(sync.Map)
}

func (e *Stage) Append(value Compatible) (ref *Stage) {
	e.selfDocument.Get("body").Call("appendChild", value.Get())
	return e
}

// Get
//
// English:
//
//  Returns the document.
//
// Português:
//
//  Retorna o documento.
func (e *Stage) Get() js.Value {
	return e.selfDocument
}

// MouseAuto
//
// English:
//
//  Sets the mouse pointer to auto.
//
// Português:
//
//  Define o ponteiro do mouse como automático.
func (e *Stage) MouseAuto() (ref *Stage) {
	e.selfDocument.Get("body").Set("style", mouse.KCursorAuto.String())
	return e
}

// MouseHide
//
// English:
//
//  Sets the mouse pointer to hide.
//
// Português:
//
//  Define o ponteiro do mouse como oculto.
func (e *Stage) MouseHide() (ref *Stage) {
	e.selfDocument.Get("body").Set("style", mouse.KCursorNone.String())
	return e
}

// SetMouse
//
// English:
//
//  Defines the shape of the mouse pointer.
//
//   Input:
//     value: mouse pointer shape.
//       Example: SetMouse(mouse.KCursorCell) // Use mouse.K... and let autocomplete do the
//                rest
//
// Português:
//
//  Define o formato do ponteiro do mouse.
//
//   Entrada:
//     value: formato do ponteiro do mouse.
//       Exemplo: SetMouse(mouse.KCursorCell) // Use mouse.K... e deixe o autocompletar fazer
//                o resto
func (e *Stage) SetMouse(value mouse.CursorType) (ref *Stage) {
	e.selfDocument.Get("body").Set("style", value.String())
	return e
}

// Add
//
// English:
//
//  Adds an element to the document.
//
//   Input:
//     value: js.Value element containing an html document.
//
// Português:
//
//  Adiciona um elemento ao documento.
//
//   Entrada:
//     value: elemento js.Value contendo um documento html.
func (e *Stage) Add(value interface{}) (ref *Stage) {
	e.selfDocument.Get("body").Call("appendChild", value)
	return e
}

// Remove
//
// English:
//
//  Removes an html element from the document.
//
//   Input:
//     value: js.Value element containing an html document.
//
// Português:
//
//  Remove um elemento html do documento.
//
//   Entrada:
//     value: elemento js.Value contendo um documento html.
func (e *Stage) Remove(value interface{}) (ref *Stage) {
	e.selfDocument.Get("body").Call("removeChild", value)
	return e
}

// GetWidth
//
// English:
//
//  Returns the width of the document in pixels.
//
//   Output:
//     width: document size in pixels.
//
// Português:
//
//  Retorna o comprimento do documento em pixels.
//
//   Saída:
//     width: tamanho do documento em pixels.
func (e Stage) GetWidth() (width int) {
	return e.selfDocument.Get("body").Get("clientWidth").Int()
}

// GetHeight
//
// English:
//
//  Returns the length of the document in pixels.
//
//   Output:
//     width: document size in pixels.
//
// Português:
//
//  Retorna a altura do documento em pixels.
//
//   Saída:
//     width: tamanho do documento em pixels.
func (e Stage) GetHeight() (height int) {
	return e.selfDocument.Get("body").Get("clientHeight").Int()
}

// ResizeToScreen
//
// English:
//
//  Resizes the document to the size of the main document.
//
// Português:
//
//  Redimensiona o documento para o tamanho do documento principal.
func (e *Stage) ResizeToScreen() (ref *Stage) {
	e.selfDocument.Get("body").Set("width", js.Global().Get("document").Get("body").Get("clientWidth").Int())
	e.selfDocument.Get("body").Set("height", js.Global().Get("document").Get("body").Get("clientHeight").Int())
	return e
}

// GetById
//
// Português:
//
//  Retorna a referência do elemento através do seu ID.
//
//   Entrada:
//     id: string que diferência maiúsculas e minúsculas representando o ID único do elemento sendo
//         procurado.
//   Nota:
//     * Elemento é uma referência a um objeto Element, ou null se um elemento com o ID especificado
//       não estiver contido neste documento.
//     * Se não existe um elemento com o id fornecido, esta função retorna null. Note, o parâmetro ID
//       diferência maiúsculas e minúsculas. Assim document.getElementById("Main") retornará null ao
//       invés do elemento <div id="main">, devido a "M" e "m" diferirem para o objetivo deste método;
//     * Elementos que não estão no documento não são procurados por getElementById. Quando criar um
//       elemento e atribuir um ID ao mesmo, você deve inserir o elemento na árvore do documento com
//       insertBefore ou método similar antes que você possa acessá-lo com getElementById:
//
//         var elemento = document.createElement("div");
//         elemento.id = 'testqq';
//         var el = document.getElementById('testqq'); // el será null!
//
//     * Documentos não-HTML, a implementação do DOM deve ter informações que diz quais atributos são
//       do tipo ID.  Atributos com o nome "id" não são do tipo ID a menos que assim sejam definidos
//       nos documentos DTD. O atributo id é definido para ser um tipo ID em casos comuns de  XHTML,
//       XUL, e outros. Implementações que não reconhecem se os atributos são do tipo ID, ou não são
//       esperados retornam null.
func (e Stage) GetById(id string) (element interface{}) {
	elementRet := js.Global().Get("document").Call("getElementById", id)
	if elementRet.IsUndefined() == true || elementRet.IsNull() {
		log.Printf("getElementById(%v).undefined", id)
		return nil
	}

	return elementRet
}

// AddListener
//
// English:
//
//  Associates a function with a mouse event.
//
//   Example:
//     stage.AddListener(browserMouse.KEventMouseOver, onMouseEvent)
//     timer := time.NewTimer(10 * time.Second)
//     go func() {
//       select {
//         case <-timer.C:
//         stage.RemoveListener(mouse.KEventMouseOver)
//       }
//     }()
//
// Português:
//
//  Associa uma função a um evento do mouse.
//
//   Exemplo:
//     stage.AddListener(browserMouse.KEventMouseOver, onMouseEvent)
//     timer := time.NewTimer(10 * time.Second)
//     go func() {
//       select {
//         case <-timer.C:
//         stage.RemoveListener(mouse.KEventMouseOver)
//       }
//     }()
//
//     func onMouseEvent(event browserMouse.MouseEvent) {
//       isNull, target := event.GetRelatedTarget()
//       if isNull == false {
//         log.Print("id: ", target.Get("id"))
//         log.Print("tagName: ", target.Get("tagName"))
//       }
//       log.Print(event.GetScreenX())
//       log.Print(event.GetScreenY())
//     }
func (e *Stage) AddListener(eventType interface{}, manager mouse.SimpleManager) (ref *Stage) {

	mouseMoveEvt := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var mouseEvent = mouse.MouseEvent{}

		if len(args) > 0 {
			mouseEvent.Object = args[0]
		}

		if manager != nil {
			manager(mouseEvent)
		}

		return nil
	})

	switch converted := eventType.(type) {
	case event.Event:
		e.listener.Store(converted.String(), mouseMoveEvt)
		e.selfDocument.Call("addEventListener", converted.String(), mouseMoveEvt)

	case eventAnimation.EventAnimation:
		e.listener.Store(converted.String(), mouseMoveEvt)
		e.selfDocument.Call("addEventListener", converted.String(), mouseMoveEvt)

	case eventClipBoard.EventClipBoard:
		e.listener.Store(converted.String(), mouseMoveEvt)
		e.selfDocument.Call("addEventListener", converted.String(), mouseMoveEvt)

	case eventDrag.EventDrag:
		e.listener.Store(converted.String(), mouseMoveEvt)
		e.selfDocument.Call("addEventListener", converted.String(), mouseMoveEvt)

	case eventFocus.EventFocus:
		e.listener.Store(converted.String(), mouseMoveEvt)
		e.selfDocument.Call("addEventListener", converted.String(), mouseMoveEvt)

	case eventHashChange.EventHashChange:
		e.listener.Store(converted.String(), mouseMoveEvt)
		e.selfDocument.Call("addEventListener", converted.String(), mouseMoveEvt)

	case eventInput.EventInput:
		e.listener.Store(converted.String(), mouseMoveEvt)
		e.selfDocument.Call("addEventListener", converted.String(), mouseMoveEvt)

	case eventKeyboard.EventKeyboard:
		e.listener.Store(converted.String(), mouseMoveEvt)
		e.selfDocument.Call("addEventListener", converted.String(), mouseMoveEvt)

	case mouse.Event:
		e.listener.Store(converted.String(), mouseMoveEvt)
		e.selfDocument.Call("addEventListener", converted.String(), mouseMoveEvt)

	case eventPageTransition.EventPageTransition:
		e.listener.Store(converted.String(), mouseMoveEvt)
		e.selfDocument.Call("addEventListener", converted.String(), mouseMoveEvt)

	case eventUi.EventUi:
		e.listener.Store(converted.String(), mouseMoveEvt)
		e.selfDocument.Call("addEventListener", converted.String(), mouseMoveEvt)

	case eventWheel.EventWheel:
		e.listener.Store(converted.String(), mouseMoveEvt)
		e.selfDocument.Call("addEventListener", converted.String(), mouseMoveEvt)

	default:
		log.Fatalf("event must be a event type")
	}

	return e
}

func (e *Stage) RemoveListener(eventType interface{}) (ref *Stage) {

	switch converted := eventType.(type) {
	case event.Event:
		f, _ := e.listener.Load(converted.String())
		e.selfDocument.Call("removeEventListener", converted.String(), f)

	case eventAnimation.EventAnimation:
		f, _ := e.listener.Load(converted.String())
		e.selfDocument.Call("removeEventListener", converted.String(), f)

	case eventClipBoard.EventClipBoard:
		f, _ := e.listener.Load(converted.String())
		e.selfDocument.Call("removeEventListener", converted.String(), f)

	case eventDrag.EventDrag:
		f, _ := e.listener.Load(converted.String())
		e.selfDocument.Call("removeEventListener", converted.String(), f)

	case eventFocus.EventFocus:
		f, _ := e.listener.Load(converted.String())
		e.selfDocument.Call("removeEventListener", converted.String(), f)

	case eventHashChange.EventHashChange:
		f, _ := e.listener.Load(converted.String())
		e.selfDocument.Call("removeEventListener", converted.String(), f)

	case eventInput.EventInput:
		f, _ := e.listener.Load(converted.String())
		e.selfDocument.Call("removeEventListener", converted.String(), f)

	case eventKeyboard.EventKeyboard:
		f, _ := e.listener.Load(converted.String())
		e.selfDocument.Call("removeEventListener", converted.String(), f)

	case mouse.Event:
		f, _ := e.listener.Load(converted.String())
		e.selfDocument.Call("removeEventListener", converted.String(), f)

	case eventPageTransition.EventPageTransition:
		f, _ := e.listener.Load(converted.String())
		e.selfDocument.Call("removeEventListener", converted.String(), f)

	case eventUi.EventUi:
		f, _ := e.listener.Load(converted.String())
		e.selfDocument.Call("removeEventListener", converted.String(), f)

	case eventWheel.EventWheel:
		f, _ := e.listener.Load(converted.String())
		e.selfDocument.Call("removeEventListener", converted.String(), f)

	default:
		log.Fatalf("event must be a event type")
	}

	return e
}
