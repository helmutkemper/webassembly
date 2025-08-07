package html

type CanvasData struct {
	// Width
	//
	//  English:
	//
	//   Redefines the width of the copied figure to canvas (when greater than zero)
	//
	//  Português:
	//
	//   Redefine o comprimento da imagem copiada para o canvas (quando maior do que zero)
	Width int

	// Height
	//
	//  English:
	//
	//   Redefines the height of the copied figure to canvas (when larger than zero)
	//
	//  Português:
	//
	//   Redefine a altura da imagem copiada para o canvas (quando maior do que zero)
	Height int

	// ZoomVertical
	//
	//  English:
	//
	//   Horizontal zoom applied to figure (when greater than zero)
	//
	//  Português:
	//
	//   Zoom horizontal aplicado a figura (quando maior do que zero)
	ZoomVertical float64

	// ZoomHorizontal
	//
	//  English:
	//
	//   Horizontal zoom applied to figure (when greater than zero)
	//
	//  Português:
	//
	//   Zoom horizontal aplicado a figura (quando maior do que zero)
	ZoomHorizontal float64

	// Alpha
	//
	//  English:
	//
	//   Alpha applied to the figure (when greater than zero)
	//
	//  Português:
	//
	//   Alpha aplicado a figura (quando maior do que zero)
	Alpha float64

	// AlphaForce
	//
	//  English:
	//
	//   Force alpha when it is zero
	//
	//  Português:
	//
	//   Força o alpha quando este é igual a zero
	AlphaForce bool

	// ZoomVerticalForce
	//
	//  English:
	//
	//   Force zoom when it is zero
	//
	//  Português:
	//
	//   Força o zoom quando este é igual a zero
	ZoomVerticalForce bool

	// ZoomHorizontalForce
	//
	//  English:
	//
	//   Force zoom when it is zero
	//
	//  Português:
	//
	//   Força o zoom quando este é igual a zero
	ZoomHorizontalForce bool
}
