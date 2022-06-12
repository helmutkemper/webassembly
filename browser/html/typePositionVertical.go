package html

type PositionVertical string

func (e PositionVertical) String() string {
	return string(e)
}

const (
	// KPositionVerticalTop
	//
	// English:
	//
	// The reference point of the marker is placed at the top edge of the shape.
	//
	// Português:
	//
	// O ponto de referência do marcador é colocado na borda superior da forma.
	KPositionVerticalTop PositionVertical = "top"

	// KPositionVerticalCenter
	//
	// English:
	//
	// The reference point of the marker is placed at the vertical center of the shape.
	//
	// Português:
	//
	// O ponto de referência do marcador é colocado no centro vertical da forma.
	KPositionVerticalCenter PositionVertical = "center"

	// KPositionVerticalBottom
	//
	// English:
	//
	// The reference point of the marker is placed at the bottom edge of the shape.
	//
	// Português:
	//
	// O ponto de referência do marcador é colocado na borda inferior da forma.
	KPositionVerticalBottom PositionVertical = "bottom"
)
