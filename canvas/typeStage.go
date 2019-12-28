package canvas

import (
	"github.com/helmutkemper/iotmaker.platform/fps"
)

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
	el.DrawImage(el.ScratchPad.GetCanvas(), 0, 0)
	el.ScratchPad.ClearRect(0, 0, el.Width, el.Height)
}

func (el *Stage) Add(drawFunc func()) string {
	return fps.AddToRunner(drawFunc, true)
}

func (el *Stage) AddSync(drawFunc func()) string {
	return fps.AddToRunner(drawFunc, false)
}
