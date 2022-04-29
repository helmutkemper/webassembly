package html

type FillRule string

func (e FillRule) String() string {
	return string(e)
}

const (
	// KFillRuleNonZero
	//
	// English:
	//
	//  In two-dimensional computer graphics, the non-zero winding rule is a means of determining
	//  whether a given point falls within an enclosed curve. Unlike the similar even-odd rule, it
	//  relies on knowing the direction of stroke for each part of the curve.
	//
	// Português:
	//
	//  Em computação gráfica bidimensional, a regra de enrolamento diferente de zero é um meio de
	//  determinar se um determinado ponto está dentro de uma curva fechada. Ao contrário da regra
	//  par-ímpar semelhante, ela depende do conhecimento da direção do curso para cada parte da curva.
	KFillRuleNonZero FillRule = "nonzero"

	// KFillRuleEvenOdd
	//
	// English:
	//
	//  The even–odd rule is an algorithm implemented in vector-based graphic software,[1] like the
	//  PostScript language and Scalable Vector Graphics (SVG), which determines how a graphical shape
	//  with more than one closed outline will be filled. Unlike the nonzero-rule algorithm, this
	//  algorithm will alternatively color and leave uncolored shapes defined by nested closed paths
	//  irrespective of their winding.
	//
	// Português:
	//
	//  A regra par-ímpar é um algoritmo implementado em software gráfico baseado em vetor,[1] como a
	//  linguagem PostScript e Scalable Vector Graphics (SVG), que determina como uma forma gráfica com
	//  mais de um contorno fechado será preenchida. Ao contrário do algoritmo de regra diferente de
	//  zero, esse algoritmo alternadamente colorirá e deixará formas não coloridas definidas por
	//  caminhos fechados aninhados, independentemente de seu enrolamento.
	KFillRuleEvenOdd FillRule = "evenodd"
)
