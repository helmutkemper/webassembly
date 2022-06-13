package html

type SvgSpacing string

func (e SvgSpacing) String() string {
	return string(e)
}

const (
	// KSvgSpacingAuto
	//
	// English:
	//
	// This value indicates that the user agent should use text-on-a-path layout algorithms to adjust the spacing between
	// typographic characters in order to achieve visually appealing results.
	//
	// Português:
	//
	// Esse valor indica que o agente do usuário deve usar algoritmos de layout de texto em um caminho para ajustar o
	// espaçamento entre caracteres tipográficos para obter resultados visualmente atraentes.
	KSvgSpacingAuto SvgSpacing = "auto"

	// KSvgSpacingExact
	//
	// English:
	//
	// This value indicates that the typographic characters should be rendered exactly according to the spacing rules as
	// specified by the layout rules for text-on-a-path.
	//
	// Português:
	//
	// Esse valor indica que os caracteres tipográficos devem ser renderizados exatamente de acordo com as regras de
	// espaçamento especificadas pelas regras de layout para texto em um caminho.
	KSvgSpacingExact SvgSpacing = "exact"
)
