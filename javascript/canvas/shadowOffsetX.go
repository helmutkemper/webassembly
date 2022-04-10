package canvas

// ShadowOffsetX
// en: Sets the horizontal distance of the shadow from the shape
//     shadowOffsetX = 0 indicates that the shadow is right behind the shape.
//     shadowOffsetX = 20 indicates that the shadow starts 20 pixels to the right (from the shape's left position).
//     shadowOffsetX = -20 indicates that the shadow starts 20 pixels to the left (from the shape's left position).
//     Tip: To adjust the vertical distance of the shadow from the shape, use the shadowOffsetY property.
//     Default value: 0
//
// pt_br: Define a distância horizontal entre a forma e a sua sombra
//     shadowOffsetX = 0 indica que a forma e sua sombra estão alinhadas uma em cima da outra.
//     shadowOffsetX = 20 indica que a forma e a sua sombra estão 20 pixels afastadas a direita (em relação a parte mais a esquerda da forma)
//     shadowOffsetX = -20 indica que a forma e a sua sombra estão 20 pixels afastadas a esquerda (em relação a parte mais a esquerda da forma)
//     Dica: Para ajustar a distância vertical, use a propriedade shadowOffsetY
//     Valor padrão: 0
func (el *Canvas) ShadowOffsetX(value int) {
	el.SelfContext.Set("shadowOffsetX", value)
}
