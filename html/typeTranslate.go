package html

type Translate string

func (e Translate) String() (element string) {
	return string(e)
}

const (
	// KTranslateYes
	//
	// English:
	//
	//  The translate attribute specifies whether the content of an element should be translated.
	//
	// Português:
	//
	//  O atributo translate especifica se o conteúdo de um elemento deve ser traduzido.
	KTranslateYes Translate = "yes"

	// KTranslateNo
	//
	// English:
	//
	//  The translate attribute specifies whether the content of an element should not be translated.
	//
	// Português:
	//
	//  O atributo translate especifica se o conteúdo de um elemento não deve ser traduzido.
	KTranslateNo Translate = "no"
)
