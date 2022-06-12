package html

type PositionHorizontal string

func (e PositionHorizontal) String() string {
	return string(e)
}

const (
	// KPositionHorizontalLeft
	//
	// English:
	//
	// The reference point of the marker is placed at the left edge of the shape.
	//
	// Português:
	//
	// O ponto de referência do marcador é colocado na borda esquerda da forma.
	KPositionHorizontalLeft PositionHorizontal = "left"

	// KPositionHorizontalCenter
	//
	// English:
	//
	// The reference point of the marker is placed at the horizontal center of the shape.
	//
	// Português:
	//
	// O ponto de referência do marcador é colocado no centro horizontal da forma.
	KPositionHorizontalCenter PositionHorizontal = "center"

	// KPositionHorizontalRight
	//
	// English:
	//
	// The reference point of the marker is placed at the right edge of the shape.
	//
	// Português:
	//
	// O ponto de referência do marcador é colocado na borda direita da forma.
	KPositionHorizontalRight PositionHorizontal = "right"
)
