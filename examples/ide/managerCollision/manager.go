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

func (e *collision) DetectBoxContained(element manager.BBox) (list []manager.BBox) {
	list = make([]manager.BBox, 0)
	aId := element.GetID()
	for _, value := range manager.Manager.GetBBox() {
		bId := value.GetID()
		if aId == bId {
			continue
		}

		if e.contained(element, value) {
			list = append(list, value)
		}
	}

	return
}

func (e *collision) DetectBoxCollision(element manager.BBox) (list []manager.BBox) {
	list = make([]manager.BBox, 0)

	aId := element.GetID()
	for _, value := range manager.Manager.GetBBox() {
		bId := value.GetID()
		if aId == bId {
			continue
		}

		if e.collision(element, value) {
			list = append(list, value)
		}
	}

	return
}

func (e *collision) DetectBoxCollisionNotContained(element manager.BBox) (list []manager.BBox) {
	list = make([]manager.BBox, 0)

	aId := element.GetID()
	for _, value := range manager.Manager.GetBBox() {
		bId := value.GetID()
		if aId == bId {
			continue
		}

		if e.contained(element, value) {
			continue
		}

		if e.collision(element, value.(manager.BBox)) {
			list = append(list, value.(manager.BBox))
		}
	}

	return
}

// detectTotal Checks if boxA is entirely contained within boxB.
func (e *collision) contained(b, a manager.BBox) bool {
	xA, yA := a.GetX(), a.GetY()
	widthA, heightA := a.GetWidth(), a.GetHeight()

	xB, yB := b.GetX(), b.GetY()
	widthB, heightB := b.GetWidth(), b.GetHeight()

	return xA >= xB && xA+widthA <= xB+widthB &&
		yA >= yB && yA+heightA <= yB+heightB
}

func (e *collision) collision(a, b manager.BBox) bool {
	xA, yA := a.GetX(), a.GetY()
	widthA, heightA := a.GetWidth(), a.GetHeight()

	xB, yB := b.GetX(), b.GetY()
	widthB, heightB := b.GetWidth(), b.GetHeight()

	return xA < xB+widthB && xA+widthA > xB &&
		yA < yB+heightB && yA+heightA > yB
}
