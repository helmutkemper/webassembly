package main

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/examples/ide/devices"
	"github.com/helmutkemper/webassembly/examples/ide/utils"
)

func main() {

	sequentialId := new(utils.SequentialId)

	device := new(devices.GenericDevice)
	device.SetSequentialId(sequentialId)
	device.SetPosition(50, 50)
	device.SetSize(400, 400)
	device.SetFatherId("container")
	device.SetName("loop")
	device.Init()

	//wm.SetWarning(true)
	//wm.flashMark()

	stage := factoryBrowser.NewStage()
	stage.Append(device.Get())

	done := make(chan struct{})
	<-done
}
