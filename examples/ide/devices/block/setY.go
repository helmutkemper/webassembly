package block

import "fmt"

// SetY Set the coordinate X of the browser screen
func (e *Block) SetY(y int) {
	if !e.initialized {
		e.y = y
		return
	}

	e.block.AddStyle("top", fmt.Sprintf("%dpx", y))
}
