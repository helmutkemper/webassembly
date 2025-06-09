package block

import "fmt"

// SetX Define a coordenada x da tela do navegador
func (e *Block) SetX(x int) {
	if !e.initialized {
		e.x = x
		return
	}
	e.block.AddStyle("left", fmt.Sprintf("%dpx", x))
}
