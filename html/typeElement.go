package html

import (
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/css"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/document"
	"log"
	"syscall/js"
	"time"
)

const (
	KIdToAppendNotFound    = "thml.AppendById().error: id to append not found:"
	KCssListNameNotFound   = "thml.CssToggle().error: css list name not found:"
	KNewElementIsUndefined = "div.NewDiv().error: new element is undefined:"
)

type Tag struct {
	id          string
	selfElement js.Value
	document    document.Document

	css          css.Class
	cssCounter   int
	cssMax       int
	cssList      []string
	cssTimeout   *time.Ticker
	cssInterval  time.Duration
	cssDone      chan struct{}
	cssLoop      int
	cssLoopFlag  bool
	cssOnLoopEnd string
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

// Css
//
// English:
//
//  Add the css classes to the created element.
//
//   Input:
//     classes: css class list.
//
//   Note:
//     * This function is equivalent to html.CssAddList("current", classes...)
//     * Css has a feature that allows you to easily change the list of css classes of an html tag,
//       with the functions CssAddList(), CssToggle() and CssToggleTime();
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
//     * Esta função equivale a CssAddList("current", classes...);
//     * Css tem uma funcionalidade que permite trocar a lista de classes css de uma tag html de forma
//       fácil, com as funções CssAddList(), CssToggle() e CssToggleTime();
//     * Equivale a <... css="name1 name2 nameN">
func (e *Tag) Css(classes ...string) (ref *Tag) {
	e.css.SetList("current", classes...)
	e.selfElement.Set("classList", e.css.String())

	return e
}

// CssAddList
//
// English:
//
//  Adds a new list of css classes.
//
//   Input:
//     name: css list name;
//     classes: css class list.
//
//   Note:
//     * This function is equivalent to html.CssAddList("current", classes...)
//     * Css has a feature that allows you to easily change the list of css classes of an html tag,
//       with the functions CssAddList(), CssToggle() and CssToggleTime();
//     * Is the equivalent of <... css="name1 name2 nameN">
//
// Português:
//
//  Adiciona uma nova lista de classes css.
//
//   Entrada:
//     name: nome da lista css;
//     classes: lista de classes css.
//
//   Nota:
//     * Esta função equivale a CssAddList("current", classes...);
//     * Css tem uma funcionalidade que permite trocar a lista de classes css de uma tag html de forma
//       fácil, com as funções CssAddList(), CssToggle() e CssToggleTime();
//     * Equivale a <... css="name1 name2 nameN">
func (e *Tag) CssAddList(name string, classes ...string) (ref *Tag) {
	e.css.SetList(name, classes...)

	return e
}

// CssRemoveFromList
//
// English:
//
//  Removes a class name from a given list of classes.
//
//   Input:
//     name: name of the class list;
//     class: name of the class to be removed.
//
//   Note:
//     * This function is equivalent to html.CssAddList("current", classes...)
//     * Css has a feature that allows you to easily change the list of css classes of an html tag,
//       with the functions CssAddList(), CssToggle() and CssToggleTime();
//     * Is the equivalent of <... css="name1 name2 nameN">
//
// Português:
//
//  Remove o nome de uma classe de uma determinada lista de classes.
//
//   Entrada:
//     name: nome da lista de classes;
//     class: nome da classe a ser removida.
//
//   Nota:
//     * Esta função equivale a CssAddList("current", classes...);
//     * Css tem uma funcionalidade que permite trocar a lista de classes css de uma tag html de forma
//       fácil, com as funções CssAddList(), CssToggle() e CssToggleTime();
//     * Equivale a <... css="name1 name2 nameN">
func (e *Tag) CssRemoveFromList(name, class string) (ref *Tag) {
	e.css.RemoveFromList(name, class)

	return e
}

// CssDeleteList
//
// English:
//
//  Removes a list of classes.
//
//   Input:
//     name: name of the list of classes to be removed.
//
//   Note:
//     * This function is equivalent to html.CssAddList("current", classes...)
//     * Css has a feature that allows you to easily change the list of css classes of an html tag,
//       with the functions CssAddList(), CssToggle() and CssToggleTime();
//     * Is the equivalent of <... css="name1 name2 nameN">
//
// Português:
//
//  Remove uma lista de classes.
//
//   Entrada:
//     name: nome da lista de classes a ser removida.
//
//   Nota:
//     * Esta função equivale a CssAddList("current", classes...);
//     * Css tem uma funcionalidade que permite trocar a lista de classes css de uma tag html de forma
//       fácil, com as funções CssAddList(), CssToggle() e CssToggleTime();
//     * Equivale a <... css="name1 name2 nameN">
func (e *Tag) CssDeleteList(name string) (ref *Tag) {
	e.css.DeleteList(name)

	return e
}

// CssToggle
//
// English:
//
//  Swap the element's css list of classes.
//
//   Entrada:
//     name: name of the list of classes to use.
//
//   Note:
//     * This function is equivalent to html.CssAddList("current", classes...)
//     * Css has a feature that allows you to easily change the list of css classes of an html tag,
//       with the functions CssAddList(), CssToggle() and CssToggleTime();
//     * Is the equivalent of:
//         <... css="name1 name2 nameN">
//         var element = document.getElementById("myDIV");
//         element.classList.toggle("mystyle");
//
// Português:
//
//  Troca a lista de classes css do elemento.
//
//   Entrada:
//     name: nome da lista de classes a ser usada.
//
//   Nota:
//     * Esta função equivale a CssAddList("current", classes...);
//     * Css tem uma funcionalidade que permite trocar a lista de classes css de uma tag html de forma
//       fácil, com as funções CssAddList(), CssToggle() e CssToggleTime();
//     * Equivale a:
//         <... css="name1 name2 nameN">
//         var element = document.getElementById("myDIV");
//         element.classList.toggle("mystyle");
func (e *Tag) CssToggle(name string) (ref *Tag) {
	ok := e.css.Toggle(name)
	if ok == false {
		log.Print(KCssListNameNotFound, name)
	}

	js.Global().Get("document").Call("getElementById", e.id).Set("classList", e.css.String())

	return e
}

// CssToggleTime
//
// English:
//
//  Swap the element's css list of classes.
//
//   Input:
//     interval: time interval between toggle;
//     name: names of the lists to be used in the toggle.
//
//   Note:
//     * This function is equivalent to html.CssAddList("current", classes...)
//     * Css has a feature that allows you to easily change the list of css classes of an html tag,
//       with the functions CssAddList(), CssToggle() and CssToggleTime();
//     * Is the equivalent of:
//         <... css="name1 name2 nameN">
//         var element = document.getElementById("myDIV");
//         element.classList.toggle("mystyle");
//
// Português:
//
//  Troca a lista de classes css do elemento.
//
//   Entrada:
//     interval: intervalo de tempo entre as trocas;
//     list: nome das listas a serem usadas na trocas.
//
//   Nota:
//     * Esta função equivale a CssAddList("current", classes...);
//     * Css tem uma funcionalidade que permite trocar a lista de classes css de uma tag html de forma
//       fácil, com as funções CssAddList(), CssToggle() e CssToggleTime();
//     * Equivale a:
//         <... css="name1 name2 nameN">
//         var element = document.getElementById("myDIV");
//         element.classList.toggle("mystyle");
func (e *Tag) CssToggleTime(interval time.Duration, list ...string) (ref *Tag) {
	e.cssInterval = interval
	e.cssList = append(e.cssList, list...)
	e.cssMax = len(list) - 1
	e.cssDone = make(chan struct{})

	return e
}

// CssToggleList
//
// English:
//
//  Defines a new list.
//
//   Input:
//     name: names of the lists to be used in the toggle.
//
//   Note:
//     * This function is equivalent to html.CssAddList("current", classes...)
//     * Css has a feature that allows you to easily change the list of css classes of an html tag,
//       with the functions CssAddList(), CssToggle() and CssToggleTime();
//     * Is the equivalent of:
//         <... css="name1 name2 nameN">
//         var element = document.getElementById("myDIV");
//         element.classList.toggle("mystyle");
//
// Português:
//
//  Define uma nova lista.
//
//   Entrada:
//     list: nome das listas a serem usadas na trocas.
//
//   Nota:
//     * Esta função equivale a CssAddList("current", classes...);
//     * Css tem uma funcionalidade que permite trocar a lista de classes css de uma tag html de forma
//       fácil, com as funções CssAddList(), CssToggle() e CssToggleTime();
//     * Equivale a:
//         <... css="name1 name2 nameN">
//         var element = document.getElementById("myDIV");
//         element.classList.toggle("mystyle");
func (e *Tag) CssToggleList(list ...string) (ref *Tag) {
	e.cssList = append(e.cssList, list...)
	e.cssMax = len(list) - 1
	e.cssDone = make(chan struct{})

	return e
}

// CssToggleStop
//
// English:
//
//  Breaks the toggle loop between css classes.
//
//   Note:
//     * This function is equivalent to html.CssAddList("current", classes...)
//     * Css has a feature that allows you to easily change the list of css classes of an html tag,
//       with the functions CssAddList(), CssToggle() and CssToggleTime();
//     * Is the equivalent of <... css="name1 name2 nameN">
//
// Português:
//
//  Interrompe o laço de trocas entre classes css.
//
//   Nota:
//     * Esta função equivale a CssAddList("current", classes...);
//     * Css tem uma funcionalidade que permite trocar a lista de classes css de uma tag html de forma
//       fácil, com as funções CssAddList(), CssToggle() e CssToggleTime();
//     * Equivale a <... css="name1 name2 nameN">
func (e *Tag) CssToggleStop() (ref *Tag) {
	e.cssDone <- struct{}{}
	return e
}

// CssToggleStartInterval
//
// English:
//
//  Start css class toggle functionality.
//
//   Entrada:
//     interval: intervalo de tempo entre as trocas.
//
//   Note:
//     * This function is equivalent to html.CssAddList("current", classes...)
//     * Css has a feature that allows you to easily change the list of css classes of an html tag,
//       with the functions CssAddList(), CssToggle() and CssToggleTime();
//     * Is the equivalent of <... css="name1 name2 nameN">
//
// Português:
//
//  Reinicializa a funcionalidade de troca de classes css após usar a função CssToggleStop().
//
//   Entrada:
//     interval: intervalo de tempo entre as trocas.
//
//   Nota:
//     * Esta função equivale a CssAddList("current", classes...);
//     * Css tem uma funcionalidade que permite trocar a lista de classes css de uma tag html de forma
//       fácil, com as funções CssAddList(), CssToggle() e CssToggleTime();
//     * Equivale a <... css="name1 name2 nameN">
func (e *Tag) CssToggleStartInterval(interval time.Duration) (ref *Tag) {
	e.cssInterval = interval
	e.CssToggleStart()

	return e
}

// CssToggleLoop
//
// English:
//
//  Defines a finite number of interactions.
//
//   Input:
//     loop: number of interactions.
//
//   Note:
//     * This function is equivalent to html.CssAddList("current", classes...)
//     * Css has a feature that allows you to easily change the list of css classes of an html tag,
//       with the functions CssAddList(), CssToggle() and CssToggleTime();
//     * Is the equivalent of <... css="name1 name2 nameN">
//
// Português:
//
//  Define um número finito de interações.
//
//   Entrada:
//     loop: número de interações.
//
//   Nota:
//     * Esta função equivale a CssAddList("current", classes...);
//     * Css tem uma funcionalidade que permite trocar a lista de classes css de uma tag html de forma
//       fácil, com as funções CssAddList(), CssToggle() e CssToggleTime();
//     * Equivale a <... css="name1 name2 nameN">
func (e *Tag) CssToggleLoop(loop int) (ref *Tag) {
	e.cssLoop = loop
	e.cssLoopFlag = true

	return e
}

// CssOnLoopEnd
//
// English:
//
//  Defines the name of the css list to be used at the end of the loop.
//
//   Input:
//     name: css list name.
//
//   Note:
//     * This function is equivalent to html.CssAddList("current", classes...)
//     * Css has a feature that allows you to easily change the list of css classes of an html tag,
//       with the functions CssAddList(), CssToggle() and CssToggleTime();
//     * Is the equivalent of <... css="name1 name2 nameN">
//
// Português:
//
//  Define o nome da lista css a ser usada no final do loop.
//
//   Entrada:
//     name: nome da lista css.
//
//   Nota:
//     * Esta função equivale a CssAddList("current", classes...);
//     * Css tem uma funcionalidade que permite trocar a lista de classes css de uma tag html de forma
//       fácil, com as funções CssAddList(), CssToggle() e CssToggleTime();
//     * Equivale a <... css="name1 name2 nameN">
func (e *Tag) CssOnLoopEnd(name string) (ref *Tag) {
	e.cssOnLoopEnd = name

	return e
}

// CssToggleStart
//
// English:
//
//  Start css class toggle functionality.
//
//   Note:
//     * This function is equivalent to html.CssAddList("current", classes...)
//     * Css has a feature that allows you to easily change the list of css classes of an html tag,
//       with the functions CssAddList(), CssToggle() and CssToggleTime();
//     * Is the equivalent of <... css="name1 name2 nameN">
//
// Português:
//
//  Reinicializa a funcionalidade de troca de classes css após usar a função CssToggleStop().
//
//   Nota:
//     * Esta função equivale a CssAddList("current", classes...);
//     * Css tem uma funcionalidade que permite trocar a lista de classes css de uma tag html de forma
//       fácil, com as funções CssAddList(), CssToggle() e CssToggleTime();
//     * Equivale a <... css="name1 name2 nameN">
func (e *Tag) CssToggleStart() (ref *Tag) {

	e.cssTimeout = time.NewTicker(e.cssInterval)
	go func(e *Tag) {
		for {
			select {
			case <-e.cssDone:
				e.cssTimeout.Stop()
				return

			case <-e.cssTimeout.C:

				name := e.cssList[e.cssCounter]

				ok := e.css.Toggle(name)
				if ok == false {
					log.Print(KCssListNameNotFound, name)
				}

				js.Global().Get("document").Call("getElementById", e.id).Set("classList", e.css.String())

				if e.cssCounter == e.cssMax {
					e.cssCounter = 0
				} else {
					e.cssCounter += 1
				}

				if e.cssLoopFlag == true {
					if e.cssLoop == 0 {

						e.cssLoopFlag = false
						e.cssTimeout.Stop()
						if e.cssOnLoopEnd != "" {
							e.CssToggle(e.cssOnLoopEnd)
						}
						return

					} else {
						e.cssLoop -= 1
					}

				}
			}
		}
	}(e)

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
//         CssAddList("red", "user", "red").
//         CssAddList("yellow", "user", "yellow").
//         CssAddList("normal", "user").
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
//         CssAddList("red", "user", "red").
//         CssAddList("yellow", "user", "yellow").
//         CssAddList("normal", "user").
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

	return e
}
