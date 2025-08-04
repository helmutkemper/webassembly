package manager

import (
	"github.com/helmutkemper/webassembly/examples/ide/interfaces"
	"github.com/helmutkemper/webassembly/examples/ide/rulesDensity"
	"log"
	"syscall/js"
	"time"
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

type IconStatus int

const (
	KPipeLineNormal IconStatus = iota
	KPipeLineDisabled
	KPipeLineSelected
	KPipeLineAttention1
	KPipeLineAttention2
	KPipeLineAlert
)

type RegisterIcon struct {
	status   IconStatus
	icon     []js.Value
	name     string
	category string
	time     time.Time
}

func (e *RegisterIcon) SetStatus(status int) {
	e.status = IconStatus(status)
}

func (e *RegisterIcon) GetStatus() (staus int) {
	return int(e.status)
}

func (e *RegisterIcon) SetName(name string) {
	e.name = name
}

func (e *RegisterIcon) SetCategory(category string) {
	e.category = category
}

func (e *RegisterIcon) SetIcon(icon []js.Value) {
	e.icon = icon
	e.time = time.Now()
}

func (e *RegisterIcon) GetIconName() (name string) {
	return e.name
}

func (e *RegisterIcon) GetIconCategory() (category string) {
	return e.category
}

func (e *RegisterIcon) GetIcon() (icon js.Value) {
	interval := time.Duration(500)
	elapsed := time.Since(e.time)
	cycle := elapsed % (time.Millisecond * 2 * interval)
	switch e.status {
	case KPipeLineAlert:
		if cycle < time.Millisecond*interval {
			return e.icon[KPipeLineAttention1]
		}
		return e.icon[KPipeLineAttention2]
	default:
		return e.icon[e.status]
	}
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
