package main

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/examples/ide/devices"
	"github.com/helmutkemper/webassembly/examples/ide/ornament"
)

func main() {

	// graphicGopherIde

	stmLoop := new(devices.StatementLoop)
	stmLoop.SetPosition(50, 50)
	_ = stmLoop.Init()
	//stmLoop.SetWarning(true)

	stmAdd := new(devices.StatementAdd)
	stmAdd.SetPosition(200, 200)
	_ = stmAdd.Init()

	stage := factoryBrowser.NewStage()
	stage.Append(stmLoop.Get())
	stage.Append(stmAdd.Get())

	wire := new(ornament.WireFrame)
	wire.Init()

	spider := new(ornament.Connections)
	spider.SetOrnament(wire)
	spider.SetFatherId("graphicGopherIde")
	spider.Init()

	done := make(chan struct{})
	<-done
}
