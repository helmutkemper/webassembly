package html

import (
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/browserMouse"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/css"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/globalDocument"
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
	tag         Tag
	id          string
	selfElement js.Value
	document    globalDocument.Document
	cursor      browserMouse.CursorType
	cssClass    *css.Class
}
