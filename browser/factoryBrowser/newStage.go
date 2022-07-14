package factoryBrowser

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/stage"
	"github.com/helmutkemper/iotmaker.webassembly/platform/engine"
)

func NewStage() (ref *stage.Stage) {
	ref = &stage.Stage{}
	ref.Setengine(&engine.Engine{})
	ref.Init()

	return ref
}
