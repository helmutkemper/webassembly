package html

type AutoCapitalize string

func (e AutoCapitalize) String() string {
	return string(e)
}

const (

	// KAutoCapitalizeOff
	//
	// English:
	//
	// No autocapitalization is applied (all letters default to lowercase)
	//
	// Português:
	//
	// Nenhuma capitalização automática é aplicada (todas as letras padrão para minúsculas)
	KAutoCapitalizeOff AutoCapitalize = "off"

	// KAutoCapitalizeNone
	//
	// English:
	//
	// No autocapitalization is applied (all letters default to lowercase)
	//
	// Português:
	//
	// Nenhuma capitalização automática é aplicada (todas as letras padrão para minúsculas)
	KAutoCapitalizeNone AutoCapitalize = "none"

	// KAutoCapitalizeOn
	//
	// English:
	//
	// The first letter of each sentence defaults to a capital letter; all other letters default to lowercase
	//
	// Português:
	//
	// A primeira letra de cada frase tem como padrão uma letra maiúscula; todas as outras letras padrão para minúsculas
	KAutoCapitalizeOn AutoCapitalize = "on"

	// KAutoCapitalizeSentences
	//
	// English:
	//
	// The first letter of each sentence defaults to a capital letter; all other letters default to lowercase
	//
	// Português:
	//
	// A primeira letra de cada frase tem como padrão uma letra maiúscula; todas as outras letras padrão para minúsculas
	KAutoCapitalizeSentences AutoCapitalize = "sentences"

	// KAutoCapitalizeWords
	//
	// English:
	//
	// The first letter of each word defaults to a capital letter; all other letters default to lowercase
	//
	// Português:
	//
	// A primeira letra de cada palavra tem como padrão uma letra maiúscula; todas as outras letras padrão para minúsculas
	KAutoCapitalizeWords AutoCapitalize = "words"

	// KAutoCapitalizeCharacters
	//
	// English:
	//
	// All letters should default to uppercase
	//
	// Português:
	//
	// Todas as letras devem ser maiúsculas por padrão
	KAutoCapitalizeCharacters AutoCapitalize = "characters"
)
