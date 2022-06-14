package html

type SvgTextRendering string

func (e SvgTextRendering) String() string {
	return string(e)
}

const (
	// KSvgTextRenderingAuto
	//
	// English:
	//
	// This value indicates that the user agent shall make appropriate tradeoffs to balance speed, legibility and
	// geometric precision, but with legibility given more importance than speed and geometric precision.
	//
	// Português:
	//
	// Este valor indica que o agente do usuário deve fazer compensações apropriadas para equilibrar velocidade,
	// legibilidade e precisão geométrica, mas com legibilidade dada mais importância do que velocidade e precisão
	// geométrica.
	KSvgTextRenderingAuto SvgTextRendering = "auto"

	// KSvgTextRenderingOptimizeSpeed
	//
	// English:
	//
	// This value indicates that the user agent shall emphasize rendering speed over legibility and geometric precision.
	// This option will sometimes cause some user agents to turn off text anti-aliasing.
	//
	// Português:
	//
	// Esse valor indica que o agente do usuário deve enfatizar a velocidade de renderização sobre a legibilidade e a
	// precisão geométrica.
	// Essa opção às vezes fará com que alguns agentes do usuário desativem a suavização de serrilhado de texto.
	KSvgTextRenderingOptimizeSpeed SvgTextRendering = "optimizeSpeed"

	// KSvgTextRenderingOptimizeLegibility
	//
	// English:
	//
	// This value indicates that the user agent shall emphasize legibility over rendering speed and geometric precision.
	// The user agent will often choose whether to apply anti-aliasing techniques, built-in font hinting or both to
	// produce the most legible text.
	//
	// Português:
	//
	// Esse valor indica que o agente do usuário deve enfatizar a legibilidade sobre a velocidade de renderização e a
	// precisão geométrica.
	// O agente do usuário geralmente escolherá aplicar técnicas de anti-aliasing, dicas de fonte incorporadas ou ambos
	// para produzir o texto mais legível.
	KSvgTextRenderingOptimizeLegibility SvgTextRendering = "optimizeLegibility"

	// KSvgTextRenderingGeometricPrecision
	//
	// English:
	//
	// This value indicates that the user agent shall emphasize geometric precision over legibility and rendering speed.
	// This option will usually cause the user agent to suspend the use of hinting so that glyph outlines are drawn with
	// comparable geometric precision to the rendering of path data.
	//
	// Português:
	//
	// Esse valor indica que o agente do usuário deve enfatizar a precisão geométrica sobre a legibilidade e a velocidade
	// de renderização.
	// Essa opção geralmente fará com que o agente do usuário suspenda o uso de dicas para que os contornos dos glifos
	// sejam desenhados com precisão geométrica comparável à renderização dos dados do caminho.
	KSvgTextRenderingGeometricPrecision SvgTextRendering = "geometricPrecision"
)
