package html

type SvgShapeRendering string

func (e SvgShapeRendering) String() string {
	return string(e)
}

const (
	// KSvgShapeRenderingAuto
	//
	// English:
	//
	// This value indicates that the user agent shall make appropriate tradeoffs to balance speed, crisp edges and
	// geometric precision, but with geometric precision given more importance than speed and crisp edges.
	//
	// Português:
	//
	// Esse valor indica que o agente do usuário deve fazer compensações apropriadas para equilibrar velocidade, bordas
	// nítidas e precisão geométrica, mas com precisão geométrica dada mais importância do que velocidade e bordas nítidas.
	KSvgShapeRenderingAuto SvgShapeRendering = "auto"

	// KSvgShapeRenderingOptimizeSpeed
	//
	// English:
	//
	// This value indicates that the user agent shall emphasize rendering speed over geometric precision and crisp edges.
	// This option will sometimes cause the user agent to turn off shape anti-aliasing.
	//
	// Português:
	//
	// Esse valor indica que o agente do usuário deve enfatizar a velocidade de renderização sobre a precisão geométrica
	// e as bordas nítidas.
	// Essa opção às vezes fará com que o agente do usuário desative o anti-aliasing de forma.
	KSvgShapeRenderingOptimizeSpeed SvgShapeRendering = "optimizeSpeed"

	// KSvgShapeRenderingCrispEdges
	//
	// English:
	//
	// This value indicates that the user agent shall attempt to emphasize the contrast between clean edges of artwork
	// over rendering speed and geometric precision. To achieve crisp edges, the user agent might turn off anti-aliasing
	// for all lines and curves or possibly just for straight lines which are close to vertical or horizontal.
	// Also, the user agent might adjust line positions and line widths to align edges with device pixels.
	//
	// Português:
	//
	// Esse valor indica que o agente do usuário deve tentar enfatizar o contraste entre as bordas limpas do trabalho
	// artístico sobre a velocidade de renderização e a precisão geométrica. Para obter bordas nítidas, o agente do
	// usuário pode desativar o anti-aliasing para todas as linhas e curvas ou possivelmente apenas para linhas retas
	// próximas à vertical ou horizontal. Além disso, o agente do usuário pode ajustar as posições e larguras de linha
	// para alinhar as bordas com os pixels do dispositivo.
	KSvgShapeRenderingCrispEdges SvgShapeRendering = "crispEdges"

	// KSvgShapeRenderingGeometricPrecision
	//
	// English:
	//
	// Indicates that the user agent shall emphasize geometric precision over speed and crisp edges.
	//
	// Português:
	//
	// Indica que o agente do usuário deve enfatizar a precisão geométrica sobre a velocidade e as bordas nítidas.
	KSvgShapeRenderingGeometricPrecision SvgShapeRendering = "geometricPrecision"
)
