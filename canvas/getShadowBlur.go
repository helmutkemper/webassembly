package canvas

// GetShadowBlur
// en: Return the blur level for shadows
//     Default value: 0
//
// pt_br: Retorna o valor de borrão da sombra
//     Valor padrão: 0
func (el *Canvas) GetShadowBlur() int {
	return el.SelfContext.Get("shadowBlur").Int()
}
