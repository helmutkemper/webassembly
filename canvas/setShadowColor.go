package canvas

import (
	"image/color"
)

// en: Sets the color to use for shadows
//     Note: Use the shadowColor property together with the shadowBlur property to create a shadow.
//     Tip: Adjust the shadow by using the shadowOffsetX and shadowOffsetY properties.
//     Default value: #000000
//
// pt_br: Define a cor da sombra
//     Nota: Use a propriedade shadowColor em conjunto com a propriedade shadowBlur para criar a sombra
//     Dica: Ajuste o local da sombra usando as propriedades shadowOffsetX e shadowOffsetY
//     Valor padrão: #000000
func (el *Canvas) SetShadowColor(value color.RGBA) {
	el.SelfContext.Set("shadowColor", RGBAToJs(value))
}
