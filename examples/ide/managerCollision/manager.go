package managerCollision

import (
	"errors"
	"github.com/helmutkemper/webassembly/examples/ide/rulesDensity"
)

var Collision *collision

func init() {
	Collision = new(collision)
	Collision.init()
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

type collision struct {
	elements []BBox

	ruleBook map[string]func()
}

func (e *collision) init() {
	e.elements = make([]BBox, len(e.elements))
}

func (e *collision) Unregister(element BBox) (err error) {
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

func (e *collision) Register(element BBox) {
	e.elements = append(e.elements, element)
}

func (e *collision) Detect(element BBox) (partial, total []BBox) {
	partial = make([]BBox, 0)
	total = make([]BBox, 0)

	aId := element.GetID()
	for _, value := range e.elements {
		bId := value.GetID()
		if aId == bId {
			continue
		}

		if e.detectTotal(element, value) {
			total = append(total, value)
			continue
		}

		if e.detectOnlyPartial(element, value) {
			partial = append(partial, value)
		}
	}

	return
}

func (e *collision) detectTotal(a, b BBox) bool {
	// Calculates the limits of a
	leftA, topA := a.GetX(), a.GetY()
	rightA, bottomA := leftA+a.GetWidth(), topA+a.GetHeight()

	// Calculates the limits of b
	leftB, topB := b.GetX(), b.GetY()
	rightB, bottomB := leftB+b.GetWidth(), topB+b.GetHeight()

	// There is partial or total collision if there is *overlap* both on the x axis and the y axis:
	horizontalOverlap := leftA < rightB && rightA > leftB
	verticalOverlap := topA < bottomB && bottomA > topB

	return horizontalOverlap && verticalOverlap
}

func (e *collision) detectOnlyPartial(a, b BBox) bool {
	if !e.detectTotal(a, b) {
		return false
	}

	// Detects total containment A in B
	if a.GetX() >= b.GetX() &&
		a.GetY() >= b.GetY() &&
		a.GetX()+a.GetWidth() <= b.GetX()+b.GetWidth() &&
		a.GetY()+a.GetHeight() <= b.GetY()+b.GetHeight() {
		return false
	}

	// Detects total containment B in a
	if b.GetX() >= a.GetX() &&
		b.GetY() >= a.GetY() &&
		b.GetX()+b.GetWidth() <= a.GetX()+a.GetWidth() &&
		b.GetY()+b.GetHeight() <= a.GetY()+a.GetHeight() {
		return false
	}

	return true
}
