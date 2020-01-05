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

	SetCursor      func(cursorType mouse.CursorType)
	cursorStageId  string
	cursorDrawFunc func()

	addToRunnerFunc         func(func()) string
	addToCacheRunnerFunc    func(func()) string
	addToRunnerPriorityFunc func(func()) string
	addLowLatencyFunc       func(func()) string

	deleteFromRunnerFunc         func(string)
	deleteFromCacheRunnerFunc    func(string)
	deleteFromRunnerPriorityFunc func(string)
	deleteLowLatencyFunc         func(string)
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
	el.deleteFromRunnerFunc(el.cursorStageId)
}

func (el *Stage) CursorShow() {
	el.cursorStageId = el.addToRunnerFunc(el.cursorDrawFunc)
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
	if el.CacheEnable == true {
		el.ScratchPad.DrawImage(el.Cache.GetCanvas(), 0, 0)
	}
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

func (el *Stage) Add(drawFunc func()) string {
	if el.addToRunnerFunc == nil {
		return ""
	}

	return el.addToRunnerFunc(drawFunc)
}

func (el *Stage) Remove(id string) {
	if el.deleteFromRunnerFunc == nil {
		return
	}

	el.deleteFromRunnerFunc(id)
}

func (el *Stage) AddToCache(drawFunc func()) string {
	if el.addToCacheRunnerFunc == nil {
		return ""
	}

	return el.addToCacheRunnerFunc(drawFunc)
}

func (el *Stage) RemoveFromCache(id string) {
	if el.deleteFromCacheRunnerFunc == nil {
		return
	}

	el.deleteFromCacheRunnerFunc(id)
}
