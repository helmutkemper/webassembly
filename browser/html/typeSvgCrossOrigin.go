package html

type SvgCrossOrigin string

func (e SvgCrossOrigin) String() string {
	return string(e)
}

const (
	// KSvgCrossOriginAnonymous
	//
	// English:
	//
	//  CORS requests for this element will have the credentials flag set to 'same-origin'.
	//
	// Português:
	//
	//  As solicitações CORS para este elemento terão o sinalizador de credenciais definido como 'mesma origem'.
	KSvgCrossOriginAnonymous SvgCrossOrigin = "anonymous"

	// KSvgCrossOriginUseCredentials
	//
	// English:
	//
	//  CORS requests for this element will have the credentials flag set to 'include'.
	//
	// Português:
	//
	//  As solicitações CORS para este elemento terão o sinalizador de credenciais definido como 'incluir'.
	KSvgCrossOriginUseCredentials SvgCrossOrigin = "use-credentials"
)
