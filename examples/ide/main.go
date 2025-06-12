package main

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/examples/ide/devices"
	"github.com/helmutkemper/webassembly/examples/ide/wireFrame"
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

	manager := new(wireFrame.Manager)
	manager.SetFatherId("graphicGopherIde")
	manager.Init()

	done := make(chan struct{})
	<-done
}
