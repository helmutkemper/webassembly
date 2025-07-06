package interfaces

import (
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/examples/ide/rulesDensity"
	"github.com/helmutkemper/webassembly/examples/ide/rulesStage"
)

type ToStage interface {
	// SetResizerButton(resizeButton block.ResizeButton)
	// SetDraggerButton(draggerButton block.ResizeButton)

	SetGridAdjust(gridAdjust rulesStage.GridAdjust)
	SetMainSvg(svg *html.TagSvg)
	SetPosition(x, y rulesDensity.Density)
	Init() (err error)
}
