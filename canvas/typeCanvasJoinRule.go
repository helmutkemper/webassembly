package canvas

type CanvasJoinRule int

var CanvasJoinRules = [...]string{
	"",
	"bevel",
	"round",
	"miter",
}

func (el CanvasJoinRule) String() string {
	return CanvasJoinRules[el]
}

const (
	// en: Creates a beveled corner
	KJoinRuleBevel CanvasJoinRule = iota + 1

	// en: A Creates a rounded corner
	KJoinRuleRound

	// en: Default. Creates a sharp corner
	KJoinRuleMiter
)
