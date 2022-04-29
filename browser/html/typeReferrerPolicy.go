package html

type ReferrerPolicy string

func (e ReferrerPolicy) String() string {
	return string(e)
}

const (
	// KRefPolicyNoReferrer
	//
	// English:
	//
	//  The Referer header will not be sent.
	//
	// Português:
	//
	//  O cabeçalho Referer não será enviado.
	KRefPolicyNoReferrer ReferrerPolicy = "no-referrer"

	// KRefPolicyNoReferrerWhenDowngrade
	//
	// English:
	//
	//  The Referer header will not be sent to origins without TLS (HTTPS).
	//
	// Português:
	//
	//  O cabeçalho Referer não será enviado para origens sem TLS (HTTPS).
	KRefPolicyNoReferrerWhenDowngrade ReferrerPolicy = "no-referrer-when-downgrade"

	// KRefPolicyOrigin
	//
	// English:
	//
	//  The sent referrer will be limited to the origin of the referring page: its scheme, host, and
	//  port.
	//
	// Português:
	//
	//  O referenciador enviado será limitado à origem da página de referência: seu esquema, host e
	//  porta.
	KRefPolicyOrigin ReferrerPolicy = "origin"

	// KRefPolicyOriginWhenCrossOrigin
	//
	// English:
	//
	//  The referrer sent to other origins will be limited to the scheme, the host, and the port.
	//  Navigations on the same origin will still include the path.
	//
	// Português:
	//
	//  O referenciador enviado para outras origens será limitado ao esquema, ao host e à porta.
	//  As navegações na mesma origem ainda incluirão o caminho.
	KRefPolicyOriginWhenCrossOrigin ReferrerPolicy = "origin-when-cross-origin"

	// KRefPolicySameOrigin
	//
	// English:
	//
	//  A referrer will be sent for same origin, but cross-origin requests will contain no referrer
	//  information.
	//
	// Português:
	//
	//  Um referenciador será enviado para a mesma origem, mas as solicitações de origem cruzada não
	//  conterão informações de referenciador.
	KRefPolicySameOrigin ReferrerPolicy = "same-origin"

	// KRefPolicyStrictOrigin
	//
	// English:
	//
	//  Only send the origin of the document as the referrer when the protocol security level stays the
	//  same (HTTPS→HTTPS), but don't send it to a less secure destination (HTTPS→HTTP).
	//
	// Português:
	//
	//  Só envie a origem do documento como referenciador quando o nível de segurança do protocolo
	//  permanecer o mesmo (HTTPS→HTTPS), mas não envie para um destino menos seguro (HTTPS→HTTP).
	KRefPolicyStrictOrigin ReferrerPolicy = "strict-origin"

	// KRefPolicyStrictOriginWhenCrossOrigin
	//
	// English:
	//
	//  Send a full URL when performing a same-origin request, only send the origin when the protocol
	//  security level stays the same (HTTPS→HTTPS), and send no header to a less secure destination
	//  (HTTPS→HTTP). (default)
	//
	// Português:
	//
	//  Envie uma URL completa ao realizar uma solicitação de mesma origem, envie a origem apenas quando
	//  o nível de segurança do protocolo permanecer o mesmo (HTTPS→HTTPS) e não envie nenhum cabeçalho
	//  para um destino menos seguro (HTTPS→HTTP). (padrão)
	KRefPolicyStrictOriginWhenCrossOrigin ReferrerPolicy = "strict-origin-when-cross-origin"

	// KRefPolicyUnsafeUrl
	//
	// English:
	//
	//  The referrer will include the origin and the path (but not the fragment, password, or username).
	//  This value is unsafe, because it leaks origins and paths from TLS-protected resources to
	//  insecure origins.
	//
	// Português:
	//
	//  O referenciador incluirá a origem e o caminho (mas não o fragmento, a senha ou o nome de
	//  usuário).
	//  Esse valor não é seguro porque vaza origens e caminhos de recursos protegidos por TLS para
	//  origens inseguras.
	KRefPolicyUnsafeUrl ReferrerPolicy = "unsafe-url"
)
