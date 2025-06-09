package block

// SetDragEnabled Enables the device's drag tool
func (e *Block) SetDragEnabled(enabled bool) {
	e.dragEnabled = enabled

	if !e.initialized {
		return
	}

	if e.dragBlocked {
		e.dragEnabled = false
	}

	e.dragCursorChange()

	if e.dragEnabled {
		e.SetResize(true)
	}

	if e.dragEnabled && e.selected {
		e.SetSelected(false)
	}
}
