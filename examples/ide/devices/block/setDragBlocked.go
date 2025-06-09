package block

// SetDragBlocked Disables the use of the drag tool
func (e *Block) SetDragBlocked(blocked bool) {
	e.dragBlocked = blocked
}
