package html

type Translate int

func (e Translate) String() (element string) {
	if e == 0 {
		return
	}
	return translateString[e]
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
	KTranslateYes Translate = iota + 1

	// KTranslateNo
	//
	// English:
	//
	//  The translate attribute specifies whether the content of an element should not be translated.
	//
	// Português:
	//
	//  O atributo translate especifica se o conteúdo de um elemento não deve ser traduzido.
	KTranslateNo
)

var translateString = [...]string{
	"",
	"yes",
	"no",
}
