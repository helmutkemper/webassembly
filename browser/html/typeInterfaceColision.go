package html

type Collision interface {
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
