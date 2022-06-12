package html

type SvgOperatorFeComposite string

func (e SvgOperatorFeComposite) String() string {
	return string(e)
}

const (
	// KSvgOperatorFeCompositeOver
	//
	// English:
	//
	// This value indicates that the source graphic defined in the in attribute is placed over the destination graphic
	// defined in the in2 attribute.
	//
	// Português:
	//
	// Este valor indica que o gráfico de origem definido no atributo in é colocado sobre o gráfico de destino definido
	// no atributo in2.
	KSvgOperatorFeCompositeOver SvgOperatorFeComposite = "over"

	// KSvgOperatorFeCompositeIn
	//
	// English:
	//
	// This value indicates that the parts of the source graphic defined in the in attribute that overlap the destination
	// graphic defined in the in2 attribute, replace the destination graphic.
	//
	// Português:
	//
	// Esse valor indica que as partes do gráfico de origem definidas no atributo in que se sobrepõem ao gráfico de
	// destino definido no atributo in2 substituem o gráfico de destino.
	KSvgOperatorFeCompositeIn SvgOperatorFeComposite = "in"

	// KSvgOperatorFeCompositeOut
	//
	// English:
	//
	// This value indicates that the parts of the source graphic defined in the in attribute that fall outside the
	// destination graphic defined in the in2 attribute, are displayed.
	//
	// Português:
	//
	// Esse valor indica que as partes do gráfico de origem definidas no atributo in que estão fora do gráfico de destino
	// definido no atributo in2 são exibidas.
	KSvgOperatorFeCompositeOut SvgOperatorFeComposite = "out"

	// KSvgOperatorFeCompositeAtop
	//
	// English:
	//
	// This value indicates that the parts of the source graphic defined in the in attribute, which overlap the
	// destination graphic defined in the in2 attribute, replace the destination graphic. The parts of the destination
	// graphic that do not overlap with the source graphic stay untouched.
	//
	// Português:
	//
	// Este valor indica que as partes do gráfico de origem definidas no atributo in, que se sobrepõem ao gráfico de
	// destino definido no atributo in2, substituem o gráfico de destino. As partes do gráfico de destino que não se
	// sobrepõem ao gráfico de origem permanecem intactas.
	KSvgOperatorFeCompositeAtop SvgOperatorFeComposite = "atop"

	// KSvgOperatorFeCompositeXor
	//
	// English:
	//
	// This value indicates that the non-overlapping regions of the source graphic defined in the in attribute and the
	// destination graphic defined in the in2 attribute are combined.
	//
	// Português:
	//
	// Este valor indica que as regiões não sobrepostas do gráfico de origem definido no atributo in e o gráfico de
	// destino definido no atributo in2 são combinados.
	KSvgOperatorFeCompositeXor SvgOperatorFeComposite = "xor"

	// KSvgOperatorFeCompositeLighter
	//
	// English:
	//
	// This value indicates that the sum of the source graphic defined in the in attribute and the destination graphic
	// defined in the in2 attribute is displayed.
	//
	// Português:
	//
	// Este valor indica que a soma do gráfico de origem definido no atributo in e o gráfico de destino definido no
	// atributo in2 é exibido.
	KSvgOperatorFeCompositeLighter SvgOperatorFeComposite = "lighter"

	// KSvgOperatorFeCompositeArithmetic
	//
	// English:
	//
	// This value indicates that the source graphic defined in the in attribute and the destination graphic defined in the
	// in2 attribute are combined using the following formula:
	//
	// result = k1*i1*i2 + k2*i1 + k3*i2 + k4
	//
	// where: i1 and i2 indicate the corresponding pixel channel values of the input image, which map to in and in2
	// respectively, and k1,k2,k3,and k4 indicate the values of the attributes with the same name.
	//
	// Português:
	//
	// Esse valor indica que o gráfico de origem definido no atributo in e o gráfico de destino definido no atributo in2
	// são combinados usando a seguinte fórmula:
	//
	// resultado = k1*i1*i2 + k2*i1 + k3*i2 + k4
	//
	// onde: i1 e i2 indicam os valores de canal de pixel correspondentes da imagem de entrada, que mapeiam para in e in2
	// respectivamente, e k1,k2,k3 e k4 indicam os valores dos atributos com o mesmo nome.
	KSvgOperatorFeCompositeArithmetic SvgOperatorFeComposite = "arithmetic"
)
