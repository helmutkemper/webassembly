package html

type TextBaseLineRule string

func (e TextBaseLineRule) String() string {
	return string(e)
}

const (
	// KTextBaseLineRuleAlphabetic
	//
	// English:
	//
	//  (Default) The text baseline is the normal alphabetic baseline.
	//
	// Português:
	//
	//  (Padrão) A linha de base do texto é a linha de base alfabética normal.
	KTextBaseLineRuleAlphabetic TextBaseLineRule = "alphabetic"

	// KTextBaseLineRuleTop
	//
	// English:
	//
	//  The text baseline is the top of the em square.
	//
	// Português:
	//
	//  A linha de base do texto é a parte superior do quadrado em.
	KTextBaseLineRuleTop TextBaseLineRule = "top"

	// KTextBaseLineRuleHanging
	//
	// English:
	//
	//  The text baseline is the hanging baseline.
	//
	// Português:
	//
	//  A linha de base do texto é a linha de base suspensa.
	KTextBaseLineRuleHanging TextBaseLineRule = "hanging"

	// KTextBaseLineRuleMiddle
	//
	// English:
	//
	//  The text baseline is the middle of the em square.
	//
	// Português:
	//
	//  A linha de base do texto é o meio do quadrado em.
	KTextBaseLineRuleMiddle TextBaseLineRule = "middle"

	// KTextBaseLineRuleIdeographic
	//
	// English:
	//
	//  The text baseline is the ideographic baseline.
	//
	// Português:
	//
	//  The text baseline is the ideographic baseline.
	KTextBaseLineRuleIdeographic TextBaseLineRule = "ideographic"

	// KTextBaseLineRuleBottom
	//
	// English:
	//
	//  The text baseline is the bottom of the bounding box.
	//
	// Português:
	//
	//  A linha de base do texto é a parte inferior da caixa delimitadora.
	KTextBaseLineRuleBottom TextBaseLineRule = "bottom"
)
