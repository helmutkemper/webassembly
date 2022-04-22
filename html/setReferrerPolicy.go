package html

import "log"

// SetReferrerPolicy
//
// English:
//
//  How much of the referrer to send when following the link.
//
//   KRefPolicyNoReferrer: The Referer header will not be sent.
//   KRefPolicyNoReferrerWhenDowngrade: The Referer header will not be sent to origins without TLS
//     (HTTPS).
//   KRefPolicyOrigin: The sent referrer will be limited to the origin of the referring page: its
//     scheme, host, and port.
//   KRefPolicyOriginWhenCrossOrigin: The referrer sent to other origins will be limited to the
//     scheme, the host, and the port. Navigations on the same origin will still include the path.
//   KRefPolicySameOrigin: A referrer will be sent for same origin, but cross-origin requests will
//     contain no referrer information.
//   KRefPolicyStrictOrigin: Only send the origin of the document as the referrer when the protocol
//     security level stays the same (HTTPS→HTTPS), but don't send it to a less secure destination
//     (HTTPS→HTTP).
//   KRefPolicyStrictOriginWhenCrossOrigin (default): Send a full URL when performing a same-origin
//     request, only send the origin when the protocol security level stays the same (HTTPS→HTTPS),
//     and send no header to a less secure destination (HTTPS→HTTP).
//   KRefPolicyUnsafeUrl: The referrer will include the origin and the path (but not the fragment,
//     password, or username). This value is unsafe, because it leaks origins and paths from
//     TLS-protected resources to insecure origins.
//
//   Note:
//     * Experimental. Expect behavior to change in the future. (04/2022)
//
// Português:
//
//  Quanto do referenciador enviar ao seguir o link.
//
//   KRefPolicyNoReferrer: O cabeçalho Referer não será enviado.
//   KRefPolicyNoReferrerWhenDowngrade: O cabeçalho Referer não será enviado para origens sem TLS
//     (HTTPS).
//   KRefPolicyOrigin: O referenciador enviado será limitado à origem da página de referência: seu
//     esquema, host e porta.
//   KRefPolicyOriginWhenCrossOrigin: O referenciador enviado para outras origens será limitado ao
//     esquema, ao host e à porta. As navegações na mesma origem ainda incluirão o caminho.
//   KRefPolicySameOrigin: Um referenciador será enviado para a mesma origem, mas as solicitações de
//     origem cruzada não conterão informações de referenciador.
//   KRefPolicyStrictOrigin: Só envie a origem do documento como referenciador quando o nível de
//     segurança do protocolo permanecer o mesmo (HTTPS→HTTPS), mas não envie para um destino menos
//     seguro (HTTPS→HTTP).
//   KRefPolicyStrictOriginWhenCrossOrigin (padrão): Envie uma URL completa ao realizar uma
//     solicitação de mesma origem, envie a origem apenas quando o nível de segurança do protocolo
//     permanecer o mesmo (HTTPS→HTTPS) e não envie nenhum cabeçalho para um destino menos seguro
//     (HTTPS→HTTP).
//   KRefPolicyUnsafeUrl: O referenciador incluirá a origem e o caminho (mas não o fragmento, a senha
//     ou o nome de usuário). Esse valor não é seguro porque vaza origens e caminhos de recursos
//     protegidos por TLS para origens inseguras.
//
//   Note:
//     * Experimental. Expect behavior to change in the future. (04/2022)
func (e *GlobalAttributes) SetReferrerPolicy(referrerPolicy ReferrerPolicy) (ref *GlobalAttributes) {
	switch e.tag {
	case KTagA:
	default:
		log.Printf("tag " + e.tag.String() + " does not support referrerpolicy property")
	}

	e.selfElement.Set("referrerpolicy", referrerPolicy)
	return e
}
