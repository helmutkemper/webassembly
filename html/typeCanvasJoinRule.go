package html

type CanvasJoinRule string

func (e CanvasJoinRule) String() string {
	return string(e)
}

const (
	// KJoinRuleBevel
	//
	// English:
	//
	//  Creates a beveled corner.
	//
	// Português:
	//
	//  Creates a beveled corner.
	KJoinRuleBevel CanvasJoinRule = "bevel"

	// KJoinRuleRound
	//
	// English:
	//
	//  A Creates a rounded corner.
	//
	// Português:
	//
	//  A Cria um canto arredondado.
	KJoinRuleRound CanvasJoinRule = "round"

	// KJoinRuleMiter
	//
	// English:
	//
	//  (Default) Creates a sharp corner.
	//
	// Português:
	//
	//  (Default) Cria um canto afiado.
	KJoinRuleMiter CanvasJoinRule = "miter"
)
