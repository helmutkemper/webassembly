package document

import (
	"errors"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/css"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/event"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/eventAnimation"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/eventClipBoard"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/eventDrag"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/eventFocus"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/eventHashChange"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/eventInput"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/eventKeyboard"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/eventMouse"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/eventPageTransition"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/eventUi"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/eventWheel"
	"log"
	"syscall/js"
)

type GenericElementTypes string

type Controlled interface {
	Property | css.Class
}

type Property struct {
	Property string
	Value    any
}

type Document struct {
	hasInitialized bool
	SelfDocument   js.Value
}

func (el *Document) Init() {
	el.hasInitialized = true
	el.SelfDocument = js.Global().Get("document")
}

func (el *Document) Get() js.Value {

	if el.hasInitialized == false {
		el.Init()
	}

	return el.SelfDocument
}

func (el *Document) HideMousePointer() {
	if el.hasInitialized == false {
		el.Init()
	}

	el.SelfDocument.Get("body").Set("style", "cursor: none")
}

func (el *Document) AppendChildToDocumentBody(value interface{}) {

	if el.hasInitialized == false {
		el.Init()
	}

	el.SelfDocument.Get("body").Call("appendChild", value)
}

func (el *Document) RemoveChildFromDocumentBody(value interface{}) {

	if el.hasInitialized == false {
		el.Init()
	}

	el.SelfDocument.Get("body").Call("removeChild", value)
}

func (el Document) GetDocumentWidth() int {
	return el.SelfDocument.Get("body").Get("clientWidth").Int()
}

// GetElementById
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
func (el Document) GetElementById(document Document, id string) (element interface{}) {
	elementRet := document.SelfDocument.Call("getElementById", id)
	if elementRet.IsUndefined() == true || elementRet.IsNull() {
		log.Printf("getElementById(%v).undefined", id)
		return nil
	}

	return elementRet
}

func (el Document) CreateElement(document Document, appendId interface{}, name string, args ...interface{}) (err error) {

	// Ordem de criação para funcionar:
	//t := js.Global().Get("document").Call("createElement", "div")
	//t.Set("id", "vivo2")
	//t.Set("classList", "animate")
	//
	//js.Global().Get("document").Call("getElementById", "palco").Call("appendChild", t)

	newElement := document.SelfDocument.Call("createElement", name)
	if newElement.IsUndefined() == true || newElement.IsNull() == true {
		err = errors.New("ls.createElement(" + name + ").error: new element is undefined")
		return
	}

	for _, genericElement := range args {
		switch converted := genericElement.(type) {
		case Property:
			newElement.Set(converted.Property, converted.Value)

		case css.Class:
			newElement.Set("classList", converted.String())
		}
	}

	switch appendId.(type) {
	case nil:

	case string:
		toAppend := document.SelfDocument.Call("getElementById", appendId)
		if toAppend.IsUndefined() == true || toAppend.IsNull() == true {
			log.Printf("CreateElement().error: id not found")
			return
		}
		toAppend.Call("appendChild", newElement)

	default:
		log.Printf("CreateElement().error: appendId must be a string or nil")
	}

	return
}

func (el Document) GetElementStyle(element interface{}, style string) (value interface{}) {
	var ok bool
	var jsValue js.Value
	jsValue, ok = element.(js.Value)
	if ok == false {
		log.Printf("GetElementStyle().error: element is not a js.value")
		return
	}

	return jsValue.Get("style").Get(style)
}

func (el Document) SetElementStyle(element interface{}, style string, value interface{}) {
	var ok bool
	var jsValue js.Value
	jsValue, ok = element.(js.Value)
	if ok == false {
		log.Printf("GetElementStyle().error: element is not a js.value")
		return
	}

	jsValue.Get("style").Set(style, value)
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
		el.Init()
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
