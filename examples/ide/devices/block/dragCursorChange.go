package block

// dragCursorChange Change the cursor when the device is being dragged
func (e *Block) dragCursorChange() {
	if !e.initialized {
		return
	}

	e.block.AddStyleConditional(e.dragEnabled, "cursor", "grab", "")
}
