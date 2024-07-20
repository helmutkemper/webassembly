package globalDocument

import (
	"errors"
	"github.com/helmutkemper/webassembly/browser/css"
	"github.com/helmutkemper/webassembly/event"
	"github.com/helmutkemper/webassembly/eventAnimation"
	"github.com/helmutkemper/webassembly/eventClipBoard"
	"github.com/helmutkemper/webassembly/eventDrag"
	"github.com/helmutkemper/webassembly/eventFocus"
	"github.com/helmutkemper/webassembly/eventHashChange"
	"github.com/helmutkemper/webassembly/eventInput"
	"github.com/helmutkemper/webassembly/eventKeyboard"
	"github.com/helmutkemper/webassembly/eventPageTransition"
	"github.com/helmutkemper/webassembly/eventUi"
	"github.com/helmutkemper/webassembly/eventWheel"
	"log"
	"syscall/js"
)

type GenericElementTypes string

type P struct {
	P string
	V any
}

type Document struct {
	hasInitialized bool
	selfDocument   js.Value
}

// Init
//
// English:
//
//	Initializes the document with the browser's main document.
//
// Português:
//
//	Inicializa o documento com o documento principal do navegador.
func (el *Document) Init() {
	el.hasInitialized = true
	el.selfDocument = js.Global().Get("document")
}

// Get
//
// English:
//
//	Returns the document.
//
// Português:
//
//	Retorna o documento.
func (el *Document) Get() js.Value {
	return el.selfDocument
}

// MousePointerAuto
//
// English:
//
//	Sets the mouse pointer to auto.
//
// Português:
//
//	Define o ponteiro do mouse como automático.
func (el *Document) MousePointerAuto() {
	el.selfDocument.Get("body").Set("style", mouse.KCursorAuto.String())
}

// MousePointerHide
//
// English:
//
//	Sets the mouse pointer to hide.
//
// Português:
//
//	Define o ponteiro do mouse como oculto.
func (el *Document) MousePointerHide() {
	el.selfDocument.Get("body").Set("style", mouse.KCursorNone.String())
}

// SetMousePointer
//
// English:
//
//	Defines the shape of the mouse pointer.
//
//	 Input:
//	   value: mouse pointer shape.
//	     Example: SetMousePointer(mouse.KCursorCell) // Use mouse.K... and let autocomplete do the
//	              rest
//
// Português:
//
//	Define o formato do ponteiro do mouse.
//
//	 Entrada:
//	   V: formato do ponteiro do mouse.
//	     Exemplo: SetMousePointer(mouse.KCursorCell) // Use mouse.K... e deixe o autocompletar fazer
//	              o resto
func (el *Document) SetMousePointer(value mouse.CursorType) {
	el.selfDocument.Get("body").Set("style", value.String())
}

// AppendToDocument
//
// English:
//
//	Adds an element to the document.
//
//	 Input:
//	   value: js.Value element containing an html document.
//
// Português:
//
//	Adiciona um elemento ao documento.
//
//	 Entrada:
//	   value: elemento js.Value contendo um documento html.
func (el *Document) AppendToDocument(value interface{}) {
	el.selfDocument.Get("body").Call("appendChild", value)
}

// RemoveFromDocument
//
// English:
//
//	Removes an html element from the document.
//
//	 Input:
//	   value: js.Value element containing an html document.
//
// Português:
//
//	Remove um elemento html do documento.
//
//	 Entrada:
//	   value: elemento js.Value contendo um documento html.
func (el *Document) RemoveFromDocument(value interface{}) {
	el.selfDocument.Get("body").Call("removeChild", value)
}

// GetDocumentWidth
//
// English:
//
//	Returns the width of the document in pixels.
//
//	 Output:
//	   width: document size in pixels.
//
// Português:
//
//	Retorna o comprimento do documento em pixels.
//
//	 Saída:
//	   width: tamanho do documento em pixels.
func (el Document) GetDocumentWidth() (width int) {
	return el.selfDocument.Get("body").Get("clientWidth").Int()
}

// GetDocumentHeight
//
// English:
//
//	Returns the length of the document in pixels.
//
//	 Output:
//	   width: document size in pixels.
//
// Português:
//
//	Retorna a altura do documento em pixels.
//
//	 Saída:
//	   width: tamanho do documento em pixels.
func (el Document) GetDocumentHeight() (height int) {
	return el.selfDocument.Get("body").Get("clientHeight").Int()
}

// ResizeToScreen
//
// English:
//
//	Resizes the document to the size of the main document.
//
// Português:
//
//	Redimensiona o documento para o tamanho do documento principal.
func (el Document) ResizeToScreen() {
	el.selfDocument.Get("body").Set("width", js.Global().Get("document").Get("body").Get("clientWidth").Int())
	el.selfDocument.Get("body").Set("height", js.Global().Get("document").Get("body").Get("clientHeight").Int())
}

// GetElementById
//
// Português:
//
//	Retorna a referência do elemento através do seu ID.
//
//	 Entrada:
//	   id: string que diferência maiúsculas e minúsculas representando o ID único do elemento sendo
//	       procurado.
//	 Nota:
//	   * Elemento é uma referência a um objeto Element, ou null se um elemento com o ID especificado
//	     não estiver contido neste documento.
//	   * Se não existe um elemento com o id fornecido, esta função retorna null. Note, o parâmetro ID
//	     diferência maiúsculas e minúsculas. Assim document.getElementById("Main") retornará null ao
//	     invés do elemento <div id="main">, devido a "M" e "m" diferirem para o objetivo deste método;
//	   * Elementos que não estão no documento não são procurados por getElementById. Quando criar um
//	     elemento e atribuir um ID ao mesmo, você deve inserir o elemento na árvore do documento com
//	     insertBefore ou método similar antes que você possa acessá-lo com getElementById:
//
//	       var elemento = document.createElement("div");
//	       elemento.id = 'testqq';
//	       var el = document.getElementById('testqq'); // el será null!
//
//	   * Documentos não-HTML, a implementação do DOM deve ter informações que diz quais atributos são
//	     do tipo ID.  Atributos com o nome "id" não são do tipo ID a menos que assim sejam definidos
//	     nos documentos DTD. O atributo id é definido para ser um tipo ID em casos comuns de  XHTML,
//	     XUL, e outros. Implementações que não reconhecem se os atributos são do tipo ID, ou não são
//	     esperados retornam null.
func (el Document) GetElementById(document Document, id string) (element interface{}) {
	elementRet := document.selfDocument.Call("getElementById", id)
	if elementRet.IsUndefined() == true || elementRet.IsNull() {
		log.Printf("getElementById(%v).undefined", id)
		return nil
	}

	return elementRet
}

// CreateElement
//
// English:
//
//	In an HTML document, the document.createElement() method creates the HTML element specified by
//	tagName, or an HTMLUnknownElement if tagName isn't recognized.
//
//	 Note:
//	   * A new HTMLElement is returned if the document is an HTMLDocument, which is the most common
//	     case. Otherwise a new Element is returned.
//
// Português:
//
//	Em um documento HTML, o método document.createElement() cria o elemento HTML especificado por
//	tagName ou um HTMLUnknownElement se tagName não for reconhecido.
//
//	 Note:
//	   * A new HTMLElement is returned if the document is an HTMLDocument, which is the most common
//	     case. Otherwise a new Element is returned.
func (el Document) CreateElement(name string, cssClass []string, properties ...P) (element js.Value, err error) {

	// Ordem de criação para funcionar:
	//t := js.Global().Get("document").Call("createElement", "div")
	//t.Set("id", "vivo2")
	//t.Set("classList", "animate")
	//
	//js.Global().Get("document").Call("getElementById", "palco").Call("appendChild", t)

	element = js.Global().Get("document").Call("createElement", name)
	if element.IsUndefined() == true || element.IsNull() == true {
		err = errors.New("ls.createElement(" + name + ").error: new element is undefined")
		return
	}

	var class = css.Class{}
	class.SetList("current", cssClass...)

	for _, p := range properties {
		element.Set(p.P, p.V)
	}

	return
}

// CreateElementAndAppend
//
// English:
//
//	In an HTML document, the document.createElement() method creates the HTML element specified by
//	tagName, or an HTMLUnknownElement if tagName isn't recognized.
//
//	 Note:
//	   * A new HTMLElement is returned if the document is an HTMLDocument, which is the most common
//	     case. Otherwise a new Element is returned.
//
// Português:
//
//	Em um documento HTML, o método document.createElement() cria o elemento HTML especificado por
//	tagName ou um HTMLUnknownElement se tagName não for reconhecido.
//
//	 Note:
//	   * A new HTMLElement is returned if the document is an HTMLDocument, which is the most common
//	     case. Otherwise a new Element is returned.
func (el Document) CreateElementAndAppend(appendId string, name string, cssClass []string, properties ...P) (element js.Value, err error) {

	// Ordem de criação para funcionar:
	//t := js.Global().Get("document").Call("createElement", "div")
	//t.Set("id", "vivo2")
	//t.Set("classList", "animate")
	//
	//js.Global().Get("document").Call("getElementById", "palco").Call("appendChild", t)

	element = js.Global().Get("document").Call("createElement", name)
	if element.IsUndefined() == true || element.IsNull() == true {
		err = errors.New("ls.createElement(" + name + ").error: new element is undefined")
		return
	}

	var class = css.Class{}
	class.SetList("current", cssClass...)
	element.Set("classList", class.String())

	for _, p := range properties {
		element.Set(p.P, p.V)
	}

	toAppend := js.Global().Get("document").Call("getElementById", appendId)
	if toAppend.IsUndefined() == true || toAppend.IsNull() == true {
		err = errors.New("id to append not found")
		return
	}

	toAppend.Call("appendChild", element)

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

func (el *Document) AppendChild(element string, value interface{}) {

	if el.hasInitialized == false {
		el.Init()
	}

	el.selfDocument.Get(element).Call("appendChild", value)
}

func (el *Document) AddEventListener(eventType interface{}, mouseMoveEvt js.Func) {
	switch converted := eventType.(type) {
	case event.Event:
		el.selfDocument.Call("addEventListener", converted.String(), mouseMoveEvt)

	case eventAnimation.EventAnimation:
		el.selfDocument.Call("addEventListener", converted.String(), mouseMoveEvt)

	case eventClipBoard.EventClipBoard:
		el.selfDocument.Call("addEventListener", converted.String(), mouseMoveEvt)

	case eventDrag.EventDrag:
		el.selfDocument.Call("addEventListener", converted.String(), mouseMoveEvt)

	case eventFocus.EventFocus:
		el.selfDocument.Call("addEventListener", converted.String(), mouseMoveEvt)

	case eventHashChange.EventHashChange:
		el.selfDocument.Call("addEventListener", converted.String(), mouseMoveEvt)

	case eventInput.EventInput:
		el.selfDocument.Call("addEventListener", converted.String(), mouseMoveEvt)

	case eventKeyboard.EventKeyboard:
		el.selfDocument.Call("addEventListener", converted.String(), mouseMoveEvt)

	case mouse.Event:
		el.selfDocument.Call("addEventListener", converted.String(), mouseMoveEvt)

	case eventPageTransition.EventPageTransition:
		el.selfDocument.Call("addEventListener", converted.String(), mouseMoveEvt)

	case eventUi.EventUi:
		el.selfDocument.Call("addEventListener", converted.String(), mouseMoveEvt)

	case eventWheel.EventWheel:
		el.selfDocument.Call("addEventListener", converted.String(), mouseMoveEvt)

	default:
		log.Fatalf("event must be a event type")
	}
}
