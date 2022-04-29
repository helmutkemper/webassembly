package html

type CrossOrigin string

func (e CrossOrigin) String() string {
	return string(e)
}

const (
	// KCrossOriginAnonymous
	//
	// English:
	//
	//  A CORS request is sent with credentials omitted (that is, no cookies, X.509 certificates, or
	//  Authorization request header).
	//
	// Português:
	//
	//  Uma solicitação CORS é enviada com credenciais omitidas (ou seja, sem cookies, certificados
	//  X.509 ou cabeçalho de solicitação de autorização).
	KCrossOriginAnonymous CrossOrigin = "anonymous"

	// KCrossOriginUseCredentials
	//
	// English:
	//
	//  The CORS request is sent with any credentials included (that is, cookies, X.509 certificates,
	//  and the Authorization request header).
	//
	// If the server does not opt into sharing credentials with the origin site (by sending back the
	// Access-Control-Allow-Credentials: true response header), then the browser marks the image as
	// tainted and restricts access to its image data.
	//
	// If the attribute has an invalid value, browsers handle it as if the anonymous value was used.
	// See CORS settings attributes for additional information.
	//
	// Português:
	//
	//  A solicitação CORS é enviada com todas as credenciais incluídas (ou seja, cookies, certificados
	//  X.509 e o cabeçalho da solicitação de autorização).
	//
	// Se o servidor não optar por compartilhar credenciais com o site de origem (enviando de volta o
	// cabeçalho Access-Control-Allow-Credentials: true response), o navegador marcará a imagem como
	// contaminada e restringirá o acesso aos dados da imagem.
	//
	// Se o atributo tiver um valor inválido, os navegadores o tratarão como se o valor anônimo fosse
	// usado. Consulte Atributos de configurações do CORS para obter informações adicionais.
	KCrossOriginUseCredentials CrossOrigin = "use-credentials"
)
