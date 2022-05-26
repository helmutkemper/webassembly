package html

// Ratio
//
// English:
//
//  The alignment value indicates whether to force uniform scaling and, if so, the alignment method to use in case the
//  aspect ratio of the viewBox doesn't match the aspect ratio of the viewport.
//
// Português:
//
//  O valor de alinhamento indica se deve-se forçar o dimensionamento uniforme e, em caso afirmativo, o método de
//  alinhamento a ser usado caso a proporção da viewBox não corresponda à proporção da janela de visualização.
type Ratio string

func (e Ratio) String() string {
	return string(e)
}

const (
	// KRatioNone
	//
	// English:
	//
	//  Do not force uniform scaling.
	//
	// Scale the graphic content of the given element non-uniformly if necessary such that the element's bounding box
	// exactly matches the viewport rectangle.
	//
	// Note that if <align> is none, then the optional <meetOrSlice> value is ignored.
	//
	// Português:
	//
	//  Não force a escala uniforme.
	//
	// Dimensione o conteúdo gráfico do elemento fornecido de forma não uniforme, se necessário, de modo que a caixa
	// delimitadora do elemento corresponda exatamente ao retângulo da janela de visualização.
	//
	// Observe que se <align> for none, o valor opcional <meetOrSlice> será ignorado.
	KRatioNone Ratio = "none"

	// KRatioXMinYMin
	//
	// English:
	//
	//  Force uniform scaling.
	//
	// Align the <min-x> of the element's viewBox with the smallest X value of the viewport.
	//
	// Align the <min-y> of the element's viewBox with the smallest Y value of the viewport.
	//
	// Português:
	//
	//  Forçar escala uniforme.
	//
	// Alinhe o <min-x> da viewBox do elemento com o menor valor X da viewport.
	//
	// Alinhe o <min-y> da viewBox do elemento com o menor valor Y da viewport.
	KRatioXMinYMin Ratio = "xMinYMin"

	// KRatioXMidYMin
	//
	// English:
	//
	//  Force uniform scaling.
	//
	// Align the midpoint X value of the element's viewBox with the midpoint X value of the viewport.
	//
	// Align the <min-y> of the element's viewBox with the smallest Y value of the viewport.
	//
	// Português:
	//
	//  Forçar escala uniforme.
	//
	// Alinhe o valor X do ponto médio da viewBox do elemento com o valor X do ponto médio da viewport.
	//
	// Alinhe o <min-y> da viewBox do elemento com o menor valor Y da viewport.
	KRatioXMidYMin Ratio = "xMidYMin"

	// KRatioXMaxYMin
	//
	// English:
	//
	//  Force uniform scaling.
	//
	// Align the <min-x>+<width> of the element's viewBox with the maximum X value of the viewport.
	//
	// Align the <min-y> of the element's viewBox with the smallest Y value of the viewport.
	//
	// Português:
	//
	//  Forçar escala uniforme.
	//
	// Alinhe o <min-x>+<width> da viewBox do elemento com o valor X máximo da viewport.
	//
	// Alinhe o <min-y> da viewBox do elemento com o menor valor Y da viewport.
	KRatioXMaxYMin Ratio = "xMaxYMin"

	// KRatioXMinYMid
	//
	// English:
	//
	//  Force uniform scaling.
	//
	// Align the <min-x> of the element's viewBox with the smallest X value of the viewport.
	//
	// Align the midpoint Y value of the element's viewBox with the midpoint Y value of the viewport.
	//
	// Português:
	//
	//  Forçar escala uniforme.
	//
	// Alinhe o <min-x> da viewBox do elemento com o menor valor X da viewport.
	//
	// Alinhe o valor Y do ponto médio da viewBox do elemento com o valor Y do ponto médio da viewport.
	KRatioXMinYMid Ratio = "xMinYMid"

	// KRatioXMidYMid
	//
	// English:
	//
	//  Force uniform scaling.
	//
	// Align the midpoint X value of the element's viewBox with the midpoint X value of the viewport.
	//
	// Align the midpoint Y value of the element's viewBox with the midpoint Y value of the viewport.
	//
	// Português:
	//
	//  Forçar escala uniforme.
	//
	// Alinhe o valor X do ponto médio da viewBox do elemento com o valor X do ponto médio da viewport.
	//
	// Alinhe o valor Y do ponto médio da viewBox do elemento com o valor Y do ponto médio da viewport.
	KRatioXMidYMid Ratio = "xMidYMid"

	// KRatioXMaxYMid
	//
	// English:
	//
	//  Force uniform scaling.
	//
	// Align the <min-x>+<width> of the element's viewBox with the maximum X value of the viewport.
	//
	// Align the midpoint Y value of the element's viewBox with the midpoint Y value of the viewport.
	//
	// Português:
	//
	//  Forçar escala uniforme.
	//
	// Alinhe o <min-x>+<width> da viewBox do elemento com o valor X máximo da viewport.
	//
	// Alinhe o valor Y do ponto médio da viewBox do elemento com o valor Y do ponto médio da viewport.
	KRatioXMaxYMid Ratio = "xMaxYMid"

	// KRatioXMinYMax
	//
	// English:
	//
	//  Force uniform scaling.
	//
	// Align the <min-x> of the element's viewBox with the smallest X value of the viewport.
	//
	// Align the <min-y>+<height> of the element's viewBox with the maximum Y value of the viewport.
	//
	// Português:
	//
	//  Forçar escala uniforme.
	//
	// Alinhe o <min-x> da viewBox do elemento com o menor valor X da viewport.
	//
	// Alinhe o <min-y>+<height> da viewBox do elemento com o valor Y máximo da viewport.
	KRatioXMinYMax Ratio = "xMinYMax"

	// KRatioXMidYMax
	//
	// English:
	//
	//  Force uniform scaling.
	//
	// Align the midpoint X value of the element's viewBox with the midpoint X value of the viewport.
	//
	// Align the <min-y>+<height> of the element's viewBox with the maximum Y value of the viewport.
	//
	// Português:
	//
	//  Forçar escala uniforme.
	//
	// Alinhe o valor X do ponto médio da viewBox do elemento com o valor X do ponto médio da viewport.
	//
	// Alinhe o <min-y>+<height> da viewBox do elemento com o valor Y máximo da viewport.
	KRatioXMidYMax Ratio = "xMidYMax"

	// KRatioXMaxYMax
	//
	// English:
	//
	//  Force uniform scaling.
	//
	// Align the <min-x>+<width> of the element's viewBox with the maximum X value of the viewport.
	//
	// Align the <min-y>+<height> of the element's viewBox with the maximum Y value of the viewport.
	//
	// Português:
	//
	//  Forçar escala uniforme.
	//
	// Alinhe o <min-x>+<width> da viewBox do elemento com o valor X máximo da viewport.
	//
	// Alinhe o <min-y>+<height> da viewBox do elemento com o valor Y máximo da viewport.
	KRatioXMaxYMax Ratio = "xMaxYMax"
)
