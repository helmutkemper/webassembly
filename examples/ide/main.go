package main

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/examples/ide/devices"
	"github.com/helmutkemper/webassembly/examples/ide/utils"
)

func main() {

	sequentialId := new(utils.SequentialId)

	device := new(devices.StatementLoop)
	device.SetSequentialId(sequentialId)
	device.SetPosition(50, 50)
	device.SetFatherId("container")
	_ = device.SetName("loop")
	_ = device.Init()

	//wm.SetWarning(true)
	//wm.flashMark()

	stage := factoryBrowser.NewStage()
	stage.Append(device.Get())

	done := make(chan struct{})
	<-done
}
