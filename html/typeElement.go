package html

import (
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/css"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/document"
	"log"
	"strings"
	"syscall/js"
)

const (
	KIdToAppendNotFound    = "html.AppendById().error: id to append not found:"
	KNewElementIsUndefined = "div.NewDiv().error: new element is undefined:"
)

type Tag struct {
	id          string
	selfElement js.Value
	document    document.Document

	css      string
	cssClass *css.Class
}

// Css
//
// English:
//
//  Is the equivalent of <... css="name1 name2 nameN">
//
//   Input:
//     classes: list of css class.
//
// Português:
//
//  Equivale a <... css="name1 name2 nameN">
//
//   Entrada:
//     classes: lista de classes css.
func (e *Tag) Css(classes ...string) (ref *Tag) {
	e.css = strings.Join(classes, " ")

	// English:
	// When a tag is called, for example NewDiv(), it already assembles the data. Therefore, this
	// function must be called whenever an html tag changes.
	//
	// Português
	// Quando uma tag é chamada, por exemplo, NewDiv(), ela já monta os dados. Por isto, esta função
	// deve ser chamada sempre que uma tag html muda.
	e.mount()

	return e
}

// mount
//
// English:
//
//  Assemble the HTML tags of each element.
//
//  Note:
//	  * When a tag is called, for example NewDiv(), it already assembles the data. Therefore, this
//	    function must be called whenever an html tag changes.
//
// Português:
//
//  Monta as tags HTML de cada elemento.
//
//  Nota:
//	  * Quando uma tag é chamada, por exemplo, NewDiv(), ela já monta os dados. Por isto, esta função
//	    deve ser chamada sempre que uma tag html muda.
func (e *Tag) mount() {
	log.Print("css:", e.css)
	if e.css != "" {
		e.selfElement.Set("classList", e.css)
	}
}

// SetCss
//
// English:
//
//  Add the css classes to the created element.
//
//   Input:
//     value: object pointer to css.Class initialized
//
//   Note:
//     * This function is equivalent to css.SetList("current", classes...)
//     * Css has a feature that allows you to easily change the list of css classes of an html tag,
//       with the functions SetList(), CssToggle() and CssToggleTime();
//     * Is the equivalent of <... css="name1 name2 nameN">
//
//
// Português:
//
//  Adiciona as classes css ao elemento criado.
//
//   Entrada:
//     classes: lista de classes css.
//
//   Nota:
//     * Esta função equivale a SetList("current", classes...);
//     * Css tem uma funcionalidade que permite trocar a lista de classes css de uma tag html de forma
//       fácil, com as funções SetList(), CssToggle() e CssToggleTime();
//     * Equivale a <... css="name1 name2 nameN">
func (e *Tag) SetCss(value *css.Class) (ref *Tag) {
	e.cssClass = value
	e.cssClass.SetRef(e.id, &e.selfElement)

	return e
}

// todo: um elemento "remove" que pare todas threads e coisas do tipo
// todo: css tem que ser ponteiro?

// AppendById
//
// English:
//
// Adds a node to the end of the list of children of a specified parent node. If the node already exists in the document, it is removed from its current parent node before being added to the new parent.
//
//   Input:
//     appendId: id of parent element.
//
//   Note:
//     * The equivalent of:
//         var p = document.createElement("p");
//         document.body.appendChild(p);
//
// Português:
//
//  Adiciona um nó ao final da lista de filhos de um nó pai especificado. Se o nó já existir no documento, ele é removido de seu nó pai atual antes de ser adicionado ao novo pai.
//
//   Entrada:
//     appendId: id do elemento pai.
//
//   Nota:
//     * Equivale a:
//         var p = document.createElement("p");
//         document.body.appendChild(p);
func (e Tag) AppendById(appendId string) (ref *Tag) {

	toAppend := js.Global().Get("document").Call("getElementById", appendId)
	if toAppend.IsUndefined() == true || toAppend.IsNull() == true {
		log.Print(KIdToAppendNotFound, appendId)
		return
	}

	toAppend.Call("appendChild", e.selfElement)
	return
}

// Append
//
// English:
//
// Adds a node to the end of the list of children of a specified parent node. If the node already exists in the document, it is removed from its current parent node before being added to the new parent.
//
//   Input:
//     append: element in js.Value format.
//
//   Note:
//     * The equivalent of:
//         var p = document.createElement("p");
//         document.body.appendChild(p);
//
// Português:
//
//  Adiciona um nó ao final da lista de filhos de um nó pai especificado. Se o nó já existir no documento, ele é removido de seu nó pai atual antes de ser adicionado ao novo pai.
//
//   Entrada:
//     appendId: elemento no formato js.Value.
//
//   Nota:
//     * Equivale a:
//         var p = document.createElement("p");
//         document.body.appendChild(p);
func (e *Tag) Append(append interface{}) (ref *Tag) {
	append.(Tag).selfElement.Call("appendChild", e.selfElement)
	return e
}

// Div
//
// English:
//
//  Tag html div
//
// Português:
//
//  Tag html div
type Div struct {
	Tag
}

// NewDiv
//
// English:
//
//  Create a html div.
//
//   Note:
//     * Div inherits Tag, so see Tag documentation for all functions.
//
//   Example:
//       // basic
//       var a html.Div
//       a.NewDiv("example1").
//         Css("user").
//         AppendById("stage")
//
//       // css
//       var b html.Div
//       b.NewDiv("example2").
//         Css("user").
//         SetList("red", "user", "red").
//         SetList("yellow", "user", "yellow").
//         SetList("normal", "user").
//         CssToggleTime(time.Second, "red", "yellow").
//         CssToggleLoop(10).
//         CssOnLoopEnd("normal").
//         CssToggleStart().
//         AppendById("stage")
//
// Português:
//
//  Cria uma div html.
//
//   Nota:
//     * Div herda Tag, por isto, veja a documentação de Tag para vê todas as funções.
//
//   Exemplo:
//       // basic
//       var a html.Div
//       a.NewDiv("example1").
//         Css("user").
//         AppendById("stage")
//
//       // css
//       var b html.Div
//       b.NewDiv("example2").
//         Css("user").
//         SetList("red", "user", "red").
//         SetList("yellow", "user", "yellow").
//         SetList("normal", "user").
//         CssToggleTime(time.Second, "red", "yellow").
//         CssToggleLoop(10).
//         CssOnLoopEnd("normal").
//         CssToggleStart().
//         AppendById("stage")
func (e *Div) NewDiv(id string) (ref *Div) {
	e.id = id
	e.selfElement = js.Global().Get("document").Call("createElement", "div")
	if e.selfElement.IsUndefined() == true || e.selfElement.IsNull() == true {
		log.Print(KNewElementIsUndefined, id)
		return
	}

	e.selfElement.Set("id", id)
	e.mount()
	return e
}
