package block

// SetSelected Defines if the device selection tool is active
func (e *Block) SetSelected(selected bool) {
	e.selected = selected

	if !e.initialized {
		return
	}

	if e.selectBlocked {
		e.selected = false
	}

	e.selectDiv.AddStyleConditional(selected, "display", "block", "none")
	e.SetResize(false)
}
