package canvas

import "image/color"

// SetFillStyle
// en: Sets the color, gradient, or pattern used to fill the drawing
//     value: a valid JavaScript value or a color.RGBA{} struct
//     Default value:	#000000
//
// pt_br: Define a cor, gradiente ou padrão usado para preencher o desenho
//     value: um valor JavaScript valido ou um struct color.RGBA{}
//     Valor padrão: #000000
func (el *Canvas) SetFillStyle(value interface{}) {
	switch value.(type) {
	case color.RGBA:
		el.SelfContext.Set("fillStyle", RGBAToJs(value.(color.RGBA)))

	default:
		el.SelfContext.Set("fillStyle", value)
	}
}
