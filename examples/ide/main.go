package main

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/examples/ide/devices"
	"github.com/helmutkemper/webassembly/examples/ide/utils"
)

func main() {

	sequentialId := new(utils.SequentialId)

	stmLoop := new(devices.StatementLoop)
	stmLoop.SetSequentialId(sequentialId)
	stmLoop.SetPosition(50, 50)
	stmLoop.SetFatherId("container")
	_ = stmLoop.SetName("stmLoop")
	_ = stmLoop.Init()

	stmAdd := new(devices.StatementAdd)
	stmAdd.SetSequentialId(sequentialId)
	stmAdd.SetPosition(20, 20)
	stmAdd.SetFatherId("container")
	_ = stmAdd.SetName("stmAdd")
	_ = stmAdd.Init()

	stage := factoryBrowser.NewStage()
	stage.Append(stmLoop.Get())
	stage.Append(stmAdd.Get())

	done := make(chan struct{})
	<-done
}
