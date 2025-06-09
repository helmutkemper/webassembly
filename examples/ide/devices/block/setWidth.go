package block

import "fmt"

// SetWidth Defines the width property of the device
func (e *Block) SetWidth(width int) {
	if !e.initialized {
		e.width = width
		return
	}

	e.block.AddStyle("width", fmt.Sprintf("%dpx", width))
}
