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
	elements []any
}

func (e *manager) Get() (elements []any) {
	return e.elements
}

func (e *manager) init() {
	e.elements = make([]any, len(e.elements))
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

func (e *manager) Register(element any) {
	e.elements = append(e.elements, element)
}

func (e *manager) GetMenu() (menu map[string]map[string]Control) {
	menu = make(map[string]map[string]Control)
	for _, element := range e.elements {
		category := element.(Icon).GetIconCategory()
		name := element.(Icon).GetIconName()
		icon := element.(Icon).GetIcon(false)

		if menu[category] == nil {
			menu[category] = make(map[string]Control)
		}

		menu[category][name] = Control{
			Icon:    icon,
			ToStage: element.(interfaces.ToStage),
		}
	}

	return
}

type Control struct {
	Icon    js.Value
	ToStage interfaces.ToStage
}
