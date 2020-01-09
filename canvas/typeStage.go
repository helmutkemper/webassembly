package canvas

import (
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/mouse"
)

// todo: density
type Stage struct {
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

	addCursorRunnerFunc     func(func()) string
	addToRunnerFunc         func(func()) string
	addToCacheRunnerFunc    func(func()) string
	addToRunnerPriorityFunc func(func()) string
	addLowLatencyFunc       func(func()) string

	deleteCursorFromRunnerFunc   func(string)
	deleteFromRunnerFunc         func(string)
	deleteFromCacheRunnerFunc    func(string)
	deleteFromRunnerPriorityFunc func(string)
	deleteLowLatencyFunc         func(string)

	fpsFunc      func(int)
	fpsCacheFunc func(int)
}

func (el *Stage) SetFps(value int) {
	if el.fpsFunc == nil {
		return
	}

	el.fpsFunc(value)
}

func (el *Stage) SetCacheFps(value int) {
	if el.fpsCacheFunc == nil {
		return
	}

	el.fpsCacheFunc(value)
}

func (el *Stage) AddToFpsFunc(function func(int)) {
	el.fpsFunc = function
}

func (el *Stage) AddToFpsCacheFunc(function func(int)) {
	el.fpsCacheFunc = function
}

func (el *Stage) AddCursorRunnerFunc(function func(func()) string) {
	el.addCursorRunnerFunc = function
}

func (el *Stage) DeleteCursorFromRunnerFunc(function func(string)) {
	el.deleteCursorFromRunnerFunc = function
}

func (el *Stage) AddToRunnerFunc(function func(func()) string) {
	el.addToRunnerFunc = function
}

func (el *Stage) DeleteFromRunnerFunc(function func(string)) {
	el.deleteFromRunnerFunc = function
}

func (el *Stage) AddToCacheRunnerFunc(function func(func()) string) {
	el.addToCacheRunnerFunc = function
}

func (el *Stage) DeleteFromCacheRunnerFunc(function func(string)) {
	el.deleteFromCacheRunnerFunc = function
}

func (el *Stage) AddToRunnerPriorityFunc(function func(func()) string) {
	el.addToRunnerPriorityFunc = function
}

func (el *Stage) DeleteFromRunnerPriorityFunc(function func(string)) {
	el.deleteFromRunnerPriorityFunc = function
}

func (el *Stage) AddLowLatencyFunc(function func(func()) string) {
	el.addLowLatencyFunc = function
}

func (el *Stage) DeleteLowLatencyFunc(function func(string)) {
	el.deleteLowLatencyFunc = function
}

func (el *Stage) CursorHide() {
	if el.cursorIsVisible == false {
		el.CursorShow()
		return
	}

	el.cursorIsVisible = false
	el.deleteCursorFromRunnerFunc(el.cursorStageId)
}

func (el *Stage) CursorShow() {
	el.cursorIsVisible = true
	el.cursorStageId = el.addCursorRunnerFunc(el.cursorDrawFunc)
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

func (el *Stage) AddWidthLowLatency(drawFunc func()) string {
	if el.addLowLatencyFunc == nil {
		return ""
	}

	return el.addLowLatencyFunc(drawFunc)
}

func (el *Stage) RemoveFromLowLatency(id string) {
	if el.deleteLowLatencyFunc == nil {
		return
	}

	el.deleteLowLatencyFunc(id)
}

func (el *Stage) AddWidthPriority(drawFunc func()) string {
	if el.addToRunnerPriorityFunc == nil {
		return ""
	}

	return el.addToRunnerPriorityFunc(drawFunc)
}

func (el *Stage) RemoveFromPriority(id string) {
	if el.deleteFromRunnerPriorityFunc == nil {
		return
	}

	el.deleteFromRunnerPriorityFunc(id)
}

func (el *Stage) AddToStage(drawFunc func()) string {
	if el.addToRunnerFunc == nil {
		return ""
	}

	return el.addToRunnerFunc(drawFunc)
}

func (el *Stage) RemoveFromStage(id string) {
	if el.deleteFromRunnerFunc == nil {
		return
	}

	el.deleteFromRunnerFunc(id)
}

func (el *Stage) AddCursor(drawFunc func()) string {
	if el.addCursorRunnerFunc == nil {
		return ""
	}

	el.cursorIsVisible = true
	return el.addCursorRunnerFunc(drawFunc)
}

func (el *Stage) RemoveCursor(id string) {
	if el.deleteCursorFromRunnerFunc == nil {
		return
	}

	el.deleteCursorFromRunnerFunc(id)
}

func (el *Stage) AddToStageAndCache(drawFunc func()) string {
	if el.addToCacheRunnerFunc == nil {
		return ""
	}

	return el.addToCacheRunnerFunc(drawFunc)
}

func (el *Stage) RemoveFromStageAndCache(id string) {
	if el.deleteFromCacheRunnerFunc == nil {
		return
	}

	el.deleteFromCacheRunnerFunc(id)
}
