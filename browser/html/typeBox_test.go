package html

import (
	"testing"
)

func TestBox_Collision(t *testing.T) {
	boxElement := Box{X: 10, Y: 10, Width: 10, Height: 10}
	element2 := Box{X: 10 - 5, Y: 10, Width: 10, Height: 10}

	var a, b, c, d bool

	a, b, c, d = boxElement.Quadrant(element2)
	if a != boxElement.collisionUpLeft(element2) || a != true {
		t.Fatal("test A, collision A fail")
	}
	if b != boxElement.collisionUpRight(element2) || b != false {
		t.Fatal("test A, collision B fail")
	}
	if c != boxElement.collisionDownLeft(element2) || c != true {
		t.Fatal("test A, collision C fail")
	}
	if d != boxElement.collisionDownRight(element2) || d != false {
		t.Fatal("test A, collision D fail")
	}

	element2 = Box{X: 10 - 4, Y: 10, Width: 10, Height: 10}

	a, b, c, d = boxElement.Quadrant(element2)
	if a != boxElement.collisionUpLeft(element2) || a != true {
		t.Fatal("test B, collision A fail")
	}
	if b != boxElement.collisionUpRight(element2) || b != true {
		t.Fatal("test B, collision B fail")
	}
	if c != boxElement.collisionDownLeft(element2) || c != true {
		t.Fatal("test B, collision C fail")
	}
	if d != boxElement.collisionDownRight(element2) || d != true {
		t.Fatal("test B, collision D fail")
	}

	element2 = Box{X: 10 - 7, Y: 10, Width: 10, Height: 10}

	a, b, c, d = boxElement.Quadrant(element2)
	if a != boxElement.collisionUpLeft(element2) || a != true {
		t.Fatal("test C, collision A fail")
	}
	if b != boxElement.collisionUpRight(element2) || b != false {
		t.Fatal("test C, collision B fail")
	}
	if c != boxElement.collisionDownLeft(element2) || c != true {
		t.Fatal("test C, collision C fail")
	}
	if d != boxElement.collisionDownRight(element2) || d != false {
		t.Fatal("test C, collision D fail")
	}

	element2 = Box{X: 10 - 7, Y: 10 - 7, Width: 10, Height: 10}

	a, b, c, d = boxElement.Quadrant(element2)
	if a != boxElement.collisionUpLeft(element2) || a != true {
		t.Fatal("test D, collision A fail")
	}
	if b != boxElement.collisionUpRight(element2) || b != false {
		t.Fatal("test D, collision B fail")
	}
	if c != boxElement.collisionDownLeft(element2) || c != false {
		t.Fatal("test D, collision C fail")
	}
	if d != boxElement.collisionDownRight(element2) || d != false {
		t.Fatal("test D, collision D fail")
	}

	element2 = Box{X: 10 - 4, Y: 10 - 7, Width: 10, Height: 10}

	a, b, c, d = boxElement.Quadrant(element2)
	if a != boxElement.collisionUpLeft(element2) || a != true {
		t.Fatal("test E, collision A fail")
	}
	if b != boxElement.collisionUpRight(element2) || b != true {
		t.Fatal("test E, collision B fail")
	}
	if c != boxElement.collisionDownLeft(element2) || c != false {
		t.Fatal("test E, collision C fail")
	}
	if d != boxElement.collisionDownRight(element2) || d != false {
		t.Fatal("test E, collision D fail")
	}

	element2 = Box{X: 10 + 7, Y: 10 + 7, Width: 10, Height: 10}

	a, b, c, d = boxElement.Quadrant(element2)
	if a != boxElement.collisionUpLeft(element2) || a != false {
		t.Fatal("test F, collision A fail")
	}
	if b != boxElement.collisionUpRight(element2) || b != false {
		t.Fatal("test F, collision B fail")
	}
	if c != boxElement.collisionDownLeft(element2) || c != false {
		t.Fatal("test F, collision C fail")
	}
	if d != boxElement.collisionDownRight(element2) || d != true {
		t.Fatal("test F, collision D fail")
	}

	element2 = Box{X: 10 + 7, Y: 10, Width: 10, Height: 10}

	a, b, c, d = boxElement.Quadrant(element2)
	if a != boxElement.collisionUpLeft(element2) || a != false {
		t.Fatal("test G, collision A fail")
	}
	if b != boxElement.collisionUpRight(element2) || b != true {
		t.Fatal("test G, collision B fail")
	}
	if c != boxElement.collisionDownLeft(element2) || c != false {
		t.Fatal("test G, collision C fail")
	}
	if d != boxElement.collisionDownRight(element2) || d != true {
		t.Fatal("test G, collision D fail")
	}
}
