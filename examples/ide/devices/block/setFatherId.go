package block

import "github.com/helmutkemper/webassembly/browser/factoryBrowser"

// SetFatherId Receives the div ID used as a stage for the IDE and puts it to occupy the entire browser area
func (e *Block) SetFatherId(fatherId string) {
	e.fatherId = fatherId

	e.ideStage = factoryBrowser.NewTagDiv().
		AddStyle("position", "relative").
		AddStyle("width", "100vw").
		AddStyle("height", "100vh")

	e.ideStage.AppendById(fatherId)
}
