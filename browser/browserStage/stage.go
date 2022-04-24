package browserStage

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

type Stage struct {
	selfDocument js.Value
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
	e.selfDocument.Get("body").Set("style", browserMouse.KCursorAuto.String())
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
	e.selfDocument.Get("body").Set("style", browserMouse.KCursorNone.String())
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
func (e *Stage) SetMouse(value browserMouse.CursorType) (ref *Stage) {
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
//
// Português:
//
//  Associa uma função a um evento do mouse.
//
//   Exemplo:
//     stage.AddListener(browserMouse.KEventMouseOver, onMouseEvent)
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
func (e *Stage) AddListener(eventType interface{}, manager browserMouse.SimpleManager) (ref *Stage) {

	mouseMoveEvt := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var mouseEvent = browserMouse.MouseEvent{}

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
		e.selfDocument.Call("addEventListener", converted.String(), mouseMoveEvt)

	case eventAnimation.EventAnimation:
		e.selfDocument.Call("addEventListener", converted.String(), mouseMoveEvt)

	case eventClipBoard.EventClipBoard:
		e.selfDocument.Call("addEventListener", converted.String(), mouseMoveEvt)

	case eventDrag.EventDrag:
		e.selfDocument.Call("addEventListener", converted.String(), mouseMoveEvt)

	case eventFocus.EventFocus:
		e.selfDocument.Call("addEventListener", converted.String(), mouseMoveEvt)

	case eventHashChange.EventHashChange:
		e.selfDocument.Call("addEventListener", converted.String(), mouseMoveEvt)

	case eventInput.EventInput:
		e.selfDocument.Call("addEventListener", converted.String(), mouseMoveEvt)

	case eventKeyboard.EventKeyboard:
		e.selfDocument.Call("addEventListener", converted.String(), mouseMoveEvt)

	case browserMouse.Event:
		e.selfDocument.Call("addEventListener", converted.String(), mouseMoveEvt)

	case eventPageTransition.EventPageTransition:
		e.selfDocument.Call("addEventListener", converted.String(), mouseMoveEvt)

	case eventUi.EventUi:
		e.selfDocument.Call("addEventListener", converted.String(), mouseMoveEvt)

	case eventWheel.EventWheel:
		e.selfDocument.Call("addEventListener", converted.String(), mouseMoveEvt)

	default:
		log.Fatalf("event must be a event type")
	}

	return e
}
