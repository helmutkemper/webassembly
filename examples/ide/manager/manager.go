package manager

import (
	"github.com/helmutkemper/webassembly/examples/ide/interfaces"
	"github.com/helmutkemper/webassembly/examples/ide/rulesDensity"
	"log"
	"syscall/js"
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

// English:
//
// Português:
var Manager *manager

// English:
//
// Português:
func init() {
	Manager = new(manager)
	Manager.init()
}

// English:
//
// Português:
type manager struct {
	bbox     []BBox
	icons    []Icon
	mapIcons map[string]map[string]Icon
}

// English:
//
// Português:
func (e *manager) init() {
	e.bbox = make([]BBox, 0)
	e.icons = make([]Icon, 0)
	e.mapIcons = make(map[string]map[string]Icon)
}

// English:
//
// Português:
func (e *manager) Register(element any) {
	var ok bool
	var bbox BBox
	var icon Icon

	if bbox, ok = element.(BBox); ok {
		e.bbox = append(e.bbox, bbox)
	}

	if icon, ok = element.(Icon); ok {
		e.icons = append(e.icons, icon)
		e.registerIcon(icon)
	}
}

// English:
//
// Português:
func (e *manager) registerIcon(element Icon) {
	category := element.GetIconCategory()
	name := element.GetIconName()

	if e.mapIcons[category] == nil {
		e.mapIcons[category] = make(map[string]Icon)
	}

	e.mapIcons[category][name] = element
}

// English:
//
// Português:
func (e *manager) GetBBox() (elements []BBox) {
	log.Printf("bbox: %+v", e.bbox)
	return e.bbox
}

// English:
//
// Português:
func (e *manager) GetIcons() (elements []Icon) {
	return e.icons
}

// English:
//
// Português:
func (e *manager) GetMapIcons() (iconList map[string]map[string]Icon) {
	return e.mapIcons
}

// English:
//
// Português:
type Control struct {
	Icon    js.Value
	ToStage interfaces.ToStage
}
