package canvas

type CanvasRepeatRule int

var CanvasRepeatRules = [...]string{
	"",
	"repeat",
	"repeat-x",
	"repeat-y",
	"no-repeat",
}

func (el CanvasRepeatRule) String() string {
	return CanvasFillRules[el]
}

const (
	// en: Default. The pattern repeats both horizontally and vertically
	KRepeatRuleRepeat CanvasRepeatRule = iota + 1

	// en: The pattern repeats only horizontally
	KRepeatRuleRepeatX

	// en: The pattern repeats only vertically
	KRepeatRuleRepeatY

	// en: The pattern will be displayed only once (no repeat)
	KRepeatRuleNoRepeat
)
