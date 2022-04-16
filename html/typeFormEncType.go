package html

type FormEncType string

func (e FormEncType) String() string {
	return string(e)
}

const (
	// KFormEncTypeApplication
	//
	// English:
	//
	//  The default if the attribute is not used.
	//
	// Português:
	//
	//  O padrão se o atributo não for usado.
	KFormEncTypeApplication FormEncType = "application/x-www-form-urlencoded"

	// KFormEncTypeMultiPart
	//
	// English:
	//
	//  Use to submit <input> elements with their type attributes set to file.
	//
	// Português:
	//
	//  Use para enviar elementos <input> com seus atributos de tipo definidos para arquivo.
	KFormEncTypeMultiPart FormEncType = "multipart/form-data"

	// KFormEncTypeText
	//
	// English:
	//
	//  Specified as a debugging aid; shouldn't be used for real form submission.
	//
	// Português:
	//
	//  Especificado como auxiliar de depuração; não deve ser usado para envio de formulário real.
	KFormEncTypeText FormEncType = "text/plain"
)
