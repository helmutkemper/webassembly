package cursor

import "github.com/helmutkemper/webassembly/examples/ide/rulesStage"

var Manager *manager

func init() {
	Manager = new(manager)
	Manager.SetFatherId(rulesStage.KStageId)
	Manager.Init()
}
