package html

import (
	"log"
	"strconv"
	"syscall/js"
)

// TagHead
//
// English:
//
// Português:
type TagHead struct {
	commonEvents commonEvents

	// selfElement
	//
	// English:
	//
	//  Reference to self element as js.Value.
	//
	// Português:
	//
	//  Referencia ao próprio elemento na forma de js.Value.
	selfElement js.Value

	// stage
	//
	// English:
	//
	//  Browser main document reference captured at startup.
	//
	// Português:
	//
	//  Referencia do documento principal do navegador capturado na inicialização.
	stage js.Value

	css map[string]map[string]any
}

// Reference
//
// English:
//
// Pass the object reference to an external variable.
//
// Português:
//
// Passa a referencia do objeto para uma variável externa.
//
//	Example: / Exemplo:
//	  var circle *html.TagSvgCircle
//	  factoryBrowser.NewTagSvgCircle().Reference(&circle).R(5).Fill(factoryColor.NewRed())
//	  log.Printf("x: %v, y: %v", circle.GetX(), circle.GetY())
func (e *TagHead) Reference(reference **TagHead) (ref *TagHead) {
	*reference = e
	return e
}

// Lang
//
// English:
//
//	Specifies the language of the element's content.
//
// The lang attribute specifies the language of the element's content.
//
// Common examples are KLanguageEnglish for English, KLanguageSpanish for Spanish, KLanguageFrench
// for French, and so on.
//
// Português:
//
//	Especifica o idioma do conteúdo do elemento.
//
// O atributo lang especifica o idioma do conteúdo do elemento.
//
// Exemplos comuns são KLanguageEnglish para inglês, KLanguageSpanish para espanhol, KLanguageFrench
// para francês e assim por diante.
func (e *TagHead) Lang(language Language) (ref *TagHead) {
	e.selfElement.Set("lang", language.String())
	return e
}

// CreateElement
//
// English:
//
//	In an HTML document, the Document.createElement() method creates the specified HTML element or an
//	HTMLUnknownElement if the given element name is not known.
//
// Português:
//
//	Em um documento HTML, o método Document.createElement() cria o elemento HTML especificado ou um
//	HTMLUnknownElement se o nome do elemento dado não for conhecido.
func (e *TagHead) CreateElement() (ref *TagHead) {
	e.selfElement = js.Global().Get("document").Call("createElement", "h1")
	if e.selfElement.IsUndefined() == true || e.selfElement.IsNull() == true {
		log.Print(KNewElementIsUndefined)
		return
	}

	return e
}

// Append
//
// English:
//
//	Adds a node to the end of the list of children of a specified parent node. If the node already
//	exists in the document, it is removed from its current parent node before being added to the new
//	parent.
//
//	 Input:
//	   append: element in js.Value format.
//
//	 Note:
//	   * The equivalent of:
//	       var p = document.createElement("p");
//	       document.body.appendChild(p);
//
// Português:
//
//	Adiciona um nó ao final da lista de filhos de um nó pai especificado. Se o nó já existir no
//	documento, ele é removido de seu nó pai atual antes de ser adicionado ao novo pai.
//
//	 Entrada:
//	   appendId: elemento no formato js.Value.
//
//	 Nota:
//	   * Equivale a:
//	       var p = document.createElement("p");
//	       document.body.appendChild(p);
//
// fixme: fazer append() assim em todas as tags html, exceto svg
func (e *TagHead) Append(elements ...Compatible) (ref *TagHead) {
	fragment := js.Global().Get("document").Call("createDocumentFragment")
	for _, element := range elements {
		fragment.Call("appendChild", element.Get())
	}

	e.selfElement.Call("appendChild", fragment)
	return e
}

func (e *TagHead) Get() (el js.Value) {
	return e.selfElement
}

// Init
//
// English:
//
//	Initializes the object correctly.
//
// Português:
//
//	Inicializa o objeto corretamente.
func (e *TagHead) Init() (ref *TagHead) {
	e.CreateElement()
	e.prepareStageReference()

	return e
}

// prepareStageReference
//
// English:
//
//	Prepares the stage reference at initialization.
//
// Português:
//
//	Prepara à referencia do stage na inicialização.
func (e *TagHead) prepareStageReference() {
	e.stage = js.Global().Get("document").Get("body")
}

// Text
//
// English:
//
// Adds plain text to the tag's content.
//
// Text:
//
// Adiciona um texto simples ao conteúdo da tag.
func (e *TagHead) Text(value string) (ref *TagHead) {
	e.selfElement.Set("textContent", value)
	return e
}

// Html
//
// English:
//
// Adds HTML to the tag's content.
//
// Text:
//
// Adiciona HTML ao conteúdo da tag.
func (e *TagHead) Html(value string) (ref *TagHead) {
	e.selfElement.Set("innerHTML", value)
	return e
}

func (e *TagHead) Css(selector, placeholder string, value any) (ref *TagHead) {
	if e.css == nil {
		e.css = make(map[string]map[string]any)
	}
	if e.css[selector] == nil {
		e.css[selector] = make(map[string]any)
	}

	e.css[selector][placeholder] = value
	return e
}

func (e *TagHead) CssAppend() (ref *TagHead) {
	if e.css == nil {
		return
	}

	document := js.Global().Get("document")

	style := document.Call("createElement", "style")
	style.Call("setAttribute", "type", "text/css")

	css := ""

	for selector, properties := range e.css {
		css += selector + " {\n"
		for placeholder, value := range properties {
			switch converted := value.(type) {
			case string:
				css += "	" + placeholder + ": " + converted + ";\n"
			case bool:
				css += "	" + placeholder + ": " + strconv.FormatBool(converted) + ";\n"
			case int:
				css += "	" + placeholder + ": " + strconv.FormatInt(int64(converted), 10) + ";\n"
			case int64:
				css += "	" + placeholder + ": " + strconv.FormatInt(converted, 10) + ";\n"
			case float64:
				css += "	" + placeholder + ": " + strconv.FormatFloat(converted, 'g', -1, 64) + ";\n"
			}
		}
		css += "}\n"
	}
	style.Set("innerHTML", css)

	document.Get("head").Call("appendChild", style)
	return e
}

// Remove
//
// English:
//
//	Removes a child node from the DOM and returns the removed node.
//
// Português:
//
//	Remove um nó filho do DOM e retorna o nó removido.
//
// # Remove
//
// English:
//
//	Removes a child node from the DOM and returns the removed node.
//
// Português:
//
//	Remove um nó filho do DOM e retorna o nó removido.
func (e *TagHead) Remove(elements ...Compatible) (ref *TagHead) {
	for _, element := range elements {
		e.selfElement.Call("removeChild", element.Get())
	}

	return e
}
