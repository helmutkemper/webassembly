package html

// todo: rever
type CollisionBoundingBox interface {
	// GetBoundingBox
	//
	// English:
	//
	// Returns the element's bounding box.
	//
	// Português:
	//
	// Retorna o bounding box do elemnto.
	GetBoundingBox() (x, y, width, height int)
}

type CollisionBox interface {
	// GetCollisionBox
	//
	// English:
	//
	// Returns the element's collision box.
	//
	// Português:
	//
	// Retorna o elemento collison box.
	GetCollisionBox() (box Box)
	AdjustBox(dx, dy int)
}
