package block

// updateOrnament Resize the device's SVG design
func (e *Block) updateOrnament() (err error) {
	width := e.block.GetOffsetWidth()
	height := e.block.GetOffsetHeight()
	_ = e.ornament.Update(width, height)
	return
}
