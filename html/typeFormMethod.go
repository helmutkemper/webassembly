package html

type FormMethod string

func (e FormMethod) String() string {
	return string(e)
}

const (
	// KFormMethodPost
	//
	// English:
	//
	//  The data from the form are included in the body of the HTTP request when sent to the server.
	//  Use when the form contains information that shouldn't be public, like login credentials.
	//
	// Português:
	//
	//  Os dados do formulário são incluídos no corpo da solicitação HTTP quando enviados ao servidor.
	//  Use quando o formulário contém informações que não devem ser públicas, como credenciais de
	//  login.
	KFormMethodPost FormMethod = "post"

	// KFormMethodGet
	//
	// English:
	//
	//  The form data are appended to the form's action URL, with a ? as a separator, and the resulting
	//  URL is sent to the server. Use this method when the form has no side effects, like search forms.
	//
	// Português:
	//
	//  Os dados do formulário são anexados à URL de ação do formulário, com um ? como separador e a URL
	//  resultante é enviada ao servidor. Use este método quando o formulário não tiver efeitos
	//  colaterais, como formulários de pesquisa.
	KFormMethodGet FormMethod = "get"
)
