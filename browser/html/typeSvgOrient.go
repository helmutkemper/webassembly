package html

type SvgOrient string

func (e SvgOrient) String() string {
	return string(e)
}

const (
	// KSvgOrientAuto
	//
	// English:
	//
	// This value indicates that the marker is oriented such that its positive x-axis is pointing in a direction relative
	// to the path at the position the marker is placed.
	//
	// Português:
	//
	// Este valor indica que o marcador está orientado de forma que seu eixo x positivo esteja apontando em uma direção
	// relativa ao caminho na posição em que o marcador é colocado.
	KSvgOrientAuto SvgOrient = "auto"

	// KSvgOrientAutoStartReverse
	//
	// English:
	//
	// If placed by marker-start, the marker is oriented 180° different from the orientation that would be used if auto
	// where specified. For all other markers, auto-start-reverse means the same as auto.
	//
	//   Notes:
	//     * This allows a single arrowhead marker to be defined that can be used for both the start and end of a path,
	//       i.e. which points outwards from both ends.
	//
	// Português:
	//
	// Se colocado pelo início do marcador, o marcador é orientado 180° diferente da orientação que seria usada se auto,
	// onde especificado. Para todos os outros marcadores, auto-start-reverse significa o mesmo que auto.
	//
	//   Notas:
	//     * Isso permite definir um único marcador de ponta de seta que pode ser usado tanto para o início quanto para o
	//       final de um caminho, ou seja, que aponta para fora de ambas as extremidades.
	KSvgOrientAutoStartReverse SvgOrient = "auto-start-reverse"
)
