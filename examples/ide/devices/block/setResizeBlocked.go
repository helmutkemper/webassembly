package block

// SetResizeBlocked Disables the use of the resize tool
func (e *Block) SetResizeBlocked(blocked bool) {
	e.resizeBlocked = blocked
}
