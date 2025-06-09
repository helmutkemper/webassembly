package block

// GetHeight returns the current height of the device.
func (e *Block) GetHeight() (height int) {
	return e.height
}
