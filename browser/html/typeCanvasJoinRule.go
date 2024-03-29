package html

type JoinRule string

func (e JoinRule) String() string {
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
	KJoinRuleBevel JoinRule = "bevel"

	// KJoinRuleRound
	//
	// English:
	//
	//  A Creates a rounded corner.
	//
	// Português:
	//
	//  A Cria um canto arredondado.
	KJoinRuleRound JoinRule = "round"

	// KJoinRuleMiter
	//
	// English:
	//
	//  (Default) Creates a sharp corner.
	//
	// Português:
	//
	//  (Default) Cria um canto afiado.
	KJoinRuleMiter JoinRule = "miter"
)
