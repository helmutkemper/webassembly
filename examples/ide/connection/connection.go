package connection

import (
	"fmt"
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/examples/ide/utils"
	"github.com/helmutkemper/webassembly/platform/factoryColor"
	"image/color"
	"log"
	"reflect"
	"syscall/js"
)

type Identity struct {
	Father               *html.TagDiv
	Id                   string
	Name                 string
	IsInput              bool
	ConnType             reflect.Kind
	IsBlocked            bool
	AcceptedNotConnected bool
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
	father               *html.TagDiv
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
	e.father = father
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

func (e *Connection) SetAcceptedNotConnected(accept bool) {
	e.acceptedNotConnected = accept
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

	switch e.dataType {
	case reflect.Bool:
		e.color = factoryColor.NewGreen()
	case reflect.Int:
		e.color = factoryColor.NewBlue()
	case reflect.Int8:
		e.color = factoryColor.NewBlue()
	case reflect.Int16:
		e.color = factoryColor.NewBlue()
	case reflect.Int32:
		e.color = factoryColor.NewBlue()
	case reflect.Int64:
		e.color = factoryColor.NewBlue()
	case reflect.Uint:
		e.color = factoryColor.NewBlueViolet()
	case reflect.Uint8:
		e.color = factoryColor.NewBlueViolet()
	case reflect.Uint16:
		e.color = factoryColor.NewBlueViolet()
	case reflect.Uint32:
		e.color = factoryColor.NewBlueViolet()
	case reflect.Uint64:
		e.color = factoryColor.NewBlueViolet()
	case reflect.Uintptr:
		e.color = factoryColor.NewBlueViolet()
	case reflect.Float32:
		e.color = factoryColor.NewYellowGreen()
	case reflect.Float64:
		e.color = factoryColor.NewYellowGreen()
	case reflect.Complex64:
		e.color = factoryColor.NewAntiqueWhite()
	case reflect.Complex128:
		e.color = factoryColor.NewAntiqueWhite()
	case reflect.Array:
		e.color = factoryColor.NewDarkMagenta()
	case reflect.Chan:
	case reflect.Func:
	case reflect.Interface:
	case reflect.Map:
	case reflect.Pointer:
	case reflect.Slice:
		e.color = factoryColor.NewDarkMagenta()
	case reflect.String:
		e.color = factoryColor.NewMediumTurquoise()
	case reflect.Struct:
		e.color = factoryColor.NewGainsboro()
	case reflect.UnsafePointer:
	}

}

func (e *Connection) GetDataType() (dataType reflect.Kind) {
	return e.dataType
}

func (e *Connection) GetIdentity() (identity Identity) {
	return Identity{
		Father:               e.father,
		Id:                   e.containerId,
		Name:                 e.name,
		IsInput:              e.isInput,
		ConnType:             e.dataType,
		IsBlocked:            e.isBlocked,
		AcceptedNotConnected: e.acceptedNotConnected,
	}
}

func (e *Connection) SetX(x int) {
	e.x = x
	e.mouseArea.AddStyle("left", fmt.Sprintf("%dpx", x-e.spaceAreaHorizontal))
}

func (e *Connection) SetY(y int) {
	e.y = y
	e.mouseArea.AddStyle("top", fmt.Sprintf("%dpx", y-e.spaceAreaVertical))
}

func (e *Connection) Create(x, y int) {
	e.x = x
	e.y = y
	e.width = 6
	e.height = 4
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

	e.father.Append(e.mouseArea)

	if e.isBlocked {
		e.mouseArea.AddStyle("cursor", "not-allowed")
		return
	}

	e.mouseArea.AddStyle("cursor", "crosshair")
}
