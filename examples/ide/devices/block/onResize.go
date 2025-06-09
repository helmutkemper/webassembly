package block

import "syscall/js"

// OnResize cannot be shadowed by the main instance, so the function in SetOnResize
func (e *Block) OnResize(element js.Value, width, height int) {
	if e.onResizeFunc != nil {
		e.onResizeFunc(element, width, height)
	}
}
