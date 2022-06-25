package factoryBrowser

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/stage"
)

func NewStage() (ref *stage.Stage) {
	ref = &stage.Stage{}
	ref.Init()

	return ref
}
