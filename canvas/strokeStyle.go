package canvas

import "image/color"

// en: Sets the color, gradient, or pattern used for strokes
//     value: a valid JavaScript value or a color.RGBA{} struct
//     Default value: #000000
//
// pt_br: Define a cor, gradiente ou padrão usado para o contorno
//     value: um valor JavaScript valido ou um struct color.RGBA{}
//     Valor padrão: #000000
func (el *Canvas) SetStrokeStyle(value interface{}) {
	switch value.(type) {
	case color.RGBA:
		el.SelfContext.Set("strokeStyle", RGBAToJs(value.(color.RGBA)))

	default:
		el.SelfContext.Set("strokeStyle", value)
	}
}
