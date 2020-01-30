package canvas

// en: Sets the blur level for shadows
//     Default value: 0
//
// pt_br: Define o valor de borrão da sombra
//     Valor padrão: 0
func (el *Canvas) SetShadowBlur(value int) {
	el.SelfContext.Set("shadowBlur", value)
}
