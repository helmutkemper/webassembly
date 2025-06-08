package devices

import (
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/examples/ide/connection"
	"github.com/helmutkemper/webassembly/examples/ide/ornament/doubleLoopArrow"
	"github.com/helmutkemper/webassembly/examples/ide/utils"
	"github.com/helmutkemper/webassembly/platform/factoryColor"
	"log"
	"reflect"
	"syscall/js"
)

type GenericDevice struct {
	block Block

	defaultWidth          int
	defaultHeight         int
	horizontalMinimumSize int
	verticalMinimumSize   int
	ornamentDraw          *doubleLoopArrow.DoubleLoopArrow
	id                    string
	connStop              *connection.Connection
	debugMode             bool

	sequentialId utils.SequentialInterface
}

func (e *GenericDevice) Get() (container *html.TagDiv) {
	return e.block.container
}

func (e *GenericDevice) SetFatherId(fatherId string) {
	e.block.SetFatherId(fatherId)
}

func (e *GenericDevice) SetName(name string) {
	e.block.SetName(name)
}

func (e *GenericDevice) SetPosition(x, y int) {
	e.block.SetPosition(x, y)
}

func (e *GenericDevice) SetSize(wight, height int) {
	e.block.SetSize(wight, height)
}

func (e *GenericDevice) SetSequentialId(sequentialId utils.SequentialInterface) {
	e.sequentialId = sequentialId
}

func (e *GenericDevice) Init() (err error) {
	e.defaultWidth = 200
	e.defaultHeight = 200
	e.horizontalMinimumSize = 100
	e.verticalMinimumSize = 100

	if e.block.width == 0 {
		e.block.width = e.defaultWidth
	}

	if e.block.height == 0 {
		e.block.height = e.defaultHeight
	}

	err = e.block.SetName("GenericDevice")
	if err != nil {
		log.Printf("retornou 1")
		return
	}

	e.id, err = e.sequentialId.GetId(e.block.GetName())
	if err != nil {
		log.Printf("retornou 2")
		return
	}

	//e.block.SetPosition(e.x, e.y)
	//e.block.SetSize(e.width, e.height)
	e.block.SetDragEnabled(false)
	e.block.SetResizeEnabled(true) // todo: false
	e.block.SetHorizontalMinimumSize(e.horizontalMinimumSize)
	e.block.SetVerticalMinimumSize(e.verticalMinimumSize)

	e.ornamentDraw = new(doubleLoopArrow.DoubleLoopArrow)
	e.ornamentDraw.SetWarningMarkMargin(20)
	_ = e.ornamentDraw.Init()

	e.block.SetOrnament(e.ornamentDraw)

	_ = e.block.Init()

	e.block.id = e.id

	e.connStop = new(connection.Connection)
	e.connStop.Create(e.block.width-50-4, e.block.height-40-2, 4, 3, factoryColor.NewRed())
	e.connStop.SetFather(e.block.GetElement())
	e.connStop.SetAsInput()
	e.connStop.SetName("stop")
	e.connStop.SetDataType(reflect.Bool)
	//e.connStop.SetAcceptedNotConnected(true)
	e.connStop.SetBlocked(false)
	e.connStop.Init()

	e.block.SetOnResize(func(element js.Value, width, height int) {
		e.connStop.SetX(width - 50 - 8 - 4)
		e.connStop.SetY(height - 40 - 8 - 2)
	})

	return nil
}
