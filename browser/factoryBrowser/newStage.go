package factoryBrowser

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/stage"
	"github.com/helmutkemper/iotmaker.webassembly/platform/globalEngine"
)

func NewStage() (ref *stage.Stage) {
	ref = &stage.Stage{}
	ref.SetEngine(globalEngine.Engine)
	ref.Init()

	return ref
}
