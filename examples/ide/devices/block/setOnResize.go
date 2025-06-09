package block

import "syscall/js"

// SetOnResize Receives the pointer to a function to be invoked during resizing
//
//	This function is due to the fact that the OnResize function cannot be shadowed by the main instance
func (e *Block) SetOnResize(f func(element js.Value, width, height int)) {
	e.onResizeFunc = f
}
