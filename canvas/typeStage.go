package canvas

import (
	iotmaker_santa_isabel_theater_interfaces "github.com/helmutkemper/iotmaker.santa_isabel_theater.interfaces"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/engine"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/mouse"
)

type Index struct {
	Id        string
	Index     int
	Collision func(x, y int) bool
}

// todo: density
type Stage struct {
	Engine    engine.IEngine
	IndexList []Index
	Canvas
	ScratchPad  Canvas
	Cache       Canvas
	CacheEnable bool
	Width       int
	Height      int
	Id          string

	// drag gera o evento de mouse out, porém, não gera o vento de mouse enter
	// por isto de haver este flag
	cursorIsVisible bool

	SetCursor      func(cursorType mouse.CursorType)
	cursorStageId  string
	cursorDrawFunc func()
}

func (el *Stage) AddEngine(engine engine.IEngine) {
	el.Engine = engine
}

func (el *Stage) GetEngine() engine.IEngine {
	return el.Engine
}

func (el *Stage) AddToIndexList(id string, index int, collision func(x, y int) bool) {
	if len(el.IndexList) == 0 {
		el.IndexList = make([]Index, 0)
	}

	el.IndexList = append([]Index{{Id: id, Index: index, Collision: collision}}, el.IndexList...)
}

func (el *Stage) RemoveFromIndexList(id string) {
	for k, v := range el.IndexList {
		if v.Id == id {
			el.IndexList = append(el.IndexList[:k], el.IndexList[k+1:]...)
		}
	}
}

func (el *Stage) IsDraggable(x, y int) string {
	for _, v := range el.IndexList {

		if v.Collision(x, y) == true {
			return v.Id
		}
	}

	return ""
}

func (el *Stage) AddToDraw(sprite iotmaker_santa_isabel_theater_interfaces.ISpriteDraw) (string, int) {
	_, index := el.Engine.DrawAddToFunctions(sprite.Draw)
	el.AddToIndexList(sprite.GetId(), index, sprite.GetCollisionBox)
	return sprite.GetId(), index
}

func (el *Stage) RemoveFromDraw(id string) {
	el.Engine.DrawDeleteFromFunctions(id)
}

func (el *Stage) AddToCalc(f func()) (string, int) {
	return el.Engine.MathAddToFunctions(f)
}

func (el *Stage) RemoveFromCalc(id string) {
	el.Engine.MathDeleteFromFunctions(id)
}

func (el *Stage) AddToHighLatency(f func()) (string, int) {
	return el.Engine.HighLatencyAddToFunctions(f)
}

func (el *Stage) RemoveFromHighLatency(id string) {
	el.Engine.HighLatencyDeleteFromFunctions(id)
}

func (el *Stage) AddToSystem(f func()) (string, int) {
	return el.Engine.SystemAddToFunctions(f)
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

func (el *Stage) SetWidth(width int) {
	el.Canvas.SetWidth(width)
	el.ScratchPad.SetWidth(width)
	el.Cache.SetWidth(width)
	el.Width = width
}

func (el *Stage) SetHeight(height int) {
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
