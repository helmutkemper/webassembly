package html

import (
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/browserMouse"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/css"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/globalDocument"
	"log"
	"strconv"
	"syscall/js"
)

func (e *GlobalAttributes) SetXY(x, y int) (ref *GlobalAttributes) {
	px := strconv.FormatInt(int64(x), 10) + "px"
	py := strconv.FormatInt(int64(y), 10) + "px"

	e.selfElement.Get("style").Set("left", px)
	e.selfElement.Get("style").Set("top", py)

	return e
}

func (e *GlobalAttributes) SetX(x int) (ref *GlobalAttributes) {
	px := strconv.FormatInt(int64(x), 10) + "px"
	e.selfElement.Get("style").Set("left", px)

	return e
}

func (e *GlobalAttributes) SetY(y int) (ref *GlobalAttributes) {
	py := strconv.FormatInt(int64(y), 10) + "px"
	e.selfElement.Get("style").Set("top", py)

	return e
}

func (e *GlobalAttributes) GetXY() (x, y int) {
	x = e.selfElement.Get("style").Get("left").Int()
	y = e.selfElement.Get("style").Get("top").Int()

	return
}

const (
	KIdToAppendNotFound    = "html.AppendById().error: id to append not found:"
	KNewElementIsUndefined = "div.NewDiv().error: new element is undefined:"
)

type GlobalAttributes struct {
	id          string
	selfElement js.Value
	document    globalDocument.Document
	cursor      browserMouse.CursorType
	cssClass    *css.Class
}

// NewDiv //fixme: mover para fábrica?
//
// English:
//
//  Creates a new html DIV element.
//
//   Note:
//     * Div Extends GlobalAttributes
//     * By default, browsers always place a line break before and after the <div> element;
//     * The <div> tag is used as a container for HTML elements - which is then styled with CSS or
//       manipulated with JavaScript;
//     * The <div> tag is easily styled by using the class or id attribute;
//     * Any sort of content can be put inside the <div> tag.
//
//  The <div> tag defines a division or a section in an HTML document.
//
// Português:
//
//  Cria um novo elemento html DIV.
//
//   Nota:
//     * Div estende GlobalAttributes
//     * Por padrão, os navegadores sempre colocam uma quebra de linha antes e depois do elemento
//       <div>;
//     * A tag <div> é usada como um contêiner para elementos HTML - que são estilizados com CSS ou
//       manipulados com JavaScript
//     * A tag <div> é facilmente estilizada usando o atributo class ou id;
//     * Qualquer tipo de conteúdo pode ser colocado dentro da tag <div>.
//
//  A tag <div> define uma divisão ou uma seção em um documento HTML.
func NewDiv(id string) (ref *Div) {
	ref = &Div{}
	ref.id = id
	ref.selfElement = js.Global().Get("document").Call("createElement", "div")
	if ref.selfElement.IsUndefined() == true || ref.selfElement.IsNull() == true {
		log.Print(KNewElementIsUndefined, id)
		return
	}

	ref.selfElement.Set("id", id)

	return ref
}
