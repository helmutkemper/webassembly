package factoryBrowser

import (
	"github.com/helmutkemper/webassembly/browser/stage"
	"github.com/helmutkemper/webassembly/platform/globalEngine"
)

func NewStage() (ref *stage.Stage) {
	ref = &stage.Stage{}
	ref.Engine(globalEngine.Engine)
	ref.Init()

	return ref
}
