package manager

import (
	"errors"
	"github.com/helmutkemper/webassembly/examples/ide/interfaces"
	"github.com/helmutkemper/webassembly/examples/ide/rulesDensity"
	"syscall/js"
)

type Id interface {
	GetID() (id string)
	GetName() (name string)
}

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
	elements []Icon
	icons    map[string]map[string]Icon
}

func (e *manager) Get() (elements []Icon) {
	return e.elements
}

func (e *manager) init() {
	e.elements = make([]Icon, len(e.elements))
	e.icons = make(map[string]map[string]Icon)
}

func (e *manager) Unregister(element any) (err error) {
	id := element.(Id).GetID()
	for key, value := range e.elements {
		if id == value.(Id).GetID() {
			e.elements = append(e.elements[:key], e.elements[key+1:]...)
			return
		}
	}

	err = errors.New("element not found")
	return
}

func (e *manager) Register(element Icon) {
	e.elements = append(e.elements, element)
	e.RegisterIcon(element)
}

func (e *manager) RegisterIcon(element Icon) {
	category := element.GetIconCategory()
	name := element.GetIconName()

	if e.icons[category] == nil {
		e.icons[category] = make(map[string]Icon)
	}

	e.icons[category][name] = element
}

func (e *manager) GetIcons() (iconList map[string]map[string]Icon) {
	return e.icons
}

type Control struct {
	Icon    js.Value
	ToStage interfaces.ToStage
}
