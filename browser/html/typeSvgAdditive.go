package html

type SvgAdditive string

func (e SvgAdditive) String() string {
	return string(e)
}

const (

	// KSvgAdditiveSum
	//
	// English:
	//
	//  Specifies that the animation will add to the underlying value of the attribute and other lower priority
	//  animations.
	//
	// Português:
	//
	//  Especifica que a animação será adicionada ao valor subjacente do atributo e outras animações de prioridade mais
	//  baixa.
	KSvgAdditiveSum SvgAdditive = "sum"

	// KSvgAdditiveReplace
	//
	// English:
	//
	//  Specifies that the animation will override the underlying value of the attribute and other lower priority
	//  animations.
	//
	// This is the default, however the behavior is also affected by the animation value attributes by and to, as
	// described in SMIL Animation: How from, to and by attributes affect additive behavior.
	//
	// Português:
	//
	//  Especifica que a animação substituirá o valor subjacente do atributo e outras animações de prioridade mais baixa.
	//
	// Este é o padrão, no entanto, o comportamento também é afetado pelos atributos de valor de animação por e para,
	// conforme descrito em Animação SMIL: Como os atributos de, para e por afetam o comportamento aditivo.
	KSvgAdditiveReplace SvgAdditive = "replace"
)
