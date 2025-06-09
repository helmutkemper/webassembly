package block

import "github.com/helmutkemper/webassembly/examples/ide/utils"

// SetID Define the device's div ID
func (e *Block) SetID(id string) (err error) {
	e.id, err = utils.VerifyUniqueId(id)
	return
}
