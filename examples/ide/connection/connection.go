package connection

import (
	"fmt"
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/examples/ide/utils"
	"image/color"
	"log"
	"reflect"
	"syscall/js"
)

type Identity struct {
	FatherID  string
	Id        string
	Name      string
	IsInput   bool
	ConnType  reflect.Kind
	IsBlocked bool
}

type Connection struct {
	x                    int
	y                    int
	width                int
	height               int
	color                color.RGBA
	name                 string
	containerId          string
	isInput              bool
	isBlocked            bool
	container            *html.TagDiv
	connection           *html.TagDiv
	mouseArea            *html.TagDiv
	dataType             reflect.Kind
	acceptedNotConnected bool
	spaceAreaVertical    int
	spaceAreaHorizontal  int
	spaceAreaColor       color.RGBA
	autoIDObj            utils.SequentialId
	autoID               string
}

func (e *Connection) SetFather(father *html.TagDiv) {
	e.container = father
	e.containerId = father.GetId()
}

func (e *Connection) SetName(name string) (err error) {
	e.name, err = utils.VerifyName(name)
	return
}

func (e *Connection) SetAsInput() {
	e.isInput = true
}

func (e *Connection) GetAsInput() (isInput bool) {
	return e.isInput
}

func (e *Connection) SetBlocked(isBlocked bool) {
	e.isBlocked = isBlocked

	if e.mouseArea == nil {
		return
	}

	if e.isBlocked {
		e.mouseArea.AddStyle("cursor", "not-allowed")
		return
	}

	e.mouseArea.AddStyle("cursor", "crosshair")
	return
}

func (e *Connection) GetBlocked() (isBlocked bool) {
	return e.isBlocked
}

func (e *Connection) SetDataType(dataType reflect.Kind) {
	e.dataType = dataType
}

func (e *Connection) GetDataType() (dataType reflect.Kind) {
	return e.dataType
}

func (e *Connection) GetIdentity() (identity Identity) {
	return Identity{} // todo: fazer
}

func (e *Connection) SetX(x int) {
	e.x = x
	e.mouseArea.AddStyle("left", fmt.Sprintf("%dpx", x))
}

func (e *Connection) SetY(y int) {
	e.y = y
	e.mouseArea.AddStyle("top", fmt.Sprintf("%dpx", y))
}

func (e *Connection) Create(x, y, width, height int, color color.RGBA) {
	e.x = x
	e.y = y
	e.width = width
	e.height = height
	e.color = color
}

func (e *Connection) Init() {
	id := e.containerId + "_" + e.name

	e.spaceAreaHorizontal = 8
	e.spaceAreaVertical = 8

	e.autoID = utils.GetRandomId()

	e.mouseArea = factoryBrowser.NewTagDiv().
		Id(id+"_space").
		AddStyle("position", "absolute").
		AddStyle("left", fmt.Sprintf("%dpx", e.x-e.spaceAreaHorizontal)).
		AddStyle("top", fmt.Sprintf("%dpx", e.y-e.spaceAreaVertical)).
		AddStyle("width", fmt.Sprintf("%dpx", e.width+2*e.spaceAreaHorizontal)).
		AddStyle("height", fmt.Sprintf("%dpx", e.height+2*e.spaceAreaVertical)).
		AddStyle("backgroundColor", e.spaceAreaColor)

	e.connection = factoryBrowser.NewTagDiv().
		Id(id+"_connection").
		AddStyle("position", "absolute").
		AddStyle("left", fmt.Sprintf("%dpx", e.spaceAreaHorizontal)).
		AddStyle("top", fmt.Sprintf("%dpx", e.spaceAreaVertical)).
		AddStyle("width", fmt.Sprintf("%dpx", e.width)).
		AddStyle("height", fmt.Sprintf("%dpx", e.height)).
		AddStyle("backgroundColor", e.color)
	e.mouseArea.Append(e.connection)

	e.mouseArea.Get().Call("addEventListener", "click", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		log.Printf("%+v", e.GetIdentity())
		return nil
	}))

	e.container.Append(e.mouseArea)

	if e.isBlocked {
		e.mouseArea.AddStyle("cursor", "not-allowed")
		return
	}

	e.mouseArea.AddStyle("cursor", "crosshair")
}
