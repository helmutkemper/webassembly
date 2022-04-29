package _global

import (
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/browser/css"
	"strconv"
	"syscall/js"
)

// SetXY
//
// English:
//
//  Sets the X and Y axes in pixels.
//
// Português:
//
//  Define os eixos X e Y em pixels.
func (e *GlobalAttributes) SetXY(x, y int) (ref *GlobalAttributes) {
	px := strconv.FormatInt(int64(x), 10) + "px"
	py := strconv.FormatInt(int64(y), 10) + "px"

	e.selfElement.Get("style").Set("left", px)
	e.selfElement.Get("style").Set("top", py)

	return e
}

// SetX
//
// English:
//
//  Sets the X axe in pixels.
//
// Português:
//
//  Define o eixo X em pixels.
func (e *GlobalAttributes) SetX(x int) (ref *GlobalAttributes) {
	px := strconv.FormatInt(int64(x), 10) + "px"
	e.selfElement.Get("style").Set("left", px)

	return e
}

// SetY
//
// English:
//
//  Sets the Y axe in pixels.
//
// Português:
//
//  Define o eixo Y em pixels.
func (e *GlobalAttributes) SetY(y int) (ref *GlobalAttributes) {
	py := strconv.FormatInt(int64(y), 10) + "px"
	e.selfElement.Get("style").Set("top", py)

	return e
}

// GetXY
//
// English:
//
//  Returns the X and Y axes in pixels.
//
// Português:
//
//  Retorna os eixos X e Y em pixels.
func (e *GlobalAttributes) GetXY() (x, y int) {
	x = e.selfElement.Get("style").Get("left").Int()
	y = e.selfElement.Get("style").Get("top").Int()

	return
}

// GetX
//
// English:
//
//  Returns the X axe in pixels.
//
// Português:
//
//  Retorna o eixo X em pixels.
func (e *GlobalAttributes) GetX() (x int) {
	x = e.selfElement.Get("style").Get("left").Int()

	return
}

// GetY
//
// English:
//
//  Returns the Y axe in pixels.
//
// Português:
//
//  Retorna o eixo Y em pixels.
func (e *GlobalAttributes) GetY() (y int) {
	y = e.selfElement.Get("style").Get("top").Int()

	return
}

const (
	KIdToAppendNotFound    = "html.AppendById().error: id to append not found:"
	KNewElementIsUndefined = "div.NewDiv().error: new element is undefined:"
)

type GlobalAttributes struct {
	tag         Tag
	id          string
	selfElement js.Value
	cursor      mouse.CursorType
	cssClass    *css.Class
}
