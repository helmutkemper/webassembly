package block

// SetSelectBlocked Disables the use of the selection tool
func (e *Block) SetSelectBlocked(blocked bool) {
	e.selectBlocked = blocked
}
