package manager

import (
	"errors"
	"github.com/helmutkemper/webassembly/examples/ide/rulesDensity"
)

type BBox interface {
	GetID() (id string)
	GetName() (name string)
	GetWidth() (width rulesDensity.Density)
	GetHeight() (height rulesDensity.Density)
	GetX() (x rulesDensity.Density)
	GetY() (y rulesDensity.Density)

	SetX(x rulesDensity.Density)
	SetY(y rulesDensity.Density)
	SetWidth(width rulesDensity.Density)
	SetHeight(height rulesDensity.Density)

	SetSelected(selected bool)
	GetSelected() (selected bool)

	SetDragEnable(enabled bool)
	GetDragEnable() (enabled bool)

	GetResizeEnable() (enabled bool)
	SetResizeEnable(enabled bool)

	GetZIndex() (zIndex int)
}

type Flags interface {
	GetInitialized() (initialized bool)
	GetWarning() (warning bool)
	GetDragBlocked() (blocked bool)
	GetDragEnable() (enabled bool)
	GetResize() (enabled bool)
	GetResizeBlocked() (blocked bool)
	GetSelectBlocked() (blocked bool)
	GetSelected() (selected bool)
}

var Manager *manager

func init() {
	Manager = new(manager)
	Manager.init()
}

type manager struct {
	elements []BBox
}

func (e *manager) Get() (elements []BBox) {
	return e.elements
}

func (e *manager) init() {
	e.elements = make([]BBox, len(e.elements))
}

func (e *manager) Unregister(element BBox) (err error) {
	id := element.GetID()
	for key, value := range e.elements {
		if id == value.GetID() {
			e.elements = append(e.elements[:key], e.elements[key+1:]...)
			return
		}
	}

	err = errors.New("element not found")
	return
}

func (e *manager) Register(element BBox) {
	e.elements = append(e.elements, element)
}
