package block

// SetPosition Defines the coordinates (x, y) of the device
func (e *Block) SetPosition(x, y int) {
	e.SetX(x)
	e.SetY(y)
}
