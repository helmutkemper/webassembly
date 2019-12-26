package canvas

// todo: density
type Stage struct {
	Canvas
	ScratchPad Canvas
	Density    float64
	Width      float64
	Height     float64
}

func (el *Stage) SetWidth(width float64) {
	el.Width = width
}

func (el *Stage) SetHeight(height float64) {
	el.Height = height
}

func (el *Stage) Clear() {
	el.ClearRect(0, 0, el.Width, el.Height)
}
