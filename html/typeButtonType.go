package html

type ButtonType string

func (e ButtonType) String() string {
	return string(e)
}

const (
	// KButtonTypeSubmit
	//
	// English:
	//
	//  The button submits the form data to the server. This is the default if the attribute is not
	//  specified for buttons associated with a <form>, or if the attribute is an empty or invalid
	//  value.
	//
	// Português:
	//
	//  O botão envia os dados do formulário para o servidor. Este é o padrão se o atributo não for
	//  especificado para botões associados a um <form> ou se o atributo for um valor vazio ou inválido.
	KButtonTypeSubmit ButtonType = "submit"

	// KButtonTypeReset
	//
	// English:
	//
	//  The button resets all the controls to their initial values, like <input type="reset">.
	//  (This behavior tends to annoy users.)
	//
	// Português:
	//
	//  O botão redefine todos os controles para seus valores iniciais, como <input type="reset">.
	//  (Esse comportamento tende a incomodar os usuários.)
	KButtonTypeReset ButtonType = "reset"

	// KButtonTypeButton
	//
	// English:
	//
	//  The button has no default behavior, and does nothing when pressed by default. It can have
	//  client-side scripts listen to the element's events, which are triggered when the events occur.
	//
	// Português:
	//
	//  O botão não tem comportamento padrão e não faz nada quando pressionado por padrão. Ele pode
	//  fazer com que os scripts do lado do cliente escutem os eventos do elemento, que são acionados
	//  quando os eventos ocorrem.
	KButtonTypeButton ButtonType = "button"
)
