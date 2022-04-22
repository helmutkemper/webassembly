package html

type EnterKeyHint string

func (e EnterKeyHint) String() string {
	return string(e)
}

const (
	// EnterKeyHintEnter
	//
	// English:
	//
	//  typically indicating inserting a new line.
	//
	// Português:
	//
	//  normalmente indicando a inserção de uma nova linha.
	EnterKeyHintEnter EnterKeyHint = "enter"

	// EnterKeyHintDone
	//
	// English:
	//
	//  Typically meaning there is nothing more to input and the input method editor (IME) will be
	//  closed.
	//
	// Português:
	//
	//  Normalmente, significa que não há mais nada para inserir e o editor de método de entrada (IME)
	//  será fechado.
	EnterKeyHintDone EnterKeyHint = "done"

	// EnterKeyHintGo
	//
	// English:
	//
	//  Typically meaning to take the user to the target of the text they typed.
	//
	// Português:
	//
	//  Normalmente significa levar o usuário ao destino do texto digitado.
	EnterKeyHintGo EnterKeyHint = "go"

	// EnterKeyHintNext
	//
	// English:
	//
	//  Typically taking the user to the next field that will accept text.
	//
	// Português:
	//
	//  Normalmente levando o usuário para o próximo campo que aceitará texto.
	EnterKeyHintNext EnterKeyHint = "next"

	// EnterKeyHintPrevious
	//
	// English:
	//
	//  Typically taking the user to the previous field that will accept text.
	//
	// Português:
	//
	//  Normalmente levando o usuário ao campo anterior que aceitará texto.
	EnterKeyHintPrevious EnterKeyHint = "previous"

	// EnterKeyHintSearch
	//
	// English:
	//
	//  Typically taking the user to the results of searching for the text they have typed.
	//
	// Português:
	//
	//  Normalmente, leva o usuário aos resultados da pesquisa do texto que digitou.
	EnterKeyHintSearch EnterKeyHint = "search"

	// EnterKeyHintSend
	//
	// English:
	//
	//  Typically delivering the text to its target.
	//
	// Português:
	//
	//  Normalmente entregando o texto ao seu destino.
	EnterKeyHintSend EnterKeyHint = "send"
)
