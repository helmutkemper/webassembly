package canvas

import (
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/engine"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/mouse"
)

// todo: density
type Stage struct {
	Engine engine.IEngine
	Canvas
	ScratchPad  Canvas
	Cache       Canvas
	CacheEnable bool
	Width       float64
	Height      float64
	Id          string

	// drag gera o evento de mouse out, porém, não gera o vento de mouse enter
	// por isto de haver este flag
	cursorIsVisible bool

	SetCursor      func(cursorType mouse.CursorType)
	cursorStageId  string
	cursorDrawFunc func()
}

func (el *Stage) AddToDraw(f func()) {
	el.Engine.DrawAddToFunctions(f)
}

func (el *Stage) RemoveFromDraw(id string) {
	el.Engine.DrawDeleteFromFunctions(id)
}

func (el *Stage) AddToCalc(f func()) {
	el.Engine.MathAddToFunctions(f)
}

func (el *Stage) RemoveFromCalc(id string) {
	el.Engine.MathDeleteFromFunctions(id)
}

func (el *Stage) AddToHighLatency(f func()) {
	el.Engine.HighLatencyAddToFunctions(f)
}

func (el *Stage) RemoveFromHighLatency(id string) {
	el.Engine.HighLatencyDeleteFromFunctions(id)
}

func (el *Stage) AddToSystem(f func()) {
	el.Engine.SystemAddToFunctions(f)
}

func (el *Stage) RemoveFromSystem(id string) {
	el.Engine.SystemDeleteFromFunctions(id)
}

func (el *Stage) CursorHide() {
	if el.cursorIsVisible == false {
		el.CursorShow()
		return
	}

	el.cursorIsVisible = false
	el.Engine.CursorRemoveDrawFunction(el.cursorStageId)
}

func (el *Stage) CursorShow() {
	el.cursorIsVisible = true
	el.cursorStageId = el.Engine.CursorAddDrawFunction(el.cursorDrawFunc)
}

func (el *Stage) SetCursorDrawFunc(function func()) {
	el.cursorDrawFunc = function
}

func (el *Stage) SetCursorStageId(id string) {
	el.cursorStageId = id
}

func (el *Stage) SetWidth(width float64) {
	el.Canvas.SetWidth(width)
	el.ScratchPad.SetWidth(width)
	el.Cache.SetWidth(width)
	el.Width = width
}

func (el *Stage) SetHeight(height float64) {
	el.Canvas.SetHeight(height)
	el.ScratchPad.SetHeight(height)
	el.Cache.SetHeight(height)
	el.Height = height
}

func (el *Stage) Clear() {
	el.ClearRect(0, 0, el.Width, el.Height)
	//if el.CacheEnable == true {
	//el.ScratchPad.DrawImage(el.Cache.GetCanvas(), 0, 0)
	//}
	el.DrawImage(el.ScratchPad.GetCanvas(), 0, 0)
	el.ScratchPad.ClearRect(0, 0, el.Width, el.Height)
}
