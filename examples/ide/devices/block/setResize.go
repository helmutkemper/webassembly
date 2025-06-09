package block

// SetResize Defines if the device resize tool is active
func (e *Block) SetResize(enabled bool) {
	e.resizeEnabled = enabled

	if !e.initialized {
		return
	}

	if e.resizeBlocked {
		e.resizeEnabled = false
	}

	e.resizeEnabledDraw()
	if enabled && e.selected {
		e.SetSelected(false)
	}
}
