package html

type CanvasCapRule string

func (e CanvasCapRule) String() string {
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
	KCapRuleButt CanvasCapRule = "butt"

	// KCapRuleRound
	//
	// English:
	//
	//  A rounded end cap is added to each end of the line.
	//
	// Português:
	//
	//  Uma tampa de extremidade arredondada é adicionada a cada extremidade da linha.
	KCapRuleRound CanvasCapRule = "round"

	// KCapRuleSquare
	//
	// English:
	//
	//  A square end cap is added to each end of the line.
	//
	// Português:
	//
	//  Uma tampa de extremidade quadrada é adicionada a cada extremidade da linha.
	KCapRuleSquare CanvasCapRule = "square"
)
