package block

import "github.com/helmutkemper/webassembly/browser/html"

// GetIdeStage Returns to Div where IDE is drawn
func (e *Block) GetIdeStage() (ideStage *html.TagDiv) {
	return e.ideStage
}
