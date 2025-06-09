package block

// SetSize Defines the height and width of the device
func (e *Block) SetSize(width, height int) {
	e.SetWidth(width)
	e.SetHeight(height)
}
