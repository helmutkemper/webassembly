package html

type CapRule string

func (e CapRule) String() string {
	return string(e)
}

const (
	// KCapRuleButt
	//
	// English:
	//
	//  (Default) A flat edge is added to each end of the line.
	//
	// Português:
	//
	//  (Padrão) Uma aresta plana é adicionada a cada extremidade da linha.
	KCapRuleButt CapRule = "butt"

	// KCapRuleRound
	//
	// English:
	//
	//  A rounded end cap is added to each end of the line.
	//
	// Português:
	//
	//  Uma tampa de extremidade arredondada é adicionada a cada extremidade da linha.
	KCapRuleRound CapRule = "round"

	// KCapRuleSquare
	//
	// English:
	//
	//  A square end cap is added to each end of the line.
	//
	// Português:
	//
	//  Uma tampa de extremidade quadrada é adicionada a cada extremidade da linha.
	KCapRuleSquare CapRule = "square"
)
