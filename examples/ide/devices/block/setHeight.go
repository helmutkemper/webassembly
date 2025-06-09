package block

import "fmt"

// SetHeight Defines the height property of the device
func (e *Block) SetHeight(height int) {
	if !e.initialized {
		e.height = height
		return
	}

	e.block.AddStyle("height", fmt.Sprintf("%dpx", height))
}
