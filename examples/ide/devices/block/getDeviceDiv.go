package block

import "github.com/helmutkemper/webassembly/browser/html"

// GetDeviceDiv Returns the div from device
func (e *Block) GetDeviceDiv() (element *html.TagDiv) {
	return e.block
}
