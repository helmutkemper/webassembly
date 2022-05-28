package html

type SvgAlignmentBaseline string

func (e SvgAlignmentBaseline) String() string {
	return string(e)
}

const (
	// KSvgAlignmentBaselineBaseline
	//
	// English:
	//
	//  Uses the dominant baseline choice of the parent. Matches the box's corresponding baseline to that of its parent.
	//
	// Português:
	//
	//  Usa a escolha de linha de base dominante do pai. Corresponde à linha de base correspondente da caixa à de seu pai.
	KSvgAlignmentBaselineBaseline SvgAlignmentBaseline = "baseline"

	// KSvgAlignmentBaselineTextBottom
	//
	// English:
	//
	//  Matches the bottom of the box to the top of the parent's content area.
	//
	// Português:
	//
	//  Corresponde a parte inferior da caixa à parte superior da área de conteúdo do pai.
	KSvgAlignmentBaselineTextBottom SvgAlignmentBaseline = "text-bottom"

	// KSvgAlignmentBaselineTextBeforeEdge
	//
	// English:
	//
	//  The alignment-point of the object being aligned is aligned with the "text-before-edge" baseline of the parent text content element.
	//
	//   Notes:
	//
	//     * This keyword may be mapped to text-top.
	//
	// Português:
	//
	//  O ponto de alinhamento do objeto que está sendo alinhado é alinhado com a linha de base "text-before-edge" do elemento de conteúdo de texto pai.
	//
	//   Notas:
	//
	//     * Esta palavra-chave pode ser mapeada para text-top.
	KSvgAlignmentBaselineTextBeforeEdge SvgAlignmentBaseline = "text-before-edge"

	// KSvgAlignmentBaselineMiddle
	//
	// English:
	//
	//  Aligns the vertical midpoint of the box with the baseline of the parent box plus half the x-height of the parent.
	//
	// Português:
	//
	//  Alinha o ponto médio vertical da caixa com a linha de base da caixa pai mais metade da altura x do pai.
	KSvgAlignmentBaselineMiddle SvgAlignmentBaseline = "middle"

	// KSvgAlignmentBaselineCentral
	//
	// English:
	//
	//  Matches the box's central baseline to the central baseline of its parent.
	//
	// Português:
	//
	//  Corresponde a linha de base central da caixa à linha de base central de seu pai.
	KSvgAlignmentBaselineCentral SvgAlignmentBaseline = "central"

	// KSvgAlignmentBaselineTextTop
	//
	// English:
	//
	//  Matches the top of the box to the top of the parent's content area.
	//
	// Português:
	//
	//  Corresponde a parte superior da caixa à parte superior da área de conteúdo do pai.
	KSvgAlignmentBaselineTextTop SvgAlignmentBaseline = "text-top"

	// KSvgAlignmentBaselineTextAfterEdge
	//
	// English:
	//
	//  The alignment-point of the object being aligned is aligned with the "text-after-edge" baseline of the parent text content element.
	//
	//   Notes:
	//     * This keyword may be mapped to text-bottom.
	//
	// Português:
	//
	//  O ponto de alinhamento do objeto que está sendo alinhado é alinhado com a linha de base "text-after-edge" do elemento de conteúdo de texto pai.
	//
	//   Notas:
	//     * Esta palavra-chave pode ser mapeada para text-bottom.
	KSvgAlignmentBaselineTextAfterEdge SvgAlignmentBaseline = "text-after-edge"

	// KSvgAlignmentBaselineIdeographic
	//
	// English:
	//
	//  Matches the box's ideographic character face under-side baseline to that of its parent.
	//
	// Português:
	//
	//  Corresponde à linha de base do lado inferior da face do caractere ideográfico da caixa com a de seu pai.
	KSvgAlignmentBaselineIdeographic SvgAlignmentBaseline = "ideographic"

	// KSvgAlignmentBaselineAlphabetic
	//
	// English:
	//
	//  Matches the box's alphabetic baseline to that of its parent.
	//
	// Português:
	//
	//  Corresponde a linha de base alfabética da caixa à de seu pai.
	KSvgAlignmentBaselineAlphabetic SvgAlignmentBaseline = "alphabetic"

	// KSvgAlignmentBaselineHanging
	//
	// English:
	//
	//  The alignment-point of the object being aligned is aligned with the "hanging" baseline of the parent text content element.
	//
	// Português:
	//
	//  O ponto de alinhamento do objeto que está sendo alinhado é alinhado com a linha de base "suspensa" do elemento de conteúdo de texto pai.
	KSvgAlignmentBaselineHanging SvgAlignmentBaseline = "hanging"

	// KSvgAlignmentBaselineMathematical
	//
	// English:
	//
	//  Matches the box's mathematical baseline to that of its parent.
	//
	// Português:
	//
	//  Corresponde a linha de base matemática da caixa à de seu pai.
	KSvgAlignmentBaselineMathematical SvgAlignmentBaseline = "mathematical"

	// KSvgAlignmentBaselineTop
	//
	// English:
	//
	//  Aligns the top of the aligned subtree with the top of the line box.
	//
	// Português:
	//
	//  Alinha o topo da subárvore alinhada com o topo da caixa de linha.
	KSvgAlignmentBaselineTop SvgAlignmentBaseline = "top"

	// KSvgAlignmentBaselineCenter
	//
	// English:
	//
	//  Aligns the center of the aligned subtree with the center of the line box.
	//
	// Português:
	//
	//  Alinha o centro da subárvore alinhada com o centro da caixa de linha.
	KSvgAlignmentBaselineCenter SvgAlignmentBaseline = "center"

	// KSvgAlignmentBaselineBottom
	//
	// English:
	//
	//  Aligns the bottom of the aligned subtree with the bottom of the line box.
	//
	// Português:
	//
	//  Alinha a parte inferior da subárvore alinhada com a parte inferior da caixa de linha.
	KSvgAlignmentBaselineBottom SvgAlignmentBaseline = "bottom"
)
