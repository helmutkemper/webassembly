package html

type SvgDominantBaseline string

func (e SvgDominantBaseline) String() string {
	return string(e)
}

const (
	// KSvgDominantBaselineAuto
	//
	// English:
	//
	//  If this property occurs on a <text> element, then the computed value depends on the value of the writing-mode
	//  attribute.
	//
	// If the writing-mode is horizontal, then the value of the dominant-baseline component is alphabetic. Otherwise, if
	// the writing-mode is vertical, then the value of the dominant-baseline component is central.
	//
	// If this property occurs on a <tspan>, <tref>, <altGlyph>, or <textPath> element, then the dominant-baseline and
	// the baseline-table components remain the same as those of the parent text content element.
	//
	// If the computed baseline-shift value actually shifts the baseline, then the baseline-table font-size component is
	// set to the value of the font-size attribute on the element on which the dominant-baseline attribute occurs,
	// otherwise the baseline-table font-size remains the same as that of the element.
	//
	// If there is no parent text content element, the scaled-baseline-table value is constructed as above for <text>
	// elements.
	//
	// Português:
	//
	// Se esta propriedade ocorrer em um elemento <text>, então o valor calculado depende do valor do atributo write-mode.
	//
	// Se o modo de escrita for horizontal, então o valor do componente da linha de base dominante será alfabético.
	// Caso contrário, se o modo de escrita for vertical, então o valor do componente da linha de base dominante
	// é central.
	//
	// Se essa propriedade ocorrer em um elemento <tspan>, <tref>, <altGlyph> ou <textPath>, os componentes de linha de
	// base dominante e tabela de linha de base permanecerão iguais aos do elemento de conteúdo de texto pai.
	//
	// Se o valor calculado de mudança de linha de base realmente mudar a linha de base, o componente tamanho da fonte
	// da tabela de linha de base será definido como o valor do atributo tamanho da fonte no elemento no qual o atributo
	// da linha de base dominante ocorre, caso contrário, a fonte da tabela de linha de base -size permanece igual ao
	// do elemento.
	//
	// Se não houver nenhum elemento de conteúdo de texto pai, o valor scaled-baseline-table será construído como acima
	// para elementos <text>.
	KSvgDominantBaselineAuto SvgDominantBaseline = "auto"

	// KSvgDominantBaselineIdeographic
	//
	// English:
	//
	//  The baseline-identifier for the dominant-baseline is set to be ideographic, the derived baseline-table is
	//  constructed using the ideographic baseline-table in the font, and the baseline-table font-size is changed to the
	//  value of the font-size attribute on this element.
	//
	// Portuguese
	//
	//  O identificador de linha de base para a linha de base dominante é definido como ideográfico, a tabela de linha de
	//  base derivada é construída usando a tabela de linha de base ideográfica na fonte e o tamanho da fonte da tabela
	//  de linha de base é alterado para o valor do tamanho da fonte atributo neste elemento.
	KSvgDominantBaselineIdeographic SvgDominantBaseline = "ideographic"

	// KSvgDominantBaselineAlphabetic
	//
	// English:
	//
	//  The baseline-identifier for the dominant-baseline is set to be alphabetic, the derived baseline-table is
	//  constructed using the alphabetic baseline-table in the font, and the baseline-table font-size is changed to the
	//  value of the font-size attribute on this element.
	//
	// Portuguese
	//
	//  O identificador de linha de base para a linha de base dominante é definido como alfabético, a tabela de linha de
	//  base derivada é construída usando a tabela de linha de base alfabética na fonte e o tamanho da fonte da tabela de
	//  linha de base é alterado para o valor do tamanho da fonte atributo neste elemento.
	KSvgDominantBaselineAlphabetic SvgDominantBaseline = "alphabetic"

	// KSvgDominantBaselineHanging
	//
	// English:
	//
	//  The baseline-identifier for the dominant-baseline is set to be hanging, the derived baseline-table is constructed
	//  using the hanging baseline-table in the font, and the baseline-table font-size is changed to the value of the
	//  font-size attribute on this element.
	//
	// Portuguese
	//
	//  O identificador de linha de base para a linha de base dominante é definido como suspenso, a tabela de linha de
	//  base derivada é construída usando a tabela de linha de base suspensa na fonte e o tamanho da fonte da tabela de
	//  linha de base é alterado para o valor do tamanho da fonte atributo neste elemento.
	KSvgDominantBaselineHanging SvgDominantBaseline = "hanging"

	// KSvgDominantBaselineMathematical
	//
	// English:
	//
	//  The baseline-identifier for the dominant-baseline is set to be mathematical, the derived baseline-table is
	//  constructed using the mathematical baseline-table in the font, and the baseline-table font-size is changed to the
	//  value of the font-size attribute on this element.
	//
	// Portuguese
	//
	//  The baseline-identifier for the dominant-baseline is set to be mathematical, the derived baseline-table is
	//  constructed using the mathematical baseline-table in the font, and the baseline-table font-size is changed to the
	//  value of the font-size attribute on this element.
	KSvgDominantBaselineMathematical SvgDominantBaseline = "mathematical"

	// KSvgDominantBaselineCentral
	//
	// English:
	//
	//  The baseline-identifier for the dominant-baseline is set to be central. The derived baseline-table is constructed
	//  from the defined baselines in a baseline-table in the font. That font baseline-table is chosen using the following
	//  priority order of baseline-table names: ideographic, alphabetic, hanging, mathematical. The baseline-table
	//  font-size is changed to the value of the font-size attribute on this element.
	//
	// Portuguese
	//
	//  O identificador de linha de base para a linha de base dominante é definido como central. A tabela de linha de base
	//  derivada é construída a partir das linhas de base definidas em uma tabela de linha de base na fonte. Essa tabela
	//  de linha de base de fonte é escolhida usando a seguinte ordem de prioridade de nomes de tabela de linha de base:
	//  ideográfica, alfabética, suspensa, matemática. O font-size da baseline-table é alterado para o valor do atributo
	//  font-size neste elemento.
	KSvgDominantBaselineCentral SvgDominantBaseline = "central"

	// KSvgDominantBaselineMiddle
	//
	// English:
	//
	//  The baseline-identifier for the dominant-baseline is set to be middle. The derived baseline-table is constructed
	//  from the defined baselines in a baseline-table in the font. That font baseline-table is chosen using the following
	//  priority order of baseline-table names: alphabetic, ideographic, hanging, mathematical. The baseline-table
	//  font-size is changed to the value of the font-size attribute on this element.
	//
	// Portuguese
	//
	//  O identificador de linha de base para a linha de base dominante é definido como meio. A tabela de linha de base
	//  derivada é construída a partir das linhas de base definidas em uma tabela de linha de base na fonte. Essa tabela
	//  de linha de base de fonte é escolhida usando a seguinte ordem de prioridade de nomes de tabela de linha de base:
	//  alfabética, ideográfica, suspensa, matemática. O font-size da baseline-table é alterado para o valor do atributo
	//  font-size neste elemento.
	KSvgDominantBaselineMiddle SvgDominantBaseline = "middle"

	// KSvgDominantBaselineTextAfterEdge
	//
	// English:
	//
	//  The baseline-identifier for the dominant-baseline is set to be text-after-edge. The derived baseline-table is
	//  constructed from the defined baselines in a baseline-table in the font. The choice of which font baseline-table
	//  to use from the baseline-tables in the font is browser dependent. The baseline-table font-size is changed to the
	//  value of the font-size attribute on this element.
	//
	// Portuguese
	//
	//  O identificador de linha de base para a linha de base dominante é definido como texto após borda. A tabela de
	//  linha de base derivada é construída a partir das linhas de base definidas em uma tabela de linha de base na fonte.
	//  A escolha de qual tabela de linha de base de fonte usar a partir das tabelas de linha de base na fonte depende do
	//  navegador. O font-size da baseline-table é alterado para o valor do atributo font-size neste elemento.
	KSvgDominantBaselineTextAfterEdge SvgDominantBaseline = "text-after-edge"

	// KSvgDominantBaselineTextBeforeEdge
	//
	// English:
	//
	//  The baseline-identifier for the dominant-baseline is set to be text-before-edge. The derived baseline-table is
	//  constructed from the defined baselines in a baseline-table in the font. The choice of which baseline-table to use
	//  from the baseline-tables in the font is browser dependent. The baseline-table font-size is changed to the value of
	//  the font-size attribute on this element.
	//
	// Portuguese
	//
	//  O identificador de linha de base para a linha de base dominante é definido como text-before-edge. A tabela de
	//  linha de base derivada é construída a partir das linhas de base definidas em uma tabela de linha de base na fonte.
	//  A escolha de qual tabela de linha de base usar a partir das tabelas de linha de base na fonte depende do
	//  navegador. O font-size da baseline-table é alterado para o valor do atributo font-size neste elemento.
	KSvgDominantBaselineTextBeforeEdge SvgDominantBaseline = "text-before-edge"

	// KSvgDominantBaselineTextTop
	//
	// English:
	//
	//  This value uses the top of the em box as the baseline.
	//
	// Portuguese
	//
	//  Esse valor usa a parte superior da caixa em como linha de base.
	KSvgDominantBaselineTextTop SvgDominantBaseline = "text-top"
)
