package canvas

import (
	"github.com/helmutkemper/iotmaker.platform/fps"
)

// todo: density
type Stage struct {
	Canvas
	ScratchPad  Canvas
	Cache       Canvas
	CacheEnable bool
	Density     float64
	Width       float64
	Height      float64
}

func (el *Stage) SetWidth(width float64) {
	el.Width = width
}

func (el *Stage) SetHeight(height float64) {
	el.Height = height
}

func (el *Stage) Clear() {
	el.ClearRect(0, 0, el.Width, el.Height)
	if el.CacheEnable == true {
		el.ScratchPad.DrawImage(el.Cache.GetCanvas(), 0, 0)
	}
	el.DrawImage(el.ScratchPad.GetCanvas(), 0, 0)
	el.ScratchPad.ClearRect(0, 0, el.Width, el.Height)
}

func (el *Stage) Add(drawFunc func()) string {
	return fps.AddToRunner(drawFunc)
}

func (el *Stage) Remove(id string) {
	fps.DeleteFromRunner(id)
}

func (el *Stage) AddToCache(drawFunc func()) string {
	return fps.AddToCacheRunner(drawFunc)
}
