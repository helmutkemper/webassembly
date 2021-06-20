package canvas

// ShadowOffsetY
// en: Sets or returns the vertical distance of the shadow from the shape
//     The shadowOffsetY property sets or returns the vertical distance of the shadow from the shape.
//     shadowOffsetY = 0 indicates that the shadow is right behind the shape.
//     shadowOffsetY = 20 indicates that the shadow starts 20 pixels below the shape's top position.
//     shadowOffsetY = -20 indicates that the shadow starts 20 pixels above the shape's top position.
//     Tip: To adjust the horizontal distance of the shadow from the shape, use the shadowOffsetX property.
//     Default value: 0
//
// pt_br: Define a distância vertical entre a forma e a sua sombra
//     shadowOffsetY = 0 indica que a forma e sua sombra estão alinhadas uma em cima da outra.
//     shadowOffsetY = 20 indica que a forma e a sua sombra estão 20 pixels afastadas para baixo (em relação a parte mais elevada da forma)
//     shadowOffsetY = -20 indica que a forma e a sua sombra estão 20 pixels afastadas para cima (em relação a parte mais elevada da forma)
//     Dica: Para ajustar a distância horizontal, use a propriedade shadowOffsetX
//     Valor padrão: 0
func (el *Canvas) ShadowOffsetY(value int) {
	el.SelfContext.Set("shadowOffsetY", value)
}
