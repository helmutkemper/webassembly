package html

type CanvasRepeatRule string

func (e CanvasRepeatRule) String() string {
	return string(e)
}

const (
	// KRepeatRuleRepeat
	//
	// English:
	//
	//  (Default) The pattern repeats both horizontally and vertically.
	//
	// Português:
	//
	//  (Padrão) O padrão se repete horizontal e verticalmente.
	KRepeatRuleRepeat CanvasRepeatRule = "repeat"

	// KRepeatRuleRepeatX
	//
	// English:
	//
	//  The pattern repeats only horizontally.
	//
	// Português:
	//
	//  O padrão se repete apenas horizontalmente.
	KRepeatRuleRepeatX CanvasRepeatRule = "repeat-x"

	// KRepeatRuleRepeatY
	//
	// English:
	//
	//  The pattern repeats only vertically.
	//
	// Português:
	//
	//  O padrão se repete apenas verticalmente.
	KRepeatRuleRepeatY CanvasRepeatRule = "repeat-y"

	// KRepeatRuleNoRepeat
	//
	// English:
	//
	//  The pattern will be displayed only once (no repeat).
	//
	// Português:
	//
	//  The pattern will be displayed only once (no repeat).
	KRepeatRuleNoRepeat CanvasRepeatRule = "no-repeat"
)
