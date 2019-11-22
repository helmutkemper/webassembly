package canvas

type CanvasCapRule int

var CanvasCapRules = [...]string{
	"",
	"butt",
	"round",
	"square",
}

func (el CanvasCapRule) String() string {
	return CanvasCapRules[el]
}

const (
	// en: Default. A flat edge is added to each end of the line
	KCapRuleButt CanvasCapRule = iota + 1

	// en: A rounded end cap is added to each end of the line
	KCapRuleRound

	// en: A square end cap is added to each end of the line
	KCapRuleSquare
)
