package globalEngine

import "github.com/helmutkemper/iotmaker.webassembly/platform/engine"

var Engine *engine.Engine

func init() {
	Engine = &engine.Engine{}
	Engine.Init()
}
