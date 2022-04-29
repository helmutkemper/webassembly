package html

type FontAlignRule string

func (e FontAlignRule) String() string {
	return string(e)
}

const (
	// KFontAlignRuleStart
	//
	// English:
	//
	//  (Default) The text starts at the specified position.
	//
	// Português:
	//
	//  (Padrão) O texto começa na posição especificada.
	KFontAlignRuleStart FontAlignRule = "start"

	// KFontAlignRuleEnd
	//
	// English:
	//
	//  The text ends at the specified position.
	//
	// Português:
	//
	//  O texto termina na posição especificada.
	KFontAlignRuleEnd FontAlignRule = "end"

	// KFontAlignRuleCenter
	//
	// English:
	//
	//  The center of the text is placed at the specified position.
	//
	// Português:
	//
	//  O centro do texto é colocado na posição especificada.
	KFontAlignRuleCenter FontAlignRule = "center"

	// KFontAlignRuleLeft
	//
	// English:
	//
	//  The text starts at the specified position.
	//
	// Português:
	//
	//  O texto começa na posição especificada.
	KFontAlignRuleLeft FontAlignRule = "left"

	// KFontAlignRuleRight
	//
	// English:
	//
	//  The text ends at the specified position.
	//
	// Português:
	//
	//  O texto termina na posição especificada.
	KFontAlignRuleRight FontAlignRule = "right"
)
