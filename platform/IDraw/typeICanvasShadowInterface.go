package iotmaker_platform_IDraw

import "image/color"

type ICanvasShadow interface {

	// SetShadowBlur
	// en: Sets the blur level for shadows
	//     Default value: 0
	//
	// pt_br: Define o valor de borrão da sombra
	//     Valor padrão: 0
	SetShadowBlur(value interface{})

	// SetShadowColor
	// en: Sets the color to use for shadows
	//     Note: Use the shadowColor property together with the shadowBlur property to
	//     create a shadow.
	//     Tip: Adjust the shadow by using the shadowOffsetX and shadowOffsetY
	//     properties.
	//     Default value: #000000
	//
	// pt_br: Define a cor da sombra
	//     Nota: Use a propriedade shadowColor em conjunto com a propriedade shadowBlur
	//     para criar a sombra
	//     Dica: Ajuste o local da sombra usando as propriedades shadowOffsetX e
	//     shadowOffsetY
	//     Valor padrão: #000000
	SetShadowColor(value color.RGBA)

	// ShadowOffsetX
	// en: Sets the horizontal distance of the shadow from the shape
	//     shadowOffsetX = 0 indicates that the shadow is right behind the shape.
	//     shadowOffsetX = 20 indicates that the shadow starts 20 pixels to the right
	//     (from the shape's left position).
	//     shadowOffsetX = -20 indicates that the shadow starts 20 pixels to the left
	//     (from the shape's left position).
	//     Tip: To adjust the vertical distance of the shadow from the shape, use the
	//     shadowOffsetY property.
	//     Default value: 0
	//
	// pt_br: Define a distância horizontal entre a forma e a sua sombra
	//     shadowOffsetX = 0 indica que a forma e sua sombra estão alinhadas uma em
	//     cima da outra.
	//     shadowOffsetX = 20 indica que a forma e a sua sombra estão 20 pixels
	//     afastadas a direita (em relação a parte mais a esquerda da forma)
	//     shadowOffsetX = -20 indica que a forma e a sua sombra estão 20 pixels
	//     afastadas a esquerda (em relação a parte mais a esquerda da forma)
	//     Dica: Para ajustar a distância vertical, use a propriedade shadowOffsetY
	//     Valor padrão: 0
	ShadowOffsetX(value int)

	// ShadowOffsetY
	// en: Sets or returns the vertical distance of the shadow from the shape
	//     The shadowOffsetY property sets or returns the vertical distance of the
	//     shadow from the shape.
	//     shadowOffsetY = 0 indicates that the shadow is right behind the shape.
	//     shadowOffsetY = 20 indicates that the shadow starts 20 pixels below the
	//     shape's top position.
	//     shadowOffsetY = -20 indicates that the shadow starts 20 pixels above the
	//     shape's top position.
	//     Tip: To adjust the horizontal distance of the shadow from the shape, use the
	//     shadowOffsetX property.
	//     Default value: 0
	//
	// pt_br: Define a distância vertical entre a forma e a sua sombra
	//     shadowOffsetY = 0 indica que a forma e sua sombra estão alinhadas uma em
	//     cima da outra.
	//     shadowOffsetY = 20 indica que a forma e a sua sombra estão 20 pixels
	//     afastadas para baixo (em relação a parte mais elevada da forma)
	//     shadowOffsetY = -20 indica que a forma e a sua sombra estão 20 pixels
	//     afastadas para cima (em relação a parte mais elevada da forma)
	//     Dica: Para ajustar a distância horizontal, use a propriedade shadowOffsetX
	//     Valor padrão: 0
	ShadowOffsetY(value int)
}
