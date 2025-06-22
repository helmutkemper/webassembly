package connection

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/examples/ide/rulesConnection"
	"github.com/helmutkemper/webassembly/examples/ide/rulesDesity"
	"log"
	"strconv"
	"syscall/js"
)

// Connection
//
//	Note: GetError() pega todos os erros de todas as conexões, porém, este código foi feito para ser usado durante o setup da IDE, não foi feito para ser usado pelo usuário final
type Connection struct {
	connection *html.TagSvgPath

	clickFunc    js.Func
	fatherId     string
	name         string
	dataType     string
	notConnected string
	lookedUp     string
	isADataInput string
	x            rulesDesity.Density
	y            rulesDesity.Density
}

type Setup struct {
	FatherId           string
	Name               string
	DataType           string
	AcceptNotConnected bool
	LookedUp           bool
	IsADataInput       bool
	ClickFunc          js.Func
}

func (e Setup) Verify() (err error) {
	return rulesConnection.TypeVerify(e.DataType)
}

func (e *Connection) Init(markEnd string) {
	e.connection = factoryBrowser.NewTagSvgPath().
		DataKey(rulesConnection.KConnectionPrefix+"Name", e.name).
		DataKey(rulesConnection.KConnectionPrefix+"DataType", e.dataType).
		DataKey(rulesConnection.KConnectionPrefix+"AcceptNoConnection", e.notConnected).
		DataKey(rulesConnection.KConnectionPrefix+"LookedUp", e.lookedUp).
		DataKey(rulesConnection.KConnectionPrefix+"IsDataInput", e.isADataInput).
		DataKey(rulesConnection.KConnectionPrefix+"FatherId", e.fatherId).
		Fill("transparent").
		Stroke("none").
		AddStyle("cursor", "crosshair").
		MarkerEnd(markEnd)

	if !e.clickFunc.IsNull() {
		e.connection.Get().Set("getConnData", js.FuncOf(e.getConnectionFunc))
		e.connection.Get().Call("addEventListener", "click", e.clickFunc)
	} else {
		log.Printf("não deveria ter entrado aqui!")
	}
}

func (e *Connection) SetXY(x, y rulesDesity.Density) {
	e.connection.DataKey(rulesConnection.KConnectionPrefix+"Top", strconv.FormatInt(int64(x), 10))
	e.connection.DataKey(rulesConnection.KConnectionPrefix+"Left", strconv.FormatInt(int64(y), 10))
}

func (e *Connection) GetX() (x rulesDesity.Density) {
	xI64, _ := strconv.ParseInt(e.connection.GetData(rulesConnection.KConnectionPrefix+"Top"), 10, 64)
	return rulesDesity.Density(xI64)
}

func (e *Connection) GetY() (y rulesDesity.Density) {
	yI64, _ := strconv.ParseInt(e.connection.GetData(rulesConnection.KConnectionPrefix+"Left"), 10, 64)
	return rulesDesity.Density(yI64)
}

// mapToJsObject Transforms a go map into an js object accepted by wasm go
//
//	Note:
//	  * The js.FuncOf(...).Call(...) receives and return interface{}, but only types that runtime go knows how to
//	    convert to js work correctly.
//	    Valid types for return: int, float64, bool, string, nil, e js.Value.
func (e *Connection) mapToJsObject(data map[string]interface{}) js.Value {
	obj := js.Global().Get("Object").New()
	for k, v := range data {
		obj.Set(k, v)
	}
	return obj
}

func (e *Connection) getConnectionFunc(_ js.Value, _ []js.Value) interface{} {
	ret := map[string]interface{}{
		"FatherId":           e.GetFatherId(),
		"Name":               e.GetName(),
		"DataType":           e.GetDataType(),
		"AcceptNotConnected": e.GetAcceptNotConnected(),
		"LookedUp":           e.GetConnectionLockedUp(),
		"IsADataInput":       e.GetAsDataInput(),
		"Top":                e.GetX(),
		"Left":               e.GetY(),
	}
	return e.mapToJsObject(ret)
}

func (e *Connection) Setup(setup Setup) {
	e.clickFunc = setup.ClickFunc
	e.fatherId = setup.FatherId
	e.name = setup.Name
	e.dataType = setup.DataType

	e.notConnected = "trowError"
	if setup.AcceptNotConnected {
		e.notConnected = "accept"
	}

	e.lookedUp = strconv.FormatBool(setup.LookedUp)

	e.isADataInput = "output"
	if setup.IsADataInput {
		e.isADataInput = "input"
	}
}

//func (e *Connection) SetClickFunc(f js.Func) {
//	e.clickFunc = f
//}

//func (e *Connection) SetFatherId(id string) {
//	e.fatherId = id
//}

func (e *Connection) GetFatherId() (id string) {
	return e.connection.GetData(rulesConnection.KConnectionPrefix + "FatherId")
}

func (e *Connection) GetSvgPath() (svgPath *html.TagSvgPath) {
	return e.connection
}

//func (e *Connection) SetReference(reference *html.TagSvgPath) {
//	e.connection = reference
//}

func (e *Connection) GetAsDataInput() (dataInput bool) {
	if e.connection.GetData(rulesConnection.KConnectionPrefix+"IsDataInput") == "input" {
		return true
	}

	return false
}

func (e *Connection) GetConnectionLockedUp() (lockedUp bool) {
	if e.connection.GetData(rulesConnection.KConnectionPrefix+"LookedUp") == "true" {
		return true
	}

	return false
}

func (e *Connection) GetAcceptNotConnected() (accept bool) {
	if e.connection.GetData(rulesConnection.KConnectionPrefix+"AcceptNoConnection") == "accept" {
		return true
	}

	return false
}

func (e *Connection) GetName() (name string) {
	return e.connection.GetData(rulesConnection.KConnectionPrefix + "Name")
}

func (e *Connection) GetDataType() (connType string) {
	return e.connection.GetData(rulesConnection.KConnectionPrefix + "DataType")
}

func (e *Connection) GetError() (err error) {
	return rulesConnection.GetError()
}
