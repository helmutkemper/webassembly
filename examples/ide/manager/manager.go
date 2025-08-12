package manager

import (
	"github.com/helmutkemper/webassembly/examples/ide/interfaces"
	"github.com/helmutkemper/webassembly/examples/ide/rulesDensity"
	"log"
	"math"
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

	inverter         float64
	delayMS          float64
	durationMS       float64
	start            time.Time
	statusOpening    int
	percentageOfTime float64
	size             float64
	x                rulesDensity.Density
	y                rulesDensity.Density
	width            rulesDensity.Density
	height           rulesDensity.Density
}

func (e *RegisterIcon) Init() {
	e.SetOpening(1)
	e.SetDelay(5 * 1000.0)
	e.SetDuration(200.0)
	e.SetWidth(100)
	e.SetHeight(100)

	log.Printf("init name: %v, %v", e.GetIconName(), e.GetIconCategory())
}

func (e *RegisterIcon) calculateSize() {
	since := time.Since(e.start)
	e.percentageOfTime = math.Min(math.Max(float64(since.Milliseconds())-e.delayMS, 0)/e.durationMS, 1.0)
}

func (e *RegisterIcon) SetDelay(delay float64) {
	e.delayMS = delay
}

func (e *RegisterIcon) SetDuration(duration float64) {
	e.durationMS = duration
}

func (e *RegisterIcon) SetX(x rulesDensity.Density) {
	e.x = x
}

func (e *RegisterIcon) GetX() (x int) {
	return int((e.width.GetFloat() - e.width.GetFloat()*e.getPercentageOfTime()) / 2.0)
}

func (e *RegisterIcon) SetY(y rulesDensity.Density) {
	e.y = y
}

func (e *RegisterIcon) GetY() (y int) {
	return int((e.height.GetFloat() - e.height.GetFloat()*e.getPercentageOfTime()) / 2.0)
}

func (e *RegisterIcon) SetWidth(width rulesDensity.Density) {
	e.width = width
}

func (e *RegisterIcon) GetWidth() (width int) {
	e.calculateSize()
	return int(e.width.GetFloat() * e.getPercentageOfTime())
}

func (e *RegisterIcon) SetHeight(height rulesDensity.Density) {
	e.height = height
}

func (e *RegisterIcon) GetHeight() (height int) {
	e.calculateSize()
	return int(e.height.GetFloat() * e.getPercentageOfTime())
}

func (e *RegisterIcon) getPercentageOfTime() (percent float64) {
	return e.inverter - e.percentageOfTime
}

func (e *RegisterIcon) SetSize(size float64) {
	e.size = size
}

func (e *RegisterIcon) SetOpening(statusOpening int) {
	if e.statusOpening == statusOpening {
		return
	}

	e.start = time.Now()
	e.statusOpening = statusOpening

	if statusOpening == 1 {
		e.inverter = 0.0
	} else if statusOpening == -1 {
		e.inverter = 1.0
	}
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
	e.calculateSize()

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
		icon.Init()
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
