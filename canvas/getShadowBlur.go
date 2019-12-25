package canvas

// en: Return the blur level for shadows
//     Default value: 0
//
// pt_br: Retorna o valor de borrão da sombra
//     Valor padrão: 0
func (el *Canvas) GetShadowBlur() float64 {
	return el.SelfContext.Get("shadowBlur").Float()
}
