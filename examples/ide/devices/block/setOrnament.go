package block

import "github.com/helmutkemper/webassembly/examples/ide/ornament"

// SetOrnament Sets the ornament draw object
//
//	ornament draw object is the instance in charge of making the SVG design of the device
func (e *Block) SetOrnament(ornament ornament.Draw) {
	e.ornament = ornament
}
