package managerCollision

import (
	"github.com/helmutkemper/webassembly/examples/ide/manager"
)

var Collision *collision

func init() {
	Collision = new(collision)
}

type collision struct {
}

func (e *collision) Detect(element manager.BBox) (partial, total []manager.BBox) {
	partial = make([]manager.BBox, 0)
	total = make([]manager.BBox, 0)

	aId := element.GetID()
	for _, value := range manager.Manager.Get() {
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

func (e *collision) detectTotal(a, b manager.BBox) bool {
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

func (e *collision) detectOnlyPartial(a, b manager.BBox) bool {
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
