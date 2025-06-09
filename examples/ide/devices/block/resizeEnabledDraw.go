package block

// resizeEnabledDraw Show/hide the resizing blocks on the screen
func (e *Block) resizeEnabledDraw() {
	if !e.initialized {
		return
	}

	e.resizerTopLeft.AddStyleConditional(e.resizeEnabled, "display", "block", "none")
	e.resizerTopRight.AddStyleConditional(e.resizeEnabled, "display", "block", "none")
	e.resizerBottomLeft.AddStyleConditional(e.resizeEnabled, "display", "block", "none")
	e.resizerBottomRight.AddStyleConditional(e.resizeEnabled, "display", "block", "none")
}
